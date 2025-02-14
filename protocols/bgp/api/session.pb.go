// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bio-routing/bio-rd/protocols/bgp/api/session.proto

package api

import (
	fmt "fmt"
	api "github.com/bio-routing/bio-rd/net/api"
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

type Session_State int32

const (
	Session_Disabled      Session_State = 0
	Session_Idle          Session_State = 1
	Session_Connect       Session_State = 2
	Session_Active        Session_State = 3
	Session_OpenSent      Session_State = 4
	Session_OpenConfirmed Session_State = 5
	Session_Established   Session_State = 6
)

var Session_State_name = map[int32]string{
	0: "Disabled",
	1: "Idle",
	2: "Connect",
	3: "Active",
	4: "OpenSent",
	5: "OpenConfirmed",
	6: "Established",
}

var Session_State_value = map[string]int32{
	"Disabled":      0,
	"Idle":          1,
	"Connect":       2,
	"Active":        3,
	"OpenSent":      4,
	"OpenConfirmed": 5,
	"Established":   6,
}

func (x Session_State) String() string {
	return proto.EnumName(Session_State_name, int32(x))
}

func (Session_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b53032c0bb76d75, []int{0, 0}
}

type Session struct {
	LocalAddress         *api.IP       `protobuf:"bytes,1,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	NeighborAddress      *api.IP       `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	LocalAsn             uint32        `protobuf:"varint,3,opt,name=local_asn,json=localAsn,proto3" json:"local_asn,omitempty"`
	PeerAsn              uint32        `protobuf:"varint,4,opt,name=peer_asn,json=peerAsn,proto3" json:"peer_asn,omitempty"`
	Status               Session_State `protobuf:"varint,5,opt,name=status,proto3,enum=bio.bgp.Session_State" json:"status,omitempty"`
	Stats                *SessionStats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	EstablishedSince     uint64        `protobuf:"varint,7,opt,name=established_since,json=establishedSince,proto3" json:"established_since,omitempty"`
	Description          string        `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b53032c0bb76d75, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetLocalAddress() *api.IP {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Session) GetNeighborAddress() *api.IP {
	if m != nil {
		return m.NeighborAddress
	}
	return nil
}

func (m *Session) GetLocalAsn() uint32 {
	if m != nil {
		return m.LocalAsn
	}
	return 0
}

func (m *Session) GetPeerAsn() uint32 {
	if m != nil {
		return m.PeerAsn
	}
	return 0
}

func (m *Session) GetStatus() Session_State {
	if m != nil {
		return m.Status
	}
	return Session_Disabled
}

func (m *Session) GetStats() *SessionStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Session) GetEstablishedSince() uint64 {
	if m != nil {
		return m.EstablishedSince
	}
	return 0
}

func (m *Session) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type SessionStats struct {
	MessagesIn           uint64   `protobuf:"varint,1,opt,name=messages_in,json=messagesIn,proto3" json:"messages_in,omitempty"`
	MessagesOut          uint64   `protobuf:"varint,2,opt,name=messages_out,json=messagesOut,proto3" json:"messages_out,omitempty"`
	Flaps                uint64   `protobuf:"varint,3,opt,name=flaps,proto3" json:"flaps,omitempty"`
	RoutesReceived       uint64   `protobuf:"varint,4,opt,name=routes_received,json=routesReceived,proto3" json:"routes_received,omitempty"`
	RoutesImported       uint64   `protobuf:"varint,5,opt,name=routes_imported,json=routesImported,proto3" json:"routes_imported,omitempty"`
	RoutesExported       uint64   `protobuf:"varint,6,opt,name=routes_exported,json=routesExported,proto3" json:"routes_exported,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionStats) Reset()         { *m = SessionStats{} }
func (m *SessionStats) String() string { return proto.CompactTextString(m) }
func (*SessionStats) ProtoMessage()    {}
func (*SessionStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b53032c0bb76d75, []int{1}
}

func (m *SessionStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionStats.Unmarshal(m, b)
}
func (m *SessionStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionStats.Marshal(b, m, deterministic)
}
func (m *SessionStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionStats.Merge(m, src)
}
func (m *SessionStats) XXX_Size() int {
	return xxx_messageInfo_SessionStats.Size(m)
}
func (m *SessionStats) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionStats.DiscardUnknown(m)
}

var xxx_messageInfo_SessionStats proto.InternalMessageInfo

func (m *SessionStats) GetMessagesIn() uint64 {
	if m != nil {
		return m.MessagesIn
	}
	return 0
}

func (m *SessionStats) GetMessagesOut() uint64 {
	if m != nil {
		return m.MessagesOut
	}
	return 0
}

func (m *SessionStats) GetFlaps() uint64 {
	if m != nil {
		return m.Flaps
	}
	return 0
}

func (m *SessionStats) GetRoutesReceived() uint64 {
	if m != nil {
		return m.RoutesReceived
	}
	return 0
}

func (m *SessionStats) GetRoutesImported() uint64 {
	if m != nil {
		return m.RoutesImported
	}
	return 0
}

func (m *SessionStats) GetRoutesExported() uint64 {
	if m != nil {
		return m.RoutesExported
	}
	return 0
}

func init() {
	proto.RegisterEnum("bio.bgp.Session_State", Session_State_name, Session_State_value)
	proto.RegisterType((*Session)(nil), "bio.bgp.Session")
	proto.RegisterType((*SessionStats)(nil), "bio.bgp.SessionStats")
}

func init() {
	proto.RegisterFile("github.com/bio-routing/bio-rd/protocols/bgp/api/session.proto", fileDescriptor_5b53032c0bb76d75)
}

var fileDescriptor_5b53032c0bb76d75 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xeb, 0x38, 0xe9, 0x38, 0x69, 0xdc, 0x15, 0x20, 0x03, 0x07, 0x4c, 0x2e, 0x44,
	0xaa, 0xb0, 0xa1, 0x48, 0xdc, 0x38, 0x94, 0xd2, 0x43, 0x4e, 0x45, 0x9b, 0x1b, 0x97, 0xc8, 0xf6,
	0x4e, 0x9d, 0x45, 0xce, 0xae, 0xe5, 0xd9, 0x54, 0xbc, 0x1f, 0x8f, 0xc2, 0x8b, 0xa0, 0x5d, 0x3b,
	0xc1, 0x02, 0x09, 0x89, 0xdb, 0xce, 0x3f, 0xdf, 0x3f, 0x99, 0xd1, 0x1f, 0xc3, 0xc7, 0x4a, 0x9a,
	0xed, 0xbe, 0x48, 0x4b, 0xbd, 0xcb, 0x0a, 0xa9, 0xdf, 0xb4, 0x7a, 0x6f, 0xa4, 0xaa, 0xba, 0xb7,
	0xc8, 0x9a, 0x56, 0x1b, 0x5d, 0xea, 0x9a, 0xb2, 0xa2, 0x6a, 0xb2, 0xbc, 0x91, 0x19, 0x21, 0x91,
	0xd4, 0x2a, 0x75, 0x1d, 0x36, 0x2e, 0xa4, 0x4e, 0x8b, 0xaa, 0x79, 0x9e, 0xfd, 0x7b, 0x8e, 0x42,
	0xe3, 0xdc, 0x0a, 0x4d, 0xe7, 0x5c, 0xfc, 0x38, 0x85, 0xf1, 0xba, 0x9b, 0xc5, 0xde, 0xc2, 0xac,
	0xd6, 0x65, 0x5e, 0x6f, 0x72, 0x21, 0x5a, 0x24, 0x8a, 0xbd, 0xc4, 0x5b, 0x86, 0x57, 0x61, 0x6a,
	0xa7, 0x5b, 0xcb, 0xea, 0x0b, 0x9f, 0x3a, 0xe2, 0xba, 0x03, 0xd8, 0x07, 0x88, 0x14, 0xca, 0x6a,
	0x5b, 0xe8, 0xf6, 0x68, 0x3a, 0xf9, 0xdb, 0x34, 0x3f, 0x40, 0x07, 0xdf, 0x0b, 0x38, 0xeb, 0x7f,
	0x89, 0x54, 0x7c, 0x9a, 0x78, 0xcb, 0x19, 0x9f, 0x74, 0x83, 0x49, 0xb1, 0x67, 0x30, 0x69, 0x10,
	0x5b, 0xd7, 0xf3, 0x5d, 0x6f, 0x6c, 0x6b, 0xdb, 0x4a, 0x21, 0x20, 0x93, 0x9b, 0x3d, 0xc5, 0xa3,
	0xc4, 0x5b, 0x9e, 0x5f, 0x3d, 0x4d, 0xfb, 0xc3, 0xd3, 0xfe, 0x86, 0x74, 0x6d, 0x72, 0x83, 0xbc,
	0xa7, 0xd8, 0x25, 0x8c, 0xec, 0x8b, 0xe2, 0xc0, 0x2d, 0xf5, 0xe4, 0x4f, 0xdc, 0xd2, 0xc4, 0x3b,
	0x86, 0x5d, 0xc2, 0x05, 0x92, 0xc9, 0x8b, 0x5a, 0xd2, 0x16, 0xc5, 0x86, 0xa4, 0x2a, 0x31, 0x1e,
	0x27, 0xde, 0xd2, 0xe7, 0xd1, 0xa0, 0xb1, 0xb6, 0x3a, 0x4b, 0x20, 0x14, 0x48, 0x65, 0x2b, 0x1b,
	0x23, 0xb5, 0x8a, 0x27, 0x89, 0xb7, 0x3c, 0xe3, 0x43, 0x69, 0xf1, 0x0d, 0x46, 0x6e, 0x19, 0x36,
	0x85, 0xc9, 0x67, 0x49, 0x79, 0x51, 0xa3, 0x88, 0x1e, 0xb1, 0x09, 0xf8, 0x2b, 0x51, 0x63, 0xe4,
	0xb1, 0x10, 0xc6, 0x37, 0x5a, 0x29, 0x2c, 0x4d, 0x74, 0xc2, 0x00, 0x82, 0xeb, 0xd2, 0xc8, 0x07,
	0x8c, 0x4e, 0xad, 0xe1, 0xae, 0x41, 0xb5, 0x46, 0x65, 0x22, 0x9f, 0x5d, 0xc0, 0xcc, 0x56, 0x37,
	0x5a, 0xdd, 0xcb, 0x76, 0x87, 0x22, 0x1a, 0xb1, 0x39, 0x84, 0xb7, 0xbf, 0x17, 0x8a, 0x82, 0xc5,
	0x4f, 0x0f, 0xa6, 0xc3, 0x93, 0xd8, 0x4b, 0x08, 0x77, 0x48, 0x94, 0x57, 0x48, 0x1b, 0xa9, 0x5c,
	0x90, 0x3e, 0x87, 0x83, 0xb4, 0x52, 0xec, 0x15, 0x4c, 0x8f, 0x80, 0xde, 0x1b, 0x97, 0x9a, 0xcf,
	0x8f, 0xa6, 0xbb, 0xbd, 0x61, 0x8f, 0x61, 0x74, 0x5f, 0xe7, 0x0d, 0xb9, 0x80, 0x7c, 0xde, 0x15,
	0xec, 0x35, 0xcc, 0xed, 0x9f, 0x0a, 0x69, 0xd3, 0x62, 0x89, 0xf2, 0x01, 0x85, 0x0b, 0xc9, 0xe7,
	0xe7, 0x9d, 0xcc, 0x7b, 0x75, 0x00, 0xca, 0x5d, 0xa3, 0x5b, 0x83, 0xc2, 0x85, 0x76, 0x04, 0x57,
	0xbd, 0x3a, 0x00, 0xf1, 0x7b, 0x0f, 0x06, 0x43, 0xf0, 0xb6, 0x57, 0x3f, 0xbd, 0xfb, 0x9a, 0xfd,
	0xe7, 0x67, 0x52, 0x04, 0x4e, 0x7a, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x87, 0xfb, 0xe6, 0x6a,
	0x60, 0x03, 0x00, 0x00,
}
