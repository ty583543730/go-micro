// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: RenewalReadAo.proto

package daniel_proto_renewal

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

// Client API for RenewalReadAo handle

type RenewalReadAoService interface {
	GetRenewalDealInfo(ctx context.Context, in *GetRenewalRequest, opts ...client.CallOption) (*GetRenewalResponse, error)
}

type renewalReadAoService struct {
	c    client.Client
	name string
}

func NewRenewalReadAoService(name string, c client.Client) RenewalReadAoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "daniel.proto.renewal"
	}
	return &renewalReadAoService{
		c:    c,
		name: name,
	}
}

func (c *renewalReadAoService) GetRenewalDealInfo(ctx context.Context, in *GetRenewalRequest, opts ...client.CallOption) (*GetRenewalResponse, error) {
	req := c.c.NewRequest(c.name, "RenewalReadAo.GetRenewalDealInfo", in)
	out := new(GetRenewalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RenewalReadAo handle

type RenewalReadAoHandler interface {
	GetRenewalDealInfo(context.Context, *GetRenewalRequest, *GetRenewalResponse) error
}

func RegisterRenewalReadAoHandler(s server.Server, hdlr RenewalReadAoHandler, opts ...server.HandlerOption) error {
	type renewalReadAo interface {
		GetRenewalDealInfo(ctx context.Context, in *GetRenewalRequest, out *GetRenewalResponse) error
	}
	type RenewalReadAo struct {
		renewalReadAo
	}
	h := &renewalReadAoHandler{hdlr}
	return s.Handle(s.NewHandler(&RenewalReadAo{h}, opts...))
}

type renewalReadAoHandler struct {
	RenewalReadAoHandler
}

func (h *renewalReadAoHandler) GetRenewalDealInfo(ctx context.Context, in *GetRenewalRequest, out *GetRenewalResponse) error {
	return h.RenewalReadAoHandler.GetRenewalDealInfo(ctx, in, out)
}
