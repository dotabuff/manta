package manta

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encVarUint encodes v as an unsigned LEB128 varint, matching reader.readVarUint*.
func encVarUint(v uint64) []byte {
	var b []byte
	for v >= 0x80 {
		b = append(b, byte(v)|0x80)
		v >>= 7
	}
	return append(b, byte(v))
}

// encVarInt encodes v as a zig-zag signed varint, matching reader.readVarInt*.
func encVarInt(v int64) []byte {
	return encVarUint(uint64((v << 1) ^ (v >> 63)))
}

// TestDecoderRepresentations locks the exact dynamic type AND value that the
// value-changing / representation-critical decoders produce through Entity.Get.
// These are the deliberate clarity-aligned changes we are owning (see
// IMPROVEMENTS.md P4.5 / Decision A): a regression here is a downstream-visible
// behavior change, which the replay golden tests would not necessarily catch.
// assert.Equal compares both the concrete type and value of the interface{}.
func TestDecoderRepresentations(t *testing.T) {
	assert := assert.New(t)

	dec := func(d fieldDecoder, buf []byte) interface{} {
		return d(newReader(buf)).iface()
	}

	// HSequence: unsigned varint minus one, returned as a signed int32 so the
	// "none" handle (wire 0) is -1 rather than a large unsigned value.
	assert.Equal(int32(4), dec(hSequenceDecoder, encVarUint(5)))
	assert.Equal(int32(-1), dec(hSequenceDecoder, encVarUint(0)))

	// HeroID_t: signed (zig-zag) varint int32.
	assert.Equal(int32(21), dec(signedDecoder, encVarInt(21)))
	assert.Equal(int32(-5), dec(signedDecoder, encVarInt(-5)))

	// int64: full 64-bit signed varint (not truncated to int32).
	assert.Equal(int64(5_000_000_000), dec(signed64Decoder, encVarInt(5_000_000_000)))

	// BloodType: fixed 8-bit unsigned, returned as uint64.
	assert.Equal(uint64(66), dec(bloodTypeDecoder, []byte{0x42}))

	// 32-bit-range unsigned handles: stored inline, returned as uint64.
	assert.Equal(uint64(5), dec(unsignedDecoder, encVarUint(5)))
	assert.Equal(uint64(0xFFFFFFFF), dec(unsignedDecoder, encVarUint(0xFFFFFFFF)))

	// Genuinely 64-bit unsigned values (e.g. steam IDs) must round-trip without
	// truncation through the boxed path, returned as uint64.
	assert.Equal(uint64(76561198140280423), dec(unsigned64Decoder, encVarUint(76561198140280423)))
	fixedBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(fixedBuf, 0x0123456789ABCDEF)
	assert.Equal(uint64(0x0123456789ABCDEF), dec(fixed64Decoder, fixedBuf))
}

// TestValueChangingDecoderWiring locks that the field types whose representation
// changed are mapped to the intended decoders, so a future edit to the decoder
// tables cannot silently revert the behavior.
func TestValueChangingDecoderWiring(t *testing.T) {
	assert := assert.New(t)

	wired := func(typeName string, buf []byte) interface{} {
		d, ok := fieldTypeDecoders[typeName]
		if !ok {
			t.Fatalf("no decoder registered for %s", typeName)
		}
		return d(newReader(buf)).iface()
	}

	assert.Equal(int32(-1), wired("HSequence", encVarUint(0)))
	assert.Equal(int32(21), wired("HeroID_t", encVarInt(21)))
	assert.Equal(int64(5_000_000_000), wired("int64", encVarInt(5_000_000_000)))
	assert.Equal(uint64(66), wired("BloodType", []byte{0x42}))
}
