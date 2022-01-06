// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/analytics/generator.proto

package analytics

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Generator service

func NewGeneratorEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Generator service

type GeneratorService interface {
	// 个体
	Agent(ctx context.Context, in *GeneratorAgentRequest, opts ...client.CallOption) (*GeneratorAgentResponse, error)
	// 记录
	Record(ctx context.Context, in *GeneratorRecordRequest, opts ...client.CallOption) (*GeneratorRecordResponse, error)
}

type generatorService struct {
	c    client.Client
	name string
}

func NewGeneratorService(name string, c client.Client) GeneratorService {
	return &generatorService{
		c:    c,
		name: name,
	}
}

func (c *generatorService) Agent(ctx context.Context, in *GeneratorAgentRequest, opts ...client.CallOption) (*GeneratorAgentResponse, error) {
	req := c.c.NewRequest(c.name, "Generator.Agent", in)
	out := new(GeneratorAgentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorService) Record(ctx context.Context, in *GeneratorRecordRequest, opts ...client.CallOption) (*GeneratorRecordResponse, error) {
	req := c.c.NewRequest(c.name, "Generator.Record", in)
	out := new(GeneratorRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generator service

type GeneratorHandler interface {
	// 个体
	Agent(context.Context, *GeneratorAgentRequest, *GeneratorAgentResponse) error
	// 记录
	Record(context.Context, *GeneratorRecordRequest, *GeneratorRecordResponse) error
}

func RegisterGeneratorHandler(s server.Server, hdlr GeneratorHandler, opts ...server.HandlerOption) error {
	type generator interface {
		Agent(ctx context.Context, in *GeneratorAgentRequest, out *GeneratorAgentResponse) error
		Record(ctx context.Context, in *GeneratorRecordRequest, out *GeneratorRecordResponse) error
	}
	type Generator struct {
		generator
	}
	h := &generatorHandler{hdlr}
	return s.Handle(s.NewHandler(&Generator{h}, opts...))
}

type generatorHandler struct {
	GeneratorHandler
}

func (h *generatorHandler) Agent(ctx context.Context, in *GeneratorAgentRequest, out *GeneratorAgentResponse) error {
	return h.GeneratorHandler.Agent(ctx, in, out)
}

func (h *generatorHandler) Record(ctx context.Context, in *GeneratorRecordRequest, out *GeneratorRecordResponse) error {
	return h.GeneratorHandler.Record(ctx, in, out)
}
