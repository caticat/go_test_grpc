// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package prototest is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	A
	B
*/
package prototest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type A struct {
	Data             *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *A) Reset()                    { *m = A{} }
func (m *A) String() string            { return proto.CompactTextString(m) }
func (*A) ProtoMessage()               {}
func (*A) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *A) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

type B struct {
	Data             *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *B) Reset()                    { *m = B{} }
func (m *B) String() string            { return proto.CompactTextString(m) }
func (*B) ProtoMessage()               {}
func (*B) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *B) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*A)(nil), "prototest.A")
	proto.RegisterType((*B)(nil), "prototest.B")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	AB(ctx context.Context, in *A, opts ...grpc.CallOption) (*B, error)
	ABB(ctx context.Context, in *A, opts ...grpc.CallOption) (TestService_ABBClient, error)
	AAB(ctx context.Context, opts ...grpc.CallOption) (TestService_AABClient, error)
	AABB(ctx context.Context, opts ...grpc.CallOption) (TestService_AABBClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) AB(ctx context.Context, in *A, opts ...grpc.CallOption) (*B, error) {
	out := new(B)
	err := grpc.Invoke(ctx, "/prototest.TestService/AB", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ABB(ctx context.Context, in *A, opts ...grpc.CallOption) (TestService_ABBClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/prototest.TestService/ABB", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceABBClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ABBClient interface {
	Recv() (*B, error)
	grpc.ClientStream
}

type testServiceABBClient struct {
	grpc.ClientStream
}

func (x *testServiceABBClient) Recv() (*B, error) {
	m := new(B)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) AAB(ctx context.Context, opts ...grpc.CallOption) (TestService_AABClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/prototest.TestService/AAB", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceAABClient{stream}
	return x, nil
}

type TestService_AABClient interface {
	Send(*A) error
	CloseAndRecv() (*B, error)
	grpc.ClientStream
}

type testServiceAABClient struct {
	grpc.ClientStream
}

func (x *testServiceAABClient) Send(m *A) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceAABClient) CloseAndRecv() (*B, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(B)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) AABB(ctx context.Context, opts ...grpc.CallOption) (TestService_AABBClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/prototest.TestService/AABB", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceAABBClient{stream}
	return x, nil
}

type TestService_AABBClient interface {
	Send(*A) error
	Recv() (*B, error)
	grpc.ClientStream
}

type testServiceAABBClient struct {
	grpc.ClientStream
}

func (x *testServiceAABBClient) Send(m *A) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceAABBClient) Recv() (*B, error) {
	m := new(B)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	AB(context.Context, *A) (*B, error)
	ABB(*A, TestService_ABBServer) error
	AAB(TestService_AABServer) error
	AABB(TestService_AABBServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_AB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(A)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).AB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototest.TestService/AB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).AB(ctx, req.(*A))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ABB_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(A)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ABB(m, &testServiceABBServer{stream})
}

type TestService_ABBServer interface {
	Send(*B) error
	grpc.ServerStream
}

type testServiceABBServer struct {
	grpc.ServerStream
}

func (x *testServiceABBServer) Send(m *B) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_AAB_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).AAB(&testServiceAABServer{stream})
}

type TestService_AABServer interface {
	SendAndClose(*B) error
	Recv() (*A, error)
	grpc.ServerStream
}

type testServiceAABServer struct {
	grpc.ServerStream
}

func (x *testServiceAABServer) SendAndClose(m *B) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceAABServer) Recv() (*A, error) {
	m := new(A)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_AABB_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).AABB(&testServiceAABBServer{stream})
}

type TestService_AABBServer interface {
	Send(*B) error
	Recv() (*A, error)
	grpc.ServerStream
}

type testServiceAABBServer struct {
	grpc.ServerStream
}

func (x *testServiceAABBServer) Send(m *B) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceAABBServer) Recv() (*A, error) {
	m := new(A)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prototest.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AB",
			Handler:    _TestService_AB_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ABB",
			Handler:       _TestService_ABB_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AAB",
			Handler:       _TestService_AAB_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AABB",
			Handler:       _TestService_AABB_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0x20, 0x01, 0x25, 0x71, 0x2e, 0x46,
	0x47, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0x1b, 0x24, 0xe1, 0x84, 0x4d, 0xc2, 0x68, 0x25, 0x23, 0x17, 0x77, 0x48, 0x6a, 0x71, 0x49,
	0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x12, 0x17, 0x93, 0xa3, 0x93, 0x10, 0x8f, 0x1e,
	0xdc, 0x4c, 0x3d, 0x47, 0x29, 0x64, 0x9e, 0x93, 0x12, 0x83, 0x90, 0x2a, 0x17, 0xb3, 0xa3, 0x13,
	0x01, 0x45, 0x06, 0x8c, 0x60, 0x65, 0x84, 0xcc, 0xd2, 0x60, 0x14, 0xd2, 0xe0, 0x62, 0x71, 0x24,
	0x68, 0x9c, 0x06, 0xa3, 0x01, 0x23, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x1d, 0x72, 0xde, 0xf5,
	0x00, 0x00, 0x00,
}