// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto

package envoy_config_filter_network_tcp_proxy_v2

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	v2 "github.com/altipla-consulting/envoy-api/envoy/config/filter/accesslog/v2"
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

// [#next-free-field: 11]
type TcpProxy struct {
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_tcp_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria. Only endpoints in the upstream
	// cluster with metadata matching that set in metadata_match will be
	// considered. The filter name should be specified as *envoy.lb*.
	MetadataMatch *core.Metadata `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// The idle timeout for connections managed by the TCP proxy filter. The idle timeout
	// is defined as the period in which there are no bytes sent or received on either
	// the upstream or downstream connection. If not set, the default idle timeout is 1 hour. If set
	// to 0s, the timeout will be disabled.
	//
	// .. warning::
	//   Disabling this timeout has a highly likelihood of yielding connection leaks due to lost TCP
	//   FIN packets, etc.
	IdleTimeout *types.Duration `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// [#not-implemented-hide:] The idle timeout for connections managed by the TCP proxy
	// filter. The idle timeout is defined as the period in which there is no
	// active traffic. If not set, there is no idle timeout. When the idle timeout
	// is reached the connection will be closed. The distinction between
	// downstream_idle_timeout/upstream_idle_timeout provides a means to set
	// timeout based on the last byte sent on the downstream/upstream connection.
	DownstreamIdleTimeout *types.Duration `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	// [#not-implemented-hide:]
	UpstreamIdleTimeout *types.Duration `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>`
	// emitted by the this tcp_proxy.
	AccessLog []*v2.AccessLog `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// [#not-implemented-hide:] Deprecated.
	DeprecatedV1 *TcpProxy_DeprecatedV1 `protobuf:"bytes,6,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	// The maximum number of unsuccessful connection attempts that will be made before
	// giving up. If the parameter is not specified, 1 connection attempt will be made.
	MaxConnectAttempts   *types.UInt32Value `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *types.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *types.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *types.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v2.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// Deprecated: Do not use.
func (m *TcpProxy) GetDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *types.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

// [#not-implemented-hide:] Deprecated.
// TCP Proxy filter configuration using V1 format.
//
// Deprecated: Do not use.
type TcpProxy_DeprecatedV1 struct {
	// The route table for the filter. All filter instances must have a route
	// table, even if it is empty.
	Routes               []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1) Reset()         { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0}
}

