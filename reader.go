package manta

import (
	"encoding/binary"
)

var (
	bigEndian    = binary.BigEndian
	littleEndian = binary.LittleEndian
)

// A reader holds a buffer and performs read operations against it.
type reader struct {
	buf  []byte
	size int
	pos  int
}

// Creates a new reader object with a given buffer.
func newReader(buf []byte) *reader {
	return &reader{buf, len(buf) * 8, 0}
}

// Calculates our byte position.
func (r *reader) byte_pos() int {
	return r.pos / 8
}

// Calculates how many bits are remaining.
func (r *reader) rem_bits() int {
	return r.size - r.pos
}

// Calculates how many bytes are remaining.
func (r *reader) rem_bytes() int {
	return (r.size - r.pos) / 8
}

// Seeks a given number of bits (may be negative).
func (r *reader) seek_bits(n int) {
	if r.pos+n >= r.size || r.pos+n < 0 {
		_panicf("seek overflow: %d bits requested, only %d remaining", n, r.size-r.pos)
	}
	r.pos += n
}

// Seeks a given number of bytes (may be negative).
func (r *reader) seek_bytes(n int) {
	r.seek_bits(n * 8)
}

// Reads a little-endian uint16.
func (r *reader) read_le_uint16() uint16 {
	return littleEndian.Uint16(r.read_bytes(2))
}

// Reads a little-endian uint32.
func (r *reader) read_le_uint32() uint32 {
	return littleEndian.Uint32(r.read_bytes(4))
}

// Reads a little-endian uint64.
func (r *reader) read_le_uint64() uint64 {
	return littleEndian.Uint64(r.read_bytes(8))
}

// Reads a big-endian uint16.
func (r *reader) read_be_uint16() uint16 {
	return bigEndian.Uint16(r.read_bytes(2))
}

// Reads a big-endian uint32.
func (r *reader) read_be_uint32() uint32 {
	return bigEndian.Uint32(r.read_bytes(4))
}

// Reads a big-endian uint64.
func (r *reader) read_be_uint64() uint64 {
	return bigEndian.Uint64(r.read_bytes(8))
}

// Reads an unsigned 32-bit varint.
func (r *reader) read_var_uint32() uint32 {
	var x uint
	var s uint
	for {
		b := uint(r.read_byte())
		x |= (b & 0x7F) << s
		s += 7
		if ((b & 0x80) == 0) || (s == 35) {
			break
		}
	}

	return uint32(x)
}

// Reads an unsigned 64-bit varint.
func (r *reader) read_var_uint64() uint64 {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		b := r.read_byte()
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				_panicf("read overflow: varint overflows uint64")
			}
			return x | uint64(b)<<s
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}

// Reads a signed 64-bit varint.
func (r *reader) read_var_sint64() int64 {
	ux := r.read_var_uint64()
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
}

// Reads a boolean value.
// TODO XXX: untested.
func (r *reader) read_boolean() bool {
	if r.rem_bits() < 1 {
		_panicf("read overflow: no bits left")
	}

	b := r.buf[r.pos/8]&(1<<uint(r.pos%8)) != 0
	r.pos += 1
	return b
}

// Reads the next byte (8 bits) in the buffer.
func (r *reader) read_byte() byte {
	return r.read_bytes(1)[0]
}

// Reads the given number of bytes from the buffer.
func (r *reader) read_bytes(n int) []byte {
	if r.rem_bits() < (n * 8) {
		_panicf("read overflow: %d bits requested, only %d remaining", n*8, r.size-r.pos)
	}

	// Fast path if our position is byte-aligned.
	if r.pos%8 == 0 {
		bpos := r.pos / 8
		r.pos += (n * 8)
		return r.buf[bpos : bpos+n]
	}

	// Slow path if our position isn't byte aligned.
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = byte(r.read_bits(8))
	}
	return buf
}

// Reads a string of a given length.
func (r *reader) read_string_n(n int) string {
	return string(r.read_bytes(n))
}

// Read a null terminated string.
func (r *reader) read_string() string {
	buf := make([]byte, 0)
	for {
		b := r.read_byte()
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}

	return string(buf)
}

// Reads bits as bytes.
func (r *reader) read_bits_as_bytes(n int) []byte {
	buf := make([]byte, (n+7)/8)
	i := 0
	for n > 7 {
		n -= 8
		buf[i] = byte(r.read_bits(8))
		i++
	}
	if n != 0 {
		buf[i] = byte(r.read_bits(8))
	}
	return buf
}

// Read bits of a given length as a uint, may or may not be byte-aligned.
func (r *reader) read_bits(n int) uint {
	if r.rem_bits() < n {
		_panicf("read overflow: %d bits requested, only %d remaining", n, r.size-r.pos)
	}

	if n > 32 {
		_panicf("invalid read: %d is greater than maximum read of 32 bits", n)
	}

	bitOffset := r.pos % 8
	nBitsToRead := bitOffset + n
	nBytesToRead := nBitsToRead / 8
	if nBitsToRead%8 != 0 {
		nBytesToRead += 1
	}

	var val uint64
	for i := 0; i < nBytesToRead; i++ {
		m := r.buf[(r.pos/8)+i]
		val += (uint64(m) << (uint64(i) * 8))
	}
	val >>= uint(bitOffset)
	val &= ((1 << uint64(n)) - 1)
	r.pos += n

	return uint(val)
}

// Take a peek at what's next in the buffer.
func (r *reader) _peek(depth, max int) {
	pos := r.pos

	indent := func(n int) string {
		s := ""
		for i := 0; i < n; i++ {
			s += "-- "
		}
		return s
	}

	defer func() {
		r.pos = pos
	}()

	if r.rem_bits() == 0 || depth > max {
		return
	}

	str := r.read_string()
	_debugf("%s [%d/%d] str = %s", indent(depth), pos, r.size, str)
	r._peek(depth+1, max)
	r.pos = pos

	buf := r.read_bytes(8)
	_debugf("%s [%d/%d] buf = %v (%s)", indent(depth), pos, r.size, buf, string(buf))
	r._peek(depth+1, max)
	r.pos = pos

	u32 := r.read_le_uint32()
	_debugf("%s [%d/%d] u32 = %d", indent(depth), pos, r.size, u32)
	r._peek(depth+1, max)
	r.pos = pos

	bln := r.read_boolean()
	_debugf("%s [%d/%d] bln = %v", indent(depth), pos, r.size, bln)
	r._peek(depth+1, max)
	r.pos = pos
}
