// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/external_auth.proto

package v2alpha

import (
	context "context"
	fmt "fmt"
	v2 "github.com/altipla-consulting/envoy-api/envoy/service/auth/v2"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/external_auth.proto", fileDescriptor_878c0ddb0c43de8d)
}

var fileDescriptor_878c0ddb0c43de8d = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0xb1, 0x0e, 0x82, 0x30,
	0x14, 0x45, 0xc3, 0xa0, 0x26, 0x4d, 0x5c, 0x18, 0x59, 0x75, 0xd0, 0xe5, 0x35, 0xc1, 0x2f, 0x10,
	0xfc, 0x00, 0xc2, 0xe8, 0x62, 0x2a, 0x79, 0x49, 0x1b, 0x49, 0x5f, 0x6d, 0x1f, 0x04, 0xfc, 0x02,
	0x3f, 0xdb, 0x50, 0x19, 0xd1, 0xf5, 0xde, 0x93, 0x7b, 0x8f, 0x00, 0xb4, 0x3d, 0x8d, 0x32, 0xa0,
	0xef, 0x4d, 0x83, 0x52, 0x75, 0xac, 0x65, 0x9f, 0xab, 0xd6, 0x69, 0x25, 0x71, 0x60, 0xf4, 0x56,
	0xb5, 0xb7, 0x29, 0x05, 0xe7, 0x89, 0x29, 0xcd, 0x22, 0x0f, 0x33, 0x0f, 0xb1, 0x99, 0xf9, 0xec,
	0xb8, 0xb8, 0xb5, 0x34, 0x93, 0x37, 0x62, 0x7b, 0xee, 0x58, 0x93, 0x37, 0x2f, 0xc5, 0x86, 0x6c,
	0x5a, 0x8b, 0x55, 0xa9, 0xb1, 0x79, 0xa4, 0x3b, 0x58, 0x7c, 0x80, 0xd8, 0xd6, 0xf8, 0xec, 0x30,
	0x70, 0xb6, 0xff, 0x0f, 0x05, 0x47, 0x36, 0x60, 0x71, 0x11, 0x07, 0x43, 0x5f, 0xd2, 0x79, 0x1a,
	0x46, 0xf8, 0xed, 0x5e, 0x88, 0x12, 0x3d, 0x87, 0x6a, 0x92, 0xab, 0x92, 0xeb, 0x66, 0x8e, 0xdf,
	0x49, 0x72, 0x5f, 0x47, 0xe3, 0xd3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x92, 0x3c, 0xbb, 0xcf, 0x2a,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(ctx context.Context, in *v2.CheckRequest, opts ...grpc.CallOption) (*v2.CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *v2.CheckRequest, opts ...grpc.CallOption) (*v2.CheckResponse, error) {
	out := new(v2.CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v2alpha.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *v2.CheckRequest) (*v2.CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*v2.CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2alpha/external_auth.proto",
}