func (m *TcpProxy_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Size(m)
}
func (m *TcpProxy_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1 proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

// A TCP proxy route consists of a set of optional L4 criteria and the
// name of a cluster. If a downstream connection matches all the
// specified criteria, the cluster in the route is used for the
// corresponding upstream connection. Routes are tried in the order
// specified until a match is found. If no match is found, the connection
// is closed. A route with no criteria is valid and always produces a
// match.
// [#next-free-field: 6]
type TcpProxy_DeprecatedV1_TCPRoute struct {
	// The cluster to connect to when a the downstream network connection
	// matches the specified criteria.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// An optional list of IP address subnets in the form
	// “ip_address/xx”. The criteria is satisfied if the destination IP
	// address of the downstream connection is contained in at least one of
	// the specified subnets. If the parameter is not specified or the list
	// is empty, the destination IP address is ignored. The destination IP
	// address of the downstream connection might be different from the
	// addresses on which the proxy is listening if the connection has been
	// redirected.
	DestinationIpList []*core.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList,proto3" json:"destination_ip_list,omitempty"`
	// An optional string containing a comma-separated list of port numbers
	// or ranges. The criteria is satisfied if the destination port of the
	// downstream connection is contained in at least one of the specified
	// ranges. If the parameter is not specified, the destination port is
	// ignored. The destination port address of the downstream connection
	// might be different from the port on which the proxy is listening if
	// the connection has been redirected.
	DestinationPorts string `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	// An optional list of IP address subnets in the form
	// “ip_address/xx”. The criteria is satisfied if the source IP address
	// of the downstream connection is contained in at least one of the
	// specified subnets. If the parameter is not specified or the list is
	// empty, the source IP address is ignored.
	SourceIpList []*core.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList,proto3" json:"source_ip_list,omitempty"`
	// An optional string containing a comma-separated list of port numbers
	// or ranges. The criteria is satisfied if the source port of the
	// downstream connection is contained in at least one of the specified
	// ranges. If the parameter is not specified, the source port is
	// ignored.
	SourcePorts          string   `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Size(m)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*core.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*core.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

// Allows for specification of multiple upstream clusters along with weights
// that indicate the percentage of traffic to be forwarded to each cluster.
// The router selects an upstream cluster based on these weights.
type TcpProxy_WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is
	// determined by its weight. The sum of weights across all entries in the
	// clusters array determines the total weight.
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto", fileDescriptor_1f6b35dbcbad27ba)
}

var fileDescriptor_1f6b35dbcbad27ba = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x3b, 0x69, 0x9a, 0x4c, 0x12, 0xd8, 0xce, 0xb2, 0x5a, 0x93, 0xad, 0xd8, 0x2c, 0x57,
	0x11, 0x2b, 0xd9, 0x6c, 0x2a, 0x21, 0x84, 0x90, 0x50, 0x9d, 0x5e, 0x34, 0xa8, 0x91, 0x82, 0x55,
	0xda, 0x4b, 0x6b, 0x6a, 0x4f, 0xdc, 0x01, 0xdb, 0x33, 0xcc, 0x8c, 0x93, 0xf4, 0x29, 0x40, 0x3c,
	0x1d, 0x8f, 0x82, 0x7a, 0x81, 0x90, 0x67, 0xc6, 0x89, 0xdb, 0x06, 0xb5, 0xa2, 0x57, 0xf6, 0x9c,
	0x73, 0xbe, 0xef, 0x3b, 0x73, 0x7e, 0x06, 0x7c, 0x8b, 0xf3, 0x25, 0xbd, 0xf1, 0x22, 0x9a, 0x2f,
	0x48, 0xe2, 0x2d, 0x48, 0x2a, 0x31, 0xf7, 0x72, 0x2c, 0x57, 0x94, 0xff, 0xea, 0xc9, 0x88, 0x85,
	0x8c, 0xd3, 0xf5, 0x8d, 0xb7, 0x1c, 0x6f, 0x0f, 0x2e, 0xe3, 0x54, 0x52, 0x38, 0x52, 0x48, 0x57,
	0x23, 0x5d, 0x8d, 0x74, 0x0d, 0xd2, 0xdd, 0x06, 0x2f, 0xc7, 0x83, 0x77, 0x5a, 0x03, 0x31, 0x52,
	0xf2, 0x44, 0x94, 0x63, 0x0f, 0xc5, 0x31, 0xc7, 0x42, 0x68, 0xaa, 0xc1, 0xe1, 0xc3, 0x80, 0x2b,
	0x24, 0xb0, 0xf1, 0x7e, 0xbd, 0x2b, 0x45, 0x14, 0x45, 0x58, 0x88, 0x94, 0x26, 0x25, 0x62, 0x73,
	0x30, 0x88, 0x2f, 0x12, 0x4a, 0x93, 0x14, 0x7b, 0xea, 0x74, 0x55, 0x2c, 0xbc, 0xb8, 0xe0, 0x48,
	0x12, 0x9a, 0xff, 0x97, 0x7f, 0xc5, 0x11, 0x63, 0x98, 0x57, 0xf9, 0xbc, 0x59, 0xa2, 0x94, 0xc4,
	0x48, 0x62, 0xaf, 0xfa, 0xd1, 0x8e, 0x2f, 0xff, 0x01, 0xa0, 0x7d, 0x1e, 0xb1, 0x79, 0x79, 0x33,
	0x38, 0x02, 0x5d, 0x21, 0x91, 0x0c, 0x19, 0xc7, 0x0b, 0xb2, 0x76, 0xac, 0xa1, 0x35, 0xea, 0xf8,
	0xfb, 0xb7, 0x7e, 0x93, 0xdb, 0x43, 0x2b, 0x00, 0xa5, 0x6f, 0xae, 0x5c, 0x70, 0x00, 0xf6, 0xa3,
	0xb4, 0x10, 0x12, 0x73, 0xc7, 0x2e, 0xa3, 0x4e, 0x5f, 0x04, 0x95, 0x01, 0xfe, 0x06, 0x0e, 0x56,
	0x98, 0x24, 0xd7, 0x12, 0xc7, 0xa1, 0xb1, 0x09, 0x07, 0x0c, 0xad, 0x51, 0x77, 0xec, 0xbb, 0x4f,
	0x2d, 0xb1, 0x5b, 0x25, 0xe5, 0x5e, 0x1a, 0xae, 0x89, 0xa6, 0x3a, 0x7d, 0x11, 0xbc, 0x5c, 0xdd,
	0x35, 0x09, 0xe8, 0x83, 0x4f, 0x32, 0x2c, 0x51, 0x8c, 0x24, 0x0a, 0x33, 0x24, 0xa3, 0x6b, 0xa7,
	0xa3, 0xf4, 0xde, 0x1a, 0x3d, 0xc4, 0x48, 0xc9, 0x59, 0xf6, 0xc1, 0x9d, 0x99, 0xc0, 0xa0, 0x5f,
	0x41, 0x66, 0x25, 0x02, 0x7e, 0x0f, 0x7a, 0x24, 0x4e, 0x71, 0x28, 0x49, 0x86, 0x69, 0x21, 0x9d,
	0xb6, 0x62, 0xf8, 0xdc, 0xd5, 0x95, 0x75, 0xab, 0xca, 0xba, 0x27, 0xa6, 0xf2, 0x41, 0xb7, 0x0c,
	0x3f, 0xd7, 0xd1, 0xf0, 0x27, 0xf0, 0x26, 0xa6, 0xab, 0x5c, 0x48, 0x8e, 0x51, 0x16, 0xde, 0x21,
	0x6a, 0x3c, 0x46, 0xf4, 0x7a, 0x8b, 0x9c, 0xd6, 0x28, 0x67, 0xe0, 0x75, 0xc1, 0x76, 0x11, 0x36,
	0x1f, 0x23, 0x7c, 0x55, 0xe1, 0xea, 0x74, 0x3f, 0x02, 0xa0, 0xa7, 0x2a, 0x4c, 0x69, 0xe2, 0xec,
	0x0d, 0x1b, 0xa3, 0xee, 0xf8, 0xc3, 0xce, 0x7e, 0x6c, 0x87, 0x6f, 0x39, 0x76, 0x8f, 0xd5, 0xe1,
	0x8c, 0x26, 0x41, 0x07, 0x55, 0xbf, 0xf0, 0x1a, 0xf4, 0x63, 0xcc, 0x38, 0x8e, 0x50, 0xd9, 0xe4,
	0xe5, 0x47, 0xa7, 0xa5, 0x52, 0xfa, 0xe1, 0x7f, 0xb4, 0xf7, 0x64, 0xc3, 0x73, 0xf1, 0xd1, 0xb7,
	0x1d, 0x2b, 0xe8, 0xc5, 0x35, 0x0b, 0xbc, 0x04, 0x9f, 0x65, 0x68, 0x1d, 0x46, 0x34, 0xcf, 0x71,
	0x24, 0x43, 0x24, 0x25, 0xce, 0x98, 0x14, 0xce, 0xbe, 0x12, 0x3c, 0x7c, 0x50, 0x83, 0x9f, 0xa7,
	0xb9, 0x3c, 0x1a, 0x5f, 0xa0, 0xb4, 0xc0, 0x6a, 0x72, 0xbf, 0xb2, 0x47, 0x56, 0x00, 0x33, 0xb4,
	0x9e, 0x68, 0x86, 0x63, 0x43, 0x30, 0xf8, 0xa3, 0x01, 0x7a, 0x75, 0x6d, 0xf8, 0x0b, 0x68, 0x71,
	0x5a, 0x48, 0x2c, 0x1c, 0x4b, 0xd5, 0xe6, 0xf4, 0x99, 0x97, 0x71, 0xcf, 0x27, 0xf3, 0xa0, 0x24,
	0xf4, 0xdb, 0xb7, 0xfe, 0xde, 0x9f, 0x96, 0xdd, 0xb6, 0x02, 0xa3, 0x30, 0xf8, 0xdd, 0x06, 0xed,
	0xca, 0x0d, 0xdf, 0x6f, 0x77, 0xe9, 0xde, 0xc6, 0x6d, 0x56, 0xea, 0x0c, 0xbc, 0x8a, 0xb1, 0x90,
	0x24, 0x57, 0xfd, 0x0d, 0x09, 0x0b, 0x53, 0x22, 0xa4, 0x63, 0xab, 0x44, 0x0f, 0x77, 0x0c, 0xf9,
	0x84, 0xc4, 0x3c, 0x40, 0x79, 0x82, 0x83, 0x83, 0x1a, 0x70, 0xca, 0xce, 0x88, 0x90, 0xf0, 0x03,
	0xa8, 0x1b, 0x43, 0x46, 0xb9, 0x14, 0x6a, 0x4a, 0x3b, 0xc1, 0xcb, 0x9a, 0x63, 0x5e, 0xda, 0xcb,
	0xd5, 0x12, 0xb4, 0xe0, 0x11, 0xde, 0xa8, 0x36, 0x9f, 0xa0, 0xda, 0xd3, 0x18, 0x23, 0xf8, 0x1e,
	0x98, 0xb3, 0xd1, 0xda, 0x53, 0x5a, 0x5d, 0x6d, 0x53, 0x32, 0xdf, 0xd9, 0x8e, 0x35, 0xf8, 0xcb,
	0x02, 0x9f, 0xde, 0xdb, 0x76, 0xb8, 0x04, 0xed, 0xcd, 0x1b, 0xa2, 0xfb, 0x32, 0x7f, 0xfe, 0x1b,
	0xe2, 0x9a, 0xaf, 0x36, 0xd7, 0xfa, 0xb3, 0xd1, 0x1a, 0xcc, 0x40, 0xff, 0x4e, 0x10, 0x7c, 0x0b,
	0x9a, 0x39, 0xca, 0xf0, 0xfd, 0x16, 0x29, 0x23, 0x7c, 0x07, 0x5a, 0xfa, 0x4d, 0x52, 0xaf, 0x61,
	0x7f, 0x3b, 0x79, 0xc6, 0xec, 0x3b, 0xe0, 0xc0, 0x50, 0x87, 0x82, 0xe1, 0x88, 0x2c, 0x08, 0xe6,
	0xb0, 0xf1, 0xb7, 0x6f, 0xf9, 0x53, 0xf0, 0x0d, 0xa1, 0xfa, 0x4a, 0x3a, 0xef, 0xa7, 0xde, 0xce,
	0xef, 0x57, 0xd7, 0x9b, 0x97, 0xc3, 0x3f, 0xb7, 0xae, 0x5a, 0x6a, 0x0b, 0x8e, 0xfe, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xae, 0x4a, 0xeb, 0x75, 0x02, 0x07, 0x00, 0x00,
}
