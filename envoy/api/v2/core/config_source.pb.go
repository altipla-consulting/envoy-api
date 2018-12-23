// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/config_source.proto

package core

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

// APIs may be fetched via either REST or gRPC.
type ApiConfigSource_ApiType int32

const (
	// REST-JSON legacy corresponds to the v1 API.
	ApiConfigSource_REST_LEGACY ApiConfigSource_ApiType = 0 // Deprecated: Do not use.
	// REST-JSON v2 API. The `canonical JSON encoding
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_ for
	// the v2 protos is used.
	ApiConfigSource_REST ApiConfigSource_ApiType = 1
	// gRPC v2 API.
	ApiConfigSource_GRPC ApiConfigSource_ApiType = 2
)

var ApiConfigSource_ApiType_name = map[int32]string{
	0: "REST_LEGACY",
	1: "REST",
	2: "GRPC",
}
var ApiConfigSource_ApiType_value = map[string]int32{
	"REST_LEGACY": 0,
	"REST":        1,
	"GRPC":        2,
}

func (x ApiConfigSource_ApiType) String() string {
	return proto.EnumName(ApiConfigSource_ApiType_name, int32(x))
}
func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_source_aed092169d5b6614, []int{0, 0}
}

// API configuration source. This identifies the API type and cluster that Envoy
// will use to fetch an xDS API.
type ApiConfigSource struct {
	ApiType ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	// Cluster names should be used only with REST_LEGACY/REST. If > 1
	// cluster is defined, clusters will be cycled through if any kind of failure
	// occurs.
	//
	// .. note::
	//
	//  The cluster with name ``cluster_name`` must be statically defined and its
	//  type must not be ``EDS``.
	ClusterNames []string `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	// Multiple gRPC services be provided for GRPC. If > 1 cluster is defined,
	// services will be cycled through if any kind of failure occurs.
	GrpcServices []*GrpcService `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	// For REST APIs, the delay between successive polls.
	RefreshDelay *types.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	// For REST APIs, the request timeout. If not set, a default value of 1s will be used.
	RequestTimeout *types.Duration `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// For GRPC APIs, the rate limit settings. If present, discovery requests made by Envoy will be
	// rate limited.
	RateLimitSettings    *RateLimitSettings `protobuf:"bytes,6,opt,name=rate_limit_settings,json=rateLimitSettings,proto3" json:"rate_limit_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_aed092169d5b6614, []int{0}
}
func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiConfigSource.Unmarshal(m, b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
}
func (dst *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(dst, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return xxx_messageInfo_ApiConfigSource.Size(m)
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if m != nil {
		return m.ApiType
	}
	return ApiConfigSource_REST_LEGACY
}

func (m *ApiConfigSource) GetClusterNames() []string {
	if m != nil {
		return m.ClusterNames
	}
	return nil
}

func (m *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

func (m *ApiConfigSource) GetRefreshDelay() *types.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ApiConfigSource) GetRequestTimeout() *types.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *ApiConfigSource) GetRateLimitSettings() *RateLimitSettings {
	if m != nil {
		return m.RateLimitSettings
	}
	return nil
}

