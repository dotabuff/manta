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
func (m pendingMessage) priority() int {
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
type pendingMessages []pendingMessage

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
	return p.processDemoPacket(m.GetData())
}

// processDemoPacket is the core of the CDemoPacket handler. It takes the raw
// packet payload so the fast envelope path (envelope_fast.go) can call it
// without materializing a proto message.
func (p *Parser) processDemoPacket(data []byte) error {
	// Reuse a parser-level buffer to store pending messages. Messages are read
	// first as pending messages then sorted before dispatch. onCDemoPacket is
	// never re-entrant (it is dispatched only via callByDemoType, never nested
	// within a callByPacketType call), so a single reused backing array is safe
	// and avoids a heap allocation per embedded message.
	ms := p.pendingMsgBuf[:0]

	// The inner stream is bit-shifted after the leading 6-bit readUBitVar, so
	// message bodies almost never sit on a byte boundary and must be copied out.
	// Carve those copies from a single reused arena sized to the packet instead
	// of allocating per message: the buffers only live until dispatch below (the
	// protobuf unmarshal copies what it keeps), so reusing the arena across
	// packets is safe for the same reason reusing pendingMsgBuf is. Message
	// headers take space too, so the payload total always fits in len(data).
	if cap(p.packetArena) < len(data) {
		p.packetArena = make([]byte, 0, len(data))
	}
	arena := p.packetArena[:0]

	// Read all messages from the buffer. Messages are packed serially as
	// {type, size, data}. We keep reading until until less than a byte remains.
	r := newReader(data)
	for r.remBytes() > 0 {
		t := int32(r.readUBitVar())
		size := r.readVarUint32()
		start := len(arena)
		end := start + int(size)
		if end > cap(arena) {
			_panicf("onCDemoPacket: message size %d exceeds packet buffer", size)
		}
		arena = arena[:end]
		r.readBytesInto(arena[start:end])
		ms = append(ms, pendingMessage{p.Tick, t, arena[start:end:end]})
	}

	// Sort messages to ensure dependencies are met. For example, we need to
	// process string tables before game events that may reference them. A
	// stable sort keeps equal-priority messages in their original file order
	// and avoids the reflection allocations of sort.Sort's interface path.
	sort.Stable(ms)

	// Dispatch messages in order, stopping on handler error. dispatchPacket
	// takes the fast envelope path for hot internal-only message types and
	// falls back to the full protobuf callback path otherwise.
	var err error
	for i := range ms {
		if err = p.dispatchPacket(ms[i].t, ms[i].buf); err != nil {
			break
		}
	}

	// Release the inner-packet buffer references and keep the slice at length
	// zero so the reused backing array does not retain packet data. This runs on
	// every path (success or error), before the single return.
	clear(ms)
	p.pendingMsgBuf = ms[:0]
	return err
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
		if err := p.processDemoPacket(m.GetPacket().GetData()); err != nil {
			return err
		}
	}

	return nil
}
