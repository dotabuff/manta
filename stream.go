package manta

import (
	"bufio"
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

	// mem is set for a stream over an in-memory buffer (see newBufferStream).
	// Reads then return subslices of data, advancing pos, instead of copying
	// through Reader into buf.
	mem  bool
	data []byte
	pos  int
}

// newStream creates a new stream from a given io.Reader. If the reader does not
// already provide buffered byte access (e.g. an *os.File), it is wrapped in a
// bufio.Reader so the per-byte varint reads in the outer loop do not issue one
// read syscall per byte. In-memory readers such as *bytes.Reader already
// satisfy io.ByteReader and are left unwrapped to avoid a redundant copy.
func newStream(r io.Reader) *stream {
	if _, ok := r.(io.ByteReader); !ok {
		r = bufio.NewReaderSize(r, buffer)
	}
	return &stream{Reader: r, buf: make([]byte, buffer), size: buffer}
}

// newBufferStream creates a stream that reads directly out of buf. The
// returned byte slices alias buf (cap-limited, so an append cannot write past
// them); the parser only reads them, and copies anything it retains.
func newBufferStream(buf []byte) *stream {
	return &stream{mem: true, data: buf}
}

// readBytes reads the given number of bytes from the reader
func (s *stream) readBytes(n uint32) ([]byte, error) {
	if s.mem {
		return s.sliceBytes(n)
	}

	if n > s.size {
		s.buf = make([]byte, n)
		s.size = n
	}

	if _, err := io.ReadFull(s.Reader, s.buf[:n]); err != nil {
		return nil, err
	}

	return s.buf[:n], nil
}

// sliceBytes returns the next n bytes of an in-memory stream without copying.
// It reproduces io.ReadFull's errors: io.EOF when no bytes remain, and
// io.ErrUnexpectedEOF when some but fewer than n do.
func (s *stream) sliceBytes(n uint32) ([]byte, error) {
	rem := len(s.data) - s.pos
	if uint64(n) > uint64(rem) {
		s.pos = len(s.data)
		if rem == 0 {
			return nil, io.EOF
		}
		return nil, io.ErrUnexpectedEOF
	}
	end := s.pos + int(n)
	buf := s.data[s.pos:end:end]
	s.pos = end
	return buf, nil
}

// readByte reads a single byte from the reader
func (s *stream) readByte() (byte, error) {
	if s.mem {
		if s.pos >= len(s.data) {
			return 0, io.EOF
		}
		b := s.data[s.pos]
		s.pos++
		return b, nil
	}

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
