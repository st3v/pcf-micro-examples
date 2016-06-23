// Code generated by protoc-gen-go.
// source: greeter.proto
// DO NOT EDIT!

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	greeter.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type sayClient struct {
	c           client.Client
	serviceName string
}

func NewSayClient(serviceName string, c client.Client) SayClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.greeter"
	}
	return &sayClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Say.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler) {
	s.Handle(s.NewHandler(&Say{hdlr}))
}

type Say struct {
	SayHandler
}

func (h *Say) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0xcf, 0xd7, 0xcb, 0xcd,
	0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0xca, 0x29, 0xc9, 0x72, 0xb1, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x6a, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x60, 0x5d, 0x99, 0x79, 0xe9, 0x12, 0x4c, 0x60, 0x35, 0x70,
	0xbe, 0x91, 0x3f, 0x17, 0x73, 0x70, 0x62, 0xa5, 0x90, 0x07, 0x17, 0xab, 0x47, 0x6a, 0x4e, 0x4e,
	0xbe, 0x90, 0xac, 0x1e, 0x36, 0xdb, 0xf4, 0xa0, 0x56, 0x49, 0xc9, 0xe1, 0x92, 0x86, 0x58, 0xa5,
	0xc4, 0x90, 0xc4, 0x06, 0x76, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0x55, 0x67, 0x57,
	0xc5, 0x00, 0x00, 0x00,
}
