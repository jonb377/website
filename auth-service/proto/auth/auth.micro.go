// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package go_micro_api_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...client.CallOption) (*CreateConnectionResponse, error)
	ConnectionChallenge(ctx context.Context, in *ConnectionChallengeRequest, opts ...client.CallOption) (*ConnectionChallengeResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...client.CallOption) (*ValidateTokenResponse, error)
	CloseConnection(ctx context.Context, in *CloseConnectionRequest, opts ...client.CallOption) (*Empty, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...client.CallOption) (*CreateConnectionResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateConnection", in)
	out := new(CreateConnectionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ConnectionChallenge(ctx context.Context, in *ConnectionChallengeRequest, opts ...client.CallOption) (*ConnectionChallengeResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ConnectionChallenge", in)
	out := new(ConnectionChallengeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...client.CallOption) (*ValidateTokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateToken", in)
	out := new(ValidateTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CloseConnection(ctx context.Context, in *CloseConnectionRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Auth.CloseConnection", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	CreateConnection(context.Context, *CreateConnectionRequest, *CreateConnectionResponse) error
	ConnectionChallenge(context.Context, *ConnectionChallengeRequest, *ConnectionChallengeResponse) error
	ValidateToken(context.Context, *ValidateTokenRequest, *ValidateTokenResponse) error
	CloseConnection(context.Context, *CloseConnectionRequest, *Empty) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		CreateConnection(ctx context.Context, in *CreateConnectionRequest, out *CreateConnectionResponse) error
		ConnectionChallenge(ctx context.Context, in *ConnectionChallengeRequest, out *ConnectionChallengeResponse) error
		ValidateToken(ctx context.Context, in *ValidateTokenRequest, out *ValidateTokenResponse) error
		CloseConnection(ctx context.Context, in *CloseConnectionRequest, out *Empty) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) CreateConnection(ctx context.Context, in *CreateConnectionRequest, out *CreateConnectionResponse) error {
	return h.AuthHandler.CreateConnection(ctx, in, out)
}

func (h *authHandler) ConnectionChallenge(ctx context.Context, in *ConnectionChallengeRequest, out *ConnectionChallengeResponse) error {
	return h.AuthHandler.ConnectionChallenge(ctx, in, out)
}

func (h *authHandler) ValidateToken(ctx context.Context, in *ValidateTokenRequest, out *ValidateTokenResponse) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}

func (h *authHandler) CloseConnection(ctx context.Context, in *CloseConnectionRequest, out *Empty) error {
	return h.AuthHandler.CloseConnection(ctx, in, out)
}
