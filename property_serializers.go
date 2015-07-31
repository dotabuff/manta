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
type DecodeFcn func(*Reader) interface{}

// Type for an array decoder function
type DecodeArrayFcn func(*Reader) interface{}

// PropertySerializer interface
type PropertySerializer struct {
	Decode          DecodeFcn
	IsArray         bool
	Length          uint32
	ArraySerializer *PropertySerializer
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

// Returns a serializer by name
func (pst *PropertySerializerTable) GetPropertySerializerByName(name string) *PropertySerializer {
	if ser := pst.Serializers[name]; ser != nil {
		return ser
	}

	if match := matchArray.FindStringSubmatch(name); match != nil {
		typeName := match[1]
		length, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			_panicf("Array length doesn't seem to be a number: %v", match[2])
		}

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = &PropertySerializer{}
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			IsArray:         true,
			Length:          uint32(length),
			ArraySerializer: serializer,
		}
		pst.Serializers[name] = ps
		return ps
	}

	// This function should panic at some point
	return &PropertySerializer{}
}
