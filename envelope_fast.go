package manta

import (
	"github.com/dotabuff/manta/dota"
	"google.golang.org/protobuf/encoding/protowire"
)

// Hand-rolled envelope decoders for the hottest internal messages.
//
// The reflective protobuf unmarshal allocates the message, a pointer per
// scalar field, and — dominating the profile — a fresh copy of every bytes
// field (CDemoPacket.data, CSVCMsg_PacketEntities.entity_data,
// CSVCMsg_UpdateStringTable.string_data effectively copy the whole replay a
// second time). The parser's internal handlers only need a few scalar fields
// plus the payload, and consume the payload synchronously, so when a message
// type has no user callbacks registered we can decode the envelope by hand
// with protowire and alias the payload instead of copying it.
//
// Gating: NewParser registers exactly one internal handler per relevant list
// before returning, and users can only register afterwards, so
// len(list) == 1 means "internal handler only". Any user registration makes
// the list longer and permanently reverts that type to the full protobuf
// path, where the user-visible message owns its own copies as before.
//
// Aliasing lifetimes: entity_data and string_data alias the demo-packet arena
// (p.packetArena), which is stable for the duration of the packet's dispatch
// loop; both are fully consumed before the handler returns (string-table
// values are copied out by readBitsAsBytes/readString). CDemoPacket.data
// aliases the outer-message buffer, which is stable until the next
// readOuterMessage; processDemoPacket copies message bodies into the arena
// before dispatching.

// dispatchDemo routes an outer demo message, taking the fast envelope path
// for CDemoPacket/CDemoSignonPacket when only the internal handler is
// registered.
func (p *Parser) dispatchDemo(t int32, buf []byte) error {
	switch t {
	case int32(dota.EDemoCommands_DEM_Packet):
		if len(p.Callbacks.onCDemoPacket) == 1 {
			return p.fastDemoPacket(buf)
		}
	case int32(dota.EDemoCommands_DEM_SignonPacket):
		if len(p.Callbacks.onCDemoSignonPacket) == 1 {
			return p.fastDemoPacket(buf)
		}
	}
	return p.Callbacks.callByDemoType(t, buf)
}

// dispatchPacket routes an embedded packet message, taking the fast envelope
// path for the hot types when only the internal handler is registered.
func (p *Parser) dispatchPacket(t int32, buf []byte) error {
	switch t {
	case int32(dota.NET_Messages_net_Tick):
		if len(p.Callbacks.onCNETMsg_Tick) == 1 {
			return p.fastNetTick(buf)
		}
	case int32(dota.SVC_Messages_svc_UpdateStringTable):
		if len(p.Callbacks.onCSVCMsg_UpdateStringTable) == 1 {
			return p.fastUpdateStringTable(buf)
		}
	case int32(dota.SVC_Messages_svc_PacketEntities):
		if len(p.Callbacks.onCSVCMsg_PacketEntities) == 1 {
			return p.fastPacketEntities(buf)
		}
	}
	return p.Callbacks.callByPacketType(t, buf)
}

// fastDemoPacket decodes the CDemoPacket envelope: data = field 3 (bytes).
func (p *Parser) fastDemoPacket(buf []byte) error {
	var data []byte
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			return _errorf("fastDemoPacket: invalid tag")
		}
		buf = buf[n:]
		if num == 3 && typ == protowire.BytesType {
			v, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				return _errorf("fastDemoPacket: invalid data field")
			}
			data = v
			buf = buf[n:]
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, buf)
		if n < 0 {
			return _errorf("fastDemoPacket: invalid field %d", num)
		}
		buf = buf[n:]
	}
	return p.processDemoPacket(data)
}

// fastNetTick decodes the CNETMsg_Tick envelope: tick = field 1 (varint).
func (p *Parser) fastNetTick(buf []byte) error {
	var tick uint64
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			return _errorf("fastNetTick: invalid tag")
		}
		buf = buf[n:]
		if num == 1 && typ == protowire.VarintType {
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return _errorf("fastNetTick: invalid tick field")
			}
			tick = v
			buf = buf[n:]
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, buf)
		if n < 0 {
			return _errorf("fastNetTick: invalid field %d", num)
		}
		buf = buf[n:]
	}
	p.NetTick = uint32(tick)
	return nil
}

// fastUpdateStringTable decodes the CSVCMsg_UpdateStringTable envelope:
// table_id = 1 (varint), num_changed_entries = 2 (varint),
// string_data = 3 (bytes).
func (p *Parser) fastUpdateStringTable(buf []byte) error {
	var tableId, numChanged int32
	var data []byte
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			return _errorf("fastUpdateStringTable: invalid tag")
		}
		buf = buf[n:]
		switch {
		case num == 1 && typ == protowire.VarintType:
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return _errorf("fastUpdateStringTable: invalid table_id")
			}
			tableId = int32(v)
			buf = buf[n:]
		case num == 2 && typ == protowire.VarintType:
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return _errorf("fastUpdateStringTable: invalid num_changed_entries")
			}
			numChanged = int32(v)
			buf = buf[n:]
		case num == 3 && typ == protowire.BytesType:
			v, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				return _errorf("fastUpdateStringTable: invalid string_data")
			}
			data = v
			buf = buf[n:]
		default:
			n = protowire.ConsumeFieldValue(num, typ, buf)
			if n < 0 {
				return _errorf("fastUpdateStringTable: invalid field %d", num)
			}
			buf = buf[n:]
		}
	}
	return p.processUpdateStringTable(tableId, numChanged, data)
}

// fastPacketEntities decodes the CSVCMsg_PacketEntities envelope:
// updated_entries = 2 (varint), legacy_is_delta = 3 (varint bool),
// entity_data = 7 (bytes).
func (p *Parser) fastPacketEntities(buf []byte) error {
	var updatedEntries int32
	var isDelta bool
	var entityData []byte
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			return _errorf("fastPacketEntities: invalid tag")
		}
		buf = buf[n:]
		switch {
		case num == 2 && typ == protowire.VarintType:
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return _errorf("fastPacketEntities: invalid updated_entries")
			}
			updatedEntries = int32(v)
			buf = buf[n:]
		case num == 3 && typ == protowire.VarintType:
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return _errorf("fastPacketEntities: invalid legacy_is_delta")
			}
			isDelta = v != 0
			buf = buf[n:]
		case num == 7 && typ == protowire.BytesType:
			v, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				return _errorf("fastPacketEntities: invalid entity_data")
			}
			entityData = v
			buf = buf[n:]
		default:
			n = protowire.ConsumeFieldValue(num, typ, buf)
			if n < 0 {
				return _errorf("fastPacketEntities: invalid field %d", num)
			}
			buf = buf[n:]
		}
	}
	return p.processPacketEntities(entityData, updatedEntries, isDelta)
}
