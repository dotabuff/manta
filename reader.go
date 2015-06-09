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
func (r *reader) bytePos() int {
	return r.pos / 8
}

// Calculates how many bits are remaining.
func (r *reader) remBits() int {
	return r.size - r.pos
}

// Calculates how many bytes are remaining.
func (r *reader) remBytes() int {
	return r.remBits() / 8
}

// Seeks a given number of bits (may be negative).
func (r *reader) seekBits(n int) {
	if r.pos+n >= r.size || r.pos+n < 0 {
		_panicf("seek overflow: %d bits requested, only %d remaining", n, r.size-r.pos)
	}
	r.pos += n
}

// Seeks a given number of bytes (may be negative).
func (r *reader) seekBytes(n int) {
	r.seekBits(n * 8)
}

// Reads a little-endian uint16.
func (r *reader) readLeUint16() uint16 {
	return littleEndian.Uint16(r.readBytes(2))
}

// Reads a little-endian uint32.
func (r *reader) readLeUint32() uint32 {
	return littleEndian.Uint32(r.readBytes(4))
}

// Reads a little-endian uint64.
func (r *reader) readLeUint64() uint64 {
	return littleEndian.Uint64(r.readBytes(8))
}

// Reads a big-endian uint16.
func (r *reader) readBeUint16() uint16 {
	return bigEndian.Uint16(r.readBytes(2))
}

// Reads a big-endian uint32.
func (r *reader) readBeUint32() uint32 {
	return bigEndian.Uint32(r.readBytes(4))
}

// Reads a big-endian uint64.
func (r *reader) readBeUint64() uint64 {
	return bigEndian.Uint64(r.readBytes(8))
}

// Reads an unsigned 32-bit varint.
func (r *reader) readVarUint32() uint32 {
	var x uint
	var s uint
	for {
		b := uint(r.readByte())
		x |= (b & 0x7F) << s
		s += 7
		if ((b & 0x80) == 0) || (s == 35) {
			break
		}
	}

	return uint32(x)
}

// Reads an unsigned 64-bit varint.
func (r *reader) readVarUint64() uint64 {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		b := r.readByte()
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
func (r *reader) readVarSint64() int64 {
	ux := r.readVarUint64()
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
}

// Reads a boolean value.
// TODO XXX: untested.
func (r *reader) readBoolean() bool {
	if r.remBits() < 1 {
		_panicf("read overflow: no bits left")
	}

	b := r.buf[r.pos/8]&(1<<uint(r.pos%8)) != 0
	r.pos += 1
	return b
}

// Reads the next byte (8 bits) in the buffer.
func (r *reader) readByte() byte {
	return r.readBytes(1)[0]
}

// Reads the given number of bytes from the buffer.
func (r *reader) readBytes(n int) []byte {
	if r.remBits() < (n * 8) {
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
		buf[i] = byte(r.readBits(8))
	}
	return buf
}

// Reads a string of a given length.
func (r *reader) readStringN(n int) string {
	return string(r.readBytes(n))
}

// Read a null terminated string.
func (r *reader) readString() string {
	buf := make([]byte, 0)
	for {
		b := r.readByte()
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}

	return string(buf)
}

// Reads bits as bytes.
func (r *reader) readBitsAsBytes(n int) []byte {
	buf := make([]byte, (n+7)/8)
	i := 0
	for n > 7 {
		n -= 8
		buf[i] = byte(r.readBits(8))
		i++
	}
	if n != 0 {
		buf[i] = byte(r.readBits(8))
	}
	return buf
}

// Read bits of a given length as a uint, may or may not be byte-aligned.
func (r *reader) readBits(n int) uint {
	if r.remBits() < n {
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
func (r *reader) peekAhead(depth, max int) {
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

	if r.remBits() == 0 || depth > max {
		return
	}

	buf := r.readBytes(8)
	_debugf("%s [%d/%d] buf = %v (%s)", indent(depth), pos, r.size, buf, string(buf))
	r.peekAhead(depth+1, max)
	r.pos = pos

	u32 := r.readLeUint32()
	_debugf("%s [%d/%d] u32 = %d", indent(depth), pos, r.size, u32)
	r.peekAhead(depth+1, max)
	r.pos = pos

	vui := r.readLeUint32()
	_debugf("%s [%d/%d] var = %d", indent(depth), pos, r.size, vui)
	r.peekAhead(depth+1, max)
	r.pos = pos

	bln := r.readBoolean()
	_debugf("%s [%d/%d] bln = %v", indent(depth), pos, r.size, bln)
	r.peekAhead(depth+1, max)
	r.pos = pos
}
