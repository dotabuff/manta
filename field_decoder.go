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
	"uint64":                   unsignedFactory,
	"QAngle":                   qangleFactory,
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
	"CEntityHandle":        unsignedDecoder,
	"CGameSceneNodeHandle": unsignedDecoder,
	"CHandle":              handleDecoder,
	"Color":                unsignedDecoder,
	"CPhysicsComponent":    componentDecoder,
	"CRenderComponent":     componentDecoder,
	"CStrongHandle":        unsignedDecoder,
	"CUtlStringToken":      unsignedDecoder,
	"CUtlSymbolLarge":      stringDecoder,
	"Vector2D":             vector2Decoder,
}

func unsignedFactory(f *field) fieldDecoder {
	if f.encoder == "fixed64" {
		return fixedDecoder
	}

	return unsignedDecoder
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
	dt := &dtField{
		BitCount:  f.bitCount,
		Encoder:   f.encoder,
		Flags:     f.encodeFlags,
		HighValue: f.highValue,
		LowValue:  f.lowValue,
		Name:      f.varName,
		Type:      f.varType,
	}

	return func(r *reader) interface{} {
		// _printf(" dt: %+v", dt)
		return newQuantizedFloatDecoder(dt).decode(r)
	}
}

func vectorFactory(f *field) fieldDecoder {
	switch f.encoder {
	case "normal":
		return floatNormalDecoder
	case "coord":
		return floatCoordDecoder
	}

	return func(r *reader) interface{} {
		d := floatFactory(f)
		return []float32{
			d(r).(float32),
			d(r).(float32),
			d(r).(float32),
		}
	}
}

func floatNormalDecoder(r *reader) interface{} {
	return r.read3BitNormal()
}

func fixedDecoder(r *reader) interface{} {
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
