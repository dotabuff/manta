package manta

import (
	"math"
)

type fieldDecoder func(*reader) interface{}
type fieldFactory func(*field) fieldDecoder

var fieldTypeFactories = map[string]fieldFactory{
	"float32":                  floatFactory,
	"CNetworkedQuantizedFloat": quantizedFactory,
	"Vector":                   vectorFactory,
	"uint64":                   unsigned64Factory,
	"QAngle":                   qangleFactory,
	"CHandle":                  unsignedFactory,
	"CStrongHandle":            unsigned64Factory,
	"CEntityHandle":            unsignedFactory,
}

var fieldNameDecoders = map[string]fieldDecoder{}

var fieldTypeDecoders = map[string]fieldDecoder{
	"bool":    booleanDecoder,
	"char":    stringDecoder,
	"color32": unsignedDecoder,
	"int16":   signedDecoder,
	"int32":   signedDecoder,
	"int64":   signedDecoder,
	"int8":    signedDecoder,
	"uint16":  unsignedDecoder,
	"uint32":  unsignedDecoder,
	"uint8":   unsignedDecoder,

	"CBodyComponent":       componentDecoder,
	"CGameSceneNodeHandle": unsignedDecoder,
	"Color":                unsignedDecoder,
	"CPhysicsComponent":    componentDecoder,
	"CRenderComponent":     componentDecoder,
	"CUtlString":           stringDecoder,
	"CUtlStringToken":      unsignedDecoder,
	"CUtlSymbolLarge":      stringDecoder,
	"Vector2D":             vector2Decoder,
}

func unsignedFactory(f *field) fieldDecoder {
	return unsignedDecoder
}

func unsigned64Factory(f *field) fieldDecoder {
	switch f.encoder {
	case "fixed64":
		return fixed64Decoder
	}
	return unsigned64Decoder
}

func floatFactory(f *field) fieldDecoder {
	switch f.encoder {
	case "coord":
		return floatCoordDecoder
	case "simtime":
		return simulationTimeDecoder
	}

	if f.bitCount == nil || (*f.bitCount <= 0 || *f.bitCount >= 32) {
		return noscaleDecoder
	}

	return quantizedFactory(f)
}

func quantizedFactory(f *field) fieldDecoder {
	qfd := newQuantizedFloatDecoder(f.bitCount, f.encodeFlags, f.lowValue, f.highValue)
	return func(r *reader) interface{} {
		return qfd.decode(r)
	}
}

func vectorFactory(f *field) fieldDecoder {
	switch f.encoder {
	case "normal":
		return vectorNormalDecoder
	case "coord":
		return vectorCoordDecoder
	}

	d := floatFactory(f)
	return func(r *reader) interface{} {
		return []float32{
			d(r).(float32),
			d(r).(float32),
			d(r).(float32),
		}
	}
}

func vectorNormalDecoder(r *reader) interface{} {
	return r.read3BitNormal()
}

func vectorCoordDecoder(r *reader) interface{} {
	return []float32{
		r.readCoord(),
		r.readCoord(),
		r.readCoord(),
	}
}

func fixed64Decoder(r *reader) interface{} {
	return r.readLeUint64()
}

func handleDecoder(r *reader) interface{} {
	return r.readVarUint32()
}

func booleanDecoder(r *reader) interface{} {
	return r.readBoolean()
}

func stringDecoder(r *reader) interface{} {
	return r.readString()
}

func defaultDecoder(r *reader) interface{} {
	return r.readVarUint32()
}

func signedDecoder(r *reader) interface{} {
	return r.readVarInt32()
}

func floatCoordDecoder(r *reader) interface{} {
	return r.readCoord()
}

func noscaleDecoder(r *reader) interface{} {
	return math.Float32frombits(r.readBits(32))
}

func simulationTimeDecoder(r *reader) interface{} {
	return float32(r.readVarUint32()) * (1.0 / 30)
}

func qangleFactory(f *field) fieldDecoder {
	if f.encoder == "qangle_pitch_yaw" {
		n := uint32(*f.bitCount)
		return func(r *reader) interface{} {
			return []float32{
				r.readAngle(n),
				r.readAngle(n),
				0.0,
			}
		}
	}

	if f.bitCount != nil && *f.bitCount != 0 {
		n := uint32(*f.bitCount)
		return func(r *reader) interface{} {
			return []float32{
				r.readAngle(n),
				r.readAngle(n),
				r.readAngle(n),
			}
		}
	}

	return func(r *reader) interface{} {
		ret := make([]float32, 3)
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
}

func vector2Decoder(r *reader) interface{} {
	return []float32{r.readFloat(), r.readFloat()}
}

func unsignedDecoder(r *reader) interface{} {
	return uint64(r.readVarUint32())
}

func unsigned64Decoder(r *reader) interface{} {
	return r.readVarUint64()
}

func componentDecoder(r *reader) interface{} {
	return r.readBits(1)
}

func findDecoder(f *field) fieldDecoder {
	if v, ok := fieldTypeFactories[f.fieldType.baseType]; ok {
		return v(f)
	}

	if v, ok := fieldNameDecoders[f.varName]; ok {
		return v
	}

	if v, ok := fieldTypeDecoders[f.fieldType.baseType]; ok {
		return v
	}

	return defaultDecoder
}

func findDecoderByBaseType(baseType string) fieldDecoder {
	if v, ok := fieldTypeDecoders[baseType]; ok {
		return v
	}

	return defaultDecoder
}
