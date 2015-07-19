package manta

import (
	"regexp"
	"strconv"
)

// Required for CWorld to work:
// ----------------------------
// - CHandle< CBaseEntity >
// - CBodyComponent
// - CNetworkedQuantizedFloat
// - CGameSceneNodeHandle
// - CStrongHandle< InfoForResourceTypeCTextureBase >
// - CStrongHandle< InfoForResourceTypeCModel >
// - CUtlStringToken
// - CUtlVector< CAnimationLayer >
// - CEntityIdentity*
// - CUtlSymbolLarge
// - CPhysicsComponent
// - CRenderComponent
//
// - Color
// - QAngle
// - HSequence
// - Vector
// - SolidType_t
// - SurroundingBoundsType_t
// - MoveCollide_t
// - MoveType_t
// - gender_t
// - RenderMode_t
// - RenderFx_t
//
// - bool
// - uint8
// - uint16
// - uint32
// - uint64
// - int8
// - int32
// - float32
//
// - float32[24]
// - CStrongHandle< InfoForResourceTypeIMaterial2 >[6]

// Type for a decoder function
type DecodeFcn func(*reader, *dt_field) interface{}

// Type for an array decoder function
type DecodeArrayFcn func(*reader, *dt_field) interface{}

// PropertySerializer interface
type PropertySerializer struct {
	Decode          DecodeFcn
	IsArray         bool
	Length          uint32
	ArraySerializer *PropertySerializer
	Name            string
}

// Contains a list of available property serializers
type PropertySerializerTable struct {
	Serializers map[string]*PropertySerializer
}

// Returns a table containing all know property serializers
func GetDefaultPropertySerializerTable() *PropertySerializerTable {
	return &PropertySerializerTable{Serializers: map[string]*PropertySerializer{}}
}

var matchArray = regexp.MustCompile(`([^[\]]+)\[(\d+)]`)
var matchVector = regexp.MustCompile(`CUtlVector.*`)

// Returns a serializer by name
func (pst *PropertySerializerTable) GetPropertySerializerByName(name string) *PropertySerializer {
	// Return existing serializer
	if ser := pst.Serializers[name]; ser != nil {
		return ser
	}

	// Set decoder
	var decoder DecodeFcn
	switch name {
	case "float32":
		decoder = decodeFloat
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "Color":
		decoder = decodeSigned
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		decoder = decodeUnsigned
	case "char":
		fallthrough
	case "CUtlSymbolLarge":
		decoder = decodeString
	case "bool":
		decoder = decodeBoolean
	default:
		// check for specific types
		switch {
		case hasPrefix(name, "CHandle"):
			decoder = decodeHandle
		case hasPrefix(name, "CUtlVector< "):
			decoder = pst.GetPropertySerializerByName(name[12:]).Decode
		default:
			//_debugf("No decoder for type %s", name)
		}
	}

	// create a new serializer based on it's name
	if match := matchArray.FindStringSubmatch(name); match != nil {
		typeName := match[1]
		length, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			_panicf("Array length doesn't seem to be a number: %v", match[2])
		}

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.GetPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			Decode:          serializer.Decode,
			IsArray:         true,
			Length:          uint32(length),
			ArraySerializer: serializer,
			Name:            typeName,
		}
		pst.Serializers[name] = ps
		return ps
	}

	if match := matchVector.FindStringSubmatch(name); match != nil {
		ps := &PropertySerializer{
			Decode:          decoder,
			IsArray:         true,
			Length:          uint32(128),
			ArraySerializer: &PropertySerializer{},
		}
		pst.Serializers[name] = ps
		return ps
	}

	// This function should panic at some point
	return &PropertySerializer{decoder, false, 0, nil, "unkown"}
}
