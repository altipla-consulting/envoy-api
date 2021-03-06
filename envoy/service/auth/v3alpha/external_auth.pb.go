// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v3alpha/external_auth.proto

package envoy_service_auth_v3alpha

import (
	context "context"
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/core"
	v3alpha "github.com/altipla-consulting/envoy-api/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	rpc "github.com/gogo/googleapis/google/rpc"
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

type CheckRequest struct {
	// The request attributes.
	Attributes           *AttributeContext `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// HTTP attributes for a denied response.
type DeniedHttpResponse struct {
	// This field allows the authorization service to send a HTTP response status
	// code to the downstream client other than 403 (Forbidden).
	Status *v3alpha.HttpStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// This field allows the authorization service to send HTTP response headers
	// to the downstream client.
	Headers []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// This field allows the authorization service to send a response body data
	// to the downstream client.
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeniedHttpResponse) Reset()         { *m = DeniedHttpResponse{} }
func (m *DeniedHttpResponse) String() string { return proto.CompactTextString(m) }
func (*DeniedHttpResponse) ProtoMessage()    {}
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{1}
}

func (m *DeniedHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeniedHttpResponse.Unmarshal(m, b)
}
func (m *DeniedHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeniedHttpResponse.Marshal(b, m, deterministic)
}
func (m *DeniedHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeniedHttpResponse.Merge(m, src)
}
func (m *DeniedHttpResponse) XXX_Size() int {
	return xxx_messageInfo_DeniedHttpResponse.Size(m)
}
func (m *DeniedHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeniedHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeniedHttpResponse proto.InternalMessageInfo

func (m *DeniedHttpResponse) GetStatus() *v3alpha.HttpStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeniedHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DeniedHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// HTTP attributes for an ok response.
type OkHttpResponse struct {
	// HTTP entity headers in addition to the original request headers. This allows the authorization
	// service to append, to add or to override headers from the original request before
	// dispatching it to the upstream. By setting `append` field to `true` in the `HeaderValueOption`,
	// the filter will append the correspondent header value to the matched request header. Note that
	// by Leaving `append` as false, the filter will either add a new header, or override an existing
	// one if there is a match.
	Headers              []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OkHttpResponse) Reset()         { *m = OkHttpResponse{} }
func (m *OkHttpResponse) String() string { return proto.CompactTextString(m) }
func (*OkHttpResponse) ProtoMessage()    {}
func (*OkHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{2}
}

func (m *OkHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkHttpResponse.Unmarshal(m, b)
}
func (m *OkHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkHttpResponse.Marshal(b, m, deterministic)
}
func (m *OkHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkHttpResponse.Merge(m, src)
}
func (m *OkHttpResponse) XXX_Size() int {
	return xxx_messageInfo_OkHttpResponse.Size(m)
}
func (m *OkHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OkHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OkHttpResponse proto.InternalMessageInfo

func (m *OkHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Intended for gRPC and Network Authorization servers `only`.
type CheckResponse struct {
	// Status `OK` allows the request. Any other status indicates the request should be denied.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// An message that contains HTTP response attributes. This message is
	// used when the authorization service needs to send custom responses to the
	// downstream client or, to modify/add request headers being dispatched to the upstream.
	//
	// Types that are valid to be assigned to HttpResponse:
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse         isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{3}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
}

type CheckResponse_DeniedResponse struct {
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,proto3,oneof"`
}

type CheckResponse_OkResponse struct {
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}

func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse() {}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (m *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CheckResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v3alpha.CheckRequest")
	proto.RegisterType((*DeniedHttpResponse)(nil), "envoy.service.auth.v3alpha.DeniedHttpResponse")
	proto.RegisterType((*OkHttpResponse)(nil), "envoy.service.auth.v3alpha.OkHttpResponse")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v3alpha.CheckResponse")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v3alpha/external_auth.proto", fileDescriptor_52e899d802958088)
}

