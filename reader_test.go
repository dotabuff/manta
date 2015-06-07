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
	assert.Equal(magicSource2, r.read_bytes(8))
	assert.Equal(8, r.byte_pos())

	// Unknown int32
	assert.Equal(uint32(46293484), r.read_le_uint32())
	assert.Equal(12, r.byte_pos())

	// Unknown int32
	assert.Equal(uint32(46291827), r.read_le_uint32())
	assert.Equal(16, r.byte_pos())

	// First packet begins

	// Varint message type (1 = EDemoCommands_DEM_FileHeader)
	assert.Equal(uint32(1), r.read_var_uint32())
	assert.Equal(17, r.byte_pos())

	// Varint message tick (4294967295 = before the first tick)
	assert.Equal(uint32(4294967295), r.read_var_uint32())
	assert.Equal(22, r.byte_pos())

	// Varint message size (140)
	assert.Equal(uint32(140), r.read_var_uint32())
	assert.Equal(24, r.byte_pos())
}

func TestReaderVarints(t *testing.T) {
	assert := assert.New(t)

	r := newReader([]byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x8C, 0x01})

	// Ensure that read_var_uint32 works as expected
	assert.Equal(uint32(1), r.read_var_uint32())
	assert.Equal(uint32(4294967295), r.read_var_uint32())
	assert.Equal(uint32(140), r.read_var_uint32())
	r.pos = 0

	// Ensure that read_var_uint64 works as expected
	assert.Equal(uint64(1), r.read_var_uint64())
	assert.Equal(uint64(4294967295), r.read_var_uint64())
	assert.Equal(uint64(140), r.read_var_uint64())
}

func TestReaderStrings(t *testing.T) {
	assert := assert.New(t)

	r := newReader([]byte("PBDEMS2"))

	assert.Equal("PBDEMS2", r.read_string_n(7))
}

func BenchmarkReadVarUint32(b *testing.B) {
	r := newReader([]byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x8C, 0x01})
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.read_var_uint32()
		r.read_var_uint32()
		r.read_var_uint32()
		r.pos = 0
	}
	b.ReportAllocs()
}

func BenchmarkReadVarUint64(b *testing.B) {
	r := newReader([]byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F, 0x8C, 0x01})
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.read_var_uint64()
		r.read_var_uint64()
		r.read_var_uint64()
		r.pos = 0
	}
	b.ReportAllocs()
}

func BenchmarkReadBytesAligned(b *testing.B) {
	r := newReader(makeBuffer(1024))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.read_bytes(2)
		r.seek_bytes(-2)
	}
	b.ReportAllocs()
}

func BenchmarkReadBytesUnaligned(b *testing.B) {
	r := newReader(makeBuffer(1024))
	r.seek_bits(6)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.read_bytes(2)
		r.seek_bytes(-2)
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
