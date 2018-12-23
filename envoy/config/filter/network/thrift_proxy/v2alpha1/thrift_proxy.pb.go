// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

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
	return fileDescriptor_thrift_proxy_606da48aff7eb26f, []int{0}
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
	return fileDescriptor_thrift_proxy_606da48aff7eb26f, []int{1}
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
	return fileDescriptor_thrift_proxy_606da48aff7eb26f, []int{0}
}
func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (dst *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(dst, src)
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
	// * :ref:`envoy.filters.thrift.rate_limit <config_thrift_filters_rate_limit>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated. See the supported
	// filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_Config
	//	*ThriftFilter_TypedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_thrift_proxy_606da48aff7eb26f, []int{1}
}
func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (dst *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(dst, src)
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

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_Config struct {
	Config *types.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *types.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ThriftFilter_Config) isThriftFilter_ConfigType() {}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ThriftFilter) GetConfig() *types.Struct {
	if x, ok := m.GetConfigType().(*ThriftFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ThriftFilter) GetTypedConfig() *types.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ThriftFilter_OneofMarshaler, _ThriftFilter_OneofUnmarshaler, _ThriftFilter_OneofSizer, []interface{}{
		(*ThriftFilter_Config)(nil),
		(*ThriftFilter_TypedConfig)(nil),
	}
}

func _ThriftFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ThriftFilter)
	// config_type
	switch x := m.ConfigType.(type) {
	case *ThriftFilter_Config:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *ThriftFilter_TypedConfig:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypedConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ThriftFilter.ConfigType has unexpected type %T", x)
	}
	return nil
}

