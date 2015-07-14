package manta

import (
	"encoding/binary"
	"math"
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

// Reads a signed 32-bit varint.
func (r *reader) readVarInt32() int32 {
	ux := r.readVarUint32()
	x := int32(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
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
func (r *reader) readVarInt64() int64 {
	ux := r.readVarUint64()
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
}

// Reads a boolean value.
func (r *reader) readBoolean() bool {
	if r.remBits() < 1 {
		_panicf("read overflow: no bits left")
	}

	b := r.buf[r.pos/8]&(1<<uint(r.pos%8)) != 0
	r.pos += 1
	return b
}

// Reads a bit varint
func (r *reader) readUBitVar() uint32 {
	ret := r.readBits(6)

	switch (ret & 0x30) {
	case 16:
		ret = (ret & 15) | (r.readBits(4) << 4);
		break;
	case 32:
		ret = (ret & 15) | (r.readBits(8) << 4);
		break;
	case 48:
		ret = (ret & 15) | (r.readBits(28) << 4);
		break;
	}

	return ret
}

// Reads the next byte (8 bits) in the buffer.
func (r *reader) readByte() byte {
	// Fast path if our position is byte-aligned.
	if r.pos%8 == 0 && r.pos+8 <= r.size {
		bpos := r.pos / 8
		r.pos += 8
		return r.buf[bpos]
	}

	// Slow path if our position isn't byte aligned.
	return byte(r.readBits(8))
}

// Reads the given number of bytes from the buffer.
func (r *reader) readBytes(n int) []byte {
	if r.remBits() < (n * 8) {
		_panicf("read overflow: %d bits requested, only %d remaining", n*8, r.remBits())
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

// Reads a float32 as the IEEE 754 binary representation of the next 4 bytes.
func (r *reader) readFloat32() float32 {
	return math.Float32frombits(r.readLeUint32())
}

// Reads a float32 with props
func (r *reader) readFloat32Bits(b int32, lP *float32, hP *float32) float32 {
	bits := int(b)
	lV, hV := float32(0.0), float32(0.0)
	if lP != nil {
		lV = *lP
	}
	if hP != nil {
		hV = *hP
	}

	dividend := r.readBits(bits)
	divisor := (1 << uint(bits)) - 1
	base := float32(dividend) / float32(divisor)
	diff := hV - lV
	return (base * diff) - lV
}

// Reads bits as bytes.
func (r *reader) readBitsAsBytes(n int) []byte {
	tmp := make([]byte, 0)
	for n >= 8 {
		tmp = append(tmp, r.readByte())
		n -= 8
	}
	if n > 0 {
		tmp = append(tmp, byte(r.readBits(n)))
	}
	return tmp
}

// Read bits of a given length as a uint, may or may not be byte-aligned.
func (r *reader) readBits(n int) uint32 {
	if r.remBits() < n {
		_panicf("read overflow: %d bits requested, only %d remaining", n, r.remBits())
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
		val += (uint64(m) << uint32(i*8))
	}
	val >>= uint32(bitOffset)
	val &= ((1 << uint32(n)) - 1)
	r.pos += n

	return uint32(val)
}
