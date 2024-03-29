// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test/test.proto

package go_micro_srv_test

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 用户信息
type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_79a3935f3fa86eee, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_79a3935f3fa86eee, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_79a3935f3fa86eee, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_79a3935f3fa86eee, []int{3}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.test.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.test.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.test.Response")
	proto.RegisterType((*Error)(nil), "go.micro.srv.test.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TestService service

type TestServiceClient interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type testServiceClient struct {
	c           client.Client
	serviceName string
}

func NewTestServiceClient(serviceName string, c client.Client) TestServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.test"
	}
	return &testServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *testServiceClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "TestService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "TestService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "TestService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceHandler interface {
	Create(context.Context, *Request, *Response) error
	Get(context.Context, *Request, *Response) error
	GetAll(context.Context, *Request, *Response) error
}

func RegisterTestServiceHandler(s server.Server, hdlr TestServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&TestService{hdlr}, opts...))
}

type TestService struct {
	TestServiceHandler
}

func (h *TestService) Create(ctx context.Context, in *Request, out *Response) error {
	return h.TestServiceHandler.Create(ctx, in, out)
}

func (h *TestService) Get(ctx context.Context, in *Request, out *Response) error {
	return h.TestServiceHandler.Get(ctx, in, out)
}

func (h *TestService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.TestServiceHandler.GetAll(ctx, in, out)
}

func init() { proto.RegisterFile("proto/test/test.proto", fileDescriptor_test_79a3935f3fa86eee) }

var fileDescriptor_test_79a3935f3fa86eee = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xcd, 0xdf, 0xd6, 0x29, 0x08, 0x0e, 0x8a, 0x4b, 0xbd, 0x94, 0x9c, 0x04, 0x31, 0x4a,
	0x05, 0x6f, 0x82, 0x52, 0xa4, 0xf7, 0x55, 0x1f, 0x20, 0x26, 0x83, 0x2c, 0x34, 0xd9, 0xb8, 0xb3,
	0xad, 0xf8, 0x26, 0xbe, 0x9a, 0x6f, 0x23, 0x99, 0x56, 0x11, 0x0c, 0x82, 0xf4, 0x92, 0xcc, 0xf7,
	0xcd, 0x2f, 0x33, 0xc3, 0x47, 0xe0, 0xb0, 0x75, 0xd6, 0xdb, 0x73, 0x4f, 0xec, 0xe5, 0x91, 0x8b,
	0xc6, 0xfd, 0x67, 0x9b, 0xd7, 0xa6, 0x74, 0x36, 0x67, 0xb7, 0xca, 0xbb, 0x46, 0xb6, 0x82, 0xf8,
	0x91, 0xc9, 0xe1, 0x1e, 0x84, 0xa6, 0x52, 0xc1, 0x24, 0x38, 0xd9, 0xd5, 0xa1, 0xa9, 0x10, 0x21,
	0x6e, 0x8a, 0x9a, 0x54, 0x28, 0x8e, 0xd4, 0xa8, 0x60, 0x50, 0xda, 0xba, 0x2d, 0x9a, 0x37, 0x15,
	0x89, 0xfd, 0x25, 0xf1, 0x00, 0x12, 0xaa, 0x0b, 0xb3, 0x50, 0xb1, 0xf8, 0x6b, 0x81, 0x63, 0x18,
	0xb6, 0x05, 0xf3, 0xab, 0x75, 0x95, 0x4a, 0xa4, 0xf1, 0xad, 0xb3, 0x2b, 0x18, 0x68, 0x7a, 0x59,
	0x12, 0x7b, 0x3c, 0x85, 0x78, 0xc9, 0xe4, 0x64, 0xf9, 0x68, 0x7a, 0x94, 0xff, 0x3a, 0x32, 0xef,
	0x2e, 0xd4, 0x02, 0x65, 0xef, 0x01, 0x0c, 0x35, 0x71, 0x6b, 0x1b, 0xa6, 0x7f, 0x7d, 0x89, 0x67,
	0x90, 0x74, 0x6f, 0x56, 0xe1, 0x24, 0xfa, 0x8b, 0x5e, 0x53, 0x78, 0x01, 0x29, 0x39, 0x67, 0x1d,
	0xab, 0x48, 0x78, 0xd5, 0xc3, 0xdf, 0x75, 0x80, 0xde, 0x70, 0xd9, 0x35, 0x24, 0x62, 0x74, 0xd9,
	0x95, 0xb6, 0x22, 0x39, 0x2b, 0xd1, 0x52, 0xe3, 0x04, 0x46, 0x15, 0x71, 0xe9, 0x4c, 0xeb, 0x8d,
	0x6d, 0x36, 0xb1, 0xfe, 0xb4, 0xa6, 0x1f, 0x01, 0x8c, 0x1e, 0x88, 0xfd, 0x3d, 0xb9, 0x95, 0x29,
	0x09, 0x67, 0x90, 0xce, 0x1c, 0x15, 0x9e, 0x70, 0xdc, 0xb3, 0x7a, 0x13, 0xde, 0xf8, 0xb8, 0xb7,
	0xb7, 0xce, 0x27, 0xdb, 0xc1, 0x1b, 0x88, 0xe6, 0xe4, 0xb7, 0x99, 0x30, 0x83, 0x74, 0x4e, 0xfe,
	0x76, 0xb1, 0xd8, 0x62, 0xc8, 0x53, 0x2a, 0xff, 0xdf, 0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x49, 0x76, 0x6b, 0x3c, 0x98, 0x02, 0x00, 0x00,
}
