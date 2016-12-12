package manta

import (
	"regexp"
	"strconv"

	"github.com/golang/protobuf/proto"
)

// Type for a decoder function
type decodeFn func(*reader, *dtField) interface{}

// PropertySerializer interface
type propertySerializer struct {
	Decode          decodeFn
	DecodeContainer decodeFn
	IsArray         bool
	Length          uint32
	ArraySerializer *propertySerializer
	Name            string
}

// Contains a list of available property serializers
type propertySerializerTable struct {
	Serializers map[string]*propertySerializer
}

// Returns a table containing all know property serializers
func newPropertySerializerTable() *propertySerializerTable {
	return &propertySerializerTable{
		Serializers: map[string]*propertySerializer{},
	}
}

// Regex for array and vector
var matchArray = regexp.MustCompile(`([^[\]]+)\[(\d+)]`)
var matchVector = regexp.MustCompile(`CUtlVector\<\s(.*)\s>$`)

// Fills serializer in dtField
func (pst *propertySerializerTable) FillSerializer(field *dtField) {
	// Handle special decoders that need the complete field data here
	switch field.Name {
	case "m_flSimulationTime":
		field.Serializer = &propertySerializer{decodeSimTime, nil, false, 0, nil, "unkown"}
		return
	case "m_flAnimTime":
		field.Serializer = &propertySerializer{decodeSimTime, nil, false, 0, nil, "unkown"}
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

	field.Serializer = pst.getPropertySerializerByName(field.Type)
}

// Returns a serializer by name
func (pst *propertySerializerTable) getPropertySerializerByName(name string) *propertySerializer {
	// Return existing serializer
	if ser := pst.Serializers[name]; ser != nil {
		return ser
	}

	// Set decoder
	var decoder decodeFn
	var decoderContainer decodeFn

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
	case "Vector2D":
		decoder = decodeVector2D
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
				decoder = pst.getPropertySerializerByName(match[1]).Decode
			} else {
				_panicf("Unable to read vector type for %s", name)
			}
		default:
			if v(6) {
				_debugf("no decoder for type %s", name)
			}
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
			serializer = pst.getPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &propertySerializer{
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
		ps := &propertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(1024),
			ArraySerializer: &propertySerializer{},
		}
		pst.Serializers[name] = ps
		return ps
	}

	if name == "C_DOTA_ItemStockInfo[MAX_ITEM_STOCKS]" {
		typeName := "C_DOTA_ItemStockInfo"

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.getPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &propertySerializer{
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
			serializer = pst.getPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &propertySerializer{
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

		ps := &propertySerializer{
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
		ps := &propertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(1024),
			ArraySerializer: nil,
			Name:            name,
		}

		pst.Serializers[name] = ps
		return ps
	}

	// This function should panic at some point
	return &propertySerializer{decoder, decoderContainer, false, 0, nil, "unkown"}
}
