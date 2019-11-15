// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: Order.proto

package go_micro_srv_order

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
)

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
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

// Client API for ORDER service

type ORDERService interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type oRDERService struct {
	c    client.Client
	name string
}

func NewORDERService(name string, c client.Client) ORDERService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.Usrv.order"
	}
	return &oRDERService{
		c:    c,
		name: name,
	}
}

func (c *oRDERService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ORDER.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oRDERService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ORDER.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ORDER service

type ORDERHandler interface {
	Create(context.Context, *Request, *Response) error
	GetAll(context.Context, *Request, *Response) error
}

func RegisterORDERHandler(s server.Server, hdlr ORDERHandler, opts ...server.HandlerOption) error {
	type oRDER interface {
		Create(ctx context.Context, in *Request, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
	}
	type ORDER struct {
		oRDER
	}
	h := &oRDERHandler{hdlr}
	return s.Handle(s.NewHandler(&ORDER{h}, opts...))
}

type oRDERHandler struct {
	ORDERHandler
}

func (h *oRDERHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.ORDERHandler.Create(ctx, in, out)
}

func (h *oRDERHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.ORDERHandler.GetAll(ctx, in, out)
}