var fileDescriptor_52e899d802958088 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x8e, 0xd3, 0x3a,
	0x14, 0xae, 0xdb, 0x3b, 0x73, 0xc1, 0xa5, 0x33, 0xe0, 0x0d, 0x55, 0x17, 0xa8, 0x44, 0x2c, 0x32,
	0x23, 0xe4, 0x48, 0x9d, 0x17, 0x60, 0x52, 0x90, 0xba, 0x00, 0xcd, 0x28, 0x08, 0x24, 0x24, 0x44,
	0xe5, 0x26, 0x47, 0x93, 0xa8, 0x51, 0x6c, 0xec, 0x93, 0xaa, 0xe5, 0x09, 0x10, 0x2f, 0xc3, 0xdb,
	0xb1, 0x60, 0x85, 0x62, 0xbb, 0x61, 0x0a, 0x22, 0x62, 0xc1, 0x2e, 0x8a, 0xbf, 0xbf, 0xf3, 0xd9,
	0x87, 0x72, 0xa8, 0x36, 0x72, 0x17, 0x19, 0xd0, 0x9b, 0x22, 0x85, 0x48, 0xd4, 0x98, 0x47, 0x9b,
	0x0b, 0x51, 0xaa, 0x5c, 0x44, 0xb0, 0x45, 0xd0, 0x95, 0x28, 0x97, 0xcd, 0x5f, 0xae, 0xb4, 0x44,
	0xc9, 0x26, 0x16, 0xcf, 0x3d, 0x9e, 0xdb, 0x13, 0x8f, 0x9f, 0x3c, 0x76, 0x5a, 0x42, 0x15, 0xad,
	0x44, 0x2a, 0x35, 0x44, 0x2b, 0x61, 0xc0, 0xd1, 0x27, 0xb3, 0x0e, 0x3b, 0x81, 0xa8, 0x8b, 0x55,
	0x8d, 0xb0, 0x4c, 0x65, 0x85, 0xb0, 0x45, 0xcf, 0x79, 0xe2, 0x38, 0xb8, 0x53, 0xd0, 0x62, 0x73,
	0x44, 0xb5, 0x34, 0x28, 0xb0, 0x36, 0x1e, 0xf5, 0xf0, 0x46, 0xca, 0x9b, 0x12, 0x22, 0xad, 0xd2,
	0xe8, 0xf0, 0x60, 0x23, 0xca, 0x22, 0x13, 0x08, 0xd1, 0xfe, 0xc3, 0x1d, 0x04, 0xef, 0xe9, 0xbd,
	0x79, 0x0e, 0xe9, 0x3a, 0x81, 0x8f, 0x35, 0x18, 0x64, 0x2f, 0x29, 0x6d, 0x23, 0x98, 0x31, 0x99,
	0x92, 0x70, 0x38, 0x7b, 0xca, 0xff, 0x3c, 0x2f, 0xbf, 0xdc, 0xa3, 0xe7, 0x2e, 0x6f, 0x72, 0x8b,
	0x1f, 0x7c, 0x25, 0x94, 0x3d, 0x87, 0xaa, 0x80, 0x6c, 0x81, 0xa8, 0x12, 0x30, 0x4a, 0x56, 0x06,
	0xd8, 0x33, 0x7a, 0xec, 0xd2, 0x79, 0x83, 0x47, 0xde, 0xa0, 0x99, 0xae, 0x15, 0x6e, 0x18, 0xaf,
	0x2d, 0x2a, 0xbe, 0xf3, 0x3d, 0x3e, 0xfa, 0x42, 0xfa, 0xf7, 0x49, 0xe2, 0x79, 0x6c, 0x4e, 0xff,
	0xcf, 0x41, 0x64, 0xa0, 0xcd, 0xb8, 0x3f, 0x1d, 0x84, 0xc3, 0xd9, 0x99, 0x97, 0x10, 0xaa, 0x68,
	0x15, 0x9a, 0xde, 0xf9, 0xc2, 0xc2, 0xde, 0x8a, 0xb2, 0x86, 0x2b, 0x85, 0x85, 0xac, 0x92, 0x3d,
	0x93, 0x31, 0xfa, 0xdf, 0x4a, 0x66, 0xbb, 0xf1, 0x60, 0x4a, 0xc2, 0xbb, 0x89, 0xfd, 0x0e, 0xde,
	0xd0, 0x93, 0xab, 0xf5, 0x41, 0xd8, 0x7f, 0x61, 0x15, 0x7c, 0x23, 0x74, 0xe4, 0x7b, 0xf6, 0xb2,
	0xe7, 0xbf, 0x74, 0xc0, 0xb8, 0xbb, 0x3b, 0xae, 0x55, 0xca, 0xdd, 0xdc, 0xed, 0xb4, 0xef, 0xe8,
	0x69, 0x66, 0x5b, 0x5c, 0x6a, 0x4f, 0x1f, 0xf7, 0x2d, 0x89, 0x77, 0xdd, 0xcc, 0xef, 0xc5, 0x2f,
	0x7a, 0xc9, 0x89, 0x13, 0x6a, 0x63, 0xbc, 0xa2, 0x43, 0xb9, 0xfe, 0x29, 0x3b, 0xb0, 0xb2, 0xe7,
	0x5d, 0xb2, 0x87, 0xf5, 0x2c, 0x7a, 0x09, 0x95, 0xed, 0x54, 0xf1, 0x29, 0x1d, 0xd9, 0x57, 0xb9,
	0x17, 0x9c, 0x49, 0x3a, 0xba, 0xac, 0x31, 0x97, 0xba, 0xf8, 0x24, 0x9a, 0x4a, 0xd8, 0x07, 0x7a,
	0x64, 0x8b, 0x60, 0x61, 0x97, 0xc9, 0xed, 0x37, 0x39, 0x39, 0xfb, 0x0b, 0xa4, 0xb3, 0x0b, 0x7a,
	0x71, 0x4c, 0xc3, 0x42, 0x3a, 0x82, 0xd2, 0x72, 0xbb, 0xeb, 0xe0, 0xc6, 0x0f, 0x5e, 0xf8, 0xe5,
	0x6e, 0x22, 0x5e, 0x37, 0xfb, 0x70, 0x4d, 0x3e, 0x13, 0xb2, 0x3a, 0xb6, 0xbb, 0x71, 0xf1, 0x23,
	0x00, 0x00, 0xff, 0xff, 0x49, 0x24, 0x9d, 0xce, 0x18, 0x04, 0x00, 0x00,
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
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v3alpha.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v3alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v3alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v3alpha/external_auth.proto",
}