// Aggregated Discovery Service (ADS) options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that ADS is to be used.
type AggregatedConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatedConfigSource) Reset()         { *m = AggregatedConfigSource{} }
func (m *AggregatedConfigSource) String() string { return proto.CompactTextString(m) }
func (*AggregatedConfigSource) ProtoMessage()    {}
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_aed092169d5b6614, []int{1}
}
func (m *AggregatedConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatedConfigSource.Unmarshal(m, b)
}
func (m *AggregatedConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatedConfigSource.Marshal(b, m, deterministic)
}
func (dst *AggregatedConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedConfigSource.Merge(dst, src)
}
func (m *AggregatedConfigSource) XXX_Size() int {
	return xxx_messageInfo_AggregatedConfigSource.Size(m)
}
func (m *AggregatedConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedConfigSource proto.InternalMessageInfo

// Rate Limit settings to be applied for discovery requests made by Envoy.
type RateLimitSettings struct {
	// Maximum number of tokens to be used for rate limting discovery request calls. If not set, a
	// default value of 100 will be used.
	MaxTokens *types.UInt32Value `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// Rate at which tokens will be filled per second. If not set, a default fill rate of 10 tokens
	// per second will be used.
	FillRate             *types.DoubleValue `protobuf:"bytes,2,opt,name=fill_rate,json=fillRate,proto3" json:"fill_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RateLimitSettings) Reset()         { *m = RateLimitSettings{} }
func (m *RateLimitSettings) String() string { return proto.CompactTextString(m) }
func (*RateLimitSettings) ProtoMessage()    {}
func (*RateLimitSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_aed092169d5b6614, []int{2}
}
func (m *RateLimitSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitSettings.Unmarshal(m, b)
}
func (m *RateLimitSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitSettings.Marshal(b, m, deterministic)
}
func (dst *RateLimitSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitSettings.Merge(dst, src)
}
func (m *RateLimitSettings) XXX_Size() int {
	return xxx_messageInfo_RateLimitSettings.Size(m)
}
func (m *RateLimitSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitSettings proto.InternalMessageInfo

func (m *RateLimitSettings) GetMaxTokens() *types.UInt32Value {
	if m != nil {
		return m.MaxTokens
	}
	return nil
}

func (m *RateLimitSettings) GetFillRate() *types.DoubleValue {
	if m != nil {
		return m.FillRate
	}
	return nil
}

// Configuration for :ref:`listeners <config_listeners>`, :ref:`clusters
// <config_cluster_manager>`, :ref:`routes
// <envoy_api_msg_RouteConfiguration>`, :ref:`endpoints
// <arch_overview_service_discovery>` etc. may either be sourced from the
// filesystem or from an xDS API source. Filesystem configs are watched with
// inotify for updates.
type ConfigSource struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_aed092169d5b6614, []int{3}
}
func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSource.Unmarshal(m, b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
}
func (dst *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(dst, src)
}
func (m *ConfigSource) XXX_Size() int {
	return xxx_messageInfo_ConfigSource.Size(m)
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *ConfigSource) GetPath() string {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (m *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigSource_OneofMarshaler, _ConfigSource_OneofUnmarshaler, _ConfigSource_OneofSizer, []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
	}
}

func _ConfigSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Path)
	case *ConfigSource_ApiConfigSource:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApiConfigSource); err != nil {
			return err
		}
	case *ConfigSource_Ads:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ads); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConfigSource.ConfigSourceSpecifier has unexpected type %T", x)
	}
	return nil
}

func _ConfigSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigSource)
	switch tag {
	case 1: // config_source_specifier.path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConfigSourceSpecifier = &ConfigSource_Path{x}
		return true, err
	case 2: // config_source_specifier.api_config_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApiConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_ApiConfigSource{msg}
		return true, err
	case 3: // config_source_specifier.ads
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AggregatedConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_Ads{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Path)))
		n += len(x.Path)
	case *ConfigSource_ApiConfigSource:
		s := proto.Size(x.ApiConfigSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigSource_Ads:
		s := proto.Size(x.Ads)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ApiConfigSource)(nil), "envoy.api.v2.core.ApiConfigSource")
	proto.RegisterType((*AggregatedConfigSource)(nil), "envoy.api.v2.core.AggregatedConfigSource")
	proto.RegisterType((*RateLimitSettings)(nil), "envoy.api.v2.core.RateLimitSettings")
	proto.RegisterType((*ConfigSource)(nil), "envoy.api.v2.core.ConfigSource")
	proto.RegisterEnum("envoy.api.v2.core.ApiConfigSource_ApiType", ApiConfigSource_ApiType_name, ApiConfigSource_ApiType_value)
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/config_source.proto", fileDescriptor_config_source_aed092169d5b6614)
}

