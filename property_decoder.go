package manta

func decodeHandle(r *reader, f *dt_field) interface{} {
	// So far these seem to occupy 32 bits but the value is made up only
	// out of what's present in the first 21 bits. In source 1, these only
	// occupied 21 bits of space.
	value := r.readBits(21) // a uint32
	r.seekBits(11)          // skip the rest of the 32 bits
	return value
}

func decodeUnsigned(r *reader, f *dt_field) interface{} {
	return r.readVarUint64()
}

func decodeSigned(r *reader, f *dt_field) interface{} {
	return r.readVarInt32()
}

func decodeBoolean(r *reader, f *dt_field) interface{} {
	return r.readBoolean()
}

func decodeFloat(r *reader, f *dt_field) interface{} {
	// This will be tricky as floats can be encoded in oh-so-many ways.
	var value float32

	// We haven't yet determined how to read a float with a bitcount.
	if f.BitCount != nil {
		// This uses the source 1 calculation using bits, lowValue, highValue.
		value = r.readFloat32Bits(*f.BitCount, f.LowValue, f.HighValue)
	} else {
		// This just reads a fixed length IEEE 754 float32. It might be 100%
		// wrong, and it will at least be wrong in cases where we have
		// flags, lowVal, highVal or bitcount.
		value = r.readFloat32()
	}

	return value
}

func decodeString(r *reader, f *dt_field) interface{} {
	return r.readString()
}
