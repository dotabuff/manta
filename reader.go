package manta

import (
	"encoding/binary"
	"fmt"
	"math"
)

// reader performs read operations against a buffer
type reader struct {
	buf      []byte
	size     uint32
	pos      uint32
	bitVal   uint64 // value of the remaining bits in the current byte
	bitCount uint32 // number of remaining bits in the current byte
}

// newReader creates a new reader object for the given buffer
func newReader(buf []byte) *reader {
	return &reader{buf, uint32(len(buf)), 0, 0, 0}
}

// reset reinitializes the reader to read from the given buffer, allowing a
// single reader to be reused across messages without allocating a new one.
func (r *reader) reset(buf []byte) {
	r.buf = buf
	r.size = uint32(len(buf))
	r.pos = 0
	r.bitVal = 0
	r.bitCount = 0
}

// remBits calculates the number of unread bits in the buffer, including the
// bits already buffered in the accumulator.
func (r *reader) remBits() uint32 {
	return r.remBytes()*8 + r.bitCount
}

func (r *reader) position() string {
	// Logical bit position consumed so far. The word-at-a-time refill can buffer
	// many bits (not just the current byte), so derive byte.bit from pos*8 minus
	// the still-buffered bitCount rather than assuming bitCount <= 8.
	bits := r.pos*8 - r.bitCount
	if rem := bits % 8; rem != 0 {
		return fmt.Sprintf("%d.%d", bits/8, rem)
	}
	return fmt.Sprintf("%d", bits/8)
}

// remBytes calculates the number of unread bytes in the buffer
func (r *reader) remBytes() uint32 {
	return r.size - r.pos
}

// nextByte reads the next byte from the buffer
func (r *reader) nextByte() byte {
	r.pos += 1
	if r.pos > r.size {
		_panicf("nextByte: insufficient buffer (%d of %d)", r.pos, r.size)
	}
	return r.buf[r.pos-1]
}

// readBitMasks[k] is a bitmask of the low k bits. It is used instead of an
// inline (1<<k)-1 so that k == 64 (a full word) does not overflow the shift.
var readBitMasks = func() [65]uint64 {
	var m [65]uint64
	for i := 0; i < 64; i++ {
		m[i] = (uint64(1) << uint(i)) - 1
	}
	m[64] = ^uint64(0)
	return m
}()

// readBits returns the uint32 value for the given number of sequential bits.
//
// The accumulator is refilled a full 64-bit word at a time while at least 8
// bytes remain, falling back to byte-at-a-time only for the final bytes of the
// buffer. Every read is <= 32 bits (quantized float decoders are asserted to
// stay within that), so a single word refill always yields enough bits. The
// refill lives out of line in fill so this fast path stays small enough to be
// inlined into its (many, hot) callers.
func (r *reader) readBits(n uint32) uint32 {
	if n > r.bitCount {
		r.fill(n)
	}
	x := r.bitVal
	r.bitVal = x >> n
	r.bitCount -= n
	return uint32(x) & (1<<n - 1)
}

// fill tops up the accumulator until at least n bits are buffered, panicking
// (via nextByte) if the buffer runs out first. The loaded word is masked to
// the whole bytes of headroom before being merged so no unclaimed
// (partial-byte) bits are left above bitCount.
func (r *reader) fill(n uint32) {
	for n > r.bitCount && r.pos+8 <= r.size {
		w := binary.LittleEndian.Uint64(r.buf[r.pos:])
		free := (64 - r.bitCount) >> 3 // whole bytes of accumulator headroom
		bits := free * 8
		r.bitVal |= (w & readBitMasks[bits]) << r.bitCount
		r.pos += free
		r.bitCount += bits
	}
	for n > r.bitCount {
		r.bitVal |= uint64(r.nextByte()) << r.bitCount
		r.bitCount += 8
	}
}

// peekBits returns the next n (<= 32) bits without consuming them, refilling the
// accumulator as needed. If fewer than n bits remain in the buffer the missing
// high bits are returned as zero; the buffer is never over-read past its end.
// This is required by the field-path op lookup, which inspects a fixed window
// that can extend past the final op near the end of the stream.
func (r *reader) peekBits(n uint32) uint32 {
	if n > r.bitCount {
		r.fillPeek(n)
	}
	return uint32(r.bitVal & (1<<n - 1))
}

