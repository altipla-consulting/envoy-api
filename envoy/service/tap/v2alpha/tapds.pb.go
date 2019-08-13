// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/tapds.proto

package envoy_service_tap_v2alpha

import (
	context "context"
	fmt "fmt"
	v2 "github.com/altipla-consulting/envoy-api/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/googleapis/google/api"
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

// [#not-implemented-hide:] A tap resource is essentially a tap configuration with a name
// The filter TapDS config references this name.
type TapResource struct {
	// The name of the tap configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Tap config to apply
	Config               *TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eef591ebf2a5317, []int{0}
}

func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapResource.Unmarshal(m, b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return xxx_messageInfo_TapResource.Size(m)
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v2alpha.TapResource")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v2alpha/tapds.proto", fileDescriptor_4eef591ebf2a5317)
}

var fileDescriptor_4eef591ebf2a5317 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xef, 0x34, 0x84, 0x9b, 0x3b, 0x2c, 0xb8, 0xe9, 0x5d, 0x5c, 0x40, 0x82, 0x04, 0x51,
	0x89, 0x8b, 0xa9, 0xa9, 0x0b, 0x13, 0xe2, 0xaa, 0x12, 0xd7, 0xa4, 0xb0, 0x71, 0x65, 0x0e, 0x65,
	0x84, 0x49, 0xda, 0x39, 0x63, 0x67, 0x68, 0x60, 0xeb, 0x2b, 0xf8, 0x5e, 0x6e, 0x7c, 0x05, 0x9f,
	0xc2, 0x95, 0xe9, 0x1f, 0x88, 0x68, 0xea, 0xca, 0xdd, 0x49, 0xcf, 0xf7, 0xfd, 0x7a, 0xce, 0xf9,
	0x86, 0x1e, 0x73, 0x99, 0xe0, 0xc6, 0xd1, 0x3c, 0x4e, 0x44, 0xc0, 0x1d, 0x03, 0xca, 0x49, 0x5c,
	0x08, 0xd5, 0x12, 0xd2, 0x7a, 0xae, 0x99, 0x8a, 0xd1, 0xa0, 0xdd, 0xcc, 0x64, 0xac, 0x90, 0x31,
	0x03, 0x8a, 0x15, 0xb2, 0x56, 0x3b, 0x27, 0x80, 0x12, 0x4e, 0xe2, 0x3a, 0x73, 0xa1, 0x03, 0x4c,
	0x78, 0xbc, 0xc9, 0x8d, 0xad, 0x93, 0x72, 0x7e, 0x80, 0x51, 0x84, 0xb2, 0xd0, 0xfd, 0x4f, 0x20,
	0x14, 0x73, 0x30, 0xdc, 0xd9, 0x16, 0x45, 0xa3, 0xbd, 0x40, 0x5c, 0x84, 0x3c, 0xe3, 0x83, 0x94,
	0x68, 0xc0, 0x08, 0x94, 0xc5, 0x5c, 0xbd, 0x25, 0xad, 0x4d, 0x41, 0xf9, 0x5c, 0xe3, 0x2a, 0x0e,
	0xb8, 0x7d, 0x40, 0x2b, 0x12, 0x22, 0xde, 0x20, 0x5d, 0x32, 0xf8, 0xe3, 0xfd, 0x7e, 0xf3, 0x2a,
	0xb1, 0xd5, 0x25, 0x7e, 0xf6, 0xd1, 0xbe, 0xa2, 0xd5, 0x00, 0xe5, 0xbd, 0x58, 0x34, 0xac, 0x2e,
	0x19, 0xd4, 0xdc, 0x3e, 0x2b, 0x5d, 0x8a, 0x4d, 0x41, 0x5d, 0x67, 0x5a, 0xbf, 0xf0, 0xb8, 0xcf,
	0x16, 0xfd, 0x37, 0x05, 0x35, 0xda, 0xee, 0x37, 0xc9, 0x5d, 0xf6, 0x2d, 0xfd, 0x3b, 0x31, 0x31,
	0x87, 0x68, 0x67, 0xd1, 0x76, 0xa7, 0x20, 0x83, 0x12, 0x2c, 0x71, 0xd9, 0xce, 0xe3, 0xf3, 0x87,
	0x15, 0xd7, 0xa6, 0x75, 0x58, 0xda, 0xd7, 0x0a, 0xa5, 0xe6, 0xbd, 0x5f, 0x03, 0x72, 0x4e, 0xec,
	0x19, 0xad, 0x8f, 0x78, 0x68, 0xe0, 0x03, 0xf9, 0xe8, 0x93, 0x33, 0x6d, 0x7f, 0xc1, 0xf7, 0xbf,
	0x17, 0xed, 0xfd, 0x63, 0x4d, 0xeb, 0x37, 0xdc, 0x04, 0xcb, 0x9f, 0x9c, 0xbe, 0xff, 0xf8, 0xf2,
	0xfa, 0x64, 0x75, 0x7a, 0xcd, 0xbd, 0x07, 0x31, 0x34, 0xa0, 0xee, 0xf2, 0x63, 0xea, 0x21, 0x39,
	0xf3, 0x2e, 0xe9, 0xa9, 0xc0, 0x1c, 0xa5, 0x62, 0x5c, 0x6f, 0xca, 0xd3, 0xf0, 0x68, 0x7a, 0x78,
	0x3d, 0x4e, 0x13, 0x1f, 0x93, 0x59, 0x35, 0x8b, 0xfe, 0xe2, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x19,
	0xee, 0x49, 0xd1, 0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v2alpha.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v2alpha.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tapds.proto",
}
