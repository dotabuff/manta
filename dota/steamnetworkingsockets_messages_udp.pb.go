// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steamnetworkingsockets_messages_udp.proto

package dota

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ESteamNetworkingUDPMsgID int32

const (
	ESteamNetworkingUDPMsgID_k_ESteamNetworkingUDPMsg_ChallengeRequest ESteamNetworkingUDPMsgID = 32
	ESteamNetworkingUDPMsgID_k_ESteamNetworkingUDPMsg_ChallengeReply   ESteamNetworkingUDPMsgID = 33
	ESteamNetworkingUDPMsgID_k_ESteamNetworkingUDPMsg_ConnectRequest   ESteamNetworkingUDPMsgID = 34
	ESteamNetworkingUDPMsgID_k_ESteamNetworkingUDPMsg_ConnectOK        ESteamNetworkingUDPMsgID = 35
	ESteamNetworkingUDPMsgID_k_ESteamNetworkingUDPMsg_ConnectionClosed ESteamNetworkingUDPMsgID = 36
	ESteamNetworkingUDPMsgID_k_ESteamNetworkingUDPMsg_NoConnection     ESteamNetworkingUDPMsgID = 37
)

var ESteamNetworkingUDPMsgID_name = map[int32]string{
	32: "k_ESteamNetworkingUDPMsg_ChallengeRequest",
	33: "k_ESteamNetworkingUDPMsg_ChallengeReply",
	34: "k_ESteamNetworkingUDPMsg_ConnectRequest",
	35: "k_ESteamNetworkingUDPMsg_ConnectOK",
	36: "k_ESteamNetworkingUDPMsg_ConnectionClosed",
	37: "k_ESteamNetworkingUDPMsg_NoConnection",
}

var ESteamNetworkingUDPMsgID_value = map[string]int32{
	"k_ESteamNetworkingUDPMsg_ChallengeRequest": 32,
	"k_ESteamNetworkingUDPMsg_ChallengeReply":   33,
	"k_ESteamNetworkingUDPMsg_ConnectRequest":   34,
	"k_ESteamNetworkingUDPMsg_ConnectOK":        35,
	"k_ESteamNetworkingUDPMsg_ConnectionClosed": 36,
	"k_ESteamNetworkingUDPMsg_NoConnection":     37,
}

func (x ESteamNetworkingUDPMsgID) Enum() *ESteamNetworkingUDPMsgID {
	p := new(ESteamNetworkingUDPMsgID)
	*p = x
	return p
}

func (x ESteamNetworkingUDPMsgID) String() string {
	return proto.EnumName(ESteamNetworkingUDPMsgID_name, int32(x))
}

func (x *ESteamNetworkingUDPMsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ESteamNetworkingUDPMsgID_value, data, "ESteamNetworkingUDPMsgID")
	if err != nil {
		return err
	}
	*x = ESteamNetworkingUDPMsgID(value)
	return nil
}

func (ESteamNetworkingUDPMsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{0}
}

type CMsgSteamSockets_UDP_Stats_Flags int32

const (
	CMsgSteamSockets_UDP_Stats_ACK_REQUEST_E2E           CMsgSteamSockets_UDP_Stats_Flags = 2
	CMsgSteamSockets_UDP_Stats_ACK_REQUEST_IMMEDIATE     CMsgSteamSockets_UDP_Stats_Flags = 4
	CMsgSteamSockets_UDP_Stats_NOT_PRIMARY_TRANSPORT_E2E CMsgSteamSockets_UDP_Stats_Flags = 16
)

var CMsgSteamSockets_UDP_Stats_Flags_name = map[int32]string{
	2:  "ACK_REQUEST_E2E",
	4:  "ACK_REQUEST_IMMEDIATE",
	16: "NOT_PRIMARY_TRANSPORT_E2E",
}

var CMsgSteamSockets_UDP_Stats_Flags_value = map[string]int32{
	"ACK_REQUEST_E2E":           2,
	"ACK_REQUEST_IMMEDIATE":     4,
	"NOT_PRIMARY_TRANSPORT_E2E": 16,
}

