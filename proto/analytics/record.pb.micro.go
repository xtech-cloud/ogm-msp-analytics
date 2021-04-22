// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/analytics/record.proto

package analytics

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for Record service

type RecordService interface {
	// 活跃
	Wake(ctx context.Context, in *Agent, opts ...client.CallOption) (*BlankResponse, error)
	// 活动
	Activity(ctx context.Context, in *RecordActivityRequest, opts ...client.CallOption) (*BlankResponse, error)
}

type recordService struct {
	c    client.Client
	name string
}

func NewRecordService(name string, c client.Client) RecordService {
	return &recordService{
		c:    c,
		name: name,
	}
}

func (c *recordService) Wake(ctx context.Context, in *Agent, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Record.Wake", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) Activity(ctx context.Context, in *RecordActivityRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Record.Activity", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Record service

type RecordHandler interface {
	// 活跃
	Wake(context.Context, *Agent, *BlankResponse) error
	// 活动
	Activity(context.Context, *RecordActivityRequest, *BlankResponse) error
}

func RegisterRecordHandler(s server.Server, hdlr RecordHandler, opts ...server.HandlerOption) error {
	type record interface {
		Wake(ctx context.Context, in *Agent, out *BlankResponse) error
		Activity(ctx context.Context, in *RecordActivityRequest, out *BlankResponse) error
	}
	type Record struct {
		record
	}
	h := &recordHandler{hdlr}
	return s.Handle(s.NewHandler(&Record{h}, opts...))
}

type recordHandler struct {
	RecordHandler
}

func (h *recordHandler) Wake(ctx context.Context, in *Agent, out *BlankResponse) error {
	return h.RecordHandler.Wake(ctx, in, out)
}

func (h *recordHandler) Activity(ctx context.Context, in *RecordActivityRequest, out *BlankResponse) error {
	return h.RecordHandler.Activity(ctx, in, out)
}
