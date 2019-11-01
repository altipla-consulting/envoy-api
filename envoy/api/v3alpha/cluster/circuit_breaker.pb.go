// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/cluster/circuit_breaker.proto

package envoy_api_v3alpha_cluster

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_api.v3alpha.cluster.CircuitBreakers.Thresholds>`
	// are defined with the same
	// :ref:`RoutingPriority<envoy_api_enum_api.v3alpha.core.RoutingPriority>`, the first one in the
	// list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_api.v3alpha.core.RoutingPriority>`, the default values
	// are used.
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f46fa043d540e79, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_api.v3alpha.core.RoutingPriority>`.
// [#next-free-field: 8]
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_api.v3alpha.core.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	Priority core.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.api.v3alpha.core.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *types.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *types.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *types.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *types.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools   *types.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f46fa043d540e79, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return core.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *types.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *types.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *types.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v3alpha.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v3alpha.cluster.CircuitBreakers.Thresholds")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/cluster/circuit_breaker.proto", fileDescriptor_9f46fa043d540e79)
}

var fileDescriptor_9f46fa043d540e79 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0xed, 0xee, 0x52, 0xb9, 0xa8, 0x45, 0x01, 0x89, 0x50, 0x21, 0x54, 0xb8, 0xb4,
	0x27, 0x5b, 0x6a, 0xc5, 0x11, 0x81, 0xb2, 0xe2, 0xc0, 0x01, 0x14, 0x45, 0x2c, 0xd7, 0xc8, 0x4d,
	0x87, 0xd4, 0xda, 0xc4, 0x63, 0xc6, 0x4e, 0x49, 0xaf, 0x3c, 0x24, 0x2f, 0xc1, 0x1b, 0x70, 0x42,
	0x8d, 0xd3, 0x16, 0xd0, 0x56, 0xea, 0x2d, 0xf1, 0xf8, 0xfb, 0x66, 0xfc, 0xdb, 0x4c, 0x80, 0xde,
	0xe0, 0x56, 0x48, 0xa3, 0xc4, 0x66, 0x21, 0x4b, 0xb3, 0x96, 0x22, 0x2f, 0x6b, 0xeb, 0x80, 0x44,
	0xae, 0x28, 0xaf, 0x95, 0xcb, 0x96, 0x04, 0xf2, 0x0e, 0x88, 0x1b, 0x42, 0x87, 0xe1, 0xb3, 0x16,
	0xe0, 0xd2, 0x28, 0xde, 0x01, 0xbc, 0x03, 0xc6, 0x2f, 0xef, 0x71, 0x21, 0x81, 0x58, 0x4a, 0x0b,
	0x9e, 0x1e, 0xbf, 0x28, 0x10, 0x8b, 0x12, 0x44, 0xfb, 0xb7, 0xac, 0xbf, 0x8a, 0xef, 0x24, 0x8d,
	0x01, 0xb2, 0x5d, 0xfd, 0xe9, 0x46, 0x96, 0x6a, 0x25, 0x1d, 0x88, 0xfd, 0x87, 0x2f, 0xbc, 0xfa,
	0x79, 0xc9, 0x46, 0x37, 0x7e, 0xa0, 0xd8, 0xcf, 0x63, 0xc3, 0x5b, 0xc6, 0xdc, 0x9a, 0xc0, 0xae,
	0xb1, 0x5c, 0xd9, 0x28, 0x98, 0xf4, 0x66, 0x83, 0xf9, 0x6b, 0x7e, 0x72, 0x3e, 0xfe, 0x1f, 0xcf,
	0x3f, 0x1f, 0xe0, 0xf4, 0x2f, 0xd1, 0xf8, 0x57, 0x8f, 0xb1, 0x63, 0x29, 0xfc, 0xc8, 0xfa, 0x86,
	0x14, 0x92, 0x72, 0xdb, 0x28, 0x98, 0x04, 0xb3, 0xe1, 0x7c, 0x7a, 0x5f, 0x0f, 0x24, 0xe0, 0x29,
	0xd6, 0x4e, 0xe9, 0x22, 0xe9, 0xb6, 0xc7, 0xfd, 0xdf, 0xf1, 0xd5, 0x8f, 0xe0, 0xe2, 0x51, 0x90,
	0x1e, 0x14, 0xe1, 0x7b, 0x36, 0xaa, 0x64, 0x93, 0xe5, 0xa8, 0x35, 0xe4, 0x4e, 0xa1, 0xb6, 0xd1,
	0xc5, 0x24, 0x98, 0x0d, 0xe6, 0xcf, 0xb9, 0xcf, 0x86, 0xef, 0xb3, 0xe1, 0xb7, 0x1f, 0xb4, 0x5b,
	0xcc, 0xbf, 0xc8, 0xb2, 0x86, 0x74, 0x58, 0xc9, 0xe6, 0xe6, 0xc8, 0x84, 0x9f, 0xd8, 0x93, 0x9d,
	0xc6, 0x80, 0x5e, 0x29, 0x5d, 0x64, 0x04, 0xdf, 0x6a, 0xb0, 0xce, 0x46, 0xbd, 0x33, 0x5c, 0x61,
	0x25, 0x9b, 0xc4, 0x83, 0x69, 0xc7, 0x85, 0x6f, 0xd9, 0xc3, 0x9d, 0xef, 0xe0, 0xb9, 0x3c, 0xc3,
	0x33, 0xa8, 0x64, 0x73, 0x10, 0xbc, 0x61, 0x03, 0x2f, 0x70, 0xa4, 0xc0, 0x46, 0x57, 0x67, 0xf0,
	0xac, 0xe5, 0xdb, 0xfd, 0xe1, 0x94, 0x8d, 0x1c, 0xc9, 0xfc, 0x2e, 0x23, 0xa8, 0xa4, 0xd2, 0x4a,
	0x17, 0xd1, 0xf5, 0x24, 0x98, 0xf5, 0xd3, 0x61, 0xbb, 0x9c, 0xee, 0x57, 0xf7, 0x07, 0x3f, 0xe6,
	0x97, 0x19, 0xc4, 0xd2, 0x46, 0x0f, 0xce, 0x3c, 0xf8, 0x31, 0xc4, 0x64, 0xc7, 0xc5, 0xef, 0xd8,
	0x54, 0xa1, 0xbf, 0x50, 0x43, 0xd8, 0x6c, 0x4f, 0xbf, 0x9f, 0xf8, 0xf1, 0xbf, 0x0f, 0x28, 0xd9,
	0xb5, 0x48, 0x82, 0xe5, 0x75, 0xdb, 0x6b, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x26, 0xf4, 0xee,
	0x23, 0x4b, 0x03, 0x00, 0x00,
}