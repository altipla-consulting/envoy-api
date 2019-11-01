// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/trace/v2/trace_service.proto

package envoy_service_trace_v2

import (
	context "context"
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
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

type StreamTracesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTracesResponse) Reset()         { *m = StreamTracesResponse{} }
func (m *StreamTracesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTracesResponse) ProtoMessage()    {}
func (*StreamTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{0}
}

func (m *StreamTracesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesResponse.Unmarshal(m, b)
}
func (m *StreamTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesResponse.Marshal(b, m, deterministic)
}
func (m *StreamTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesResponse.Merge(m, src)
}
func (m *StreamTracesResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTracesResponse.Size(m)
}
func (m *StreamTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesResponse proto.InternalMessageInfo

type StreamTracesMessage struct {
	// Identifier data effectively is a structured metadata.
	// As a performance optimization this will only be sent in the first message
	// on the stream.
	Identifier *StreamTracesMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// A list of Span entries
	Spans                []*v1.Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTracesMessage) Reset()         { *m = StreamTracesMessage{} }
func (m *StreamTracesMessage) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage) ProtoMessage()    {}
func (*StreamTracesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{1}
}

func (m *StreamTracesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage.Unmarshal(m, b)
}
func (m *StreamTracesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage.Merge(m, src)
}
func (m *StreamTracesMessage) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage.Size(m)
}
func (m *StreamTracesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage proto.InternalMessageInfo

func (m *StreamTracesMessage) GetIdentifier() *StreamTracesMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTracesMessage) GetSpans() []*v1.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type StreamTracesMessage_Identifier struct {
	// The node sending the access log messages over the stream.
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTracesMessage_Identifier) Reset()         { *m = StreamTracesMessage_Identifier{} }
func (m *StreamTracesMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage_Identifier) ProtoMessage()    {}
func (*StreamTracesMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{1, 0}
}

func (m *StreamTracesMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamTracesMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage_Identifier.Merge(m, src)
}
func (m *StreamTracesMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Size(m)
}
func (m *StreamTracesMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage_Identifier proto.InternalMessageInfo

func (m *StreamTracesMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamTracesResponse)(nil), "envoy.service.trace.v2.StreamTracesResponse")
	proto.RegisterType((*StreamTracesMessage)(nil), "envoy.service.trace.v2.StreamTracesMessage")
	proto.RegisterType((*StreamTracesMessage_Identifier)(nil), "envoy.service.trace.v2.StreamTracesMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/trace/v2/trace_service.proto", fileDescriptor_6feca8f22ae39b94)
}

var fileDescriptor_6feca8f22ae39b94 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0xdf, 0xf4, 0xbd, 0x3e, 0x64, 0xda, 0x85, 0x46, 0x69, 0x4b, 0x28, 0x58, 0x8a, 0x42,
	0x51, 0x99, 0xd0, 0x88, 0x2e, 0x5d, 0xc4, 0x55, 0x17, 0x4a, 0x49, 0xc5, 0xad, 0x4c, 0x93, 0x6b,
	0x19, 0x68, 0xe7, 0x0e, 0x99, 0x31, 0xd8, 0x85, 0x7b, 0xf1, 0x9f, 0xfa, 0x17, 0x5c, 0x49, 0x66,
	0x92, 0x9a, 0x45, 0x05, 0xdd, 0x5d, 0xee, 0x3d, 0xe7, 0xf0, 0xdd, 0x43, 0x4f, 0x40, 0xe6, 0xb8,
	0x0e, 0x34, 0x64, 0xb9, 0x48, 0x20, 0x30, 0x19, 0x4f, 0x20, 0xc8, 0x43, 0x37, 0x3c, 0x94, 0x6b,
	0xa6, 0x32, 0x34, 0xe8, 0x75, 0xac, 0x96, 0x55, 0x4b, 0x2b, 0x61, 0x79, 0xe8, 0xf7, 0x5d, 0x06,
	0x57, 0xa2, 0x70, 0x26, 0x98, 0x41, 0x30, 0xe7, 0xba, 0x74, 0xf9, 0xfd, 0x05, 0xe2, 0x62, 0x09,
	0xf6, 0xcc, 0xa5, 0x44, 0xc3, 0x8d, 0x40, 0xa9, 0xcb, 0xeb, 0x31, 0x2a, 0x90, 0x09, 0x48, 0xfd,
	0xa4, 0x03, 0xbb, 0xa9, 0x10, 0xc6, 0x6e, 0x28, 0x65, 0xdd, 0x9c, 0x2f, 0x45, 0xca, 0x0d, 0x04,
	0xd5, 0xe0, 0x0e, 0xc3, 0x0e, 0x3d, 0x98, 0x99, 0x0c, 0xf8, 0xea, 0xae, 0x50, 0xeb, 0x18, 0xb4,
	0x42, 0xa9, 0x61, 0xf8, 0x4e, 0xe8, 0x7e, 0xfd, 0x70, 0x03, 0x5a, 0xf3, 0x05, 0x78, 0xf7, 0x94,
	0x8a, 0x14, 0xa4, 0x11, 0x8f, 0x02, 0xb2, 0x1e, 0x19, 0x90, 0x51, 0x2b, 0xbc, 0x64, 0xdb, 0x1f,
	0x63, 0x5b, 0x02, 0xd8, 0x64, 0xe3, 0x8e, 0x6b, 0x49, 0xde, 0x05, 0x6d, 0x6a, 0xc5, 0xa5, 0xee,
	0x35, 0x06, 0x7f, 0x47, 0xad, 0xf0, 0x90, 0x7d, 0xfd, 0xe5, 0x48, 0xab, 0xd4, 0x31, 0x9b, 0x29,
	0x2e, 0x63, 0xa7, 0xf6, 0xaf, 0x29, 0x9d, 0xd4, 0x43, 0xfe, 0x49, 0x4c, 0xa1, 0xc4, 0xea, 0x96,
	0x58, 0x5c, 0x89, 0x02, 0xa6, 0xe8, 0x95, 0xdd, 0x62, 0x0a, 0xd1, 0xce, 0x47, 0xd4, 0x7c, 0x23,
	0x8d, 0x5d, 0x12, 0x5b, 0x79, 0xf8, 0x42, 0xdb, 0x96, 0x71, 0xe6, 0xf8, 0xbd, 0x15, 0x6d, 0xd7,
	0xc9, 0xbd, 0xd3, 0x5f, 0xfc, 0xe7, 0x9f, 0xfd, 0x44, 0xbc, 0xa9, 0xf9, 0xcf, 0x88, 0x44, 0x57,
	0xf4, 0x48, 0xa0, 0x73, 0xa9, 0x0c, 0x9f, 0xd7, 0xdf, 0x04, 0x44, 0x7b, 0x75, 0xc8, 0x69, 0xd1,
	0xc9, 0x94, 0xbc, 0x12, 0x32, 0xff, 0x6f, 0xfb, 0x39, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x9b, 0x8b, 0xc0, 0x8b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/envoy.service.trace.v2.TraceService/StreamTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(TraceService_StreamTracesServer) error
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v2.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v2/trace_service.proto",
}
