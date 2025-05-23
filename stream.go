package manta

import (
	"io"
	"sync"

	"github.com/dotabuff/manta/dota"
)

const (
	bufferInitial = 1024 * 100     // 100KB initial buffer
	bufferMax     = 1024 * 1024 * 4 // 4MB max buffer size for pooling
)

// Size classes for buffer pools (powers of 2 for efficient allocation)
var bufferSizeClasses = []uint32{
	1024 * 100,   // 100KB
	1024 * 200,   // 200KB  
	1024 * 400,   // 400KB
	1024 * 800,   // 800KB
	1024 * 1600,  // 1.6MB
	1024 * 3200,  // 3.2MB
}

// Size-class based buffer pools to reduce allocations
var streamBufferPools = make([]*sync.Pool, len(bufferSizeClasses))

func init() {
	// Initialize pools for each size class
	for i, size := range bufferSizeClasses {
		poolSize := size // Capture for closure
		streamBufferPools[i] = &sync.Pool{
			New: func() interface{} {
				return make([]byte, poolSize)
			},
		}
	}
}

// getBufferSizeClass returns the index of the smallest size class that can fit the requested size
func getBufferSizeClass(requestedSize uint32) int {
	for i, classSize := range bufferSizeClasses {
		if requestedSize <= classSize {
			return i
		}
	}
	return -1 // Size too large for pooling
}

// getPooledBuffer gets a buffer from the appropriate size class pool
func getPooledBuffer(requestedSize uint32) ([]byte, int) {
	classIndex := getBufferSizeClass(requestedSize)
	if classIndex == -1 {
		// Size too large for pooling, allocate directly
		return make([]byte, requestedSize), -1
	}
	
	buf := streamBufferPools[classIndex].Get().([]byte)
	return buf, classIndex
}

// returnPooledBuffer returns a buffer to the appropriate pool
func returnPooledBuffer(buf []byte, classIndex int) {
	if classIndex >= 0 && classIndex < len(streamBufferPools) {
		streamBufferPools[classIndex].Put(buf)
	}
	// If classIndex is -1, it was directly allocated and will be GC'd
}

// stream wraps an io.Reader to provide functions necessary for reading the
// outer replay structure.
type stream struct {
	io.Reader
	buf        []byte
	size       uint32
	pooledBuf  bool // tracks if buf came from pool
	classIndex int  // tracks which pool class this buffer came from (-1 if not pooled)
}

// newStream creates a new stream from a given io.Reader
func newStream(r io.Reader) *stream {
	buf, classIndex := getPooledBuffer(bufferInitial)
	return &stream{
		Reader:     r,
		buf:        buf,
		size:       uint32(len(buf)),
		pooledBuf:  classIndex >= 0,
		classIndex: classIndex,
	}
}

// Close returns the buffer to the pool if it was pooled
func (s *stream) Close() {
	if s.pooledBuf {
		returnPooledBuffer(s.buf, s.classIndex)
	}
	s.pooledBuf = false
	s.classIndex = -1
}

// readBytes reads the given number of bytes from the reader
func (s *stream) readBytes(n uint32) ([]byte, error) {
	if n > s.size {
		// Return current buffer to pool if it was pooled
		if s.pooledBuf {
			returnPooledBuffer(s.buf, s.classIndex)
		}
		
		// Grow buffer intelligently: either 2x current size or requested size, whichever is larger
		newSize := s.size * 2
		if n > newSize {
			newSize = n
		}
		
		// Get new buffer from appropriate size class pool
		newBuf, newClassIndex := getPooledBuffer(newSize)
		
		s.buf = newBuf
		s.size = uint32(len(newBuf))
		s.pooledBuf = newClassIndex >= 0
		s.classIndex = newClassIndex
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