// fillPeek is fill for peekBits: it stops at the end of the buffer instead of
// panicking, leaving the missing high bits zero.
func (r *reader) fillPeek(n uint32) {
	for n > r.bitCount && r.pos+8 <= r.size {
		w := binary.LittleEndian.Uint64(r.buf[r.pos:])
		free := (64 - r.bitCount) >> 3
		bits := free * 8
		r.bitVal |= (w & readBitMasks[bits]) << r.bitCount
		r.pos += free
		r.bitCount += bits
	}
	for n > r.bitCount && r.pos < r.size {
		r.bitVal |= uint64(r.nextByte()) << r.bitCount
		r.bitCount += 8
	}
}

// skipBits discards n bits already buffered in the accumulator. It panics if
// fewer than n bits are buffered, which happens only on a truncated or corrupt
// stream (a field-path op lookup resolving past the end of the buffer). The
// parser's top-level recover turns that into an error, rather than letting
// bitCount underflow and the field-path decode spin on garbage.
func (r *reader) skipBits(n uint32) {
	if n > r.bitCount {
		r.skipUnderflow(n)
	}
	r.bitVal >>= n
	r.bitCount -= n
}

// skipUnderflow reports a skipBits past the buffered bits. It is kept out of
// line so skipBits stays inlinable.
//
//go:noinline
func (r *reader) skipUnderflow(n uint32) {
	_panicf("skipBits: insufficient buffer (need %d, have %d bits)", n, r.bitCount)
}

// realign discards the whole bytes the word refill buffered in the accumulator
// by rewinding the read position, so byte-oriented reads can proceed directly
// from (and alias) the underlying buffer. Only valid when byte-aligned, i.e.
// bitCount is a multiple of 8.
func (r *reader) realign() {
	r.pos -= r.bitCount >> 3
	r.bitVal = 0
	r.bitCount = 0
}

// readByte reads a single byte
func (r *reader) readByte() byte {
	// Fast path if we're byte aligned
	if r.bitCount&7 == 0 {
		if r.bitCount != 0 {
			r.realign()
		}
		return r.nextByte()
	}

	return byte(r.readBits(8))
}

// readBytes reads the given number of bytes
func (r *reader) readBytes(n uint32) []byte {
	// Fast path if we're byte aligned
	if r.bitCount&7 == 0 {
		if r.bitCount != 0 {
			r.realign()
		}
		r.pos += n
		if r.pos > r.size {
			_panicf("readBytes: insufficient buffer (%d of %d)", r.pos, r.size)
		}
		return r.buf[r.pos-n : r.pos]
	}

	buf := make([]byte, n)
	r.readBytesInto(buf)
	return buf
}

// readBytesInto fills dst with the next len(dst) bytes of the stream. Unlike
// readBytes it never allocates: the caller provides the destination, and the
// unaligned path copies a 32-bit word at a time rather than byte by byte.
func (r *reader) readBytesInto(dst []byte) {
	n := uint32(len(dst))

	// Fast path if we're byte aligned
	if r.bitCount&7 == 0 {
		if r.bitCount != 0 {
			r.realign()
		}
		r.pos += n
		if r.pos > r.size {
			_panicf("readBytesInto: insufficient buffer (%d of %d)", r.pos, r.size)
		}
		copy(dst, r.buf[r.pos-n:r.pos])
		return
	}

	// Unaligned: rather than a readBits(32) per word, shift-merge straight from
	// the buffer. The accumulator holds the bits starting at logical bit
	// position p (not on a byte boundary), and they are still in buf, so output
	// byte i is the 8 bits at p+8i, which straddle buf[p/8+i] and the next byte.
	p := r.pos*8 - r.bitCount
	if uint64(n)*8 > uint64(r.size)*8-uint64(p) {
		_panicf("readBytesInto: insufficient buffer (%d bytes at bit %d of %d)", n, p, r.size*8)
	}
	shift := p & 7
	src := r.buf[p>>3 : r.size]
	i := 0
	// Eight output bytes need nine input bytes, so the word loop stops one
	// short of the end of src; the remainder takes the byte path.
	for ; i+9 <= len(src) && i+8 <= len(dst); i += 8 {
		w := binary.LittleEndian.Uint64(src[i:])
		binary.LittleEndian.PutUint64(dst[i:], w>>shift|uint64(src[i+8])<<(64-shift))
	}
	for ; i < len(dst); i++ {
		dst[i] = src[i]>>shift | src[i+1]<<(8-shift)
	}

	// Resume after the copied bytes. The position is still not byte aligned
	// (it moved by whole bytes), so reload the rest of its byte as the
	// accumulator, exactly as a byte-at-a-time reader would hold it.
	end := p>>3 + n
	r.bitVal = uint64(r.buf[end]) >> shift
	r.bitCount = 8 - shift
	r.pos = end + 1
}

