// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.5
// source: steammessages_cloud.steamworkssdk.proto

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

type CCloud_GetUploadServerInfo_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
}

func (x *CCloud_GetUploadServerInfo_Request) Reset() {
	*x = CCloud_GetUploadServerInfo_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_GetUploadServerInfo_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_GetUploadServerInfo_Request) ProtoMessage() {}

func (x *CCloud_GetUploadServerInfo_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_GetUploadServerInfo_Request.ProtoReflect.Descriptor instead.
func (*CCloud_GetUploadServerInfo_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{0}
}

func (x *CCloud_GetUploadServerInfo_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type CCloud_GetUploadServerInfo_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerUrl *string `protobuf:"bytes,1,opt,name=server_url,json=serverUrl" json:"server_url,omitempty"`
}

func (x *CCloud_GetUploadServerInfo_Response) Reset() {
	*x = CCloud_GetUploadServerInfo_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_GetUploadServerInfo_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_GetUploadServerInfo_Response) ProtoMessage() {}

func (x *CCloud_GetUploadServerInfo_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_GetUploadServerInfo_Response.ProtoReflect.Descriptor instead.
func (*CCloud_GetUploadServerInfo_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{1}
}

func (x *CCloud_GetUploadServerInfo_Response) GetServerUrl() string {
	if x != nil && x.ServerUrl != nil {
		return *x.ServerUrl
	}
	return ""
}

type CCloud_GetFileDetails_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ugcid *uint64 `protobuf:"varint,1,opt,name=ugcid" json:"ugcid,omitempty"`
	Appid *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
}

func (x *CCloud_GetFileDetails_Request) Reset() {
	*x = CCloud_GetFileDetails_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_GetFileDetails_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_GetFileDetails_Request) ProtoMessage() {}

func (x *CCloud_GetFileDetails_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_GetFileDetails_Request.ProtoReflect.Descriptor instead.
func (*CCloud_GetFileDetails_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{2}
}

func (x *CCloud_GetFileDetails_Request) GetUgcid() uint64 {
	if x != nil && x.Ugcid != nil {
		return *x.Ugcid
	}
	return 0
}

func (x *CCloud_GetFileDetails_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type CCloud_UserFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid          *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Ugcid          *uint64 `protobuf:"varint,2,opt,name=ugcid" json:"ugcid,omitempty"`
	Filename       *string `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Timestamp      *uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	FileSize       *uint32 `protobuf:"varint,5,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	Url            *string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	SteamidCreator *uint64 `protobuf:"fixed64,7,opt,name=steamid_creator,json=steamidCreator" json:"steamid_creator,omitempty"`
}

func (x *CCloud_UserFile) Reset() {
	*x = CCloud_UserFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_UserFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_UserFile) ProtoMessage() {}

func (x *CCloud_UserFile) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_UserFile.ProtoReflect.Descriptor instead.
func (*CCloud_UserFile) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{3}
}

func (x *CCloud_UserFile) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CCloud_UserFile) GetUgcid() uint64 {
	if x != nil && x.Ugcid != nil {
		return *x.Ugcid
	}
	return 0
}

func (x *CCloud_UserFile) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

func (x *CCloud_UserFile) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *CCloud_UserFile) GetFileSize() uint32 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *CCloud_UserFile) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *CCloud_UserFile) GetSteamidCreator() uint64 {
	if x != nil && x.SteamidCreator != nil {
		return *x.SteamidCreator
	}
	return 0
}

