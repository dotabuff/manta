package manta

import (
	"math"
)

func decodeLeUint64(r *reader, f *dtField) interface{} {
	return r.readLeUint64()
}

func decodeHandle(r *reader, f *dtField) interface{} {
	return r.readVarUint32()
}

func decodeByte(r *reader, f *dtField) interface{} {
	return r.readBits(8)
}

func decodeShort(r *reader, f *dtField) interface{} {
	return r.readBits(16)
}

func decodeUnsigned(r *reader, f *dtField) interface{} {
	switch f.Encoder {
	case "fixed64":
		return decodeLeUint64(r, f)
	}

	return r.readVarUint64()
}

func decodeSigned(r *reader, f *dtField) interface{} {
	return r.readVarInt32()
}

func decodeSigned64(r *reader, f *dtField) interface{} {
	return r.readVarInt64()
}

func decodeFixed32(r *reader, f *dtField) interface{} {
	return r.readBits(32)
}

func decodeFixed64(r *reader, f *dtField) interface{} {
	ret := uint64(r.readBits(32))

	ret = (ret << 32) | uint64(r.readBits(32))
	return ret
}

func decodeBoolean(r *reader, f *dtField) interface{} {
	return r.readBoolean()
}

func decodeFloat(r *reader, f *dtField) interface{} {
	// Parse specific encoders
	switch f.Encoder {
	case "coord":
		return r.readCoord()
	}

	// Decode as noscale if it has an appropriate bitcount.
	if f.BitCount == nil || (*f.BitCount <= 0 || *f.BitCount >= 32) {
		return decodeFloatNoscale(r, f)
	}

	// Otherwise fall back to quantized decoding.
	return decodeQuantized(r, f)
}

func decodeFloatNoscale(r *reader, f *dtField) interface{} {
	return math.Float32frombits(r.readBits(uint32(*f.BitCount)))
}

func decodeQuantized(r *reader, f *dtField) interface{} {
	return newQuantizedFloatDecoder(f).decode(r)
}

func decodeSimTime(r *reader, f *dtField) interface{} {
	return float32(r.readVarUint32()) * (1.0 / 30)
}

func decodeString(r *reader, f *dtField) interface{} {
	return r.readString()
}

func decodeVector(r *reader, f *dtField) interface{} {
	return r.readVarUint32()
}

func decodeClass(r *reader, f *dtField) interface{} {
	return r.readVarUint32()
}

func decodeFVector(r *reader, f *dtField) interface{} {
	// Parse specific encoders
	switch f.Encoder {
	case "normal":
		return r.read3BitNormal()
	}

	return []float32{decodeFloat(r, f).(float32), decodeFloat(r, f).(float32), decodeFloat(r, f).(float32)}
}

func decodeNop(r *reader, f *dtField) interface{} {
	return 0
}

func decodePointer(r *reader, f *dtField) interface{} {
	// Seems to be encoded as a single bit, not sure what to make of it
	if !r.readBoolean() {
		_panicf("Figure out how this works")
	}

	return 0
}

func decodeQAngle(r *reader, f *dtField) interface{} {
	ret := [3]float32{0.0, 0.0, 0.0}

	// Parse specific encoders
	switch f.Encoder {
	case "qangle_pitch_yaw":
		if f.BitCount != nil && f.Flags != nil && (*f.Flags&0x20 != 0) {
			_panicf("Special Case: Unkown for now")
		}

		ret[0] = r.readAngle(uint32(*f.BitCount))
		ret[1] = r.readAngle(uint32(*f.BitCount))
		return ret
	}

	// Parse a standard angle
	if f.BitCount != nil && *f.BitCount == 32 {
		_panicf("Special Case: Unkown for now")
	} else if f.BitCount != nil && *f.BitCount != 0 {
		ret[0] = r.readAngle(uint32(*f.BitCount))
		ret[1] = r.readAngle(uint32(*f.BitCount))
		ret[2] = r.readAngle(uint32(*f.BitCount))

		return ret
	} else {
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

	_panicf("No valid encoding determined")
	return ret
}

func decodeComponent(r *reader, f *dtField) interface{} {
	return r.readBits(1)
}

func decodeHSequence(r *reader, f *dtField) interface{} {
	// wrong, just testing
	return r.readBits(1)
}

func decodeVector2D(r *reader, f *dtField) interface{} {
	return []float32{r.readFloat(), r.readFloat()}
}
