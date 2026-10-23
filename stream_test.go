package manta

import (
	"bytes"
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBufferStreamMatchesReader verifies the zero-copy in-memory stream used by
// NewParser returns the same bytes and the same errors (io.EOF at a clean end,
// io.ErrUnexpectedEOF on a short read) as the io.Reader stream, so the outer
// message loop terminates identically on complete and truncated replays.
func TestBufferStreamMatchesReader(t *testing.T) {
	assert := assert.New(t)
	rng := rand.New(rand.NewSource(1))

	for trial := 0; trial < 500; trial++ {
		data := make([]byte, rng.Intn(64))
		rng.Read(data)

		mem := newBufferStream(data)
		rdr := newStream(bytes.NewReader(data))

		for step := 0; step < 40; step++ {
			switch rng.Intn(3) {
			case 0:
				n := uint32(rng.Intn(12))
				a, errA := mem.readBytes(n)
				b, errB := rdr.readBytes(n)
				assert.Equal(errB, errA, "readBytes(%d) trial %d step %d", n, trial, step)
				if errB == nil {
					assert.Equal(b, a, "readBytes(%d) trial %d step %d", n, trial, step)
					assert.Equal(len(a), cap(a), "readBytes result is cap-limited")
				}
			case 1:
				a, errA := mem.readByte()
				b, errB := rdr.readByte()
				assert.Equal(errB, errA, "readByte trial %d step %d", trial, step)
				assert.Equal(b, a, "readByte trial %d step %d", trial, step)
			case 2:
				a, errA := mem.readVarUint32()
				b, errB := rdr.readVarUint32()
				assert.Equal(errB, errA, "readVarUint32 trial %d step %d", trial, step)
				assert.Equal(b, a, "readVarUint32 trial %d step %d", trial, step)
			}
		}
	}

	// Explicit boundary cases.
	s := newBufferStream([]byte{1, 2, 3})
	_, err := s.readBytes(4)
	assert.Equal(io.ErrUnexpectedEOF, err)
	_, err = s.readBytes(1)
	assert.Equal(io.EOF, err)
	b, err := s.readBytes(0)
	assert.NoError(err)
	assert.Empty(b)
}