var fileDescriptor_config_source_aed092169d5b6614 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x4e, 0xdb, 0x4a,
	0x14, 0x8d, 0xe3, 0x00, 0xc9, 0x24, 0x40, 0x32, 0xa0, 0x87, 0x1f, 0x7a, 0xe2, 0xe5, 0xe5, 0x51,
	0x29, 0x65, 0x61, 0x4b, 0x66, 0x59, 0x75, 0x41, 0x12, 0x04, 0x55, 0x51, 0x45, 0x27, 0x69, 0xa5,
	0xae, 0xac, 0xc1, 0xb9, 0x31, 0xa3, 0x3a, 0x1e, 0x77, 0x66, 0x9c, 0x92, 0x6d, 0xbf, 0xa2, 0x9b,
	0xee, 0xab, 0x7e, 0x41, 0xd5, 0x15, 0x7f, 0xd1, 0x65, 0xa5, 0xee, 0xd8, 0xf7, 0x03, 0xaa, 0x99,
	0x18, 0x15, 0x48, 0x10, 0x5e, 0xdd, 0x7b, 0xe7, 0x9c, 0x33, 0xf7, 0x1e, 0xdf, 0x41, 0x8f, 0x20,
	0x99, 0xf0, 0xa9, 0x47, 0x53, 0xe6, 0x4d, 0x7c, 0x2f, 0xe4, 0x02, 0xbc, 0x90, 0x27, 0x23, 0x16,
	0x05, 0x92, 0x67, 0x22, 0x04, 0x37, 0x15, 0x5c, 0x71, 0xdc, 0x30, 0x30, 0x97, 0xa6, 0xcc, 0x9d,
	0xf8, 0xae, 0x86, 0x6d, 0xef, 0xce, 0x33, 0x23, 0x91, 0x86, 0x81, 0x04, 0x31, 0x61, 0xd7, 0xc4,
	0xed, 0x9d, 0x88, 0xf3, 0x28, 0x06, 0xcf, 0x64, 0x67, 0xd9, 0xc8, 0x1b, 0x66, 0x82, 0x2a, 0xc6,
	0x93, 0xfb, 0xce, 0xdf, 0x0b, 0x9a, 0xa6, 0x20, 0x64, 0x7e, 0xbe, 0x35, 0xa1, 0x31, 0x1b, 0x52,
	0x05, 0xde, 0x75, 0x90, 0x1f, 0x6c, 0x46, 0x3c, 0xe2, 0x26, 0xf4, 0x74, 0x34, 0xab, 0xb6, 0x7e,
	0xd9, 0x68, 0xfd, 0x20, 0x65, 0x5d, 0x33, 0x42, 0xdf, 0x4c, 0x80, 0x5f, 0xa2, 0x32, 0x4d, 0x59,
	0xa0, 0xa6, 0x29, 0x38, 0x56, 0xd3, 0x6a, 0xaf, 0xf9, 0x7b, 0xee, 0xdc, 0x38, 0xee, 0x1d, 0x96,
	0xce, 0x07, 0xd3, 0x14, 0x3a, 0xe8, 0xdb, 0xd5, 0xa5, 0xbd, 0xf4, 0xc1, 0x2a, 0xd6, 0x2d, 0xb2,
	0x42, 0x67, 0x45, 0xfc, 0x3f, 0x5a, 0x0d, 0xe3, 0x4c, 0x2a, 0x10, 0x41, 0x42, 0xc7, 0x20, 0x9d,
	0x62, 0xd3, 0x6e, 0x57, 0x48, 0x2d, 0x2f, 0xbe, 0xd0, 0x35, 0xdc, 0x45, 0xab, 0x37, 0x0d, 0x91,
	0x4e, 0xa9, 0x69, 0xb7, 0xab, 0xfe, 0xce, 0x82, 0xcb, 0x8f, 0x44, 0x1a, 0xf6, 0x67, 0x30, 0x52,
	0x8b, 0xfe, 0x24, 0x12, 0xf7, 0xd0, 0xaa, 0x80, 0x91, 0x00, 0x79, 0x1e, 0x0c, 0x21, 0xa6, 0x53,
	0xc7, 0x6e, 0x5a, 0xed, 0xaa, 0xff, 0xb7, 0x3b, 0xf3, 0xcd, 0xbd, 0xf6, 0xcd, 0xed, 0xe5, 0xbe,
	0x76, 0x4a, 0x1f, 0x7f, 0xfc, 0x6b, 0x91, 0x5a, 0xce, 0xea, 0x69, 0x12, 0x1e, 0xa0, 0x75, 0x01,
	0xef, 0x32, 0x90, 0x2a, 0x50, 0x6c, 0x0c, 0x3c, 0x53, 0xce, 0xd2, 0x43, 0x3a, 0x75, 0xad, 0xa3,
	0x87, 0x5f, 0xf9, 0x62, 0x95, 0xf6, 0x8a, 0xe5, 0x02, 0x59, 0xcb, 0x35, 0x06, 0x33, 0x09, 0x3c,
	0x40, 0x1b, 0x82, 0x2a, 0x08, 0x62, 0x36, 0x66, 0x2a, 0x90, 0xa0, 0x14, 0x4b, 0x22, 0xe9, 0x2c,
	0x1b, 0xe5, 0xdd, 0x05, 0x63, 0x12, 0xaa, 0xe0, 0x44, 0x83, 0xfb, 0x39, 0x96, 0x34, 0xc4, 0xdd,
	0x52, 0xcb, 0x47, 0x2b, 0xb9, 0xf7, 0x78, 0x03, 0x55, 0xc9, 0x61, 0x7f, 0x10, 0x9c, 0x1c, 0x1e,
	0x1d, 0x74, 0xdf, 0xd4, 0x0b, 0xdb, 0xc5, 0xb2, 0x85, 0xcb, 0xa8, 0xa4, 0x8b, 0x75, 0x13, 0x1d,
	0x91, 0xd3, 0x6e, 0xbd, 0xd8, 0x72, 0xd0, 0x5f, 0x07, 0x51, 0x24, 0x20, 0xa2, 0x0a, 0x86, 0x37,
	0x7f, 0x63, 0xeb, 0x93, 0x85, 0x1a, 0x73, 0xd7, 0xe2, 0x27, 0x08, 0x8d, 0xe9, 0x45, 0xa0, 0xf8,
	0x5b, 0x48, 0xa4, 0x59, 0x8a, 0xaa, 0xff, 0xcf, 0x9c, 0x15, 0xaf, 0x9e, 0x25, 0x6a, 0xdf, 0x7f,
	0x4d, 0xe3, 0x0c, 0x48, 0x65, 0x4c, 0x2f, 0x06, 0x06, 0x8e, 0x9f, 0xa3, 0xca, 0x88, 0xc5, 0x71,
	0xa0, 0x5b, 0x77, 0x8a, 0xf7, 0x70, 0x7b, 0x3c, 0x3b, 0x8b, 0xc1, 0x70, 0x3b, 0x75, 0xed, 0x62,
	0x15, 0x57, 0xfe, 0x2b, 0xe4, 0x1f, 0x29, 0x6b, 0x01, 0xdd, 0x56, 0xeb, 0xbb, 0x85, 0x6a, 0xb7,
	0xb6, 0x75, 0x13, 0x95, 0x52, 0xaa, 0xce, 0x4d, 0x53, 0x95, 0xe3, 0x02, 0x31, 0x19, 0x3e, 0x45,
	0x0d, 0xbd, 0xc3, 0xb7, 0x9e, 0x66, 0x7e, 0x77, 0xeb, 0xe1, 0x65, 0x3e, 0x2e, 0x90, 0x75, 0x7a,
	0xe7, 0x55, 0x3c, 0x45, 0x36, 0x1d, 0xca, 0x7c, 0x9d, 0x1e, 0x2f, 0xd2, 0x58, 0x68, 0xe8, 0x71,
	0x81, 0x68, 0x5e, 0xa7, 0x89, 0xb6, 0x6e, 0x35, 0x13, 0xc8, 0x14, 0x42, 0x36, 0x62, 0x20, 0xf0,
	0xd2, 0xd7, 0xab, 0x4b, 0xdb, 0xea, 0x94, 0x3e, 0xff, 0xdc, 0xb1, 0xce, 0x96, 0x8d, 0x23, 0xfb,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x74, 0xa1, 0xa3, 0x68, 0x04, 0x00, 0x00,
}
