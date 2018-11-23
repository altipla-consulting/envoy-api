// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto

package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Thrift transport types supported by Envoy.
type TransportType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which transport to use.
	// For upstream connections, the Thrift proxy will use same transport as the downstream
	// connection.
	TransportType_AUTO_TRANSPORT TransportType = 0
	// The Thrift proxy will use the Thrift framed transport.
	TransportType_FRAMED TransportType = 1
	// The Thrift proxy will use the Thrift unframed transport.
	TransportType_UNFRAMED TransportType = 2
	// The Thrift proxy will assume the client is using the Thrift header transport.
	TransportType_HEADER TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}

// Thrift Protocol types supported by Envoy.
type ProtocolType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which protocol to use.
	// Note that the older, non-strict (or lax) binary protocol is not included in automatic protocol
	// detection. For upstream connections, the Thrift proxy will use the same protocol as the
	// downstream connection.
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	// The Thrift proxy will use the Thrift binary protocol.
	ProtocolType_BINARY ProtocolType = 1
	// The Thrift proxy will use Thrift non-strict binary protocol.
	ProtocolType_LAX_BINARY ProtocolType = 2
	// The Thrift proxy will use the Thrift compact protocol.
	ProtocolType_COMPACT ProtocolType = 3
	// The Thrift proxy will use the Thrift "Twitter" protocol implemented by the finagle library.
	ProtocolType_TWITTER ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{1}
}

// [#comment:next free field: 6]
type ThriftProxy struct {
	// Supplies the type of transport that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_TRANSPORT<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.TransportType.AUTO_TRANSPORT>`.
	Transport TransportType `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_PROTOCOL<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.ProtocolType.AUTO_PROTOCOL>`.
	Protocol ProtocolType `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType" json:"protocol,omitempty"`
	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// A list of individual Thrift filters that make up the filter chain for requests made to the
	// Thrift proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no thrift_filters are specified, a default Thrift router filter
	// (`envoy.filters.thrift.router`) is used.
	ThriftFilters        []*ThriftFilter `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}

func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

// ThriftFilter configures a Thrift filter.
// [#comment:next free field: 3]
type ThriftFilter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(zuercher): Auto generate the following list]
	// * :ref:`envoy.filters.thrift.router <config_thrift_filters_router>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	Config               *types.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{1}
}

func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return xxx_messageInfo_ThriftFilter.Size(m)
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ThriftFilter) GetConfig() *types.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// ThriftProtocolOptions specifies Thrift upstream protocol options. This object is used in
// in :ref:`extension_protocol_options<envoy_api_field_Cluster.extension_protocol_options>`, keyed
// by the name `envoy.filters.network.thrift_proxy`.
// [#comment:next free field: 3]
type ThriftProtocolOptions struct {
	// Supplies the type of transport that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_TRANSPORT<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.TransportType.AUTO_TRANSPORT>`,
	// which is the default, causes the proxy to use the same transport as the downstream connection.
	Transport TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_PROTOCOL<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.ProtocolType.AUTO_PROTOCOL>`,
	// which is the default, causes the proxy to use the same protocol as the downstream connection.
	Protocol             ProtocolType `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{2}
}

func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_ThriftProtocolOptions.Size(m)
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto", fileDescriptor_e8fab7646d88fc90)
}

