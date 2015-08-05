package manta

import (
	"strconv"
)

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
	_debugf(
		"Bitcount: %v, Low: %v, High: %v, Flags: %v",
		saveReturnInt32(f.BitCount),
		saveReturnFloat32(f.LowValue, "nil"),
		saveReturnFloat32(f.HighValue, "nil"),
		strconv.FormatInt(int64(saveReturnInt32(f.Flags)), 2),
	)

	// Tread all the different flags
	//if (flags & 1) == 1 {
	/* Old way of reading a cell cord

	intval := r.readBits(1)
	fractval := r.readBits(1)

	if intval == 1 || fractval == 1 {
		ret := float32(0.0)
		sign := r.readBits(1)

		if intval == 1 {
			intval = r.readBits(int(bc)) + 1
		}

		if fractval == 1 {
			fractval = r.readBits(5)
		}

		ret = float32(intval) + float32(fractval) * float32(1.0 / 5)

		if sign == 1 {
			return -ret
		} else {
			return ret
		}
	}*/
	//}

	var bc int32
	bc = 10
	if f.BitCount != nil {
		bc = *f.BitCount
	}

	dividend := r.readBits(int(bc))
	divisor := (1 << uint32(bc)) - 1
	return float32(dividend) / float32(divisor)
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
}

func decodeQuantized(r *reader, f *dt_field) interface{} {
	// Lets do this for now
	return decodeFloat(r, f)
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
		_panicf("Figure out how this works")
	}

	return 0
}

func decodeQAngle(r *reader, f *dt_field) interface{} {
	if f.Flags != nil {
		// There is a flag check against 0x20 in the disasembly
		_debugf("Angle flags: %v", *f.Flags)
	}

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

func decodeComponent(r *reader, f *dt_field) interface{} {
	_debugf(
		"Bitcount: %v, Low: %v, High: %v, Flags: %v",
		saveReturnInt32(f.BitCount),
		saveReturnFloat32(f.LowValue, "nil"),
		saveReturnFloat32(f.HighValue, "nil"),
		strconv.FormatInt(int64(saveReturnInt32(f.Flags)), 2),
	)

	// might be encoded like a pointer (1 bit for set / unset, etc.)
	//if r.readBits(1) == 1 {
	//	return r.readBits(1)
	//}

	return 0
}

func decodeStrongHandle(r *reader, f *dt_field) interface{} {
	// wrong, just testing
	return r.readBits(1)
}

func decodeHSequence(r *reader, f *dt_field) interface{} {
	// wrong, just testing
	return r.readBits(1)
}
