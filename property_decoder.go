package manta

import (
	"strconv"
)

func decodeHandle(r *Reader, f *dt_field) interface{} {
	// So far these seem to occupy 32 bits but the value is made up only
	// out of what's present in the first 21 bits. In source 1, these only
	// occupied 21 bits of space.
	value := r.readBits(21) // a uint32
	r.seekBits(11)          // skip the rest of the 32 bits
	return value
}

func decodeByte(r *Reader, f *dt_field) interface{} {
	return r.readBits(8)
}

func decodeShort(r *Reader, f *dt_field) interface{} {
	return r.readBits(16)
}

func decodeUnsigned(r *Reader, f *dt_field) interface{} {
	return r.readVarUint64()
}

func decodeSigned(r *Reader, f *dt_field) interface{} {
	return r.readVarInt32()
}

func decodeSigned64(r *Reader, f *dt_field) interface{} {
	return r.readVarInt64()
}

func decodeBoolean(r *Reader, f *dt_field) interface{} {
	return r.readBoolean()
}

func decodeFloat(r *Reader, f *dt_field) interface{} {
	_debugf(
		"Bitcount: %v, Low: %v, High: %v, Flags: %v, Encoder: %v",
		saveReturnInt32(f.BitCount),
		saveReturnFloat32(f.LowValue, "nil"),
		saveReturnFloat32(f.HighValue, "nil"),
		strconv.FormatInt(int64(saveReturnInt32(f.Flags)), 2),
		f.Encoder,
	)

	// Parse specific encoders
	switch f.Encoder {
	case "coord":
		return r.readCoord()
	}

	var BitCount int
	var Low float32
	var High float32

	if f.BitCount != nil {
		BitCount = int(*f.BitCount)
	} else {
		// Maybe treated as no scale or something?
		return r.readVarUint32()
	}

	if f.LowValue != nil {
		Low = *f.LowValue
	} else {
		Low = 0.0
	}

	if f.HighValue != nil {
		High = *f.HighValue
	} else {
		High = 1.0
	}

	if f.Flags != nil {
		// Skip this case
		if *f.Flags&0x4 != 0 && f.LowValue == nil {
			// This doesn't fell right
			return r.readBits(2)
		}

		// Read raw float
		if *f.Flags&0x100 != 0 {
			return r.readBits(BitCount)
		}

		// read low value if empty
		if *f.Flags&0x10 != 0 && r.readBoolean() {
			return f.LowValue
		}

		// read high value if empty
		if *f.Flags&0x20 != 0 && r.readBoolean() {
			return f.HighValue
		}
	}

	dividend := r.readBits(BitCount)
	divisor := (1 << uint32(BitCount)) - 1
	return Low + (float32(dividend)/float32(divisor))*(High-Low)
}

func decodeString(r *Reader, f *dt_field) interface{} {
	return r.readString()
}

func decodeVector(r *Reader, f *dt_field) interface{} {
	size := r.readVarUint32()

	if size > 0 {
		_panicf("Ive been called, %v", size)
	}

	return 0
}

func decodeClass(r *Reader, f *dt_field) interface{} {
	return r.readVarUint32()
}

func decodeQuantized(r *Reader, f *dt_field) interface{} {
	// Lets do this for now
	return decodeFloat(r, f)
}

func decodeFVector(r *Reader, f *dt_field) interface{} {
	// Parse specific encoders
	switch f.Encoder {
	case "normal":
		return r.read3BitNormal()
	}

	return []float32{decodeFloat(r, f).(float32), decodeFloat(r, f).(float32), decodeFloat(r, f).(float32)}
}

func decodeNop(r *Reader, f *dt_field) interface{} {
	return 0
}

func decodePointer(r *Reader, f *dt_field) interface{} {
	// Seems to be encoded as a single bit, not sure what to make of it
	if !r.readBoolean() {
		_panicf("Figure out how this works")
	}

	return 0
}

func decodeQAngle(r *Reader, f *dt_field) interface{} {
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

func decodeComponent(r *Reader, f *dt_field) interface{} {
	_debugf(
		"Bitcount: %v, Low: %v, High: %v, Flags: %v",
		saveReturnInt32(f.BitCount),
		saveReturnFloat32(f.LowValue, "nil"),
		saveReturnFloat32(f.HighValue, "nil"),
		strconv.FormatInt(int64(saveReturnInt32(f.Flags)), 2),
	)

	return r.readBits(1)
}

func decodeHSequence(r *Reader, f *dt_field) interface{} {
	// wrong, just testing
	return r.readBits(1)
}