type CCloud_GetFileDetails_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details *CCloud_UserFile `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
}

func (x *CCloud_GetFileDetails_Response) Reset() {
	*x = CCloud_GetFileDetails_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_GetFileDetails_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_GetFileDetails_Response) ProtoMessage() {}

func (x *CCloud_GetFileDetails_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_GetFileDetails_Response.ProtoReflect.Descriptor instead.
func (*CCloud_GetFileDetails_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{4}
}

func (x *CCloud_GetFileDetails_Response) GetDetails() *CCloud_UserFile {
	if x != nil {
		return x.Details
	}
	return nil
}

type CCloud_EnumerateUserFiles_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid           *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	ExtendedDetails *bool   `protobuf:"varint,2,opt,name=extended_details,json=extendedDetails" json:"extended_details,omitempty"`
	Count           *uint32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	StartIndex      *uint32 `protobuf:"varint,4,opt,name=start_index,json=startIndex" json:"start_index,omitempty"`
}

func (x *CCloud_EnumerateUserFiles_Request) Reset() {
	*x = CCloud_EnumerateUserFiles_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_EnumerateUserFiles_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_EnumerateUserFiles_Request) ProtoMessage() {}

func (x *CCloud_EnumerateUserFiles_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_EnumerateUserFiles_Request.ProtoReflect.Descriptor instead.
func (*CCloud_EnumerateUserFiles_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{5}
}

func (x *CCloud_EnumerateUserFiles_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CCloud_EnumerateUserFiles_Request) GetExtendedDetails() bool {
	if x != nil && x.ExtendedDetails != nil {
		return *x.ExtendedDetails
	}
	return false
}

func (x *CCloud_EnumerateUserFiles_Request) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *CCloud_EnumerateUserFiles_Request) GetStartIndex() uint32 {
	if x != nil && x.StartIndex != nil {
		return *x.StartIndex
	}
	return 0
}

type CCloud_EnumerateUserFiles_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files      []*CCloud_UserFile `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	TotalFiles *uint32            `protobuf:"varint,2,opt,name=total_files,json=totalFiles" json:"total_files,omitempty"`
}

func (x *CCloud_EnumerateUserFiles_Response) Reset() {
	*x = CCloud_EnumerateUserFiles_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_EnumerateUserFiles_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_EnumerateUserFiles_Response) ProtoMessage() {}

func (x *CCloud_EnumerateUserFiles_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_EnumerateUserFiles_Response.ProtoReflect.Descriptor instead.
func (*CCloud_EnumerateUserFiles_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{6}
}

func (x *CCloud_EnumerateUserFiles_Response) GetFiles() []*CCloud_UserFile {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *CCloud_EnumerateUserFiles_Response) GetTotalFiles() uint32 {
	if x != nil && x.TotalFiles != nil {
		return *x.TotalFiles
	}
	return 0
}

type CCloud_Delete_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid    *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
}

func (x *CCloud_Delete_Request) Reset() {
	*x = CCloud_Delete_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_Delete_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_Delete_Request) ProtoMessage() {}

func (x *CCloud_Delete_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_Delete_Request.ProtoReflect.Descriptor instead.
func (*CCloud_Delete_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{7}
}

func (x *CCloud_Delete_Request) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

func (x *CCloud_Delete_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type CCloud_Delete_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CCloud_Delete_Response) Reset() {
	*x = CCloud_Delete_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCloud_Delete_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCloud_Delete_Response) ProtoMessage() {}

func (x *CCloud_Delete_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_cloud_steamworkssdk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCloud_Delete_Response.ProtoReflect.Descriptor instead.
func (*CCloud_Delete_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP(), []int{8}
}

var File_steammessages_cloud_steamworkssdk_proto protoreflect.FileDescriptor

var file_steammessages_cloud_steamworkssdk_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x74, 0x61, 0x1a,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x22, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x23, 0x43,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x22, 0x4b, 0x0a, 0x1d, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x67, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x75, 0x67, 0x63, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0xcf,
	0x01, 0x0a, 0x0f, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x67, 0x63, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x75, 0x67, 0x63, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x69, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x0e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x51, 0x0a, 0x1e, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x21, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x45,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x72, 0x0a, 0x22, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x45, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x96, 0x05, 0x0a, 0x05, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x12, 0xa6, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x2e, 0x64,
	0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3a, 0x82, 0xb5, 0x18, 0x36, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x12, 0x81, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x23, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xb5, 0x18,
	0x20, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x20, 0x6f, 0x6e, 0x20, 0x61, 0x20, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x12, 0xc4, 0x01, 0x0a, 0x12, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e,
	0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xb5, 0x18,
	0x57, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x20, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x61, 0x70,
	0x70, 0x20, 0x49, 0x44, 0x2e, 0x20, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x75, 0x70,
	0x20, 0x74, 0x6f, 0x20, 0x35, 0x30, 0x30, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x20, 0x61, 0x74,
	0x20, 0x61, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x12, 0x6e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82,
	0xb5, 0x18, 0x25, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x66, 0x69, 0x6c,
	0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27,
	0x73, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x1a, 0x29, 0x82, 0xb5, 0x18, 0x25, 0x41, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x20, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61,
	0x2f, 0x64, 0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_steammessages_cloud_steamworkssdk_proto_rawDescOnce sync.Once
	file_steammessages_cloud_steamworkssdk_proto_rawDescData = file_steammessages_cloud_steamworkssdk_proto_rawDesc
)

func file_steammessages_cloud_steamworkssdk_proto_rawDescGZIP() []byte {
	file_steammessages_cloud_steamworkssdk_proto_rawDescOnce.Do(func() {
		file_steammessages_cloud_steamworkssdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_cloud_steamworkssdk_proto_rawDescData)
	})
	return file_steammessages_cloud_steamworkssdk_proto_rawDescData
}