func (x CMsgSteamSockets_UDP_Stats_Flags) Enum() *CMsgSteamSockets_UDP_Stats_Flags {
	p := new(CMsgSteamSockets_UDP_Stats_Flags)
	*p = x
	return p
}

func (x CMsgSteamSockets_UDP_Stats_Flags) String() string {
	return proto.EnumName(CMsgSteamSockets_UDP_Stats_Flags_name, int32(x))
}

func (x *CMsgSteamSockets_UDP_Stats_Flags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMsgSteamSockets_UDP_Stats_Flags_value, data, "CMsgSteamSockets_UDP_Stats_Flags")
	if err != nil {
		return err
	}
	*x = CMsgSteamSockets_UDP_Stats_Flags(value)
	return nil
}

func (CMsgSteamSockets_UDP_Stats_Flags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{6, 0}
}

type CMsgSteamSockets_UDP_ChallengeRequest struct {
	ConnectionId         *uint32  `protobuf:"fixed32,1,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
	MyTimestamp          *uint64  `protobuf:"fixed64,3,opt,name=my_timestamp,json=myTimestamp" json:"my_timestamp,omitempty"`
	ProtocolVersion      *uint32  `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamSockets_UDP_ChallengeRequest) Reset()         { *m = CMsgSteamSockets_UDP_ChallengeRequest{} }
func (m *CMsgSteamSockets_UDP_ChallengeRequest) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_ChallengeRequest) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_ChallengeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{0}
}

func (m *CMsgSteamSockets_UDP_ChallengeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeRequest.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_ChallengeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeRequest.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_ChallengeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeRequest.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_ChallengeRequest) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeRequest.Size(m)
}
func (m *CMsgSteamSockets_UDP_ChallengeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeRequest proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_ChallengeRequest) GetConnectionId() uint32 {
	if m != nil && m.ConnectionId != nil {
		return *m.ConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ChallengeRequest) GetMyTimestamp() uint64 {
	if m != nil && m.MyTimestamp != nil {
		return *m.MyTimestamp
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ChallengeRequest) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

type CMsgSteamSockets_UDP_ChallengeReply struct {
	ConnectionId         *uint32  `protobuf:"fixed32,1,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
	Challenge            *uint64  `protobuf:"fixed64,2,opt,name=challenge" json:"challenge,omitempty"`
	YourTimestamp        *uint64  `protobuf:"fixed64,3,opt,name=your_timestamp,json=yourTimestamp" json:"your_timestamp,omitempty"`
	ProtocolVersion      *uint32  `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamSockets_UDP_ChallengeReply) Reset()         { *m = CMsgSteamSockets_UDP_ChallengeReply{} }
func (m *CMsgSteamSockets_UDP_ChallengeReply) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_ChallengeReply) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_ChallengeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{1}
}

