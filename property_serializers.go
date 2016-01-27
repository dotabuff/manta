package manta

import (
	"regexp"
	"strconv"

	"github.com/golang/protobuf/proto"
)

// Type for a decoder function
type DecodeFcn func(*Reader, *dt_field) interface{}

// PropertySerializer interface
type PropertySerializer struct {
	Decode          DecodeFcn
	DecodeContainer DecodeFcn
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

// Regex for array and vector
var matchArray = regexp.MustCompile(`([^[\]]+)\[(\d+)]`)
var matchVector = regexp.MustCompile(`CUtlVector\<\s(.*)\s>$`)

// Fills serializer in dt_field
func (pst *PropertySerializerTable) FillSerializer(field *dt_field) {
	// Handle special decoders that need the complete field data here
	switch field.Name {
	case "m_flSimulationTime":
		field.Serializer = &PropertySerializer{decodeSimTime, nil, false, 0, nil, "unkown"}
		return
	case "m_flAnimTime":
		field.Serializer = &PropertySerializer{decodeSimTime, nil, false, 0, nil, "unkown"}
		return
	}

	// Handle special fields in old replays where the low and high values of a
	// quantized float were invalid.
	if field.build < 955 {
		switch field.Name {
		case "m_flMana", "m_flMaxMana":
			field.LowValue = nil
			field.HighValue = proto.Float32(8192.0)
		}

	}

	field.Serializer = pst.GetPropertySerializerByName(field.Type)
}

// Returns a serializer by name
func (pst *PropertySerializerTable) GetPropertySerializerByName(name string) *PropertySerializer {
	// Return existing serializer
	if ser := pst.Serializers[name]; ser != nil {
		return ser
	}

	// Set decoder
	var decoder DecodeFcn
	var decoderContainer DecodeFcn

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
		decoder = decodeSigned
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "Color":
		decoder = decodeUnsigned
	case "char":
		fallthrough
	case "CUtlSymbolLarge":
		decoder = decodeString
	case "Vector":
		decoder = decodeFVector
	case "bool":
		decoder = decodeBoolean
	case "CNetworkedQuantizedFloat":
		decoder = decodeQuantized
	case "CRenderComponent":
		fallthrough
	case "CPhysicsComponent":
		fallthrough
	case "CBodyComponent":
		decoder = decodeComponent
	case "QAngle":
		decoder = decodeQAngle
	case "CGameSceneNodeHandle":
		decoder = decodeHandle
	default:
		// check for specific types
		switch {
		case hasPrefix(name, "CHandle"):
			decoder = decodeHandle
		case hasPrefix(name, "CStrongHandle"):
			decoder = decodeUnsigned
		case hasPrefix(name, "CUtlVector< "):
			if match := matchVector.FindStringSubmatch(name); match != nil {
				decoderContainer = decodeVector
				decoder = pst.GetPropertySerializerByName(match[1]).Decode
			} else {
				_panicf("Unable to read vector type for %s", name)
			}
		default:
			//_debugf("No decoder for type %s", name)
		}
	}

	// match all pointers as boolean
	if name[len(name)-1:] == "*" {
		decoder = decodeBoolean
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
			DecodeContainer: decoderContainer,
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
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(1024),
			ArraySerializer: &PropertySerializer{},
		}
		pst.Serializers[name] = ps
		return ps
	}

	if name == "C_DOTA_ItemStockInfo[MAX_ITEM_STOCKS]" {
		typeName := "C_DOTA_ItemStockInfo"

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.GetPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			Decode:          serializer.Decode,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(8),
			ArraySerializer: serializer,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	if name == "CDOTA_AbilityDraftAbilityState[MAX_ABILITY_DRAFT_ABILITIES]" {
		typeName := "CDOTA_AbilityDraftAbilityState"

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.GetPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			Decode:          serializer.Decode,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(48),
			ArraySerializer: serializer,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	// That the type does not indicate an array is somewhat bad for the way we are
	// parsing things at the moment :(
	if name == "m_SpeechBubbles" {
		typeName := "m_SpeechBubbles"

		ps := &PropertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(5),
			ArraySerializer: nil,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	if name == "DOTA_PlayerChallengeInfo" || name == "DOTA_CombatLogQueryProgress" {
		ps := &PropertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(30),
			ArraySerializer: nil,
			Name:            name,
		}

		pst.Serializers[name] = ps
		return ps
	}

	// This function should panic at some point
	return &PropertySerializer{decoder, decoderContainer, false, 0, nil, "unkown"}
}
