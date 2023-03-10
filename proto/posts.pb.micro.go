// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/posts.proto

package posts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Posts service

func NewPostsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Posts service

type PostsService interface {
	Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error)
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Posts_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Posts_PingPongService, error)
}

type postsService struct {
	c    client.Client
	name string
}

func NewPostsService(name string, c client.Client) PostsService {
	return &postsService{
		c:    c,
		name: name,
	}
}

func (c *postsService) Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error) {
	req := c.c.NewRequest(c.name, "Posts.Save", in)
	out := new(SaveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Posts.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Posts_StreamService, error) {
	req := c.c.NewRequest(c.name, "Posts.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &postsServiceStream{stream}, nil
}

type Posts_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type postsServiceStream struct {
	stream client.Stream
}

func (x *postsServiceStream) Close() error {
	return x.stream.Close()
}

func (x *postsServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postsServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postsServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postsServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postsService) PingPong(ctx context.Context, opts ...client.CallOption) (Posts_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Posts.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &postsServicePingPong{stream}, nil
}

type Posts_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type postsServicePingPong struct {
	stream client.Stream
}

func (x *postsServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *postsServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *postsServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postsServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postsServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *postsServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Posts service

type PostsHandler interface {
	Save(context.Context, *SaveRequest, *SaveResponse) error
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Posts_StreamStream) error
	PingPong(context.Context, Posts_PingPongStream) error
}

func RegisterPostsHandler(s server.Server, hdlr PostsHandler, opts ...server.HandlerOption) error {
	type posts interface {
		Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Posts struct {
		posts
	}
	h := &postsHandler{hdlr}
	return s.Handle(s.NewHandler(&Posts{h}, opts...))
}

type postsHandler struct {
	PostsHandler
}

func (h *postsHandler) Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error {
	return h.PostsHandler.Save(ctx, in, out)
}

func (h *postsHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.PostsHandler.Call(ctx, in, out)
}

func (h *postsHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PostsHandler.Stream(ctx, m, &postsStreamStream{stream})
}

type Posts_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type postsStreamStream struct {
	stream server.Stream
}

func (x *postsStreamStream) Close() error {
	return x.stream.Close()
}

func (x *postsStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postsStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postsStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postsStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *postsHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.PostsHandler.PingPong(ctx, &postsPingPongStream{stream})
}

type Posts_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type postsPingPongStream struct {
	stream server.Stream
}

func (x *postsPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *postsPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postsPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postsPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postsPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *postsPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
