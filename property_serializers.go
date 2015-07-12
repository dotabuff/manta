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
	Decode          DecodeFcn
	DecodeArray     DecodeArrayFcn
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
	// Init table
	tbl := &PropertySerializerTable{}
	tbl.Serializers = make(map[string]*PropertySerializer)

	// Append default serializers/decoders
	// For now, only arrays are added

	// CWorld
	tbl.Serializers["float32"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["float32[24]"] = &PropertySerializer{
		nil, nil, true, 24, tbl.Serializers["float32"],
	}

	tbl.Serializers["CStrongHandle< InfoForResourceTypeIMaterial2 >"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["CStrongHandle< InfoForResourceTypeIMaterial2 >[6]"] = &PropertySerializer{
		nil, nil, true, 6, tbl.Serializers["CStrongHandle< InfoForResourceTypeIMaterial2 >"],
	}

	//CDOTA_PlayerResource
	tbl.Serializers["float32[10]"] = &PropertySerializer{
		nil, nil, true, 10, tbl.Serializers["float32"],
	}

	tbl.Serializers["int32"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["int32[10]"] = &PropertySerializer{
		nil, nil, true, 10, tbl.Serializers["int32"],
	}

	tbl.Serializers["int32[20]"] = &PropertySerializer{
		nil, nil, true, 20, tbl.Serializers["int32"],
	}

	tbl.Serializers["int32[64]"] = &PropertySerializer{
		nil, nil, true, 64, tbl.Serializers["int32"],
	}

	tbl.Serializers["uint16"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["uint16[64]"] = &PropertySerializer{
		nil, nil, true, 64, tbl.Serializers["uint16"],
	}

	tbl.Serializers["uint64"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["uint64[64]"] = &PropertySerializer{
		nil, nil, true, 64, tbl.Serializers["uint64"],
	}

	tbl.Serializers["uint64[128]"] = &PropertySerializer{
		nil, nil, true, 128, tbl.Serializers["uint64"],
	}

	tbl.Serializers["uint64[256]"] = &PropertySerializer{
		nil, nil, true, 256, tbl.Serializers["uint64"],
	}

	tbl.Serializers["CUtlSymbolLarge"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["CUtlSymbolLarge[64]"] = &PropertySerializer{
		nil, nil, true, 64, tbl.Serializers["CUtlSymbolLarge"],
	}

	tbl.Serializers["CUtlSymbolLarge[6]"] = &PropertySerializer{
		nil, nil, true, 6, tbl.Serializers["CUtlSymbolLarge"],
	}

	tbl.Serializers["bool"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["bool[64]"] = &PropertySerializer{
		nil, nil, true, 64, tbl.Serializers["bool"],
	}

	tbl.Serializers["bool[10]"] = &PropertySerializer{
		nil, nil, true, 10, tbl.Serializers["bool"],
	}

	tbl.Serializers["Color"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["Color[10]"] = &PropertySerializer{
		nil, nil, true, 10, tbl.Serializers["Color"],
	}

	tbl.Serializers["CHandle< CBaseEntity >"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	tbl.Serializers["CHandle< CBaseEntity >[10]"] = &PropertySerializer{
		nil, nil, true, 10, tbl.Serializers["CHandle< CBaseEntity >"],
	}

	return tbl
}

// Returns a serializer by name
func (pst *PropertySerializerTable) GetPropertySerializerByName(name string) *PropertySerializer {
	ser := pst.Serializers[name]

	if ser == nil {
		// This function should panic at some point
		return &PropertySerializer{
			nil, nil, false, 0, nil,
		}
	}

	return ser
}
