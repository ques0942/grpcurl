// Code generated by protoc-gen-go.
// source: internal/testdata/v2/v2.proto
// DO NOT EDIT!

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	internal/testdata/v2/v2.proto

It has these top-level messages:
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import grpcurl_test "github.com/kazegusuri/grpcurl/internal/testdata"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Echo(ctx context.Context, in *grpcurl_test.EchoMessage, opts ...grpc.CallOption) (*grpcurl_test.EchoMessage, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *grpcurl_test.EchoMessage, opts ...grpc.CallOption) (*grpcurl_test.EchoMessage, error) {
	out := new(grpcurl_test.EchoMessage)
	err := grpc.Invoke(ctx, "/grpcurl.test.v2.Echo/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Echo(context.Context, *grpcurl_test.EchoMessage) (*grpcurl_test.EchoMessage, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpcurl_test.EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcurl.test.v2.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*grpcurl_test.EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcurl.test.v2.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/testdata/v2/v2.proto",
}

func init() { proto.RegisterFile("internal/testdata/v2/v2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x49, 0x2d, 0x2e, 0x49, 0x49, 0x2c, 0x49, 0xd4, 0x2f, 0x33,
	0xd2, 0x2f, 0x33, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x2f, 0x2a, 0x48, 0x2e,
	0x2d, 0xca, 0xd1, 0x03, 0xc9, 0xea, 0x95, 0x19, 0x49, 0xa9, 0x60, 0xaa, 0x4f, 0x4d, 0xce, 0xc8,
	0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x85, 0x68, 0x33, 0x72, 0xe3, 0x62, 0x71, 0x4d,
	0xce, 0xc8, 0x17, 0xb2, 0x83, 0xd2, 0x92, 0x7a, 0x28, 0xe6, 0x80, 0xc4, 0x7c, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0xa5, 0x70, 0x4b, 0x29, 0x31, 0x38, 0x99, 0x47, 0x99, 0xa6, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x67, 0x27, 0x56, 0xa5, 0xa6, 0x97, 0x16, 0x97, 0x16,
	0x65, 0xea, 0x43, 0xf5, 0xe8, 0x63, 0x73, 0xbd, 0x75, 0x99, 0x51, 0x12, 0x1b, 0xd8, 0x1d, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xc2, 0xdf, 0x70, 0xdf, 0x00, 0x00, 0x00,
}
