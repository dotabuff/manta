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
		int32(dota.SVC_Messages_svc_UpdateStringTable):
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

// Provides a sortable structure for storing packets in the same tick.
type pendingMessages []*pendingMessage

func (ms pendingMessages) Len() int      { return len(ms) }
func (ms pendingMessages) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }
func (ms pendingMessages) Less(i, j int) bool {
	if ms[i].tick < ms[j].tick {
		return true
	}

	// String tables first
	return ms[i].priority() < ms[j].priority()
}

// Internal parser for callback OnCDemoPacket, responsible for extracting
// multiple inner packets from a single CDemoPacket. This is the main payload
// for most data in the replay.
func (p *Parser) onCDemoPacket(m *dota.CDemoPacket) error {
	r := newDemoPacketReader(m.GetData())

	ms := make(pendingMessages, 0)

	// Collect all messages from the packet.
	for r.hasNext() {
		t, buf := r.readNext()
		ms = append(ms, &pendingMessage{p.Tick, t, buf})
	}

	// Sort messages to ensure context dependencies are met.
	sort.Sort(ms)

	// Dispatch messages in order.
	for _, m := range ms {
		// Skip message we don't have a definition for (yet)
		// XXX TODO: remove this when we get updated protos.
		if m.t == 547 || m.t == 400 {
			continue
		}

		// Call each packet, panic if we encounter an error.
		// XXX TODO: this should return the error up the chain. Panic for debugging.
		if err := p.CallByPacketType(m.t, m.buf); err != nil {
			panic(err)
		}
	}

	return nil
}

// Creates a new demoPacketHandler, used to read data in the expected format.
func newDemoPacketReader(buf []byte) *demoPacketReader {
	return &demoPacketReader{newReader(buf)}
}

// Reads a series of inner packets from a CDemoPacket buffer
type demoPacketReader struct {
	r *reader
}

// Determines whether or not another packet is available.
// XXX TODO: this seems wrong, we may be skipping the last packet or some
// other value at the end of the buffer.
func (r *demoPacketReader) hasNext() bool {
	return r.r.remBits() > 10
}

// Reads the next packet, returning a type and inner buffer.
// XXX TODO: detail our knowledge of the structure of this packet.
func (r *demoPacketReader) readNext() (int32, []byte) {
	t := r.r.readBits(6)

	if header := t >> 4; header != 0 {
		bits := int(header*4 + (((2 - header) >> 31) & 16))
		t = (t & 15) | (r.r.readBits(bits) << 4)
	}

	size := r.r.readVarUint32()
	buf := r.r.readBytes(int(size))

	return int32(t), buf
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
