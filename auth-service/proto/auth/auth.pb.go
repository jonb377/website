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
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x4f, 0xfa, 0x30,
	0x14, 0xc5, 0xd9, 0xff, 0x0f, 0x08, 0x37, 0x18, 0xb5, 0x02, 0x2e, 0xd3, 0x18, 0xd2, 0x07, 0x45,
	0x8c, 0x33, 0xea, 0xbb, 0xc9, 0x18, 0x26, 0x1a, 0xc2, 0xcb, 0x14, 0xe3, 0xeb, 0x84, 0x1b, 0x58,
	0xdc, 0xda, 0xb9, 0x15, 0x12, 0x3e, 0x94, 0xdf, 0xd1, 0x50, 0xb6, 0x31, 0x61, 0x28, 0x2f, 0xcb,
	0x6e, 0x77, 0xfa, 0xbb, 0x67, 0xf7, 0xb4, 0x50, 0xf3, 0x03, 0x2e, 0xf8, 0xb5, 0x3d, 0x11, 0x63,
	0xf9, 0xd0, 0x65, 0x4d, 0x0e, 0x46, 0x5c, 0xf7, 0x9c, 0x41, 0xc0, 0x75, 0xdb, 0x77, 0xf4, 0xf9,
	0x07, 0x6a, 0xc2, 0x91, 0x19, 0xa0, 0x2d, 0xd0, 0xe4, 0x8c, 0xe1, 0x40, 0x38, 0x9c, 0x59, 0xf8,
	0x39, 0xc1, 0x50, 0x90, 0x0a, 0x28, 0x86, 0xaa, 0x34, 0x94, 0x66, 0xc5, 0x52, 0x0c, 0xa2, 0x41,
	0xa9, 0x1f, 0x62, 0xc0, 0x6c, 0x0f, 0xd5, 0x7f, 0x0d, 0xa5, 0x59, 0xb6, 0x92, 0x9a, 0xde, 0x83,
	0xba, 0x0e, 0x09, 0x7d, 0xce, 0x42, 0x9c, 0x53, 0xda, 0x31, 0xa5, 0x4d, 0xaa, 0x50, 0x78, 0xe1,
	0x1f, 0xc8, 0x22, 0xc4, 0xa2, 0xa0, 0x2d, 0xd0, 0x96, 0x3b, 0xcd, 0xb1, 0xed, 0xba, 0xc8, 0x46,
	0x98, 0xf2, 0xd1, 0x8b, 0x09, 0x3d, 0x7a, 0x03, 0xc7, 0x99, 0xda, 0xa8, 0x1d, 0x81, 0xfc, 0xa3,
	0xd1, 0xeb, 0x46, 0x7a, 0xf9, 0x4e, 0xcf, 0xa0, 0x9e, 0xda, 0xe2, 0xf2, 0x10, 0xd7, 0x7e, 0xb1,
	0x64, 0x29, 0x06, 0xed, 0x40, 0xf5, 0xd5, 0x76, 0x9d, 0xa1, 0x2d, 0x50, 0xfa, 0x8a, 0x55, 0x89,
	0x69, 0x25, 0x65, 0x9a, 0xd4, 0xa1, 0xd8, 0xc1, 0xa9, 0x33, 0x88, 0xc7, 0x11, 0x55, 0xb4, 0x0f,
	0xb5, 0x15, 0x4a, 0x64, 0xed, 0x04, 0xca, 0xcf, 0x18, 0x86, 0x0e, 0x67, 0x4f, 0x9d, 0x08, 0xb5,
	0x5c, 0x20, 0xa7, 0x00, 0x51, 0xd1, 0xc5, 0x99, 0x44, 0x56, 0xac, 0xd4, 0x0a, 0xdd, 0x81, 0xc2,
	0x83, 0xe7, 0x8b, 0xd9, 0xed, 0xd7, 0x7f, 0xc8, 0x1b, 0x13, 0x31, 0x26, 0x1e, 0xec, 0xaf, 0x4e,
	0x9d, 0xb4, 0xf4, 0xb5, 0x88, 0xf5, 0x0d, 0xf9, 0x6a, 0x97, 0x5b, 0x69, 0x17, 0xe6, 0x69, 0x8e,
	0x4c, 0xe1, 0x30, 0x63, 0xf0, 0xe4, 0x2a, 0x8b, 0xb2, 0x31, 0x4c, 0x4d, 0xdf, 0x56, 0x9e, 0xf4,
	0x1d, 0xc2, 0xee, 0x8f, 0x79, 0x92, 0xf3, 0x0c, 0x44, 0x56, 0x6e, 0x5a, 0xf3, 0x6f, 0x61, 0xd2,
	0xe5, 0x0d, 0xf6, 0xe4, 0xc9, 0x48, 0xcd, 0xf2, 0xe2, 0x77, 0xab, 0xa9, 0x73, 0xa4, 0xa9, 0x19,
	0x52, 0x99, 0x16, 0xcd, 0xbd, 0x17, 0xe5, 0xdd, 0xbb, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x43, 0x13, 0x5c, 0x94, 0x03, 0x00, 0x00,
}