var fileDescriptor_e8fab7646d88fc90 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x3b, 0xb6, 0x1b, 0x9a, 0xe3, 0x24, 0x32, 0x23, 0x50, 0xa3, 0x08, 0x50, 0x94, 0x55,
	0x94, 0x85, 0xad, 0x86, 0x35, 0x02, 0xe7, 0x52, 0x51, 0xa9, 0x8d, 0xad, 0xe9, 0x54, 0x5c, 0x24,
	0x88, 0xdc, 0x60, 0x3b, 0x06, 0xe3, 0xb1, 0xc6, 0x93, 0xd0, 0x6c, 0x59, 0xf1, 0x1e, 0xbc, 0x05,
	0x2b, 0x1e, 0x85, 0x2d, 0x12, 0x0f, 0x81, 0x3c, 0x76, 0xda, 0x44, 0x5d, 0x35, 0x12, 0xec, 0x7c,
	0x2e, 0xf3, 0xfd, 0xe7, 0x1c, 0xff, 0x30, 0xf2, 0x93, 0x25, 0x5b, 0x59, 0x33, 0x96, 0x04, 0x51,
	0x68, 0x05, 0x51, 0x2c, 0x7c, 0x6e, 0x25, 0xbe, 0xf8, 0xc2, 0xf8, 0x27, 0x4b, 0xcc, 0x79, 0x14,
	0x88, 0x69, 0xca, 0xd9, 0xd5, 0xca, 0x5a, 0xf6, 0xbd, 0x38, 0x9d, 0x7b, 0x47, 0x5b, 0x59, 0x33,
	0xe5, 0x4c, 0x30, 0x7c, 0x24, 0x29, 0x66, 0x41, 0x31, 0x0b, 0x8a, 0x59, 0x52, 0xcc, 0xad, 0xfe,
	0x35, 0xa5, 0xf5, 0xec, 0xee, 0xc2, 0x9c, 0x2d, 0x84, 0x5f, 0x28, 0xb6, 0x1e, 0x85, 0x8c, 0x85,
	0xb1, 0x6f, 0xc9, 0xe8, 0x72, 0x11, 0x58, 0x99, 0xe0, 0x8b, 0x99, 0x28, 0xab, 0x87, 0x4b, 0x2f,
	0x8e, 0x3e, 0x78, 0xc2, 0xb7, 0xd6, 0x1f, 0x65, 0xe1, 0x41, 0xc8, 0x42, 0x26, 0x3f, 0xad, 0xfc,
	0xab, 0xc8, 0x76, 0x7e, 0xa9, 0xa0, 0x53, 0x29, 0xe9, 0xe6, 0x8a, 0xf8, 0x23, 0x54, 0x05, 0xf7,
	0x92, 0x2c, 0x65, 0x5c, 0x34, 0x95, 0x36, 0xea, 0x36, 0xfa, 0x2f, 0xcc, 0x3b, 0xaf, 0x68, 0xd2,
	0x35, 0x83, 0xae, 0x52, 0x7f, 0x00, 0x3f, 0x7e, 0xff, 0x54, 0xf7, 0xbf, 0x22, 0xc5, 0x40, 0xe4,
	0x06, 0x8f, 0x43, 0x38, 0x90, 0x43, 0xcc, 0x58, 0xdc, 0x54, 0xa5, 0xd4, 0xf3, 0x1d, 0xa4, 0xdc,
	0x12, 0x71, 0x4b, 0xe9, 0x1a, 0x8e, 0x7b, 0xa0, 0x67, 0xc2, 0xcb, 0x5f, 0xfa, 0x41, 0x74, 0xd5,
	0x44, 0x6d, 0xd4, 0xad, 0x0e, 0xaa, 0x79, 0xab, 0xc6, 0x95, 0x36, 0x22, 0x90, 0x57, 0x5d, 0x59,
	0xc4, 0x73, 0xa8, 0xc9, 0x63, 0x4f, 0x8b, 0x19, 0x9a, 0x5a, 0x1b, 0x75, 0xf5, 0xfe, 0x78, 0x87,
	0xc1, 0x48, 0x8e, 0x19, 0xca, 0x07, 0x0b, 0xee, 0x89, 0x88, 0x25, 0x44, 0xe7, 0x37, 0x39, 0x1c,
	0x40, 0xa3, 0x7c, 0x58, 0xe0, 0xb2, 0xe6, 0x7e, 0x5b, 0xed, 0xea, 0x3b, 0x1d, 0xa1, 0xf8, 0x85,
	0xc7, 0xb2, 0x95, 0xd4, 0xc5, 0x46, 0x94, 0x75, 0xde, 0x43, 0x6d, 0xb3, 0x8c, 0x1f, 0x83, 0x96,
	0x78, 0x9f, 0xfd, 0xdb, 0x67, 0x90, 0x69, 0x6c, 0x41, 0xa5, 0x5c, 0x5d, 0x91, 0xab, 0x1f, 0x9a,
	0x85, 0xdf, 0xcc, 0xb5, 0xdf, 0xcc, 0x73, 0xe9, 0x37, 0x52, 0xb6, 0x75, 0xfe, 0x20, 0x78, 0x78,
	0x6d, 0x21, 0x79, 0x70, 0x27, 0xcd, 0xb7, 0xcd, 0xb6, 0xcd, 0x84, 0xfe, 0x9f, 0x99, 0x94, 0x7f,
	0x68, 0xa6, 0x9e, 0x03, 0xf5, 0xad, 0x81, 0x30, 0x86, 0x86, 0x7d, 0x41, 0x9d, 0x29, 0x25, 0xf6,
	0xe4, 0xdc, 0x75, 0x08, 0x35, 0xf6, 0x30, 0x40, 0xe5, 0x98, 0xd8, 0x67, 0xe3, 0x91, 0x81, 0x70,
	0x0d, 0x0e, 0x2e, 0x26, 0x65, 0xa4, 0xe4, 0x95, 0x97, 0x63, 0x7b, 0x34, 0x26, 0x86, 0xda, 0xd2,
	0xbe, 0x7d, 0x7f, 0xb2, 0xd7, 0x7b, 0x07, 0xb5, 0x4d, 0x59, 0x7c, 0x1f, 0xea, 0x92, 0xe7, 0x12,
	0x87, 0x3a, 0x43, 0xe7, 0xb4, 0xc0, 0x0d, 0x4e, 0x26, 0x36, 0x79, 0x63, 0x20, 0xdc, 0x00, 0x38,
	0xb5, 0x5f, 0x4f, 0xcb, 0x58, 0xc1, 0x3a, 0xdc, 0x1b, 0x3a, 0x67, 0xae, 0x3d, 0xa4, 0x86, 0x9a,
	0x07, 0xf4, 0xd5, 0x09, 0xa5, 0x63, 0x62, 0x68, 0x05, 0x7e, 0xa0, 0xbd, 0x55, 0x96, 0xfd, 0xcb,
	0x8a, 0x9c, 0xff, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x63, 0x7c, 0x96, 0xf5, 0x04,
	0x00, 0x00,
}
