// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dota_hud_types.proto

package dota

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EHeroSelectionText int32

const (
	EHeroSelectionText_k_EHeroSelectionText_Invalid                              EHeroSelectionText = -1
	EHeroSelectionText_k_EHeroSelectionText_None                                 EHeroSelectionText = 0
	EHeroSelectionText_k_EHeroSelectionText_ChooseHero                           EHeroSelectionText = 1
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Planning_YouFirst           EHeroSelectionText = 2
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Planning_TheyFirst          EHeroSelectionText = 3
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Banning                     EHeroSelectionText = 4
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Ban_Waiting                 EHeroSelectionText = 5
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_PickTwo                     EHeroSelectionText = 6
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_PickOneMore                 EHeroSelectionText = 7
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_PickOne                     EHeroSelectionText = 8
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_WaitingRadiant              EHeroSelectionText = 9
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_WaitingDire                 EHeroSelectionText = 10
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TeammateRandomed            EHeroSelectionText = 11
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_YouPicking_LosingGold       EHeroSelectionText = 12
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold      EHeroSelectionText = 13
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_ChooseCaptain           EHeroSelectionText = 14
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain EHeroSelectionText = 15
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_YouSelect               EHeroSelectionText = 16
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_TheySelect              EHeroSelectionText = 17
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_YouBan                  EHeroSelectionText = 18
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_TheyBan                 EHeroSelectionText = 19
	EHeroSelectionText_k_EHeroSelectionText_RandomDraft_HeroReview               EHeroSelectionText = 20
	EHeroSelectionText_k_EHeroSelectionText_RandomDraft_RoundDisplay             EHeroSelectionText = 21
	EHeroSelectionText_k_EHeroSelectionText_RandomDraft_Waiting                  EHeroSelectionText = 22
	EHeroSelectionText_k_EHeroSelectionText_EventGame_BanPhase                   EHeroSelectionText = 23
)

var EHeroSelectionText_name = map[int32]string{
	-1: "k_EHeroSelectionText_Invalid",
	0:  "k_EHeroSelectionText_None",
	1:  "k_EHeroSelectionText_ChooseHero",
	2:  "k_EHeroSelectionText_AllDraft_Planning_YouFirst",
	3:  "k_EHeroSelectionText_AllDraft_Planning_TheyFirst",
	4:  "k_EHeroSelectionText_AllDraft_Banning",
	5:  "k_EHeroSelectionText_AllDraft_Ban_Waiting",
	6:  "k_EHeroSelectionText_AllDraft_PickTwo",
	7:  "k_EHeroSelectionText_AllDraft_PickOneMore",
	8:  "k_EHeroSelectionText_AllDraft_PickOne",
	9:  "k_EHeroSelectionText_AllDraft_WaitingRadiant",
	10: "k_EHeroSelectionText_AllDraft_WaitingDire",
	11: "k_EHeroSelectionText_AllDraft_TeammateRandomed",
	12: "k_EHeroSelectionText_AllDraft_YouPicking_LosingGold",
	13: "k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold",
	14: "k_EHeroSelectionText_CaptainsMode_ChooseCaptain",
	15: "k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain",
	16: "k_EHeroSelectionText_CaptainsMode_YouSelect",
	17: "k_EHeroSelectionText_CaptainsMode_TheySelect",
	18: "k_EHeroSelectionText_CaptainsMode_YouBan",
	19: "k_EHeroSelectionText_CaptainsMode_TheyBan",
	20: "k_EHeroSelectionText_RandomDraft_HeroReview",
	21: "k_EHeroSelectionText_RandomDraft_RoundDisplay",
	22: "k_EHeroSelectionText_RandomDraft_Waiting",
	23: "k_EHeroSelectionText_EventGame_BanPhase",
}

