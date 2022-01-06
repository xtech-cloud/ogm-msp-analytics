// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/analytics/query.proto

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

// Api Endpoints for Query service

func NewQueryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Query service

type QueryService interface {
	// 个体
	Agent(ctx context.Context, in *QueryAgentRequest, opts ...client.CallOption) (*QueryAgentResponse, error)
	// 事件
	Event(ctx context.Context, in *QueryEventRequest, opts ...client.CallOption) (*QueryEventResponse, error)
}

type queryService struct {
	c    client.Client
	name string
}

func NewQueryService(name string, c client.Client) QueryService {
	return &queryService{
		c:    c,
		name: name,
	}
}

func (c *queryService) Agent(ctx context.Context, in *QueryAgentRequest, opts ...client.CallOption) (*QueryAgentResponse, error) {
	req := c.c.NewRequest(c.name, "Query.Agent", in)
	out := new(QueryAgentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryService) Event(ctx context.Context, in *QueryEventRequest, opts ...client.CallOption) (*QueryEventResponse, error) {
	req := c.c.NewRequest(c.name, "Query.Event", in)
	out := new(QueryEventResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Query service

type QueryHandler interface {
	// 个体
	Agent(context.Context, *QueryAgentRequest, *QueryAgentResponse) error
	// 事件
	Event(context.Context, *QueryEventRequest, *QueryEventResponse) error
}

func RegisterQueryHandler(s server.Server, hdlr QueryHandler, opts ...server.HandlerOption) error {
	type query interface {
		Agent(ctx context.Context, in *QueryAgentRequest, out *QueryAgentResponse) error
		Event(ctx context.Context, in *QueryEventRequest, out *QueryEventResponse) error
	}
	type Query struct {
		query
	}
	h := &queryHandler{hdlr}
	return s.Handle(s.NewHandler(&Query{h}, opts...))
}

type queryHandler struct {
	QueryHandler
}

func (h *queryHandler) Agent(ctx context.Context, in *QueryAgentRequest, out *QueryAgentResponse) error {
	return h.QueryHandler.Agent(ctx, in, out)
}

func (h *queryHandler) Event(ctx context.Context, in *QueryEventRequest, out *QueryEventResponse) error {
	return h.QueryHandler.Event(ctx, in, out)
}
