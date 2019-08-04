// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package go_micro_api_auth

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

type CreateConnectionRequest struct {
	A                    []byte   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConnectionRequest) Reset()         { *m = CreateConnectionRequest{} }
func (m *CreateConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionRequest) ProtoMessage()    {}
func (*CreateConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *CreateConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionRequest.Unmarshal(m, b)
}
func (m *CreateConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionRequest.Merge(m, src)
}
func (m *CreateConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionRequest.Size(m)
}
func (m *CreateConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionRequest proto.InternalMessageInfo

func (m *CreateConnectionRequest) GetA() []byte {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *CreateConnectionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CreateConnectionResponse struct {
	B                    []byte   `protobuf:"bytes,1,opt,name=B,proto3" json:"B,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConnectionResponse) Reset()         { *m = CreateConnectionResponse{} }
func (m *CreateConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionResponse) ProtoMessage()    {}
func (*CreateConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *CreateConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionResponse.Unmarshal(m, b)
}
func (m *CreateConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionResponse.Marshal(b, m, deterministic)
}
func (m *CreateConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionResponse.Merge(m, src)
}
func (m *CreateConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionResponse.Size(m)
}
func (m *CreateConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionResponse proto.InternalMessageInfo

func (m *CreateConnectionResponse) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *CreateConnectionResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ConnectionChallengeRequest struct {
	M                    []byte   `protobuf:"bytes,1,opt,name=M,proto3" json:"M,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionChallengeRequest) Reset()         { *m = ConnectionChallengeRequest{} }
func (m *ConnectionChallengeRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionChallengeRequest) ProtoMessage()    {}
func (*ConnectionChallengeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *ConnectionChallengeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionChallengeRequest.Unmarshal(m, b)
}
func (m *ConnectionChallengeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionChallengeRequest.Marshal(b, m, deterministic)
}
func (m *ConnectionChallengeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionChallengeRequest.Merge(m, src)
}
func (m *ConnectionChallengeRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionChallengeRequest.Size(m)
}
func (m *ConnectionChallengeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionChallengeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionChallengeRequest proto.InternalMessageInfo

func (m *ConnectionChallengeRequest) GetM() []byte {
	if m != nil {
		return m.M
	}
	return nil
}

type ConnectionChallengeResponse struct {
	HAMK                 []byte   `protobuf:"bytes,1,opt,name=HAMK,proto3" json:"HAMK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionChallengeResponse) Reset()         { *m = ConnectionChallengeResponse{} }
func (m *ConnectionChallengeResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionChallengeResponse) ProtoMessage()    {}
func (*ConnectionChallengeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *ConnectionChallengeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionChallengeResponse.Unmarshal(m, b)
}
func (m *ConnectionChallengeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionChallengeResponse.Marshal(b, m, deterministic)
}
func (m *ConnectionChallengeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionChallengeResponse.Merge(m, src)
}
func (m *ConnectionChallengeResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionChallengeResponse.Size(m)
}
func (m *ConnectionChallengeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionChallengeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionChallengeResponse proto.InternalMessageInfo

func (m *ConnectionChallengeResponse) GetHAMK() []byte {
	if m != nil {
		return m.HAMK
	}
	return nil
}

type ConnectionCloseRequest struct {
	A                    bool     `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionCloseRequest) Reset()         { *m = ConnectionCloseRequest{} }
func (m *ConnectionCloseRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionCloseRequest) ProtoMessage()    {}
func (*ConnectionCloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *ConnectionCloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionCloseRequest.Unmarshal(m, b)
}
func (m *ConnectionCloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionCloseRequest.Marshal(b, m, deterministic)
}
func (m *ConnectionCloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionCloseRequest.Merge(m, src)
}
func (m *ConnectionCloseRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionCloseRequest.Size(m)
}
func (m *ConnectionCloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionCloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionCloseRequest proto.InternalMessageInfo

func (m *ConnectionCloseRequest) GetA() bool {
	if m != nil {
		return m.A
	}
	return false
}

type ValidateTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Device               string   `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (m *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(m, src)
}
func (m *ValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRequest.Size(m)
}
func (m *ValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRequest proto.InternalMessageInfo

func (m *ValidateTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ValidateTokenRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type ValidateTokenResponse struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	SessionKey           []byte   `protobuf:"bytes,2,opt,name=SessionKey,proto3" json:"SessionKey,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenResponse) Reset()         { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()    {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{6}
}

func (m *ValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenResponse.Unmarshal(m, b)
}
func (m *ValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenResponse.Marshal(b, m, deterministic)
}
func (m *ValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenResponse.Merge(m, src)
}
func (m *ValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenResponse.Size(m)
}
func (m *ValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenResponse proto.InternalMessageInfo

func (m *ValidateTokenResponse) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *ValidateTokenResponse) GetSessionKey() []byte {
	if m != nil {
		return m.SessionKey
	}
	return nil
}

func (m *ValidateTokenResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateConnectionRequest)(nil), "go.micro.api.auth.CreateConnectionRequest")
	proto.RegisterType((*CreateConnectionResponse)(nil), "go.micro.api.auth.CreateConnectionResponse")
	proto.RegisterType((*ConnectionChallengeRequest)(nil), "go.micro.api.auth.ConnectionChallengeRequest")
	proto.RegisterType((*ConnectionChallengeResponse)(nil), "go.micro.api.auth.ConnectionChallengeResponse")
	proto.RegisterType((*ConnectionCloseRequest)(nil), "go.micro.api.auth.ConnectionCloseRequest")
	proto.RegisterType((*ValidateTokenRequest)(nil), "go.micro.api.auth.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "go.micro.api.auth.ValidateTokenResponse")
	proto.RegisterType((*Empty)(nil), "go.micro.api.auth.Empty")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x4f, 0xe2, 0x40,
	0x14, 0xc5, 0xe9, 0xf2, 0x67, 0xe1, 0x86, 0xcd, 0xee, 0x8e, 0x80, 0x4d, 0x35, 0x86, 0xcc, 0x83,
	0x22, 0xc6, 0x1a, 0xf5, 0xdd, 0xa4, 0x14, 0x13, 0x0d, 0xe1, 0xa5, 0xfe, 0x89, 0xaf, 0x15, 0x6e,
	0xa0, 0xb1, 0x9d, 0x29, 0xed, 0x40, 0xc2, 0x87, 0xf2, 0x3b, 0x1a, 0x86, 0xb6, 0x14, 0x28, 0xca,
	0x4b, 0xd3, 0x3b, 0x3d, 0xf3, 0xbb, 0x67, 0xee, 0x99, 0x42, 0xdd, 0x0f, 0xb8, 0xe0, 0x57, 0xf6,
	0x54, 0x8c, 0xe5, 0x43, 0x97, 0x35, 0xf9, 0x3f, 0xe2, 0xba, 0xe7, 0x0c, 0x02, 0xae, 0xdb, 0xbe,
	0xa3, 0x2f, 0x3e, 0x50, 0x13, 0x0e, 0xcd, 0x00, 0x6d, 0x81, 0x26, 0x67, 0x0c, 0x07, 0xc2, 0xe1,
	0xcc, 0xc2, 0xc9, 0x14, 0x43, 0x41, 0xaa, 0xa0, 0x18, 0xaa, 0xd2, 0x54, 0x5a, 0x55, 0x4b, 0x31,
	0x88, 0x06, 0xe5, 0x97, 0x10, 0x03, 0x66, 0x7b, 0xa8, 0xfe, 0x6a, 0x2a, 0xad, 0x8a, 0x95, 0xd4,
	0xf4, 0x0e, 0xd4, 0x6d, 0x48, 0xe8, 0x73, 0x16, 0xe2, 0x82, 0xd2, 0x89, 0x29, 0x1d, 0x52, 0x83,
	0xe2, 0x33, 0xff, 0x40, 0x16, 0x21, 0x96, 0x05, 0x6d, 0x83, 0xb6, 0xda, 0x69, 0x8e, 0x6d, 0xd7,
	0x45, 0x36, 0xc2, 0x94, 0x8f, 0x7e, 0x4c, 0xe8, 0xd3, 0x6b, 0x38, 0xca, 0xd4, 0x46, 0xed, 0x08,
	0x14, 0x1e, 0x8c, 0x7e, 0x2f, 0xd2, 0xcb, 0x77, 0x7a, 0x0a, 0x8d, 0xd4, 0x16, 0x97, 0x87, 0xb8,
	0x75, 0xc4, 0xb2, 0xa5, 0x18, 0xb4, 0x0b, 0xb5, 0x57, 0xdb, 0x75, 0x86, 0xb6, 0x40, 0xe9, 0x2b,
	0x56, 0x25, 0xa6, 0x95, 0x94, 0x69, 0xd2, 0x80, 0x52, 0x17, 0x67, 0xce, 0x20, 0x1e, 0x47, 0x54,
	0xd1, 0x09, 0xd4, 0x37, 0x28, 0x91, 0xb5, 0x63, 0xa8, 0x3c, 0x61, 0x18, 0x3a, 0x9c, 0x3d, 0x76,
	0x23, 0xd4, 0x6a, 0x81, 0x9c, 0x00, 0x44, 0x45, 0x0f, 0xe7, 0x12, 0x59, 0xb5, 0x52, 0x2b, 0x6b,
	0xf3, 0xcf, 0x6f, 0xcc, 0xff, 0x37, 0x14, 0xef, 0x3d, 0x5f, 0xcc, 0x6f, 0x3e, 0xf3, 0x50, 0x30,
	0xa6, 0x62, 0x4c, 0x3c, 0xf8, 0xb7, 0x99, 0x08, 0x69, 0xeb, 0x5b, 0xf1, 0xeb, 0x3b, 0xb2, 0xd7,
	0x2e, 0xf6, 0xd2, 0x2e, 0x0f, 0x46, 0x73, 0x64, 0x06, 0x07, 0x19, 0xa1, 0x90, 0xcb, 0x2c, 0xca,
	0xce, 0xa0, 0x35, 0x7d, 0x5f, 0x79, 0xd2, 0x77, 0x08, 0x7f, 0xd6, 0x66, 0x4d, 0xce, 0x32, 0x10,
	0x59, 0x99, 0x6a, 0xad, 0x9f, 0x85, 0x49, 0x97, 0x37, 0xf8, 0x2b, 0x6f, 0x4d, 0x6a, 0x96, 0xe7,
	0xdf, 0x5b, 0x4d, 0xdd, 0x31, 0x4d, 0xcd, 0x90, 0xca, 0xb4, 0x68, 0xee, 0xbd, 0x24, 0xff, 0xcb,
	0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x9c, 0x08, 0xb6, 0xb0, 0x03, 0x00, 0x00,
}
