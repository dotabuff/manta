// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: connectionless_netmessages.proto

package dota

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type C2S_CONNECT_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostVersion       *uint32                       `protobuf:"varint,1,opt,name=host_version,json=hostVersion" json:"host_version,omitempty"`
	AuthProtocol      *uint32                       `protobuf:"varint,2,opt,name=auth_protocol,json=authProtocol" json:"auth_protocol,omitempty"`
	ChallengeNumber   *uint32                       `protobuf:"varint,3,opt,name=challenge_number,json=challengeNumber" json:"challenge_number,omitempty"`
	ReservationCookie *uint64                       `protobuf:"fixed64,4,opt,name=reservation_cookie,json=reservationCookie" json:"reservation_cookie,omitempty"`
	LowViolence       *bool                         `protobuf:"varint,5,opt,name=low_violence,json=lowViolence" json:"low_violence,omitempty"`
	EncryptedPassword []byte                        `protobuf:"bytes,6,opt,name=encrypted_password,json=encryptedPassword" json:"encrypted_password,omitempty"`
	Splitplayers      []*CCLCMsg_SplitPlayerConnect `protobuf:"bytes,7,rep,name=splitplayers" json:"splitplayers,omitempty"`
	AuthSteam         []byte                        `protobuf:"bytes,8,opt,name=auth_steam,json=authSteam" json:"auth_steam,omitempty"`
	ChallengeContext  *string                       `protobuf:"bytes,9,opt,name=challenge_context,json=challengeContext" json:"challenge_context,omitempty"`
	UseSnp            *int32                        `protobuf:"zigzag32,10,opt,name=use_snp,json=useSnp" json:"use_snp,omitempty"`
}

func (x *C2S_CONNECT_Message) Reset() {
	*x = C2S_CONNECT_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectionless_netmessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CONNECT_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CONNECT_Message) ProtoMessage() {}

func (x *C2S_CONNECT_Message) ProtoReflect() protoreflect.Message {
	mi := &file_connectionless_netmessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CONNECT_Message.ProtoReflect.Descriptor instead.
func (*C2S_CONNECT_Message) Descriptor() ([]byte, []int) {
	return file_connectionless_netmessages_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_CONNECT_Message) GetHostVersion() uint32 {
	if x != nil && x.HostVersion != nil {
		return *x.HostVersion
	}
	return 0
}

func (x *C2S_CONNECT_Message) GetAuthProtocol() uint32 {
	if x != nil && x.AuthProtocol != nil {
		return *x.AuthProtocol
	}
	return 0
}

func (x *C2S_CONNECT_Message) GetChallengeNumber() uint32 {
	if x != nil && x.ChallengeNumber != nil {
		return *x.ChallengeNumber
	}
	return 0
}

func (x *C2S_CONNECT_Message) GetReservationCookie() uint64 {
	if x != nil && x.ReservationCookie != nil {
		return *x.ReservationCookie
	}
	return 0
}

func (x *C2S_CONNECT_Message) GetLowViolence() bool {
	if x != nil && x.LowViolence != nil {
		return *x.LowViolence
	}
	return false
}

func (x *C2S_CONNECT_Message) GetEncryptedPassword() []byte {
	if x != nil {
		return x.EncryptedPassword
	}
	return nil
}

func (x *C2S_CONNECT_Message) GetSplitplayers() []*CCLCMsg_SplitPlayerConnect {
	if x != nil {
		return x.Splitplayers
	}
	return nil
}

func (x *C2S_CONNECT_Message) GetAuthSteam() []byte {
	if x != nil {
		return x.AuthSteam
	}
	return nil
}

func (x *C2S_CONNECT_Message) GetChallengeContext() string {
	if x != nil && x.ChallengeContext != nil {
		return *x.ChallengeContext
	}
	return ""
}

func (x *C2S_CONNECT_Message) GetUseSnp() int32 {
	if x != nil && x.UseSnp != nil {
		return *x.UseSnp
	}
	return 0
}

type C2S_CONNECTION_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddonName *string `protobuf:"bytes,1,opt,name=addon_name,json=addonName" json:"addon_name,omitempty"`
	UseSnp    *bool   `protobuf:"varint,2,opt,name=use_snp,json=useSnp" json:"use_snp,omitempty"`
}