// readLeUint32 reads an little-endian uint32
func (r *reader) readLeUint32() uint32 {
	// Unaligned: read straight from the bit accumulator. Bits are consumed
	// LSB-first from little-endian bytes, so this equals the LE decode of the
	// next four stream bytes without the readBytes slow-path allocation. The
	// aligned path keeps the zero-copy readBytes fast path.
	if r.bitCount&7 != 0 {
		return r.readBits(32)
	}
	return binary.LittleEndian.Uint32(r.readBytes(4))
}

// readLeUint64 reads a little-endian uint64
func (r *reader) readLeUint64() uint64 {
	// See readLeUint32: two accumulator words replace the unaligned readBytes
	// allocation (readBits is capped at 32 bits per call).
	if r.bitCount&7 != 0 {
		lo := uint64(r.readBits(32))
		return lo | uint64(r.readBits(32))<<32
	}
	return binary.LittleEndian.Uint64(r.readBytes(8))
}

// readVarUint64 reads an unsigned 32-bit varint
func (r *reader) readVarUint32() uint32 {
	// Unaligned (the common case inside entity and packet bitstreams), each
	// byte would go through readBits(8). Peek up to five bytes from the
	// accumulator at once instead and consume only the bytes the varint used:
	// the same bits, in the same order. The aligned path below reads straight
	// from the buffer and keeps the reader's exact pos.
	if r.bitCount&7 != 0 {
		if r.bitCount < 40 {
			r.fillPeek(40)
		}
		v := r.bitVal
		var x uint32
		for i := uint32(0); ; i++ {
			b := uint32(v>>(8*i)) & 0xFF
			x |= (b & 0x7F) << (7 * i)
			if b&0x80 == 0 || i == 4 {
				r.skipBits(8 * (i + 1))
				return x
			}
		}
	}

	var x, s uint32
	for {
		b := uint32(r.readByte())
		x |= (b & 0x7F) << s
		s += 7
		if ((b & 0x80) == 0) || (s == 35) {
			break
		}
	}

	return x
}

