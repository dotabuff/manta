package manta

import (
	"math"
	"strconv"
	"sync"
)

func decodeLeUint64(r *Reader, f *dt_field) interface{} {
	return r.readLeUint64()
}

func decodeHandle(r *Reader, f *dt_field) interface{} {
	return r.readVarUint32()
}

func decodeByte(r *Reader, f *dt_field) interface{} {
	return r.readBits(8)
}

func decodeShort(r *Reader, f *dt_field) interface{} {
	return r.readBits(16)
}

func decodeUnsigned(r *Reader, f *dt_field) interface{} {
	switch f.Encoder {
	case "fixed64":
		return decodeLeUint64(r, f)
	}

	return r.readVarUint64()
}

func decodeSigned(r *Reader, f *dt_field) interface{} {
	return r.readVarInt32()
}

func decodeSigned64(r *Reader, f *dt_field) interface{} {
	return r.readVarInt64()
}

func decodeFixed32(r *Reader, f *dt_field) interface{} {
	return r.readBits(32)
}

func decodeFixed64(r *Reader, f *dt_field) interface{} {
	ret := uint64(r.readBits(32))

	ret = (ret << 32) | uint64(r.readBits(32))
	return ret
}

func decodeBoolean(r *Reader, f *dt_field) interface{} {
	return r.readBoolean()
}

func decodeFloat(r *Reader, f *dt_field) interface{} {
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

func decodeFloatNoscale(r *Reader, f *dt_field) interface{} {
	return math.Float32frombits(r.readBits(int(*f.BitCount)))
}

// A list of field -> encoder types with mutex
var (
	qMap = map[*dt_field]*QuantizedFloatDecoder{}
	qMu  = sync.Mutex{}
)

func decodeQuantized(r *Reader, f *dt_field) interface{} {
	// Find or create the correct field decoder
	qMu.Lock()
	if _, ok := qMap[f]; !ok {
		qMap[f] = InitQFD(f)
	}
	q := qMap[f]
	qMu.Unlock()

	// Decode value
	_debugf(
		"Bitcount: %v, Low: %v, High: %v, Flags: %v",
		q.Bitcount,
		q.Low,
		q.High,
		strconv.FormatInt(int64(q.Flags), 2),
	)

	return q.Decode(r)
}

func decodeSimTime(r *Reader, f *dt_field) interface{} {
	return float32(r.readVarUint32()) * (1.0 / 30)
}

func decodeString(r *Reader, f *dt_field) interface{} {
	return r.readString()
}

func decodeVector(r *Reader, f *dt_field) interface{} {
	return r.readVarUint32()
}

func decodeClass(r *Reader, f *dt_field) interface{} {
	return r.readVarUint32()
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
	ret := [3]float32{0.0, 0.0, 0.0}

	// Parse specific encoders
	switch f.Encoder {
	case "qangle_pitch_yaw":
		if f.BitCount != nil && f.Flags != nil && (*f.Flags&0x20 != 0) {
			_panicf("Special Case: Unkown for now")
		}

		ret[0] = r.readAngle(uint(*f.BitCount))
		ret[1] = r.readAngle(uint(*f.BitCount))
		return ret
	}

	// Parse a standard angle
	if f.BitCount != nil && *f.BitCount == 32 {
		_panicf("Special Case: Unkown for now")
	} else if f.BitCount != nil && *f.BitCount != 0 {
		ret[0] = r.readAngle(uint(*f.BitCount))
		ret[1] = r.readAngle(uint(*f.BitCount))
		ret[2] = r.readAngle(uint(*f.BitCount))

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
