package manta

import (
	"io"

	"github.com/dotabuff/manta/dota"
)

const buffer = 1024 * 100

// stream wraps an io.Reader to provide functions necessary for reading the
// outer replay structure.
type stream struct {
	io.Reader
	buf  []byte
	size uint32
}

// newStream creates a new stream from a given io.Reader
func newStream(r io.Reader) *stream {
	return &stream{r, make([]byte, buffer), buffer}
}

// readBytes reads the given number of bytes from the reader
func (s *stream) readBytes(n uint32) ([]byte, error) {
	if n > s.size {
		s.buf = make([]byte, n)
		s.size = n
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