func _ThriftFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ThriftFilter)
	switch tag {
	case 2: // config_type.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.Struct)
		err := b.DecodeMessage(msg)
		m.ConfigType = &ThriftFilter_Config{msg}
		return true, err
	case 3: // config_type.typed_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.Any)
		err := b.DecodeMessage(msg)
		m.ConfigType = &ThriftFilter_TypedConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ThriftFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ThriftFilter)
	// config_type
	switch x := m.ConfigType.(type) {
	case *ThriftFilter_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ThriftFilter_TypedConfig:
		s := proto.Size(x.TypedConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
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
	return fileDescriptor_thrift_proxy_606da48aff7eb26f, []int{2}
}
func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (dst *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(dst, src)
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
	proto.RegisterType((*ThriftProxy)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProtocolOptions")
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType", ProtocolType_name, ProtocolType_value)
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto", fileDescriptor_thrift_proxy_606da48aff7eb26f)
}

var fileDescriptor_thrift_proxy_606da48aff7eb26f = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb3, 0xb6, 0x5b, 0xda, 0x71, 0x12, 0x99, 0x55, 0x51, 0x43, 0x04, 0x28, 0xea, 0x29,
	0xea, 0xc1, 0x56, 0xc3, 0x89, 0x03, 0x02, 0xe7, 0xa3, 0x4a, 0xa5, 0x36, 0xb6, 0xb6, 0xae, 0xf8,
	0x90, 0x50, 0xe4, 0xa6, 0xb6, 0x63, 0x08, 0x5e, 0x6b, 0xbd, 0x09, 0xcd, 0x95, 0x13, 0xef, 0xc1,
	0x8d, 0x47, 0xe0, 0xc4, 0xa3, 0x70, 0x45, 0xe2, 0x21, 0x90, 0x77, 0x9d, 0x36, 0x21, 0xa7, 0x46,
	0x82, 0xdb, 0xce, 0xd7, 0x6f, 0xfe, 0x3b, 0x33, 0xd0, 0x0d, 0x92, 0x19, 0x9d, 0x5b, 0x23, 0x9a,
	0x84, 0x71, 0x64, 0x85, 0xf1, 0x84, 0x07, 0xcc, 0x4a, 0x02, 0xfe, 0x89, 0xb2, 0x0f, 0x16, 0x1f,
	0xb3, 0x38, 0xe4, 0xc3, 0x94, 0xd1, 0xeb, 0xb9, 0x35, 0x6b, 0xf9, 0x93, 0x74, 0xec, 0x1f, 0xad,
	0x78, 0xcd, 0x94, 0x51, 0x4e, 0xf1, 0x91, 0xa0, 0x98, 0x92, 0x62, 0x4a, 0x8a, 0x59, 0x50, 0xcc,
	0x95, 0xfc, 0x05, 0xa5, 0xfe, 0xfc, 0xee, 0x8d, 0x19, 0x9d, 0xf2, 0x40, 0x76, 0xac, 0x3f, 0x8c,
	0x28, 0x8d, 0x26, 0x81, 0x25, 0xac, 0xcb, 0x69, 0x68, 0xf9, 0x49, 0x21, 0xa6, 0xfe, 0xe8, 0xef,
	0x50, 0xc6, 0xd9, 0x74, 0xc4, 0x8b, 0xe8, 0xfe, 0xcc, 0x9f, 0xc4, 0x57, 0x3e, 0x0f, 0xac, 0xc5,
	0xa3, 0x08, 0xec, 0x45, 0x34, 0xa2, 0xe2, 0x69, 0xe5, 0x2f, 0xe9, 0x3d, 0xf8, 0xa9, 0x82, 0xee,
	0x09, 0x35, 0x6e, 0x2e, 0x06, 0xbf, 0x87, 0x5d, 0xce, 0xfc, 0x24, 0x4b, 0x29, 0xe3, 0x35, 0xa5,
	0x81, 0x9a, 0xd5, 0xd6, 0x4b, 0xf3, 0xce, 0xbf, 0x37, 0xbd, 0x05, 0xc3, 0x9b, 0xa7, 0x41, 0x1b,
	0xbe, 0xff, 0xfa, 0xa1, 0x6e, 0x7d, 0x46, 0x8a, 0x81, 0xc8, 0x2d, 0x1e, 0x47, 0xb0, 0x23, 0x44,
	0x8c, 0xe8, 0xa4, 0xa6, 0x8a, 0x56, 0x2f, 0x36, 0x68, 0xe5, 0x16, 0x88, 0xb5, 0x4e, 0x37, 0x70,
	0x7c, 0x08, 0x7a, 0xc6, 0xfd, 0xbc, 0x32, 0x08, 0xe3, 0xeb, 0x1a, 0x6a, 0xa0, 0xe6, 0x6e, 0x7b,
	0x37, 0x4f, 0xd5, 0x98, 0xd2, 0x40, 0x04, 0xf2, 0xa8, 0x2b, 0x82, 0x78, 0x0c, 0x65, 0xb1, 0x87,
	0xa1, 0xd4, 0x50, 0xd3, 0x1a, 0xa8, 0xa9, 0xb7, 0x7a, 0x1b, 0x08, 0x23, 0x39, 0xa6, 0x23, 0x0a,
	0xa6, 0xcc, 0xe7, 0x31, 0x4d, 0x88, 0xce, 0x6e, 0x7d, 0x38, 0x84, 0x6a, 0x51, 0x28, 0x71, 0x59,
	0x6d, 0xab, 0xa1, 0x36, 0xf5, 0x8d, 0x86, 0x20, 0x57, 0x78, 0x2c, 0x52, 0x49, 0x85, 0x2f, 0x59,
	0xd9, 0xc1, 0x37, 0x04, 0xe5, 0xe5, 0x38, 0x7e, 0x0c, 0x5a, 0xe2, 0x7f, 0x0c, 0xd6, 0xe7, 0x20,
	0xdc, 0xf8, 0x08, 0xb6, 0x8b, 0xbf, 0x2b, 0xe2, 0xef, 0xfb, 0xa6, 0x3c, 0x38, 0x73, 0x71, 0x70,
	0xe6, 0xb9, 0x38, 0xb8, 0x7e, 0x89, 0x14, 0x89, 0xf8, 0x19, 0x94, 0xf9, 0x3c, 0x0d, 0xae, 0x16,
	0x43, 0x53, 0x45, 0xe1, 0xde, 0x5a, 0xa1, 0x9d, 0xcc, 0xfb, 0x25, 0xa2, 0x8b, 0x5c, 0x39, 0x85,
	0x76, 0x05, 0x74, 0x59, 0x34, 0xcc, 0xbd, 0x07, 0xbf, 0x11, 0x3c, 0xb8, 0xb9, 0x47, 0xb1, 0x3d,
	0x27, 0xcd, 0x47, 0x97, 0xad, 0x5e, 0x26, 0xfa, 0x7f, 0x97, 0xa9, 0xfc, 0xc3, 0xcb, 0x3c, 0x74,
	0xa0, 0xb2, 0x22, 0x08, 0x63, 0xa8, 0xda, 0x17, 0x9e, 0x33, 0xf4, 0x88, 0x3d, 0x38, 0x77, 0x1d,
	0xe2, 0x19, 0x25, 0x0c, 0xb0, 0x7d, 0x4c, 0xec, 0xb3, 0x5e, 0xd7, 0x40, 0xb8, 0x0c, 0x3b, 0x17,
	0x83, 0xc2, 0x52, 0xf2, 0x48, 0xbf, 0x67, 0x77, 0x7b, 0xc4, 0x50, 0xeb, 0xda, 0x97, 0xaf, 0x4f,
	0x4a, 0x87, 0xef, 0xa0, 0xbc, 0xdc, 0x16, 0xdf, 0x87, 0x8a, 0xe0, 0xb9, 0xc4, 0xf1, 0x9c, 0x8e,
	0x73, 0x2a, 0x71, 0xed, 0x93, 0x81, 0x4d, 0xde, 0x18, 0x08, 0x57, 0x01, 0x4e, 0xed, 0xd7, 0xc3,
	0xc2, 0x56, 0xb0, 0x0e, 0xf7, 0x3a, 0xce, 0x99, 0x6b, 0x77, 0x3c, 0x43, 0xcd, 0x0d, 0xef, 0xd5,
	0x89, 0xe7, 0xf5, 0x88, 0xa1, 0x49, 0x7c, 0x5b, 0x7b, 0xab, 0xcc, 0x5a, 0x97, 0xdb, 0x42, 0xff,
	0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xaa, 0x3e, 0xa7, 0x5d, 0x05, 0x00, 0x00,
}
