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

// TestFixed8Encoder locks the "fixed8" var encoder: the server writes these
// fields as 8 raw bits whatever their type. Read as varints, a 255 consumed
// several bytes and desynced the rest of the entity: the first CParticleSystem
// baseline of match 9032897977 (m_iServerControlPointAssignments = 255 x4)
// failed with "nextByte: insufficient buffer (380 of 379)".
func TestFixed8Encoder(t *testing.T) {
	assert := assert.New(t)

	newFixed8 := func(varName, varType string, model int) *field {
		f := &field{varName: varName, varType: varType, encoder: "fixed8", fieldType: newFieldType(varType)}
		f.setModel(model)
		return f
	}
	bitsRead := func(r *reader) uint32 { return r.pos*8 - r.bitCount }

	// uint8[4] of 255 followed by the varint invalid handle 0xFFFFFF, as in
	// that baseline: the array takes exactly 4 bytes.
	arr := newFixed8("m_iServerControlPointAssignments", "uint8[4]", fieldModelFixedArray)
	r := newReader(append([]byte{0xff, 0xff, 0xff, 0xff}, encVarUint(0xffffff)...))
	for i := 0; i < 4; i++ {
		assert.Equal(uint64(255), arr.decoder(r).iface())
	}
	assert.Equal(uint64(0xffffff), unsignedDecoder(r).iface())
	assert.Equal(uint32(64), bitsRead(r))

	// int8 is signed, other types unsigned; each reads exactly 8 bits.
	for _, c := range []struct {
		varType string
		in      byte
		want    interface{}
	}{
		{"int8", 0xff, int32(-1)},
		{"int8", 0x29, int32(41)},
		{"uint8", 0x83, uint64(131)},
		{"AnimationAlgorithm_t", 0xff, uint32(255)},
	} {
		r := newReader([]byte{c.in, 0x01})
		assert.Equal(c.want, newFixed8("m_field", c.varType, fieldModelSimple).decoder(r).iface(), c.varType)
		assert.Equal(uint32(8), bitsRead(r), c.varType)
	}

	// CNetworkUtlVectorBase< uint8 >: varint length, 8-bit elements.
	vec := newFixed8("m_vecPlayerDraftPickOrder", "CNetworkUtlVectorBase< uint8 >", fieldModelVariableArray)
	r = newReader([]byte{0x02, 0xff, 0x80})
	assert.Equal(uint64(2), vec.baseDecoder(r).iface())
	assert.Equal(uint64(255), vec.childDecoder(r).iface())
	assert.Equal(uint64(128), vec.childDecoder(r).iface())
	assert.Equal(uint32(24), bitsRead(r))
}