func (m *CMsgSteamSockets_UDP_ChallengeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeReply.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_ChallengeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeReply.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_ChallengeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeReply.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_ChallengeReply) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeReply.Size(m)
}
func (m *CMsgSteamSockets_UDP_ChallengeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeReply.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_ChallengeReply proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_ChallengeReply) GetConnectionId() uint32 {
	if m != nil && m.ConnectionId != nil {
		return *m.ConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ChallengeReply) GetChallenge() uint64 {
	if m != nil && m.Challenge != nil {
		return *m.Challenge
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ChallengeReply) GetYourTimestamp() uint64 {
	if m != nil && m.YourTimestamp != nil {
		return *m.YourTimestamp
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ChallengeReply) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

type CMsgSteamSockets_UDP_ConnectRequest struct {
	ClientConnectionId    *uint32                                  `protobuf:"fixed32,1,opt,name=client_connection_id,json=clientConnectionId" json:"client_connection_id,omitempty"`
	Challenge             *uint64                                  `protobuf:"fixed64,2,opt,name=challenge" json:"challenge,omitempty"`
	MyTimestamp           *uint64                                  `protobuf:"fixed64,5,opt,name=my_timestamp,json=myTimestamp" json:"my_timestamp,omitempty"`
	PingEstMs             *uint32                                  `protobuf:"varint,6,opt,name=ping_est_ms,json=pingEstMs" json:"ping_est_ms,omitempty"`
	Crypt                 *CMsgSteamDatagramSessionCryptInfoSigned `protobuf:"bytes,7,opt,name=crypt" json:"crypt,omitempty"`
	Cert                  *CMsgSteamDatagramCertificateSigned      `protobuf:"bytes,4,opt,name=cert" json:"cert,omitempty"`
	LegacyProtocolVersion *uint32                                  `protobuf:"varint,8,opt,name=legacy_protocol_version,json=legacyProtocolVersion" json:"legacy_protocol_version,omitempty"`
	IdentityString        *string                                  `protobuf:"bytes,10,opt,name=identity_string,json=identityString" json:"identity_string,omitempty"`
	LegacyClientSteamId   *uint64                                  `protobuf:"fixed64,3,opt,name=legacy_client_steam_id,json=legacyClientSteamId" json:"legacy_client_steam_id,omitempty"`
	LegacyIdentityBinary  *CMsgSteamNetworkingIdentityLegacyBinary `protobuf:"bytes,9,opt,name=legacy_identity_binary,json=legacyIdentityBinary" json:"legacy_identity_binary,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                 `json:"-"`
	XXX_unrecognized      []byte                                   `json:"-"`
	XXX_sizecache         int32                                    `json:"-"`
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) Reset()         { *m = CMsgSteamSockets_UDP_ConnectRequest{} }
func (m *CMsgSteamSockets_UDP_ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_ConnectRequest) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{2}
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectRequest.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_ConnectRequest.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectRequest.Size(m)
}
func (m *CMsgSteamSockets_UDP_ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_ConnectRequest proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetClientConnectionId() uint32 {
	if m != nil && m.ClientConnectionId != nil {
		return *m.ClientConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetChallenge() uint64 {
	if m != nil && m.Challenge != nil {
		return *m.Challenge
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetMyTimestamp() uint64 {
	if m != nil && m.MyTimestamp != nil {
		return *m.MyTimestamp
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetPingEstMs() uint32 {
	if m != nil && m.PingEstMs != nil {
		return *m.PingEstMs
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetCrypt() *CMsgSteamDatagramSessionCryptInfoSigned {
	if m != nil {
		return m.Crypt
	}
	return nil
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetCert() *CMsgSteamDatagramCertificateSigned {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetLegacyProtocolVersion() uint32 {
	if m != nil && m.LegacyProtocolVersion != nil {
		return *m.LegacyProtocolVersion
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetIdentityString() string {
	if m != nil && m.IdentityString != nil {
		return *m.IdentityString
	}
	return ""
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetLegacyClientSteamId() uint64 {
	if m != nil && m.LegacyClientSteamId != nil {
		return *m.LegacyClientSteamId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectRequest) GetLegacyIdentityBinary() *CMsgSteamNetworkingIdentityLegacyBinary {
	if m != nil {
		return m.LegacyIdentityBinary
	}
	return nil
}

type CMsgSteamSockets_UDP_ConnectOK struct {
	ClientConnectionId   *uint32                                  `protobuf:"fixed32,1,opt,name=client_connection_id,json=clientConnectionId" json:"client_connection_id,omitempty"`
	ServerConnectionId   *uint32                                  `protobuf:"fixed32,5,opt,name=server_connection_id,json=serverConnectionId" json:"server_connection_id,omitempty"`
	YourTimestamp        *uint64                                  `protobuf:"fixed64,3,opt,name=your_timestamp,json=yourTimestamp" json:"your_timestamp,omitempty"`
	DelayTimeUsec        *uint32                                  `protobuf:"varint,4,opt,name=delay_time_usec,json=delayTimeUsec" json:"delay_time_usec,omitempty"`
	Crypt                *CMsgSteamDatagramSessionCryptInfoSigned `protobuf:"bytes,7,opt,name=crypt" json:"crypt,omitempty"`
	Cert                 *CMsgSteamDatagramCertificateSigned      `protobuf:"bytes,8,opt,name=cert" json:"cert,omitempty"`
	IdentityString       *string                                  `protobuf:"bytes,11,opt,name=identity_string,json=identityString" json:"identity_string,omitempty"`
	LegacyServerSteamId  *uint64                                  `protobuf:"fixed64,2,opt,name=legacy_server_steam_id,json=legacyServerSteamId" json:"legacy_server_steam_id,omitempty"`
	LegacyIdentityBinary *CMsgSteamNetworkingIdentityLegacyBinary `protobuf:"bytes,10,opt,name=legacy_identity_binary,json=legacyIdentityBinary" json:"legacy_identity_binary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *CMsgSteamSockets_UDP_ConnectOK) Reset()         { *m = CMsgSteamSockets_UDP_ConnectOK{} }
func (m *CMsgSteamSockets_UDP_ConnectOK) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_ConnectOK) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_ConnectOK) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{3}
}

func (m *CMsgSteamSockets_UDP_ConnectOK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectOK.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_ConnectOK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectOK.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_ConnectOK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_ConnectOK.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_ConnectOK) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectOK.Size(m)
}
func (m *CMsgSteamSockets_UDP_ConnectOK) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_ConnectOK.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_ConnectOK proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_ConnectOK) GetClientConnectionId() uint32 {
	if m != nil && m.ClientConnectionId != nil {
		return *m.ClientConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetServerConnectionId() uint32 {
	if m != nil && m.ServerConnectionId != nil {
		return *m.ServerConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetYourTimestamp() uint64 {
	if m != nil && m.YourTimestamp != nil {
		return *m.YourTimestamp
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetDelayTimeUsec() uint32 {
	if m != nil && m.DelayTimeUsec != nil {
		return *m.DelayTimeUsec
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetCrypt() *CMsgSteamDatagramSessionCryptInfoSigned {
	if m != nil {
		return m.Crypt
	}
	return nil
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetCert() *CMsgSteamDatagramCertificateSigned {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetIdentityString() string {
	if m != nil && m.IdentityString != nil {
		return *m.IdentityString
	}
	return ""
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetLegacyServerSteamId() uint64 {
	if m != nil && m.LegacyServerSteamId != nil {
		return *m.LegacyServerSteamId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectOK) GetLegacyIdentityBinary() *CMsgSteamNetworkingIdentityLegacyBinary {
	if m != nil {
		return m.LegacyIdentityBinary
	}
	return nil
}

type CMsgSteamSockets_UDP_ConnectionClosed struct {
	ToConnectionId       *uint32  `protobuf:"fixed32,4,opt,name=to_connection_id,json=toConnectionId" json:"to_connection_id,omitempty"`
	FromConnectionId     *uint32  `protobuf:"fixed32,5,opt,name=from_connection_id,json=fromConnectionId" json:"from_connection_id,omitempty"`
	Debug                *string  `protobuf:"bytes,2,opt,name=debug" json:"debug,omitempty"`
	ReasonCode           *uint32  `protobuf:"varint,3,opt,name=reason_code,json=reasonCode" json:"reason_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamSockets_UDP_ConnectionClosed) Reset()         { *m = CMsgSteamSockets_UDP_ConnectionClosed{} }
func (m *CMsgSteamSockets_UDP_ConnectionClosed) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_ConnectionClosed) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_ConnectionClosed) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{4}
}

func (m *CMsgSteamSockets_UDP_ConnectionClosed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectionClosed.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_ConnectionClosed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectionClosed.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_ConnectionClosed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_ConnectionClosed.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_ConnectionClosed) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_ConnectionClosed.Size(m)
}
func (m *CMsgSteamSockets_UDP_ConnectionClosed) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_ConnectionClosed.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_ConnectionClosed proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_ConnectionClosed) GetToConnectionId() uint32 {
	if m != nil && m.ToConnectionId != nil {
		return *m.ToConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectionClosed) GetFromConnectionId() uint32 {
	if m != nil && m.FromConnectionId != nil {
		return *m.FromConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_ConnectionClosed) GetDebug() string {
	if m != nil && m.Debug != nil {
		return *m.Debug
	}
	return ""
}

func (m *CMsgSteamSockets_UDP_ConnectionClosed) GetReasonCode() uint32 {
	if m != nil && m.ReasonCode != nil {
		return *m.ReasonCode
	}
	return 0
}

type CMsgSteamSockets_UDP_NoConnection struct {
	FromConnectionId     *uint32  `protobuf:"fixed32,2,opt,name=from_connection_id,json=fromConnectionId" json:"from_connection_id,omitempty"`
	ToConnectionId       *uint32  `protobuf:"fixed32,3,opt,name=to_connection_id,json=toConnectionId" json:"to_connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamSockets_UDP_NoConnection) Reset()         { *m = CMsgSteamSockets_UDP_NoConnection{} }
func (m *CMsgSteamSockets_UDP_NoConnection) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_NoConnection) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_NoConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{5}
}

func (m *CMsgSteamSockets_UDP_NoConnection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_NoConnection.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_NoConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_NoConnection.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_NoConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_NoConnection.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_NoConnection) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_NoConnection.Size(m)
}
func (m *CMsgSteamSockets_UDP_NoConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_NoConnection.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_NoConnection proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_NoConnection) GetFromConnectionId() uint32 {
	if m != nil && m.FromConnectionId != nil {
		return *m.FromConnectionId
	}
	return 0
}

func (m *CMsgSteamSockets_UDP_NoConnection) GetToConnectionId() uint32 {
	if m != nil && m.ToConnectionId != nil {
		return *m.ToConnectionId
	}
	return 0
}

type CMsgSteamSockets_UDP_Stats struct {
	Stats                *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
	Flags                *uint32                             `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *CMsgSteamSockets_UDP_Stats) Reset()         { *m = CMsgSteamSockets_UDP_Stats{} }
func (m *CMsgSteamSockets_UDP_Stats) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamSockets_UDP_Stats) ProtoMessage()    {}
func (*CMsgSteamSockets_UDP_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_436d73538b11cdb4, []int{6}
}

func (m *CMsgSteamSockets_UDP_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamSockets_UDP_Stats.Unmarshal(m, b)
}
func (m *CMsgSteamSockets_UDP_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamSockets_UDP_Stats.Marshal(b, m, deterministic)
}
func (m *CMsgSteamSockets_UDP_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamSockets_UDP_Stats.Merge(m, src)
}
func (m *CMsgSteamSockets_UDP_Stats) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamSockets_UDP_Stats.Size(m)
}
func (m *CMsgSteamSockets_UDP_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamSockets_UDP_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamSockets_UDP_Stats proto.InternalMessageInfo

func (m *CMsgSteamSockets_UDP_Stats) GetStats() *CMsgSteamDatagramConnectionQuality {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *CMsgSteamSockets_UDP_Stats) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func init() {
	proto.RegisterEnum("dota.ESteamNetworkingUDPMsgID", ESteamNetworkingUDPMsgID_name, ESteamNetworkingUDPMsgID_value)
	proto.RegisterEnum("dota.CMsgSteamSockets_UDP_Stats_Flags", CMsgSteamSockets_UDP_Stats_Flags_name, CMsgSteamSockets_UDP_Stats_Flags_value)
	proto.RegisterType((*CMsgSteamSockets_UDP_ChallengeRequest)(nil), "dota.CMsgSteamSockets_UDP_ChallengeRequest")
	proto.RegisterType((*CMsgSteamSockets_UDP_ChallengeReply)(nil), "dota.CMsgSteamSockets_UDP_ChallengeReply")
	proto.RegisterType((*CMsgSteamSockets_UDP_ConnectRequest)(nil), "dota.CMsgSteamSockets_UDP_ConnectRequest")
	proto.RegisterType((*CMsgSteamSockets_UDP_ConnectOK)(nil), "dota.CMsgSteamSockets_UDP_ConnectOK")
	proto.RegisterType((*CMsgSteamSockets_UDP_ConnectionClosed)(nil), "dota.CMsgSteamSockets_UDP_ConnectionClosed")
	proto.RegisterType((*CMsgSteamSockets_UDP_NoConnection)(nil), "dota.CMsgSteamSockets_UDP_NoConnection")
	proto.RegisterType((*CMsgSteamSockets_UDP_Stats)(nil), "dota.CMsgSteamSockets_UDP_Stats")
}

func init() {
	proto.RegisterFile("steamnetworkingsockets_messages_udp.proto", fileDescriptor_436d73538b11cdb4)
}

var fileDescriptor_436d73538b11cdb4 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x9d, 0x5c, 0xbb, 0xad, 0xaf, 0xeb, 0x44, 0x60, 0xd3, 0x4d, 0x0d, 0xb6, 0xce, 0x51, 0xe6,
	0xd6, 0x59, 0xd7, 0xa0, 0xc8, 0x80, 0x3d, 0x0d, 0x03, 0x52, 0x59, 0xc3, 0x84, 0xcc, 0x89, 0x2b,
	0x39, 0x05, 0xf6, 0x44, 0xb0, 0x12, 0xa3, 0x09, 0x91, 0x44, 0x4f, 0xa4, 0x3b, 0x08, 0x7b, 0xd9,
	0xdf, 0xd8, 0xde, 0xf7, 0x17, 0xfa, 0xb0, 0x9f, 0xb0, 0x5f, 0x35, 0x88, 0x94, 0xdd, 0xd8, 0x96,
	0x3f, 0x10, 0xf4, 0xcd, 0x3a, 0xf7, 0x1c, 0xf1, 0xdc, 0x7b, 0x0f, 0x2d, 0x38, 0xe2, 0x82, 0x92,
	0x24, 0xa5, 0xe2, 0x77, 0x96, 0x5d, 0x47, 0x69, 0xc8, 0x99, 0x7f, 0x4d, 0x05, 0xc7, 0x09, 0xe5,
	0x9c, 0x84, 0x94, 0xe3, 0x49, 0x30, 0x3e, 0x1e, 0x67, 0x4c, 0x30, 0x54, 0x0f, 0x98, 0x20, 0xfb,
	0xcf, 0x37, 0x09, 0x7c, 0x9a, 0x09, 0xae, 0x24, 0xfb, 0xdd, 0x0d, 0x64, 0x45, 0x33, 0xff, 0xd6,
	0xa0, 0x6b, 0x0d, 0x78, 0xe8, 0x15, 0x6c, 0xaf, 0xe4, 0x5c, 0xf6, 0x87, 0xd8, 0xfa, 0x95, 0xc4,
	0x31, 0x4d, 0x43, 0xea, 0xd2, 0xdf, 0x26, 0x94, 0x0b, 0x74, 0x08, 0x6d, 0x9f, 0xa5, 0x29, 0xf5,
	0x45, 0xc4, 0x52, 0x1c, 0x05, 0x86, 0xd6, 0xd1, 0x7a, 0xf7, 0xdc, 0x07, 0x1f, 0x40, 0x27, 0x40,
	0x07, 0xf0, 0x20, 0xc9, 0xb1, 0x88, 0x12, 0xca, 0x05, 0x49, 0xc6, 0xc6, 0x9d, 0x8e, 0xd6, 0xbb,
	0xeb, 0xb6, 0x92, 0x7c, 0x34, 0x85, 0xd0, 0x11, 0xe8, 0xf2, 0x68, 0x9f, 0xc5, 0xf8, 0x1d, 0xcd,
	0x78, 0xc4, 0x52, 0xa3, 0xde, 0xd1, 0x7a, 0x6d, 0x77, 0x77, 0x8a, 0xbf, 0x51, 0xb0, 0xf9, 0xaf,
	0x06, 0x87, 0x9b, 0xcc, 0x8d, 0xe3, 0x7c, 0x3b, 0x6b, 0x9f, 0x43, 0xd3, 0x9f, 0xca, 0x8c, 0x9a,
	0xf4, 0xf5, 0x01, 0x40, 0x5d, 0xd8, 0xc9, 0xd9, 0x24, 0x5b, 0xb2, 0xde, 0x2e, 0xd0, 0x5b, 0x99,
	0x7f, 0x5f, 0x5f, 0x65, 0x5e, 0xb9, 0x9a, 0xce, 0xf5, 0x25, 0xec, 0xf9, 0x71, 0x44, 0x53, 0x81,
	0xab, 0x7a, 0x40, 0xaa, 0x66, 0x6d, 0xdf, 0xc9, 0xe2, 0x0a, 0x1a, 0xcb, 0x2b, 0x78, 0x02, 0xad,
	0x71, 0x94, 0x86, 0x98, 0x72, 0x81, 0x13, 0x6e, 0xdc, 0x95, 0x0d, 0x34, 0x0b, 0xc8, 0xe6, 0x62,
	0xc0, 0x91, 0x05, 0x0d, 0x3f, 0xcb, 0xc7, 0xc2, 0xb8, 0xd7, 0xd1, 0x7a, 0xad, 0x93, 0x17, 0xc7,
	0x45, 0xfc, 0x8e, 0x67, 0xcd, 0xf4, 0x89, 0x20, 0x61, 0x46, 0x12, 0x8f, 0xf2, 0xa2, 0x53, 0xab,
	0xa0, 0x3a, 0xe9, 0x15, 0xf3, 0xa2, 0x30, 0xa5, 0x81, 0xab, 0xb4, 0xe8, 0x7b, 0xa8, 0x17, 0x79,
	0x94, 0xe3, 0x69, 0x9d, 0xf4, 0x56, 0xbc, 0xc3, 0xa2, 0x99, 0x88, 0xae, 0x22, 0x9f, 0x08, 0x5a,
	0xca, 0xa5, 0x0a, 0x7d, 0x07, 0x9f, 0xc5, 0x34, 0x24, 0x7e, 0x8e, 0x97, 0xe6, 0x7d, 0x5f, 0xda,
	0x7d, 0xa4, 0xca, 0xc3, 0xf9, 0xa9, 0xa3, 0x67, 0xb0, 0x1b, 0x05, 0x34, 0x15, 0x91, 0xc8, 0x31,
	0x17, 0x59, 0x94, 0x86, 0x06, 0x74, 0xb4, 0x5e, 0xd3, 0xdd, 0x99, 0xc2, 0x9e, 0x44, 0xd1, 0xb7,
	0xf0, 0x69, 0x79, 0x40, 0x39, 0x7d, 0x79, 0x5f, 0x8a, 0xc1, 0xab, 0xc5, 0x3f, 0x54, 0x55, 0x4b,
	0x16, 0xa5, 0x65, 0x27, 0x40, 0xfe, 0x4c, 0x34, 0x3b, 0xe4, 0x6d, 0x94, 0x92, 0x2c, 0x37, 0x9a,
	0x95, 0x93, 0x3a, 0x9f, 0x5d, 0x3f, 0xa7, 0xe4, 0xff, 0x2c, 0xe5, 0xaf, 0xa4, 0xc8, 0xdd, 0x53,
	0x2f, 0x9b, 0xd6, 0x14, 0x6a, 0xfe, 0x55, 0x87, 0x27, 0xeb, 0x82, 0x73, 0x71, 0x76, 0x8b, 0xcc,
	0xbc, 0x84, 0x3d, 0x4e, 0xb3, 0x77, 0x34, 0x5b, 0x50, 0x34, 0x94, 0x42, 0xd5, 0xe6, 0x14, 0x5b,
	0xde, 0x88, 0xa7, 0xb0, 0x1b, 0xd0, 0x98, 0xa8, 0xc4, 0xe1, 0x09, 0xa7, 0x7e, 0x79, 0x21, 0xda,
	0x12, 0x2e, 0x88, 0x97, 0x9c, 0xfa, 0x1f, 0x37, 0x53, 0xf7, 0x6f, 0x95, 0xa9, 0x8a, 0x6c, 0xb4,
	0x36, 0x64, 0xa3, 0x9c, 0xd9, 0x2c, 0x1b, 0xb5, 0x9b, 0xd9, 0xf0, 0x64, 0x71, 0x73, 0x36, 0xe0,
	0xe3, 0x65, 0xe3, 0xfd, 0xca, 0xbf, 0xeb, 0xd9, 0xea, 0xac, 0x98, 0x71, 0x1a, 0xa0, 0x1e, 0xe8,
	0x82, 0x2d, 0x2c, 0xbb, 0x2e, 0x97, 0xbd, 0x23, 0xd8, 0xdc, 0xa2, 0xbf, 0x01, 0x74, 0x95, 0xb1,
	0xa4, 0x32, 0x18, 0x7a, 0x51, 0x99, 0x63, 0xef, 0x41, 0x23, 0xa0, 0x6f, 0x27, 0xa1, 0x1c, 0x45,
	0xd3, 0x55, 0x0f, 0xe8, 0x4b, 0x68, 0x65, 0x94, 0x70, 0x96, 0x62, 0x9f, 0x05, 0x54, 0x26, 0xa5,
	0xed, 0x82, 0x82, 0x2c, 0x16, 0x50, 0xf3, 0x0f, 0x38, 0xa8, 0xf4, 0x7d, 0x7e, 0xc3, 0xcb, 0x0a,
	0x27, 0xb5, 0x15, 0x4e, 0xaa, 0x3a, 0xbc, 0x53, 0xd5, 0xa1, 0xf9, 0x9f, 0x06, 0xfb, 0x95, 0xa7,
	0x7b, 0x82, 0x08, 0x8e, 0x7e, 0x80, 0x06, 0x2f, 0x7e, 0xc8, 0xeb, 0xb3, 0x26, 0x56, 0xb3, 0x57,
	0xbe, 0x9e, 0x90, 0x38, 0x12, 0xb9, 0xab, 0x64, 0xc5, 0x48, 0xae, 0x62, 0x12, 0xf2, 0xb2, 0x6d,
	0xf5, 0x60, 0xbe, 0x81, 0xc6, 0x8f, 0xc5, 0x0f, 0xf4, 0x10, 0x76, 0x4f, 0xad, 0x33, 0xec, 0xda,
	0xaf, 0x2f, 0x6d, 0x6f, 0x84, 0xed, 0x13, 0x5b, 0xaf, 0xa1, 0xc7, 0xf0, 0xe8, 0x26, 0xe8, 0x0c,
	0x06, 0x76, 0xdf, 0x39, 0x1d, 0xd9, 0x7a, 0x1d, 0x7d, 0x01, 0x8f, 0xcf, 0x2f, 0x46, 0x78, 0xe8,
	0x3a, 0x83, 0x53, 0xf7, 0x17, 0x3c, 0x72, 0x4f, 0xcf, 0xbd, 0xe1, 0x85, 0xab, 0x94, 0xfa, 0xd7,
	0xff, 0xd4, 0xc0, 0xb0, 0x17, 0x12, 0x74, 0xd9, 0x1f, 0x0e, 0x78, 0xe8, 0xf4, 0xd1, 0x0b, 0x38,
	0xba, 0xc6, 0xd5, 0xd5, 0xa5, 0x2f, 0xba, 0xde, 0x41, 0xcf, 0xe1, 0xd9, 0x36, 0xf4, 0x71, 0x9c,
	0xeb, 0x07, 0xeb, 0xc9, 0x73, 0xdf, 0x34, 0xdd, 0x44, 0x4f, 0xc1, 0xdc, 0x44, 0xbe, 0x38, 0xd3,
	0x0f, 0xd7, 0x1b, 0x5e, 0xc8, 0xb4, 0xfe, 0x15, 0x3a, 0x82, 0xee, 0x4a, 0xfa, 0xcd, 0x28, 0xe9,
	0xdd, 0x57, 0x8d, 0x9f, 0xb4, 0x3f, 0xb5, 0x4f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x1c,
	0x8b, 0xca, 0x66, 0x09, 0x00, 0x00,
}
