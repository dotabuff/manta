package manta

import (
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
