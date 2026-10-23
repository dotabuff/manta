package manta

import (
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that we properly parse the beginning of a replay.
func TestReaderReplayBeginning(t *testing.T) {
	assert := assert.New(t)

	// Buffer represents the first bytes in a sample replay
	buf := []byte{
		// Null terminated PBDEMS2 string
		0x50, 0x42, 0x44, 0x45, 0x4D, 0x53, 0x32, 0x00,
		// Unknown int32
		0xEC, 0x61, 0xC2, 0x02,
		// Unknown int32
		0x73, 0x5B, 0xC2, 0x02,
		// First packet begins
		//   Varint message type (1 = EDemoCommands_DEM_FileHeader)
		0x01,
		//   Varint message tick (4294967295 = before the first tick)
		0xFF, 0xFF, 0xFF, 0xFF, 0x0F,
		//   Varint buffer size (140)
		0x8C, 0x01,
	}

	r := newReader(buf)

	// Null terminated PBDEMS2 string
	assert.Equal(magicSource2, r.readBytes(8))
	assert.Equal(uint32(8), r.pos)

	// Unknown int32
	assert.Equal(uint32(46293484), r.readLeUint32())
	assert.Equal(uint32(12), r.pos)

	// Unknown int32
	assert.Equal(uint32(46291827), r.readLeUint32())
	assert.Equal(uint32(16), r.pos)

	// First packet begins

	// Varint message type (1 = EDemoCommands_DEM_FileHeader)
	assert.Equal(uint32(1), r.readVarUint32())
	assert.Equal(uint32(17), r.pos)

	// Varint message tick (4294967295 = before the first tick)
	assert.Equal(uint32(4294967295), r.readVarUint32())
	assert.Equal(uint32(22), r.pos)

	// Varint message size (140)
	assert.Equal(uint32(140), r.readVarUint32())
	assert.Equal(uint32(24), r.pos)
}

func TestReaderVarints(t *testing.T) {
	assert := assert.New(t)

	r := newReader([]byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x8C, 0x01})

	// Ensure that readVarUint32 works as expected
	assert.Equal(uint32(1), r.readVarUint32())
	assert.Equal(uint32(4294967295), r.readVarUint32())
	assert.Equal(uint32(140), r.readVarUint32())

	r.pos = 0
	assert.Equal(uint32(1), r.readVarUint32())
	assert.Equal(int32(-2147483648), r.readVarInt32())

	r.pos = 0

	// Ensure that readVarUint64 works as expected
	assert.Equal(uint64(1), r.readVarUint64())
	assert.Equal(uint64(4294967295), r.readVarUint64())
	assert.Equal(uint64(140), r.readVarUint64())
}

func TestReaderStrings(t *testing.T) {
	assert := assert.New(t)

	r := newReader([]byte{'P', 'B', 'D', 'E', 'M', 'S', '2', 0x0, 'E', 'X', 'T', 'R', 'A', 0x0})

	assert.Equal("PBDEMS2", r.readStringN(7))
	r.pos = 0
	assert.Equal("PBDEMS2", r.readString())
	assert.Equal("EXTRA", r.readString())
}

func TestReaderUnaligned(t *testing.T) {
	assert := assert.New(t)

	r := newReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})

	assert.Equal(uint32(0x7f), r.readBits(7))
	assert.Equal(uint32(0xff), r.readBits(8))
	assert.Equal(uint32(0xffff), r.readBits(16))
	assert.Equal(uint32(0xffffffff), r.readBits(32))
	assert.Equal(uint32(0x01), r.readBits(1))
}