// readVarInt64 reads a signed 32-bit varint
func (r *reader) readVarInt32() int32 {
	ux := r.readVarUint32()
	x := int32(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
}

// readVarUint64 reads an unsigned 64-bit varint
func (r *reader) readVarUint64() uint64 {
	var x, s uint64
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

// readVarInt64 reads a signed 64-bit varint
func (r *reader) readVarInt64() int64 {
	ux := r.readVarUint64()
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x
}

// readBoolean reads and interprets single bit as true or false
func (r *reader) readBoolean() bool {
	if r.bitCount == 0 {
		r.fill(1)
	}
	b := r.bitVal&1 != 0
	r.bitVal >>= 1
	r.bitCount--
	return b
}

// readFloat reads an IEEE 754 float
func (r *reader) readFloat() float32 {
	return math.Float32frombits(r.readLeUint32())
}

// readUBitVar reads a variable length uint32 with encoding in last to bits of 6 bit group
func (r *reader) readUBitVar() uint32 {
	// A 6-bit group whose top two bits select 0, 4, 8 or 28 more bits (at most
	// 34 in all). Peek the longest form at once and consume only what was used.
	if r.bitCount < 34 {
		r.fillPeek(34)
	}
	v := r.bitVal
	ret := uint32(v) & 0x3F

	switch ret & 0x30 {
	case 16:
		r.skipBits(10)
		return (ret & 15) | (uint32(v>>6)&0xF)<<4
	case 32:
		r.skipBits(14)
		return (ret & 15) | (uint32(v>>6)&0xFF)<<4
	case 48:
		r.skipBits(34)
		return (ret & 15) | (uint32(v>>6)&0xFFFFFFF)<<4
	}

	r.skipBits(6)
	return ret
}

// readUBitVarFP reads a variable length uint32 encoded using fieldpath encoding
func (r *reader) readUBitVarFP() uint32 {
	// Up to four flag bits select a 2, 4, 10, 17 or 31-bit value (at most 35
	// bits in all). Peek the longest form at once and consume only what was
	// used, rather than testing each flag with its own readBoolean.
	if r.bitCount < 35 {
		r.fillPeek(35)
	}
	v := r.bitVal

	switch {
	case v&1 != 0:
		r.skipBits(3)
		return uint32(v>>1) & 0x3
	case v&2 != 0:
		r.skipBits(6)
		return uint32(v>>2) & 0xF
	case v&4 != 0:
		r.skipBits(13)
		return uint32(v>>3) & 0x3FF
	case v&8 != 0:
		r.skipBits(21)
		return uint32(v>>4) & 0x1FFFF
	}

	r.skipBits(35)
	return uint32(v>>4) & 0x7FFFFFFF
}

func (r *reader) readUBitVarFieldPath() int {
	return int(r.readUBitVarFP())
}

// readStringN reads a string of a given length
func (r *reader) readStringN(n uint32) string {
	return string(r.readBytes(n))
}

// readString reads a null terminated string
func (r *reader) readString() string {
	// Most strings here are short field/table keys; a small starting capacity
	// avoids the repeated growth of a cap-0 append chain.
	buf := make([]byte, 0, 32)
	for {
		b := r.readByte()
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}

	return string(buf)
}

// readCoord reads a coord as a float32
func (r *reader) readCoord() float32 {
	value := float32(0.0)

	intval := r.readBits(1)
	fractval := r.readBits(1)
	signbit := false

	if intval != 0 || fractval != 0 {
		signbit = r.readBoolean()

		if intval != 0 {
			intval = r.readBits(14) + 1
		}

		if fractval != 0 {
			fractval = r.readBits(5)
		}

		value = float32(intval) + float32(fractval)*(1.0/(1<<5))

		// Fixup the sign if negative.
		if signbit {
			value = -value
		}
	}

	return value
}

// readAngle reads a bit angle of the given size
func (r *reader) readAngle(n uint32) float32 {
	return float32(r.readBits(n)) * 360.0 / float32(int(1<<n))
}

// readNormal reads a normalized float vector
func (r *reader) readNormal() float32 {
	isNeg := r.readBoolean()
	len := r.readBits(11)
	ret := float32(len) * float32(1.0/(float32(1<<11)-1.0))

	if isNeg {
		return -ret
	} else {
		return ret
	}
}

// read3BitNormal reads a normalized float vector
func (r *reader) read3BitNormal() []float32 {
	ret := []float32{0.0, 0.0, 0.0}

	hasX := r.readBoolean()
	haxY := r.readBoolean()

	if hasX {
		ret[0] = r.readNormal()
	}

	if haxY {
		ret[1] = r.readNormal()
	}

	negZ := r.readBoolean()
	prodsum := ret[0]*ret[0] + ret[1]*ret[1]

	if prodsum < 1.0 {
		ret[2] = float32(math.Sqrt(float64(1.0 - prodsum)))
	} else {
		ret[2] = 0.0
	}

	if negZ {
		ret[2] = -ret[2]
	}

	return ret
}

// readBitsAsBytes reads the given number of bits in groups of bytes
func (r *reader) readBitsAsBytes(n uint32) []byte {
	// Allocate the exact result size up front (the caller retains the slice, so
	// a fresh allocation is required) and fill a 32-bit word at a time instead
	// of growing a cap-0 slice byte by byte.
	tmp := make([]byte, (n+7)/8)
	i := 0
	for n >= 32 {
		binary.LittleEndian.PutUint32(tmp[i:], r.readBits(32))
		i += 4
		n -= 32
	}
	for n >= 8 {
		tmp[i] = r.readByte()
		i++
		n -= 8
	}
	if n > 0 {
		tmp[i] = byte(r.readBits(n))
	}
	return tmp
}
