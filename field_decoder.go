package manta

import (
	"math"
)

type fieldDecoder func(*reader) cell
type fieldFactory func(*field) fieldDecoder

var fieldTypeFactories = map[string]fieldFactory{
	"float32":                  floatFactory,
	"CNetworkedQuantizedFloat": quantizedFactory,
	"Vector":                   vectorFactory(3),
	"Vector2D":                 vectorFactory(2),
	"Vector4D":                 vectorFactory(4),
	"VectorWS":                 vectorFactory(3),
	"Quaternion":               vectorFactory(4),
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
	"int64":   signed64Decoder,
	"int8":    signedDecoder,
	"uint16":  unsignedDecoder,
	"uint32":  unsignedDecoder,
	"uint8":   unsignedDecoder,

	"GameTime_t":     noscaleDecoder,
	"HeroFacetKey_t": unsigned64Decoder,
	"HeroID_t":       signedDecoder,
	"HSequence":      hSequenceDecoder,
	"BloodType":      bloodTypeDecoder,

	"CBodyComponent":       componentDecoder,
	"CGameSceneNodeHandle": unsignedDecoder,
	"Color":                unsignedDecoder,
	"CPhysicsComponent":    componentDecoder,
	"CRenderComponent":     componentDecoder,
	"CUtlString":           stringDecoder,
	"CUtlStringToken":      unsignedDecoder,
	"CUtlSymbolLarge":      stringDecoder,
	"CUtlBinaryBlock":      cUtlBinaryBlockDecoder,
	"CGlobalSymbol":        stringDecoder,
	"ResourceId_t":         unsigned64Decoder,
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
	case "runetime":
		return runeTimeDecoder
	}

	if f.bitCount == nil || (*f.bitCount <= 0 || *f.bitCount >= 32) {
		return noscaleDecoder
	}

	return quantizedFactory(f)
}

func quantizedFactory(f *field) fieldDecoder {
	qfd := newQuantizedFloatDecoder(f.bitCount, f.encodeFlags, f.lowValue, f.highValue)
	return func(r *reader) cell {
		return floatCell(qfd.decode(r))
	}
}

func vectorFactory(n int) fieldFactory {
	return func(f *field) fieldDecoder {
		if n == 3 && f.encoder == "normal" {
			return vectorNormalDecoder
		}

		d := floatFactory(f)
		return func(r *reader) cell {
			x := make([]float32, n)
			for i := 0; i < n; i++ {
				x[i] = d(r).asFloat32()
			}
			return refCell(x)
		}
	}
}

func vectorNormalDecoder(r *reader) cell {
	return refCell(r.read3BitNormal())
}

func fixed64Decoder(r *reader) cell {
	return refCell(r.readLeUint64())
}

func handleDecoder(r *reader) cell {
	return uint32Cell(r.readVarUint32())
}

func booleanDecoder(r *reader) cell {
	return boolCell(r.readBoolean())
}

func stringDecoder(r *reader) cell {
	return refCell(r.readString())
}

func defaultDecoder(r *reader) cell {
	return uint32Cell(r.readVarUint32())
}

func signedDecoder(r *reader) cell {
	return int32Cell(r.readVarInt32())
}

func signed64Decoder(r *reader) cell {
	return refCell(r.readVarInt64())
}

// bloodTypeDecoder reads a fixed 8-bit value, matching clarity's
// IntUnsignedDecoder(8). This is bit-identical to an unsigned varint only while
// the value is < 128; the full golden suite gates that this holds on the corpus.
func bloodTypeDecoder(r *reader) cell {
	return smallUint64Cell(uint64(r.readBits(8)))
}

// hSequenceDecoder decodes a sequence handle as an unsigned varint minus one,
// matching clarity's IntMinusOneDecoder. It returns a signed int32 so the
// "none" handle (wire value 0) is -1 rather than wrapping to a large unsigned.
func hSequenceDecoder(r *reader) cell {
	return int32Cell(int32(r.readVarUint32()) - 1)
}

// cUtlBinaryBlockDecoder reads a length-prefixed binary blob (varint length
// followed by that many bytes), matching clarity's CUtlBinaryBlockDecoder. The
// bytes are copied so the stored value does not alias the transient read buffer.
func cUtlBinaryBlockDecoder(r *reader) cell {
	n := r.readVarUint32()
	b := r.readBytes(n)
	out := make([]byte, len(b))
	copy(out, b)
	return refCell(out)
}

func floatCoordDecoder(r *reader) cell {
	return floatCell(r.readCoord())
}

func noscaleDecoder(r *reader) cell {
	return floatCell(math.Float32frombits(r.readBits(32)))
}

func runeTimeDecoder(r *reader) cell {
	return floatCell(math.Float32frombits(r.readBits(4)))
}

func simulationTimeDecoder(r *reader) cell {
	return floatCell(float32(r.readVarUint32()) * (1.0 / 30))
}

func qangleFactory(f *field) fieldDecoder {
	bc := uint32(0)
	if f.bitCount != nil {
		bc = uint32(*f.bitCount)
	}

	if f.encoder == "qangle_pitch_yaw" {
		// Bit counts 0 and 32 carry raw 32-bit floats; otherwise bit angles.
		if bc == 0 || bc == 32 {
			return func(r *reader) cell {
				return refCell([]float32{
					math.Float32frombits(r.readBits(32)),
					math.Float32frombits(r.readBits(32)),
					0.0,
				})
			}
		}
		return func(r *reader) cell {
			return refCell([]float32{
				r.readAngle(bc),
				r.readAngle(bc),
				0.0,
			})
		}
	}

	if f.encoder == "qangle_precise" {
		return func(r *reader) cell {
			ret := make([]float32, 3)
			rX := r.readBoolean()
			rY := r.readBoolean()
			rZ := r.readBoolean()
			if rX {
				ret[0] = r.readAngle(20)
			}
			if rY {
				ret[1] = r.readAngle(20)
			}
			if rZ {
				ret[2] = r.readAngle(20)
			}
			return refCell(ret)
		}
	}

	if bc == 32 {
		return func(r *reader) cell {
			return refCell([]float32{
				math.Float32frombits(r.readBits(32)),
				math.Float32frombits(r.readBits(32)),
				math.Float32frombits(r.readBits(32)),
			})
		}
	}

	if bc != 0 {
		return func(r *reader) cell {
			return refCell([]float32{
				r.readAngle(bc),
				r.readAngle(bc),
				r.readAngle(bc),
			})
		}
	}

	return func(r *reader) cell {
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
		return refCell(ret)
	}
}

func vector2Decoder(r *reader) cell {
	return refCell([]float32{r.readFloat(), r.readFloat()})
}

func unsignedDecoder(r *reader) cell {
	// readVarUint32 yields at most 32 bits, so the uint64 value fits inline.
	return smallUint64Cell(uint64(r.readVarUint32()))
}

func unsigned64Decoder(r *reader) cell {
	return refCell(r.readVarUint64())
}

func componentDecoder(r *reader) cell {
	return uint32Cell(r.readBits(1))
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
