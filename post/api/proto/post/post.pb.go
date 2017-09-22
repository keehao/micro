// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post/api/proto/post/post.proto

/*
Package go_micro_api_post is a generated protocol buffer package.

It is generated from these files:
	post/api/proto/post/post.proto

It has these top-level messages:
	Response
*/
package go_micro_api_post

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_api "github.com/micro/go-api/proto"
import go_micro_srv_post "github.com/hb-go/micro/post/srv/proto/post"
import go_micro_srv_post1 "github.com/hb-go/micro/post/srv/proto/comment"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Response struct {
	Post     *go_micro_srv_post.RspPost       `protobuf:"bytes,1,opt,name=post" json:"post,omitempty"`
	Comments []*go_micro_srv_post1.CommentDto `protobuf:"bytes,2,rep,name=comments" json:"comments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetPost() *go_micro_srv_post.RspPost {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Response) GetComments() []*go_micro_srv_post1.CommentDto {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "go.micro.api.post.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Post service

type PostClient interface {
	Post(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error)
}

type postClient struct {
	c           client.Client
	serviceName string
}

func NewPostClient(serviceName string, c client.Client) PostClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.api.post"
	}
	return &postClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *postClient) Post(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Post.Post", in)
	out := new(go_micro_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Post service

type PostHandler interface {
	Post(context.Context, *go_micro_api.Request, *go_micro_api.Response) error
}

func RegisterPostHandler(s server.Server, hdlr PostHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Post{hdlr}, opts...))
}

type Post struct {
	PostHandler
}

func (h *Post) Post(ctx context.Context, in *go_micro_api.Request, out *go_micro_api.Response) error {
	return h.PostHandler.Post(ctx, in, out)
}

func init() { proto.RegisterFile("post/api/proto/post/post.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xc8, 0x2f, 0x2e,
	0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0x73, 0x41, 0x84, 0x1e,
	0x98, 0x2f, 0x24, 0x98, 0x9e, 0xaf, 0x97, 0x9b, 0x99, 0x5c, 0x94, 0xaf, 0x97, 0x58, 0x90, 0xa9,
	0x07, 0x92, 0x90, 0x52, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07,
	0x4b, 0xe9, 0xa7, 0xe7, 0xeb, 0x22, 0x4c, 0x00, 0xab, 0x03, 0xb1, 0xa4, 0x4c, 0x91, 0x14, 0x66,
	0x24, 0xe9, 0xa6, 0xe7, 0x43, 0x95, 0x83, 0xed, 0x28, 0x2e, 0x2a, 0xc3, 0x6e, 0xa5, 0x94, 0x35,
	0x71, 0xda, 0x92, 0xf3, 0x73, 0x73, 0x53, 0xf3, 0x4a, 0x60, 0x34, 0x44, 0xb3, 0x52, 0x29, 0x17,
	0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1e, 0x17, 0x0b, 0x48, 0x8f, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xdc, 0x2b, 0xc5, 0x45, 0x65, 0x60, 0xaf, 0xe8,
	0x05, 0x15, 0x17, 0x04, 0xe4, 0x17, 0x97, 0x04, 0x81, 0xd5, 0x09, 0x59, 0x72, 0x71, 0x40, 0x0d,
	0x2b, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xc5, 0xa2, 0xc7, 0x19, 0xa2, 0xc4, 0xa5,
	0x24, 0x3f, 0x08, 0xae, 0xdc, 0xc8, 0x9e, 0x8b, 0x05, 0x64, 0x90, 0x90, 0x39, 0x94, 0x16, 0xd5,
	0x43, 0x09, 0xb7, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x31, 0x74, 0x61, 0x88, 0x4b,
	0x95, 0x18, 0x92, 0xd8, 0xc0, 0xce, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xcb, 0xdf,
	0xb2, 0x90, 0x01, 0x00, 0x00,
}