// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/metrics/v3alpha/metrics_service.proto

package envoy_service_metrics_v3alpha

import (
	context "context"
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_go "github.com/prometheus/client_model/go"
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

type StreamMetricsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricsResponse) Reset()         { *m = StreamMetricsResponse{} }
func (m *StreamMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsResponse) ProtoMessage()    {}
func (*StreamMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f243833e1cf5b45, []int{0}
}

func (m *StreamMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsResponse.Unmarshal(m, b)
}
func (m *StreamMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsResponse.Marshal(b, m, deterministic)
}
func (m *StreamMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsResponse.Merge(m, src)
}
func (m *StreamMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsResponse.Size(m)
}
func (m *StreamMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsResponse proto.InternalMessageInfo

type StreamMetricsMessage struct {
	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamMetricsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// A list of metric entries
	EnvoyMetrics         []*_go.MetricFamily `protobuf:"bytes,2,rep,name=envoy_metrics,json=envoyMetrics,proto3" json:"envoy_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StreamMetricsMessage) Reset()         { *m = StreamMetricsMessage{} }
func (m *StreamMetricsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage) ProtoMessage()    {}
func (*StreamMetricsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f243833e1cf5b45, []int{1}
}

func (m *StreamMetricsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsMessage.Unmarshal(m, b)
}
func (m *StreamMetricsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsMessage.Marshal(b, m, deterministic)
}
func (m *StreamMetricsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage.Merge(m, src)
}
func (m *StreamMetricsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsMessage.Size(m)
}
func (m *StreamMetricsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage proto.InternalMessageInfo

func (m *StreamMetricsMessage) GetIdentifier() *StreamMetricsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamMetricsMessage) GetEnvoyMetrics() []*_go.MetricFamily {
	if m != nil {
		return m.EnvoyMetrics
	}
	return nil
}

type StreamMetricsMessage_Identifier struct {
	// The node sending metrics over the stream.
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamMetricsMessage_Identifier) Reset()         { *m = StreamMetricsMessage_Identifier{} }
func (m *StreamMetricsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage_Identifier) ProtoMessage()    {}
func (*StreamMetricsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f243833e1cf5b45, []int{1, 0}
}

func (m *StreamMetricsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamMetricsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamMetricsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage_Identifier.Merge(m, src)
}
func (m *StreamMetricsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Size(m)
}
func (m *StreamMetricsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage_Identifier proto.InternalMessageInfo

func (m *StreamMetricsMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamMetricsResponse)(nil), "envoy.service.metrics.v3alpha.StreamMetricsResponse")
	proto.RegisterType((*StreamMetricsMessage)(nil), "envoy.service.metrics.v3alpha.StreamMetricsMessage")
	proto.RegisterType((*StreamMetricsMessage_Identifier)(nil), "envoy.service.metrics.v3alpha.StreamMetricsMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/metrics/v3alpha/metrics_service.proto", fileDescriptor_3f243833e1cf5b45)
}

var fileDescriptor_3f243833e1cf5b45 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdd, 0xfa, 0x07, 0xd9, 0x5a, 0x91, 0xa8, 0xb4, 0x04, 0x85, 0xda, 0x53, 0x41, 0xd8,
	0x40, 0xe3, 0xc9, 0x83, 0x87, 0x1c, 0xac, 0x1e, 0x2a, 0x25, 0xbd, 0x5b, 0xb6, 0xc9, 0x68, 0x07,
	0x92, 0x6c, 0xd8, 0x5d, 0x83, 0xbd, 0x78, 0x56, 0x4f, 0x7e, 0x5e, 0x4f, 0xd2, 0xec, 0x26, 0x5a,
	0x90, 0x82, 0xde, 0x32, 0x99, 0x79, 0xbf, 0x79, 0x33, 0xb3, 0xd4, 0x87, 0xac, 0x10, 0x0b, 0x4f,
	0x81, 0x2c, 0x30, 0x02, 0x2f, 0x05, 0x2d, 0x31, 0x52, 0x5e, 0xe1, 0xf3, 0x24, 0x9f, 0xf3, 0x2a,
	0x9e, 0xda, 0x3c, 0xcb, 0xa5, 0xd0, 0xc2, 0x39, 0x2d, 0x45, 0xac, 0xfa, 0x69, 0x8b, 0x98, 0x15,
	0xb9, 0x67, 0x86, 0xc9, 0x73, 0xac, 0x39, 0x91, 0x90, 0xe0, 0xcd, 0xb8, 0xb2, 0x04, 0xb7, 0x55,
	0x69, 0x4c, 0xd8, 0x2e, 0x78, 0x82, 0x31, 0xd7, 0xe0, 0x55, 0x1f, 0x26, 0xd1, 0x6b, 0xd3, 0xe3,
	0x89, 0x96, 0xc0, 0xd3, 0x91, 0xa9, 0x0f, 0x41, 0xe5, 0x22, 0x53, 0xd0, 0x7b, 0x6b, 0xd0, 0xa3,
	0x95, 0xcc, 0x08, 0x94, 0xe2, 0x8f, 0xe0, 0xdc, 0x53, 0x8a, 0x31, 0x64, 0x1a, 0x1f, 0x10, 0x64,
	0x87, 0x74, 0x49, 0xbf, 0x39, 0xb8, 0x62, 0x6b, 0x0d, 0xb3, 0xdf, 0x40, 0xec, 0xb6, 0xa6, 0x84,
	0x3f, 0x88, 0xce, 0x90, 0xb6, 0x4a, 0xd8, 0xd4, 0x42, 0x3a, 0x8d, 0xee, 0x66, 0xbf, 0x39, 0xe8,
	0x31, 0x14, 0x4b, 0xcf, 0x29, 0xe8, 0x39, 0x3c, 0x29, 0x16, 0x25, 0x08, 0x99, 0x66, 0x86, 0x79,
	0xcd, 0x53, 0x4c, 0x16, 0xe1, 0x5e, 0x29, 0xb4, 0x6d, 0xdc, 0x1b, 0x4a, 0xbf, 0x5b, 0x38, 0x97,
	0x74, 0x2b, 0x13, 0x31, 0x58, 0xc3, 0x27, 0xd6, 0x30, 0xcf, 0xb1, 0x36, 0xb9, 0x5c, 0x21, 0xbb,
	0x13, 0x31, 0x04, 0xbb, 0x9f, 0xc1, 0xf6, 0x3b, 0x69, 0x1c, 0x90, 0xb0, 0xd4, 0x0c, 0x3e, 0x08,
	0xdd, 0xb7, 0xd4, 0x89, 0x99, 0xd0, 0x79, 0xa1, 0xad, 0x95, 0xa1, 0x1c, 0xff, 0x1f, 0x2b, 0x70,
	0x2f, 0xfe, 0x22, 0xaa, 0x4f, 0xb3, 0xd1, 0x27, 0xc1, 0x90, 0x9e, 0xa3, 0x30, 0xea, 0x5c, 0x8a,
	0xe7, 0xc5, 0x7a, 0x50, 0x70, 0xb8, 0x6a, 0x7f, 0xbc, 0xbc, 0xfd, 0x98, 0xbc, 0x12, 0x32, 0xdb,
	0x29, 0xdf, 0x81, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x47, 0x54, 0x8d, 0x8d, 0xa8, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/envoy.service.metrics.v3alpha.MetricsService/StreamMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceStreamMetricsClient{stream}
	return x, nil
}

type MetricsService_StreamMetricsClient interface {
	Send(*StreamMetricsMessage) error
	CloseAndRecv() (*StreamMetricsResponse, error)
	grpc.ClientStream
}

type metricsServiceStreamMetricsClient struct {
	grpc.ClientStream
}

func (x *metricsServiceStreamMetricsClient) Send(m *StreamMetricsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsClient) CloseAndRecv() (*StreamMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(MetricsService_StreamMetricsServer) error
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_StreamMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).StreamMetrics(&metricsServiceStreamMetricsServer{stream})
}

type MetricsService_StreamMetricsServer interface {
	SendAndClose(*StreamMetricsResponse) error
	Recv() (*StreamMetricsMessage, error)
	grpc.ServerStream
}

type metricsServiceStreamMetricsServer struct {
	grpc.ServerStream
}

func (x *metricsServiceStreamMetricsServer) SendAndClose(m *StreamMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsServer) Recv() (*StreamMetricsMessage, error) {
	m := new(StreamMetricsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.metrics.v3alpha.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetrics",
			Handler:       _MetricsService_StreamMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/metrics/v3alpha/metrics_service.proto",
}