var EHeroSelectionText_value = map[string]int32{
	"k_EHeroSelectionText_Invalid":                              -1,
	"k_EHeroSelectionText_None":                                 0,
	"k_EHeroSelectionText_ChooseHero":                           1,
	"k_EHeroSelectionText_AllDraft_Planning_YouFirst":           2,
	"k_EHeroSelectionText_AllDraft_Planning_TheyFirst":          3,
	"k_EHeroSelectionText_AllDraft_Banning":                     4,
	"k_EHeroSelectionText_AllDraft_Ban_Waiting":                 5,
	"k_EHeroSelectionText_AllDraft_PickTwo":                     6,
	"k_EHeroSelectionText_AllDraft_PickOneMore":                 7,
	"k_EHeroSelectionText_AllDraft_PickOne":                     8,
	"k_EHeroSelectionText_AllDraft_WaitingRadiant":              9,
	"k_EHeroSelectionText_AllDraft_WaitingDire":                 10,
	"k_EHeroSelectionText_AllDraft_TeammateRandomed":            11,
	"k_EHeroSelectionText_AllDraft_YouPicking_LosingGold":       12,
	"k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold":      13,
	"k_EHeroSelectionText_CaptainsMode_ChooseCaptain":           14,
	"k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain": 15,
	"k_EHeroSelectionText_CaptainsMode_YouSelect":               16,
	"k_EHeroSelectionText_CaptainsMode_TheySelect":              17,
	"k_EHeroSelectionText_CaptainsMode_YouBan":                  18,
	"k_EHeroSelectionText_CaptainsMode_TheyBan":                 19,
	"k_EHeroSelectionText_RandomDraft_HeroReview":               20,
	"k_EHeroSelectionText_RandomDraft_RoundDisplay":             21,
	"k_EHeroSelectionText_RandomDraft_Waiting":                  22,
	"k_EHeroSelectionText_EventGame_BanPhase":                   23,
}

func (x EHeroSelectionText) Enum() *EHeroSelectionText {
	p := new(EHeroSelectionText)
	*p = x
	return p
}

func (x EHeroSelectionText) String() string {
	return proto.EnumName(EHeroSelectionText_name, int32(x))
}

func (x *EHeroSelectionText) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EHeroSelectionText_value, data, "EHeroSelectionText")
	if err != nil {
		return err
	}
	*x = EHeroSelectionText(value)
	return nil
}

func (EHeroSelectionText) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b16d5740bf0aa997, []int{0}
}

var E_HudLocalizeToken = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50501,
	Name:          "dota.hud_localize_token",
	Tag:           "bytes,50501,opt,name=hud_localize_token",
	Filename:      "dota_hud_types.proto",
}

func init() {
	proto.RegisterEnum("dota.EHeroSelectionText", EHeroSelectionText_name, EHeroSelectionText_value)
	proto.RegisterExtension(E_HudLocalizeToken)
}

func init() {
	proto.RegisterFile("dota_hud_types.proto", fileDescriptor_b16d5740bf0aa997)
}

