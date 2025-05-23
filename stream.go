package manta

import (
	"io"
	"sync"

	"github.com/dotabuff/manta/dota"
)

const (
	bufferInitial = 1024 * 100 // 100KB initial buffer
	bufferMax     = 1024 * 1024 * 4 // 4MB max buffer size for pooling
)

// Buffer pool for stream buffers to reduce allocations
var streamBufferPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, bufferInitial)
	},
}

// stream wraps an io.Reader to provide functions necessary for reading the
// outer replay structure.
type stream struct {
	io.Reader
	buf        []byte
	size       uint32
	pooledBuf  bool // tracks if buf came from pool
}

// newStream creates a new stream from a given io.Reader
func newStream(r io.Reader) *stream {
	buf := streamBufferPool.Get().([]byte)
	return &stream{
		Reader:    r,
		buf:       buf,
		size:      uint32(len(buf)),
		pooledBuf: true,
	}
}

// Close returns the buffer to the pool if it was pooled
func (s *stream) Close() {
	if s.pooledBuf && len(s.buf) <= bufferMax {
		streamBufferPool.Put(s.buf)
	}
	s.pooledBuf = false
}

// readBytes reads the given number of bytes from the reader
func (s *stream) readBytes(n uint32) ([]byte, error) {
	if n > s.size {
		// Grow buffer intelligently: either 2x current size or requested size, whichever is larger
		newSize := s.size * 2
		if n > newSize {
			newSize = n
		}
		
		// For very large buffers, don't use pool to avoid memory pressure
		if newSize > bufferMax {
			s.buf = make([]byte, newSize)
			s.pooledBuf = false
		} else {
			// Try to get a larger buffer from pool first
			if s.pooledBuf {
				streamBufferPool.Put(s.buf)
			}
			s.buf = make([]byte, newSize) // Pool doesn't have size classes, so allocate directly
			s.pooledBuf = false // Mark as non-pooled since we made it ourselves
		}
		s.size = newSize
	}

	if _, err := io.ReadFull(s.Reader, s.buf[:n]); err != nil {
		return nil, err
	}

	return s.buf[:n], nil
}

// readByte reads a single byte from the reader
func (s *stream) readByte() (byte, error) {
	buf, err := s.readBytes(1)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

// readCommand reads a varuint32 as an EDemoCommands
func (s *stream) readCommand() (dota.EDemoCommands, error) {
	c, err := s.readVarUint32()
	return dota.EDemoCommands(c), err
}

// readVarUint32 reads an unsigned 32-bit varint
func (s *stream) readVarUint32() (uint32, error) {
	var x, y uint32
	for {
		b, err := s.readByte()
		if err != nil {
			return 0, err
		}
		u := uint32(b)
		x |= (u & 0x7F) << y
		y += 7
		if ((u & 0x80) == 0) || (y == 35) {
			break
		}
	}

	return x, nil
}
