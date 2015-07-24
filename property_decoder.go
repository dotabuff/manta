package manta

func decodeHandle(r *reader, f *dt_field) interface{} {
	// So far these seem to occupy 32 bits but the value is made up only
	// out of what's present in the first 21 bits. In source 1, these only
	// occupied 21 bits of space.
	value := r.readBits(21) // a uint32
	r.seekBits(11)          // skip the rest of the 32 bits
	return value
}

func decodeByte(r *reader, f *dt_field) interface{} {
	return r.readBits(8)
}

func decodeShort(r *reader, f *dt_field) interface{} {
	return r.readBits(16)
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

	_debugf("s2 calc, %v, %v, %v, %v", f.BitCount, f.LowValue, f.HighValue, f.Flags)

	// We haven't yet determined how to read a float with a bitcount.
	if f.BitCount != nil {
		value = r.readFloat32Bits(*f.BitCount, f.LowValue, f.HighValue)
	} else {
		dividend := r.readBits(10)
		divisor := (1 << 10) - 1

		return float32(dividend) / float32(divisor)
	}

	return value
}

func decodeString(r *reader, f *dt_field) interface{} {
	return r.readString()
}

func decodeVector(r *reader, f *dt_field) interface{} {
	size := r.readVarUint32()

	if size > 0 {
		_panicf("Ive been called")
	}

	return 0
}

func decodeClass(r *reader, f *dt_field) interface{} {
	return r.readVarUint32()
	//return r.readBits(8)
}

func decodeQuantized(r *reader, f *dt_field) interface{} {
	_debugf("Quantized calc, %v, %v, %v, %v", f.BitCount, f.LowValue, f.HighValue, f.Flags)
	return r.readBits(11)
}

func decodeFVector(r *reader, f *dt_field) interface{} {
	var r2 [3]uint32

	r2[0] = r.readBits(10) // this should probably be readFloat
	r2[1] = r.readBits(10)

	if r.readBits(1) == 1 {
		r2[2] = r.readBits(10)
	} else {
		r2[2] = 0
	}

	return r2
}

func decodeNop(r *reader, f *dt_field) interface{} {
	return 0
}

func decodePointer(r *reader, f *dt_field) interface{} {
	// Seems to be encoded as a single bit, not sure what to make of it
	if !r.readBoolean() {
		r.readBits(30)
	}

	return 0
}

func decodeQAngle(r *reader, f *dt_field) interface{} {
	if f.Flags != nil {
		// There is a flag check against 0x20 in the disasembly
		_debugf("Angle flags: %v", *f.Flags)
	}

	r.dumpBits(1024)

	ret := [3]float32{0.0, 0.0, 0.0}

	rX := r.readBoolean()
	rY := r.readBoolean()
	rZ := r.readBoolean()

	if rX {
		ret[0] = r.readCoord()
	}

	if rY {
		ret[1] = r.readCoord()
	}

	if rZ {
		ret[2] = r.readCoord()
	}

	return ret
}