// TestReaderReadBytesZeroCopy guards the invariant that a byte-aligned readBytes
// returns a slice aliasing the underlying buffer rather than a copy. Callers such
// as CDemoPacket pendingMessage parsing and sendtables rely on this aliasing, and
// the word-at-a-time reader must preserve it (via realign) even when it has
// buffered bytes ahead in the accumulator.
func TestReaderReadBytesZeroCopy(t *testing.T) {
	buf := make([]byte, 32)
	for i := range buf {
		buf[i] = byte(i + 1)
	}

	// Aligned read at the start of the buffer (bitCount == 0).
	r := newReader(buf)
	got := r.readBytes(4)
	idx := r.pos - 4
	buf[idx] ^= 0xFF
	if got[0] != buf[idx] {
		t.Fatalf("aligned readBytes returned a copy, not a zero-copy alias")
	}

	// A bit read first buffers a whole word ahead; the following byte-aligned
	// readBytes must realign and still alias the buffer (bitCount a non-zero
	// multiple of 8).
	for i := range buf {
		buf[i] = byte(i + 1)
	}
	r = newReader(buf)
	_ = r.readBits(8)
	if r.bitCount == 0 || r.bitCount%8 != 0 {
		t.Fatalf("expected non-zero byte alignment after readBits(8), bitCount=%d", r.bitCount)
	}
	got = r.readBytes(4)
	idx = r.pos - 4
	buf[idx] ^= 0xFF
	if got[0] != buf[idx] {
		t.Fatalf("realigned readBytes returned a copy, not a zero-copy alias")
	}
}

func BenchmarkReadVarUint32(b *testing.B) {
	r := newReader([]byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x8C, 0x01})
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.readVarUint32()
		r.readVarUint32()
		r.readVarUint32()
		r.pos = 0
	}
	b.ReportAllocs()
}

func BenchmarkReadVarUint64(b *testing.B) {
	r := newReader([]byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x8C, 0x01})
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.readVarUint64()
		r.readVarUint64()
		r.readVarUint64()
		r.pos = 0
	}
	b.ReportAllocs()
}

func BenchmarkReadBytesAligned(b *testing.B) {
	r := newReader(makeBuffer(1024))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.readBytes(2)
		rewindBytes(r, 2)
	}
	b.ReportAllocs()
}

func BenchmarkReadBytesUnaligned(b *testing.B) {
	r := newReader(makeBuffer(1024))
	r.readBits(6)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.readBytes(2)
		rewindBytes(r, 2)
	}
	b.ReportAllocs()
}

func makeBuffer(n int) []byte {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = byte(rand.Intn(255))
	}
	return buf
}

func rewindBytes(r *reader, n uint32) {
	r.pos -= n
	r.bitCount = 0
	r.bitVal = 0
}