var file_steammessages_cloud_steamworkssdk_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_steammessages_cloud_steamworkssdk_proto_goTypes = []interface{}{
	(*CCloud_GetUploadServerInfo_Request)(nil),  // 0: dota.CCloud_GetUploadServerInfo_Request
	(*CCloud_GetUploadServerInfo_Response)(nil), // 1: dota.CCloud_GetUploadServerInfo_Response
	(*CCloud_GetFileDetails_Request)(nil),       // 2: dota.CCloud_GetFileDetails_Request
	(*CCloud_UserFile)(nil),                     // 3: dota.CCloud_UserFile
	(*CCloud_GetFileDetails_Response)(nil),      // 4: dota.CCloud_GetFileDetails_Response
	(*CCloud_EnumerateUserFiles_Request)(nil),   // 5: dota.CCloud_EnumerateUserFiles_Request
	(*CCloud_EnumerateUserFiles_Response)(nil),  // 6: dota.CCloud_EnumerateUserFiles_Response
	(*CCloud_Delete_Request)(nil),               // 7: dota.CCloud_Delete_Request
	(*CCloud_Delete_Response)(nil),              // 8: dota.CCloud_Delete_Response
}
var file_steammessages_cloud_steamworkssdk_proto_depIdxs = []int32{
	3, // 0: dota.CCloud_GetFileDetails_Response.details:type_name -> dota.CCloud_UserFile
	3, // 1: dota.CCloud_EnumerateUserFiles_Response.files:type_name -> dota.CCloud_UserFile
	0, // 2: dota.Cloud.GetUploadServerInfo:input_type -> dota.CCloud_GetUploadServerInfo_Request
	2, // 3: dota.Cloud.GetFileDetails:input_type -> dota.CCloud_GetFileDetails_Request
	5, // 4: dota.Cloud.EnumerateUserFiles:input_type -> dota.CCloud_EnumerateUserFiles_Request
	7, // 5: dota.Cloud.Delete:input_type -> dota.CCloud_Delete_Request
	1, // 6: dota.Cloud.GetUploadServerInfo:output_type -> dota.CCloud_GetUploadServerInfo_Response
	4, // 7: dota.Cloud.GetFileDetails:output_type -> dota.CCloud_GetFileDetails_Response
	6, // 8: dota.Cloud.EnumerateUserFiles:output_type -> dota.CCloud_EnumerateUserFiles_Response
	8, // 9: dota.Cloud.Delete:output_type -> dota.CCloud_Delete_Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_steammessages_cloud_steamworkssdk_proto_init() }
func file_steammessages_cloud_steamworkssdk_proto_init() {
	if File_steammessages_cloud_steamworkssdk_proto != nil {
		return
	}
	file_steammessages_unified_base_steamworkssdk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_GetUploadServerInfo_Request); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_GetUploadServerInfo_Response); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_GetFileDetails_Request); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_UserFile); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_GetFileDetails_Response); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_EnumerateUserFiles_Request); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_EnumerateUserFiles_Response); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_Delete_Request); i {
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
		file_steammessages_cloud_steamworkssdk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCloud_Delete_Response); i {
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
			RawDescriptor: file_steammessages_cloud_steamworkssdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steammessages_cloud_steamworkssdk_proto_goTypes,
		DependencyIndexes: file_steammessages_cloud_steamworkssdk_proto_depIdxs,
		MessageInfos:      file_steammessages_cloud_steamworkssdk_proto_msgTypes,
	}.Build()
	File_steammessages_cloud_steamworkssdk_proto = out.File
	file_steammessages_cloud_steamworkssdk_proto_rawDesc = nil
	file_steammessages_cloud_steamworkssdk_proto_goTypes = nil
	file_steammessages_cloud_steamworkssdk_proto_depIdxs = nil
}