func (x *C2S_CONNECTION_Message) Reset() {
	*x = C2S_CONNECTION_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectionless_netmessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CONNECTION_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CONNECTION_Message) ProtoMessage() {}

func (x *C2S_CONNECTION_Message) ProtoReflect() protoreflect.Message {
	mi := &file_connectionless_netmessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CONNECTION_Message.ProtoReflect.Descriptor instead.
func (*C2S_CONNECTION_Message) Descriptor() ([]byte, []int) {
	return file_connectionless_netmessages_proto_rawDescGZIP(), []int{1}
}

func (x *C2S_CONNECTION_Message) GetAddonName() string {
	if x != nil && x.AddonName != nil {
		return *x.AddonName
	}
	return ""
}

func (x *C2S_CONNECTION_Message) GetUseSnp() bool {
	if x != nil && x.UseSnp != nil {
		return *x.UseSnp
	}
	return false
}

var File_connectionless_netmessages_proto protoreflect.FileDescriptor

var file_connectionless_netmessages_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6c, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x65, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x74, 0x61, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x03, 0x0a, 0x13,
	0x43, 0x32, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x77, 0x5f, 0x76, 0x69, 0x6f,
	0x6c, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6c, 0x6f, 0x77,
	0x56, 0x69, 0x6f, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x4c, 0x43, 0x4d, 0x73, 0x67, 0x5f, 0x53, 0x70, 0x6c,
	0x69, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x0c, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x0a, 0x11,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x5f, 0x73, 0x6e, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x75, 0x73, 0x65, 0x53,
	0x6e, 0x70, 0x22, 0x50, 0x0a, 0x16, 0x43, 0x32, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x5f, 0x73, 0x6e, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x53, 0x6e, 0x70, 0x42, 0x28, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x74,
	0x61, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61, 0x80, 0x01, 0x00,
}

var (
	file_connectionless_netmessages_proto_rawDescOnce sync.Once
	file_connectionless_netmessages_proto_rawDescData = file_connectionless_netmessages_proto_rawDesc
)

func file_connectionless_netmessages_proto_rawDescGZIP() []byte {
	file_connectionless_netmessages_proto_rawDescOnce.Do(func() {
		file_connectionless_netmessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectionless_netmessages_proto_rawDescData)
	})
	return file_connectionless_netmessages_proto_rawDescData
}

var file_connectionless_netmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_connectionless_netmessages_proto_goTypes = []interface{}{
	(*C2S_CONNECT_Message)(nil),        // 0: dota.C2S_CONNECT_Message
	(*C2S_CONNECTION_Message)(nil),     // 1: dota.C2S_CONNECTION_Message
	(*CCLCMsg_SplitPlayerConnect)(nil), // 2: dota.CCLCMsg_SplitPlayerConnect
}
var file_connectionless_netmessages_proto_depIdxs = []int32{
	2, // 0: dota.C2S_CONNECT_Message.splitplayers:type_name -> dota.CCLCMsg_SplitPlayerConnect
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_connectionless_netmessages_proto_init() }
func file_connectionless_netmessages_proto_init() {
	if File_connectionless_netmessages_proto != nil {
		return
	}
	file_netmessages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_connectionless_netmessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CONNECT_Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_connectionless_netmessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CONNECTION_Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connectionless_netmessages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectionless_netmessages_proto_goTypes,
		DependencyIndexes: file_connectionless_netmessages_proto_depIdxs,
		MessageInfos:      file_connectionless_netmessages_proto_msgTypes,
	}.Build()
	File_connectionless_netmessages_proto = out.File
	file_connectionless_netmessages_proto_rawDesc = nil
	file_connectionless_netmessages_proto_goTypes = nil
	file_connectionless_netmessages_proto_depIdxs = nil
}
