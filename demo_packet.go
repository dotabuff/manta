package manta

import (
	"sort"

	"github.com/dotabuff/manta/dota"
)

// A message that has been read from an outerMessage but not yet processed.
type pendingMessage struct {
	tick uint32
	t    int32
	buf  []byte
}

// Calculates the priority of the message. Lower is more important.
func (m *pendingMessage) priority() int {
	switch m.t {
	case
		// These messages provide context needed for the rest of the tick
		// and should have the highest priority.
		int32(dota.NET_Messages_net_Tick),
		int32(dota.SVC_Messages_svc_CreateStringTable),
		int32(dota.SVC_Messages_svc_UpdateStringTable),
		int32(dota.NET_Messages_net_SpawnGroup_Load):
		return -10

	case
		// These messages benefit from having context but may also need to
		// provide context in terms of delta updates.
		int32(dota.SVC_Messages_svc_PacketEntities):
		return 5

	case
		// These messages benefit from having as much context as possible and
		// should have the lowest priority.
		int32(dota.EBaseGameEvents_GE_Source1LegacyGameEvent):
		return 10
	}

	return 0
}

// Provides a sortable structure for storing messages in the same packet.
type pendingMessages []*pendingMessage

func (ms pendingMessages) Len() int      { return len(ms) }
func (ms pendingMessages) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }
func (ms pendingMessages) Less(i, j int) bool {
	if ms[i].tick > ms[j].tick {
		return false
	}
	if ms[i].tick < ms[j].tick {
		return true
	}
	return ms[i].priority() < ms[j].priority()
}

// Internal parser for callback OnCDemoPacket, responsible for extracting
// multiple inner packets from a single CDemoPacket. This is the main structure
// that contains all other data types in the demo file.
func (p *Parser) onCDemoPacket(m *dota.CDemoPacket) error {
	// Create a slice to store pending mesages. Messages are read first as
	// pending messages then sorted before dispatch.
	ms := make(pendingMessages, 0, 2)

	// Read all messages from the buffer. Messages are packed serially as
	// {type, size, data}. We keep reading until until less than a byte remains.
	r := newReader(m.GetData())
	for r.remBytes() > 0 {
		t := int32(r.readUBitVar())
		size := r.readVarUint32()
		buf := r.readBytes(size)
		ms = append(ms, &pendingMessage{p.Tick, t, buf})
	}

	// Sort messages to ensure dependencies are met. For example, we need to
	// process string tables before game events that may reference them.
	sort.Sort(ms)

	// Dispatch messages in order, returning on handler error.
	for _, m := range ms {
		if err := p.Callbacks.callByPacketType(m.t, m.buf); err != nil {
			return err
		}
	}

	return nil
}

// Internal parser for callback OnCDemoFullPacket.
func (p *Parser) onCDemoFullPacket(m *dota.CDemoFullPacket) error {
	// Per Valve docs, parse the CDemoStringTables first.
	if m.StringTable != nil {
		if err := p.onCDemoStringTables(m.GetStringTable()); err != nil {
			return err
		}
	}

	// Then the CDemoPacket.
	if m.Packet != nil {
		if err := p.onCDemoPacket(m.GetPacket()); err != nil {
			return err
		}
	}

	return nil
}
