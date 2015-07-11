package manta

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
type DecodeFcn func(*reader) interface{}

// Type for an array decoder function
type DecodeArrayFcn func(*reader) interface{}

// PropertySerializer interface
type PropertySerializer struct {
	Decode      DecodeFcn
	DecodeArray DecodeArrayFcn
	IsArray     bool
	Length      uint32
}

// Contains a list of available property serializers
type PropertySerializerTable struct {
	Serializers map[string]*PropertySerializer
}

// Returns a table containing all know property serializers
func GetDefaultPropertySerializerTable() *PropertySerializerTable {
	// Init table
	tbl := &PropertySerializerTable{}
	tbl.Serializers = make(map[string]*PropertySerializer)

	// Append default serializers/decoders
	// For now, only arrays are added

	tbl.Serializers["float32[24]"] = &PropertySerializer{
		nil, nil, true, 24,
	}

	tbl.Serializers["CStrongHandle< InfoForResourceTypeIMaterial2 >[6]"] = &PropertySerializer{
		nil, nil, true, 6,
	}

	return tbl
}

// Returns a serializer by name
func (pst *PropertySerializerTable) GetPropertySerializerByName(name string) *PropertySerializer {
	ser := pst.Serializers[name]

	if ser == nil {
		// This function should panic at some point
		return &PropertySerializer{
			nil, nil, false, 0,
		}
	}

	return ser
}
