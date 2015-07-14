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
	assert.Equal(8, r.bytePos())

	// Unknown int32
	assert.Equal(uint32(46293484), r.readLeUint32())
	assert.Equal(12, r.bytePos())

	// Unknown int32
	assert.Equal(uint32(46291827), r.readLeUint32())
	assert.Equal(16, r.bytePos())

	// First packet begins

	// Varint message type (1 = EDemoCommands_DEM_FileHeader)
	assert.Equal(uint32(1), r.readVarUint32())
	assert.Equal(17, r.bytePos())

	// Varint message tick (4294967295 = before the first tick)
	assert.Equal(uint32(4294967295), r.readVarUint32())
	assert.Equal(22, r.bytePos())

	// Varint message size (140)
	assert.Equal(uint32(140), r.readVarUint32())
	assert.Equal(24, r.bytePos())
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

func TestReaderBoolean(t *testing.T) {
	assert := assert.New(t)

	// Start with any random mixed buffer
	buf := _read_fixture("send_tables/1560315800.pbmsg")
	r := newReader(buf)

	// Iterate through each bit
	for r.pos < r.size {
		// Read it as a 1-bit uint, either 0 (false) or 1 (true)
		expect := false
		if n := r.readBits(1); n == 1 {
			expect = true
		}
		r.pos--

		// Read it as a bool
		got := r.readBoolean()

		// Check that uint 1 == true, uint 0 = false
		assert.Equal(expect, got)
	}
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
		r.seekBytes(-2)
	}
	b.ReportAllocs()
}

func BenchmarkReadBytesUnaligned(b *testing.B) {
	r := newReader(makeBuffer(1024))
	r.seekBits(6)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.readBytes(2)
		r.seekBytes(-2)
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
