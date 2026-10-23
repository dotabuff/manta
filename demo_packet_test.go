package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
)

// TestProcessDemoPacketReadsBufferedTail verifies processDemoPacket dispatches
// every embedded message even when the trailing messages have already been
// pulled into the bit reader's accumulator. The word-at-a-time refill can load
// up to 8 bytes past the current message, leaving r.pos at the end of the
// buffer while whole messages are still buffered, so the loop must test the
// logically unread bits rather than r.remBytes().
func TestProcessDemoPacketReadsBufferedTail(t *testing.T) {
	for bodyLen := 0; bodyLen < 48; bodyLen++ {
		for tail := 1; tail <= 4; tail++ {
			// One CNETMsg_SplitScreenUser (type 3) with an opaque body of
			// bodyLen bytes, then `tail` empty CNETMsg_NOP (type 0) messages.
			// Each header is a 6-bit ubitvar type plus a varint size, so every
			// body is bit-shifted.
			w := &bitWriter{}
			w.write(3, 6)
			w.write(uint32(bodyLen), 8)
			for i := 0; i < bodyLen; i++ {
				w.write(0, 8)
			}
			for i := 0; i < tail; i++ {
				w.write(0, 6)
				w.write(0, 8)
			}

			p := &Parser{Callbacks: newCallbacks()}
			got := 0
			p.Callbacks.OnCNETMsg_NOP(func(*dota.CNETMsg_NOP) error {
				got++
				return nil
			})
			if err := p.processDemoPacket(w.buf); err != nil {
				t.Fatalf("bodyLen %d tail %d: %s", bodyLen, tail, err)
			}
			if got != tail {
				t.Errorf("bodyLen %d: dispatched %d of %d trailing messages", bodyLen, got, tail)
			}
		}
	}
}
