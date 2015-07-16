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
	tbl := &PropertySerializerTable{Serializers: map[string]*PropertySerializer{}}
	srz := tbl.Serializers

	// Append default serializers/decoders
	// For now, only arrays are added

	// CWorld
	srz["float32"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["float32[24]"] = &PropertySerializer{
		nil, nil, true, 24, srz["float32"],
	}

	srz["CStrongHandle< InfoForResourceTypeIMaterial2 >"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["CStrongHandle< InfoForResourceTypeIMaterial2 >[6]"] = &PropertySerializer{
		nil, nil, true, 6, srz["CStrongHandle< InfoForResourceTypeIMaterial2 >"],
	}

	//CDOTA_PlayerResource
	srz["float32[10]"] = &PropertySerializer{
		nil, nil, true, 10, srz["float32"],
	}

	srz["int32"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["int32[10]"] = &PropertySerializer{
		nil, nil, true, 10, srz["int32"],
	}

	srz["int32[20]"] = &PropertySerializer{
		nil, nil, true, 20, srz["int32"],
	}

	srz["int32[64]"] = &PropertySerializer{
		nil, nil, true, 64, srz["int32"],
	}

	srz["uint16"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["uint16[64]"] = &PropertySerializer{
		nil, nil, true, 64, srz["uint16"],
	}

	srz["uint64"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["uint64[64]"] = &PropertySerializer{
		nil, nil, true, 64, srz["uint64"],
	}

	srz["uint64[128]"] = &PropertySerializer{
		nil, nil, true, 128, srz["uint64"],
	}

	srz["uint64[256]"] = &PropertySerializer{
		nil, nil, true, 256, srz["uint64"],
	}

	srz["CUtlSymbolLarge"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["CUtlSymbolLarge[64]"] = &PropertySerializer{
		nil, nil, true, 64, srz["CUtlSymbolLarge"],
	}

	srz["CUtlSymbolLarge[6]"] = &PropertySerializer{
		nil, nil, true, 6, srz["CUtlSymbolLarge"],
	}

	srz["bool"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["bool[64]"] = &PropertySerializer{
		nil, nil, true, 64, srz["bool"],
	}

	srz["bool[10]"] = &PropertySerializer{
		nil, nil, true, 10, srz["bool"],
	}

	srz["Color"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["Color[10]"] = &PropertySerializer{
		nil, nil, true, 10, srz["Color"],
	}

	srz["CHandle< CBaseEntity >"] = &PropertySerializer{
		nil, nil, false, 0, nil,
	}

	srz["CHandle< CBaseEntity >[10]"] = &PropertySerializer{
		nil, nil, true, 10, srz["CHandle< CBaseEntity >"],
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