var fileDescriptor_b16d5740bf0aa997 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x52, 0x13, 0x3d,
	0x1c, 0xc6, 0xe9, 0x0b, 0xef, 0x07, 0x79, 0xdf, 0x57, 0x63, 0x44, 0x8d, 0x19, 0x1d, 0x70, 0x18,
	0x06, 0xca, 0x47, 0x8b, 0x80, 0x22, 0x3a, 0x1e, 0x50, 0xca, 0x87, 0x33, 0x40, 0x6b, 0xed, 0x88,
	0x38, 0x38, 0x6b, 0xe8, 0x86, 0x76, 0xed, 0x36, 0xe9, 0xec, 0x07, 0x58, 0xc7, 0x03, 0x8f, 0xbd,
	0x94, 0xdc, 0x87, 0x77, 0x90, 0xfb, 0xf0, 0x12, 0x74, 0x42, 0xb7, 0x58, 0x64, 0xbf, 0xda, 0x83,
	0x1e, 0xfc, 0xf7, 0x79, 0x9e, 0x5f, 0xfe, 0xc9, 0x26, 0x1b, 0x30, 0x66, 0x0a, 0x8f, 0x1a, 0x0d,
	0xdf, 0x34, 0xbc, 0x4e, 0x9b, 0xb9, 0xb9, 0xb6, 0x23, 0x3c, 0x81, 0x46, 0x74, 0x95, 0x4c, 0xd4,
	0x85, 0xa8, 0xdb, 0x2c, 0x7f, 0x5e, 0x3b, 0xf6, 0x4f, 0xf2, 0x26, 0x73, 0x6b, 0x8e, 0xd5, 0xf6,
	0x84, 0xd3, 0xd5, 0xcd, 0x7e, 0x47, 0x00, 0x6d, 0xee, 0x30, 0x47, 0xbc, 0x62, 0x36, 0xab, 0x79,
	0x96, 0xe0, 0x55, 0xf6, 0xd1, 0x43, 0x59, 0x70, 0xaf, 0x69, 0x5c, 0xad, 0x1b, 0x2f, 0xf8, 0x29,
	0xb5, 0x2d, 0x13, 0xfe, 0xe8, 0xfd, 0x32, 0xe8, 0x3e, 0xb8, 0x1b, 0x2a, 0xdd, 0x17, 0x9c, 0xc1,
	0x21, 0xb4, 0x03, 0xc6, 0x43, 0x1f, 0x6f, 0x34, 0x84, 0x70, 0x99, 0xae, 0xc3, 0x0c, 0x99, 0x94,
	0x0a, 0x8f, 0x4f, 0x16, 0x4b, 0xd5, 0x75, 0x43, 0x57, 0x8c, 0x0b, 0x69, 0x9f, 0x0c, 0x35, 0x40,
	0x3e, 0x34, 0x69, 0xdd, 0xb6, 0x8b, 0x0e, 0x3d, 0xf1, 0x8c, 0xb2, 0x4d, 0x39, 0xb7, 0x78, 0xdd,
	0x38, 0x14, 0xfe, 0x96, 0xe5, 0xb8, 0x1e, 0xfc, 0x83, 0x2c, 0x4b, 0x85, 0xf3, 0xa1, 0xc9, 0xd1,
	0x36, 0xf4, 0x01, 0x2c, 0xa6, 0x24, 0x55, 0x1b, 0xac, 0xd3, 0x45, 0x0d, 0x93, 0x15, 0xa9, 0xf0,
	0x62, 0x4a, 0xd4, 0x85, 0x0f, 0x1d, 0x81, 0xa9, 0x78, 0x56, 0xa1, 0x6b, 0x81, 0x23, 0xe4, 0xa1,
	0x54, 0x78, 0x21, 0x0a, 0x50, 0xb6, 0x6a, 0xcd, 0x9e, 0xd6, 0xd8, 0x17, 0x2d, 0x8b, 0x53, 0x8f,
	0x21, 0x0a, 0xb2, 0x89, 0xe9, 0xc6, 0x01, 0xb5, 0x3c, 0x4d, 0xf8, 0x93, 0x2c, 0x49, 0x85, 0x73,
	0x03, 0x11, 0x4c, 0xf4, 0x2e, 0xa9, 0x01, 0xed, 0xab, 0x9e, 0x09, 0xf8, 0x57, 0x8a, 0x78, 0xfd,
	0x57, 0x6e, 0x50, 0x97, 0xf5, 0x5c, 0x88, 0x25, 0x75, 0xa0, 0x85, 0x25, 0xce, 0xf6, 0x84, 0xc3,
	0xe0, 0xdf, 0xe4, 0xb1, 0x54, 0x78, 0x69, 0x00, 0x44, 0xe0, 0x4c, 0xd7, 0x45, 0x89, 0x33, 0xf8,
	0xcf, 0xc0, 0x5d, 0x94, 0x38, 0x43, 0x4d, 0x30, 0x1f, 0x1f, 0x1f, 0xac, 0x41, 0x85, 0x9a, 0x16,
	0xe5, 0x1e, 0x1c, 0x25, 0x6b, 0x52, 0xe1, 0x47, 0x29, 0x29, 0x97, 0xcd, 0xc9, 0x53, 0x16, 0xe8,
	0x8b, 0x96, 0xc3, 0x20, 0x18, 0x68, 0xca, 0xfa, 0x9c, 0x48, 0x80, 0x5c, 0x3c, 0xa6, 0xca, 0x68,
	0xab, 0x45, 0x3d, 0x56, 0xa1, 0xdc, 0x14, 0x2d, 0x66, 0xc2, 0x7f, 0xc9, 0x33, 0xa9, 0xf0, 0x6a,
	0xfc, 0x1e, 0xf9, 0xdd, 0x65, 0x94, 0x29, 0x17, 0x0e, 0x6d, 0x51, 0xd4, 0x06, 0xcb, 0xf1, 0xc0,
	0x43, 0xe1, 0xeb, 0xa1, 0xea, 0xf7, 0x73, 0x57, 0xb8, 0x16, 0xaf, 0x6f, 0x0b, 0xdb, 0x84, 0xff,
	0x91, 0x55, 0xa9, 0xf0, 0x72, 0x3c, 0x35, 0xd4, 0x8a, 0x1c, 0xb0, 0x92, 0xd0, 0x62, 0x83, 0x75,
	0x42, 0x90, 0xff, 0x93, 0x27, 0x52, 0xe1, 0x95, 0x84, 0x46, 0x43, 0xbd, 0x91, 0xc7, 0xdc, 0x06,
	0x6d, 0x7b, 0xd4, 0xe2, 0xee, 0x9e, 0x30, 0x59, 0x70, 0x2c, 0x06, 0x25, 0x78, 0x2d, 0xe6, 0x98,
	0x8b, 0xb6, 0xa1, 0xcf, 0x60, 0x2d, 0x99, 0x14, 0xac, 0xf8, 0x96, 0x70, 0x2e, 0x33, 0xaf, 0x93,
	0xe7, 0x52, 0xe1, 0xb5, 0x64, 0x66, 0x44, 0x00, 0x7a, 0x0f, 0xe6, 0x92, 0xe9, 0x87, 0xc2, 0xef,
	0x3e, 0x85, 0x90, 0xe4, 0xa5, 0xc2, 0x73, 0xc9, 0xbc, 0x0b, 0x0b, 0x3a, 0x8e, 0xd8, 0x74, 0x97,
	0xe4, 0x7a, 0x15, 0x02, 0xc4, 0x0d, 0xb2, 0x28, 0x15, 0x9e, 0x4f, 0x46, 0xfc, 0xf2, 0xa0, 0xb7,
	0x60, 0x26, 0x55, 0x17, 0x05, 0xca, 0x21, 0x22, 0xf3, 0x52, 0xe1, 0x99, 0x54, 0x2d, 0x14, 0x28,
	0x47, 0x47, 0x11, 0xfb, 0xf8, 0xca, 0x58, 0x74, 0xf8, 0x4d, 0xb2, 0x20, 0x15, 0xce, 0xa6, 0x1b,
	0xbc, 0x4e, 0x8f, 0x9a, 0xff, 0xee, 0xbe, 0xeb, 0xbe, 0xa2, 0xfa, 0x69, 0x85, 0x9d, 0x5a, 0xec,
	0x0c, 0x8e, 0xc5, 0xcc, 0x7f, 0xb8, 0x05, 0x51, 0xb0, 0x90, 0x48, 0xa8, 0x08, 0x9f, 0x9b, 0x45,
	0xcb, 0x6d, 0xdb, 0xb4, 0x03, 0x6f, 0x91, 0x9c, 0x54, 0x78, 0x36, 0x7e, 0xdb, 0xf4, 0x3b, 0xd0,
	0x41, 0xc4, 0xf4, 0xf7, 0x23, 0x7a, 0x9f, 0xb7, 0xdb, 0x24, 0x2b, 0x15, 0x9e, 0x8a, 0x4f, 0x0f,
	0xc4, 0xe8, 0x0d, 0x98, 0x0e, 0x0d, 0xde, 0x3c, 0x65, 0xdc, 0xdb, 0xa6, 0x2d, 0xa6, 0xbf, 0x84,
	0xe7, 0x67, 0x22, 0xbc, 0x43, 0xe6, 0xa4, 0xc2, 0xd3, 0xa1, 0xb9, 0x57, 0xe5, 0x4f, 0x5f, 0x02,
	0xa4, 0x2f, 0x6b, 0xb6, 0xa8, 0x51, 0xdb, 0xfa, 0xc4, 0x0c, 0x4f, 0x34, 0x19, 0x47, 0x0f, 0x72,
	0xdd, 0xab, 0x5a, 0xae, 0x77, 0x55, 0xcb, 0x6d, 0x72, 0xbf, 0xf5, 0x9a, 0xda, 0x3e, 0x2b, 0xb5,
	0x75, 0x94, 0x8b, 0xbf, 0x7d, 0x1d, 0x9e, 0xc8, 0xcc, 0x8c, 0x56, 0x60, 0xc3, 0x37, 0x77, 0x03,
	0x77, 0x55, 0x9b, 0x0b, 0xc3, 0x5f, 0x32, 0x43, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x56, 0xde,
	0x60, 0x79, 0x08, 0x0a, 0x00, 0x00,
}