// Reference implementations of the variable-length integer reads, as they were
// written before decoding straight from the accumulator: one readByte per
// varint byte and one readBoolean/readBits per flag group.
func refReadVarUint32(r *reader) uint32 {
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

func refReadUBitVar(r *reader) uint32 {
	ret := r.readBits(6)
	switch ret & 0x30 {
	case 16:
		ret = (ret & 15) | (r.readBits(4) << 4)
	case 32:
		ret = (ret & 15) | (r.readBits(8) << 4)
	case 48:
		ret = (ret & 15) | (r.readBits(28) << 4)
	}
	return ret
}

func refReadUBitVarFP(r *reader) uint32 {
	if r.readBoolean() {
		return r.readBits(2)
	}
	if r.readBoolean() {
		return r.readBits(4)
	}
	if r.readBoolean() {
		return r.readBits(10)
	}
	if r.readBoolean() {
		return r.readBits(17)
	}
	return r.readBits(31)
}

// TestReaderVarLengthMatchesReference checks readVarUint32, readUBitVar and
// readUBitVarFP against the reference implementations on random buffers at
// random bit offsets, including reads that run off the end of the buffer
// (both must panic). Results, the logical bits consumed and a follow-up read
// must agree; byte-aligned varints must also leave pos exact.
func TestReaderVarLengthMatchesReference(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	type pair struct {
		name     string
		fast     func(*reader) uint32
		ref      func(*reader) uint32
		exactPos bool // compare pos/bitCount exactly when byte aligned
	}
	pairs := []pair{
		{"readVarUint32", (*reader).readVarUint32, refReadVarUint32, true},
		{"readUBitVar", (*reader).readUBitVar, refReadUBitVar, false},
		{"readUBitVarFP", (*reader).readUBitVarFP, refReadUBitVarFP, false},
	}

	call := func(f func(*reader) uint32, r *reader) (v uint32, panicked bool) {
		defer func() { panicked = recover() != nil }()
		return f(r), false
	}

	for _, pr := range pairs {
		for trial := 0; trial < 20000; trial++ {
			buf := make([]byte, rng.Intn(24))
			for i := range buf {
				// Bias towards set high bits so long varints are common.
				buf[i] = byte(rng.Intn(256)) | byte(rng.Intn(2)<<7)
			}
			skip := uint32(0)
			if len(buf) > 0 {
				skip = uint32(rng.Intn(len(buf)*8 + 1))
			}
			ra, rb := readerAtBit(buf, skip), readerAtBit(buf, skip)

			va, pa := call(pr.fast, ra)
			vb, pb := call(pr.ref, rb)
			if pa != pb {
				t.Fatalf("%s buf=%x skip=%d: panicked=%v, reference panicked=%v", pr.name, buf, skip, pa, pb)
			}
			if pa {
				continue
			}
			if va != vb {
				t.Fatalf("%s buf=%x skip=%d: got %d, want %d", pr.name, buf, skip, va, vb)
			}
			if a, b := ra.pos*8-ra.bitCount, rb.pos*8-rb.bitCount; a != b {
				t.Fatalf("%s buf=%x skip=%d: consumed to bit %d, want %d", pr.name, buf, skip, a, b)
			}
			if pr.exactPos && skip%8 == 0 && (ra.pos != rb.pos || ra.bitCount != rb.bitCount) {
				t.Fatalf("%s buf=%x skip=%d: aligned pos/bitCount %d/%d, want %d/%d", pr.name, buf, skip, ra.pos, ra.bitCount, rb.pos, rb.bitCount)
			}
			if rem := ra.remBits(); rem > 0 {
				n := uint32(min(rem, 32))
				if a, b := ra.readBits(n), rb.readBits(n); a != b {
					t.Fatalf("%s buf=%x skip=%d: follow-up read %x, want %x", pr.name, buf, skip, a, b)
				}
			}
		}
	}
}

// refReadBytesInto is readBytesInto as written before the unaligned
// shift-merge copy: one readBits(32) per word, then readBits(8) per byte.
func refReadBytesInto(r *reader, dst []byte) {
	if r.bitCount&7 == 0 {
		r.readBytesInto(dst)
		return
	}
	i := 0
	for len(dst)-i >= 4 {
		binary.LittleEndian.PutUint32(dst[i:], r.readBits(32))
		i += 4
	}
	for ; i < len(dst); i++ {
		dst[i] = byte(r.readBits(8))
	}
}

// TestReaderReadBytesIntoMatchesReference checks readBytesInto against the
// reference copy at random bit offsets and lengths, including copies that run
// off the end of the buffer (both must panic): same bytes, same bits consumed,
// and the same follow-up reads.
func TestReaderReadBytesIntoMatchesReference(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	call := func(f func(*reader, []byte), r *reader, dst []byte) (panicked bool) {
		defer func() { panicked = recover() != nil }()
		f(r, dst)
		return false
	}
	for trial := 0; trial < 20000; trial++ {
		buf := make([]byte, rng.Intn(48))
		rng.Read(buf)
		skip := uint32(0)
		if len(buf) > 0 {
			skip = uint32(rng.Intn(len(buf)*8 + 1))
		}
		n := rng.Intn(len(buf) + 2)
		ra, rb := readerAtBit(buf, skip), readerAtBit(buf, skip)
		da, db := make([]byte, n), make([]byte, n)

		pa := call((*reader).readBytesInto, ra, da)
		pb := call(refReadBytesInto, rb, db)
		if pa != pb {
			t.Fatalf("buf=%x skip=%d n=%d: panicked=%v, reference panicked=%v", buf, skip, n, pa, pb)
		}
		if pa {
			continue
		}
		if string(da) != string(db) {
			t.Fatalf("buf=%x skip=%d n=%d: got %x, want %x", buf, skip, n, da, db)
		}
		if a, b := ra.pos*8-ra.bitCount, rb.pos*8-rb.bitCount; a != b {
			t.Fatalf("buf=%x skip=%d n=%d: consumed to bit %d, want %d", buf, skip, n, a, b)
		}
		for rem := ra.remBits(); rem > 0; rem = ra.remBits() {
			k := uint32(min(rem, uint64(1+rng.Intn(32))))
			if a, b := ra.readBits(k), rb.readBits(k); a != b {
				t.Fatalf("buf=%x skip=%d n=%d: follow-up read %x, want %x", buf, skip, n, a, b)
			}
		}
	}
}

// refReadBitsAsBytes is readBitsAsBytes as written before it copied whole
// bytes with readBytesInto.
func refReadBitsAsBytes(r *reader, n uint32) []byte {
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

// TestReaderReadBitsAsBytesMatchesReference checks readBitsAsBytes against the
// reference at random bit offsets and bit lengths, including overruns.
func TestReaderReadBitsAsBytesMatchesReference(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	call := func(f func(*reader, uint32) []byte, r *reader, n uint32) (b []byte, panicked bool) {
		defer func() { panicked = recover() != nil }()
		return f(r, n), false
	}
	for trial := 0; trial < 20000; trial++ {
		buf := make([]byte, rng.Intn(48))
		rng.Read(buf)
		skip := uint32(0)
		if len(buf) > 0 {
			skip = uint32(rng.Intn(len(buf)*8 + 1))
		}
		n := uint32(rng.Intn(len(buf)*8 + 16))
		ra, rb := readerAtBit(buf, skip), readerAtBit(buf, skip)
		a, pa := call((*reader).readBitsAsBytes, ra, n)
		b, pb := call(refReadBitsAsBytes, rb, n)
		if pa != pb {
			t.Fatalf("buf=%x skip=%d n=%d: panicked=%v, reference panicked=%v", buf, skip, n, pa, pb)
		}
		if pa {
			continue
		}
		if string(a) != string(b) {
			t.Fatalf("buf=%x skip=%d n=%d: got %x, want %x", buf, skip, n, a, b)
		}
		if x, y := ra.pos*8-ra.bitCount, rb.pos*8-rb.bitCount; x != y {
			t.Fatalf("buf=%x skip=%d n=%d: consumed to bit %d, want %d", buf, skip, n, x, y)
		}
		if rem := ra.remBits(); rem > 0 {
			k := uint32(min(rem, 32))
			if x, y := ra.readBits(k), rb.readBits(k); x != y {
				t.Fatalf("buf=%x skip=%d n=%d: follow-up read %x, want %x", buf, skip, n, x, y)
			}
		}
	}
}

// TestReaderLargeBufferOffsets reads beyond the 512 MiB mark, where a byte
// position times 8 no longer fits in uint32. Bit positions and remaining bit
// counts must be computed in 64 bits, or the unaligned copy reads from near the
// start of the buffer and remBits wraps.
func TestReaderLargeBufferOffsets(t *testing.T) {
	if testing.Short() {
		t.Skip("allocates a 512 MiB buffer")
	}

	const off = 1<<29 + 64
	buf := make([]byte, off+64)
	for i := off; i < len(buf); i++ {
		buf[i] = byte(i*37 + 11)
	}

	if got, want := newReader(buf).remBits(), uint64(len(buf))*8; got != want {
		t.Fatalf("remBits at start = %d, want %d", got, want)
	}

	// Jump straight to the region of interest, then go unaligned.
	r := newReader(buf)
	r.pos = off
	if got := r.readBits(3); got != uint32(buf[off]&7) {
		t.Fatalf("readBits(3) = %d, want %d", got, buf[off]&7)
	}

	dst := make([]byte, 20)
	r.readBytesInto(dst)
	for i := range dst {
		if want := buf[off+i]>>3 | buf[off+i+1]<<5; dst[i] != want {
			t.Fatalf("byte %d = %02x, want %02x", i, dst[i], want)
		}
	}

	if got, want := r.bitPos(), uint64(off)*8+3+20*8; got != want {
		t.Fatalf("bitPos = %d, want %d", got, want)
	}
	if got, want := r.remBits(), uint64(len(buf))*8-(uint64(off)*8+3+20*8); got != want {
		t.Fatalf("remBits = %d, want %d", got, want)
	}
}
