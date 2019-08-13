// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v2/stats.proto

package v2

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	matcher "github.com/altipla-consulting/envoy-api/envoy/type/matcher"
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

// Configuration for pluggable stats sinks.
type StatsSink struct {
	// The name of the stats sink to instantiate. The name must match a supported
	// stats sink. The built-in stats sinks are:
	//
	// * :ref:`envoy.statsd <envoy_api_msg_config.metrics.v2.StatsdSink>`
	// * :ref:`envoy.dog_statsd <envoy_api_msg_config.metrics.v2.DogStatsdSink>`
	// * :ref:`envoy.metrics_service <envoy_api_msg_config.metrics.v2.MetricsServiceConfig>`
	// * :ref:`envoy.stat_sinks.hystrix <envoy_api_msg_config.metrics.v2.HystrixSink>`
	//
	// Sinks optionally support tagged/multiple dimensional metrics.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Stats sink specific configuration which depends on the sink being instantiated. See
	// :ref:`StatsdSink <envoy_api_msg_config.metrics.v2.StatsdSink>` for an example.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*StatsSink_Config
	//	*StatsSink_TypedConfig
	ConfigType           isStatsSink_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StatsSink) Reset()         { *m = StatsSink{} }
func (m *StatsSink) String() string { return proto.CompactTextString(m) }
func (*StatsSink) ProtoMessage()    {}
func (*StatsSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{0}
}

func (m *StatsSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsSink.Unmarshal(m, b)
}
func (m *StatsSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsSink.Marshal(b, m, deterministic)
}
func (m *StatsSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsSink.Merge(m, src)
}
func (m *StatsSink) XXX_Size() int {
	return xxx_messageInfo_StatsSink.Size(m)
}
func (m *StatsSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsSink proto.InternalMessageInfo

func (m *StatsSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isStatsSink_ConfigType interface {
	isStatsSink_ConfigType()
}

type StatsSink_Config struct {
	Config *types.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type StatsSink_TypedConfig struct {
	TypedConfig *types.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*StatsSink_Config) isStatsSink_ConfigType() {}

func (*StatsSink_TypedConfig) isStatsSink_ConfigType() {}

func (m *StatsSink) GetConfigType() isStatsSink_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *StatsSink) GetConfig() *types.Struct {
	if x, ok := m.GetConfigType().(*StatsSink_Config); ok {
		return x.Config
	}
	return nil
}

func (m *StatsSink) GetTypedConfig() *types.Any {
	if x, ok := m.GetConfigType().(*StatsSink_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsSink_Config)(nil),
		(*StatsSink_TypedConfig)(nil),
	}
}

// Statistics configuration such as tagging.
type StatsConfig struct {
	// Each stat name is iteratively processed through these tag specifiers.
	// When a tag is matched, the first capture group is removed from the name so
	// later :ref:`TagSpecifiers <envoy_api_msg_config.metrics.v2.TagSpecifier>` cannot match that
	// same portion of the match.
	StatsTags []*TagSpecifier `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags,proto3" json:"stats_tags,omitempty"`
	// Use all default tag regexes specified in Envoy. These can be combined with
	// custom tags specified in :ref:`stats_tags
	// <envoy_api_field_config.metrics.v2.StatsConfig.stats_tags>`. They will be processed before
	// the custom tags.
	//
	// .. note::
	//
	//   If any default tags are specified twice, the config will be considered
	//   invalid.
	//
	// See :repo:`well_known_names.h <source/common/config/well_known_names.h>` for a list of the
	// default tags in Envoy.
	//
	// If not provided, the value is assumed to be true.
	UseAllDefaultTags *types.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags,proto3" json:"use_all_default_tags,omitempty"`
	// Inclusion/exclusion matcher for stat name creation. If not provided, all stats are instantiated
	// as normal. Preventing the instantiation of certain families of stats can improve memory
	// performance for Envoys running especially large configs.
	StatsMatcher         *StatsMatcher `protobuf:"bytes,3,opt,name=stats_matcher,json=statsMatcher,proto3" json:"stats_matcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StatsConfig) Reset()         { *m = StatsConfig{} }
func (m *StatsConfig) String() string { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()    {}
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{1}
}

func (m *StatsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsConfig.Unmarshal(m, b)
}
func (m *StatsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsConfig.Marshal(b, m, deterministic)
}
func (m *StatsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsConfig.Merge(m, src)
}
func (m *StatsConfig) XXX_Size() int {
	return xxx_messageInfo_StatsConfig.Size(m)
}
func (m *StatsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StatsConfig proto.InternalMessageInfo

func (m *StatsConfig) GetStatsTags() []*TagSpecifier {
	if m != nil {
		return m.StatsTags
	}
	return nil
}

func (m *StatsConfig) GetUseAllDefaultTags() *types.BoolValue {
	if m != nil {
		return m.UseAllDefaultTags
	}
	return nil
}

func (m *StatsConfig) GetStatsMatcher() *StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

// Configuration for disabling stat instantiation.
type StatsMatcher struct {
	// Types that are valid to be assigned to StatsMatcher:
	//	*StatsMatcher_RejectAll
	//	*StatsMatcher_ExclusionList
	//	*StatsMatcher_InclusionList
	StatsMatcher         isStatsMatcher_StatsMatcher `protobuf_oneof:"stats_matcher"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StatsMatcher) Reset()         { *m = StatsMatcher{} }
func (m *StatsMatcher) String() string { return proto.CompactTextString(m) }
func (*StatsMatcher) ProtoMessage()    {}
func (*StatsMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{2}
}

func (m *StatsMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsMatcher.Unmarshal(m, b)
}
func (m *StatsMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsMatcher.Marshal(b, m, deterministic)
}
func (m *StatsMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsMatcher.Merge(m, src)
}
func (m *StatsMatcher) XXX_Size() int {
	return xxx_messageInfo_StatsMatcher.Size(m)
}
func (m *StatsMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StatsMatcher proto.InternalMessageInfo

type isStatsMatcher_StatsMatcher interface {
	isStatsMatcher_StatsMatcher()
}

type StatsMatcher_RejectAll struct {
	RejectAll bool `protobuf:"varint,1,opt,name=reject_all,json=rejectAll,proto3,oneof"`
}

type StatsMatcher_ExclusionList struct {
	ExclusionList *matcher.ListStringMatcher `protobuf:"bytes,2,opt,name=exclusion_list,json=exclusionList,proto3,oneof"`
}

type StatsMatcher_InclusionList struct {
	InclusionList *matcher.ListStringMatcher `protobuf:"bytes,3,opt,name=inclusion_list,json=inclusionList,proto3,oneof"`
}

func (*StatsMatcher_RejectAll) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_ExclusionList) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_InclusionList) isStatsMatcher_StatsMatcher() {}

func (m *StatsMatcher) GetStatsMatcher() isStatsMatcher_StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

func (m *StatsMatcher) GetRejectAll() bool {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_RejectAll); ok {
		return x.RejectAll
	}
	return false
}

func (m *StatsMatcher) GetExclusionList() *matcher.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_ExclusionList); ok {
		return x.ExclusionList
	}
	return nil
}

func (m *StatsMatcher) GetInclusionList() *matcher.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_InclusionList); ok {
		return x.InclusionList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsMatcher_RejectAll)(nil),
		(*StatsMatcher_ExclusionList)(nil),
		(*StatsMatcher_InclusionList)(nil),
	}
}

// Designates a tag name and value pair. The value may be either a fixed value
// or a regex providing the value via capture groups. The specified tag will be
// unconditionally set if a fixed value, otherwise it will only be set if one
// or more capture groups in the regex match.
type TagSpecifier struct {
	// Attaches an identifier to the tag values to identify the tag being in the
	// sink. Envoy has a set of default names and regexes to extract dynamic
	// portions of existing stats, which can be found in :repo:`well_known_names.h
	// <source/common/config/well_known_names.h>` in the Envoy repository. If a :ref:`tag_name
	// <envoy_api_field_config.metrics.v2.TagSpecifier.tag_name>` is provided in the config and
	// neither :ref:`regex <envoy_api_field_config.metrics.v2.TagSpecifier.regex>` or
	// :ref:`fixed_value <envoy_api_field_config.metrics.v2.TagSpecifier.fixed_value>` were specified,
	// Envoy will attempt to find that name in its set of defaults and use the accompanying regex.
	//
	// .. note::
	//
	//   It is invalid to specify the same tag name twice in a config.
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	// Types that are valid to be assigned to TagValue:
	//	*TagSpecifier_Regex
	//	*TagSpecifier_FixedValue
	TagValue             isTagSpecifier_TagValue `protobuf_oneof:"tag_value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TagSpecifier) Reset()         { *m = TagSpecifier{} }
func (m *TagSpecifier) String() string { return proto.CompactTextString(m) }
func (*TagSpecifier) ProtoMessage()    {}
func (*TagSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{3}
}

func (m *TagSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagSpecifier.Unmarshal(m, b)
}
func (m *TagSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagSpecifier.Marshal(b, m, deterministic)
}
func (m *TagSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagSpecifier.Merge(m, src)
}
func (m *TagSpecifier) XXX_Size() int {
	return xxx_messageInfo_TagSpecifier.Size(m)
}
func (m *TagSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TagSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_TagSpecifier proto.InternalMessageInfo

func (m *TagSpecifier) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

type isTagSpecifier_TagValue interface {
	isTagSpecifier_TagValue()
}

type TagSpecifier_Regex struct {
	Regex string `protobuf:"bytes,2,opt,name=regex,proto3,oneof"`
}

type TagSpecifier_FixedValue struct {
	FixedValue string `protobuf:"bytes,3,opt,name=fixed_value,json=fixedValue,proto3,oneof"`
}

func (*TagSpecifier_Regex) isTagSpecifier_TagValue() {}

func (*TagSpecifier_FixedValue) isTagSpecifier_TagValue() {}

func (m *TagSpecifier) GetTagValue() isTagSpecifier_TagValue {
	if m != nil {
		return m.TagValue
	}
	return nil
}

func (m *TagSpecifier) GetRegex() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *TagSpecifier) GetFixedValue() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_FixedValue); ok {
		return x.FixedValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagSpecifier_Regex)(nil),
		(*TagSpecifier_FixedValue)(nil),
	}
}

// Stats configuration proto schema for built-in *envoy.statsd* sink. This sink does not support
// tagged metrics.
type StatsdSink struct {
	// Types that are valid to be assigned to StatsdSpecifier:
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
	// Optional custom prefix for StatsdSink. If
	// specified, this will override the default prefix.
	// For example:
	//
	// .. code-block:: json
	//
	//   {
	//     "prefix" : "envoy-prod"
	//   }
	//
	// will change emitted stats to
	//
	// .. code-block:: cpp
	//
	//   envoy-prod.test_counter:1|c
	//   envoy-prod.test_timer:5|ms
	//
	// Note that the default prefix, "envoy", will be used if a prefix is not
	// specified.
	//
	// Stats with default prefix:
	//
	// .. code-block:: cpp
	//
	//   envoy.test_counter:1|c
	//   envoy.test_timer:5|ms
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsdSink) Reset()         { *m = StatsdSink{} }
func (m *StatsdSink) String() string { return proto.CompactTextString(m) }
func (*StatsdSink) ProtoMessage()    {}
func (*StatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{4}
}

func (m *StatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsdSink.Unmarshal(m, b)
}
func (m *StatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsdSink.Marshal(b, m, deterministic)
}
func (m *StatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsdSink.Merge(m, src)
}
func (m *StatsdSink) XXX_Size() int {
	return xxx_messageInfo_StatsdSink.Size(m)
}
func (m *StatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsdSink proto.InternalMessageInfo

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
}

type StatsdSink_Address struct {
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type StatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier() {}

func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (m *StatsdSink) GetAddress() *core.Address {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *StatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

func (m *StatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
}

// Stats configuration proto schema for built-in *envoy.dog_statsd* sink.
// The sink emits stats with `DogStatsD <https://docs.datadoghq.com/guides/dogstatsd/>`_
// compatible tags. Tags are configurable via :ref:`StatsConfig
// <envoy_api_msg_config.metrics.v2.StatsConfig>`.
// [#comment:next free field: 3]
type DogStatsdSink struct {
	// Types that are valid to be assigned to DogStatsdSpecifier:
	//	*DogStatsdSink_Address
	DogStatsdSpecifier isDogStatsdSink_DogStatsdSpecifier `protobuf_oneof:"dog_statsd_specifier"`
	// Optional custom metric name prefix. See :ref:`StatsdSink's prefix field
	// <envoy_api_field_config.metrics.v2.StatsdSink.prefix>` for more details.
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogStatsdSink) Reset()         { *m = DogStatsdSink{} }
func (m *DogStatsdSink) String() string { return proto.CompactTextString(m) }
func (*DogStatsdSink) ProtoMessage()    {}
func (*DogStatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{5}
}

func (m *DogStatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogStatsdSink.Unmarshal(m, b)
}
func (m *DogStatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogStatsdSink.Marshal(b, m, deterministic)
}
func (m *DogStatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogStatsdSink.Merge(m, src)
}
func (m *DogStatsdSink) XXX_Size() int {
	return xxx_messageInfo_DogStatsdSink.Size(m)
}
func (m *DogStatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_DogStatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_DogStatsdSink proto.InternalMessageInfo

type isDogStatsdSink_DogStatsdSpecifier interface {
	isDogStatsdSink_DogStatsdSpecifier()
}

type DogStatsdSink_Address struct {
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

func (*DogStatsdSink_Address) isDogStatsdSink_DogStatsdSpecifier() {}

func (m *DogStatsdSink) GetDogStatsdSpecifier() isDogStatsdSink_DogStatsdSpecifier {
	if m != nil {
		return m.DogStatsdSpecifier
	}
	return nil
}

func (m *DogStatsdSink) GetAddress() *core.Address {
	if x, ok := m.GetDogStatsdSpecifier().(*DogStatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *DogStatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DogStatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DogStatsdSink_Address)(nil),
	}
}

// Stats configuration proto schema for built-in *envoy.stat_sinks.hystrix* sink.
// The sink emits stats in `text/event-stream
// <https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events>`_
// formatted stream for use by `Hystrix dashboard
// <https://github.com/Netflix-Skunkworks/hystrix-dashboard/wiki>`_.
//
// Note that only a single HystrixSink should be configured.
//
// Streaming is started through an admin endpoint :http:get:`/hystrix_event_stream`.
type HystrixSink struct {
	// The number of buckets the rolling statistical window is divided into.
	//
	// Each time the sink is flushed, all relevant Envoy statistics are sampled and
	// added to the rolling window (removing the oldest samples in the window
	// in the process). The sink then outputs the aggregate statistics across the
	// current rolling window to the event stream(s).
	//
	// rolling_window(ms) = stats_flush_interval(ms) * num_of_buckets
	//
	// More detailed explanation can be found in `Hystrix wiki
	// <https://github.com/Netflix/Hystrix/wiki/Metrics-and-Monitoring#hystrixrollingnumber>`_.
	NumBuckets           int64    `protobuf:"varint,1,opt,name=num_buckets,json=numBuckets,proto3" json:"num_buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HystrixSink) Reset()         { *m = HystrixSink{} }
func (m *HystrixSink) String() string { return proto.CompactTextString(m) }
func (*HystrixSink) ProtoMessage()    {}
func (*HystrixSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e6d6532209c486, []int{6}
}

func (m *HystrixSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HystrixSink.Unmarshal(m, b)
}
func (m *HystrixSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HystrixSink.Marshal(b, m, deterministic)
}
func (m *HystrixSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HystrixSink.Merge(m, src)
}
func (m *HystrixSink) XXX_Size() int {
	return xxx_messageInfo_HystrixSink.Size(m)
}
func (m *HystrixSink) XXX_DiscardUnknown() {
	xxx_messageInfo_HystrixSink.DiscardUnknown(m)
}

var xxx_messageInfo_HystrixSink proto.InternalMessageInfo

func (m *HystrixSink) GetNumBuckets() int64 {
	if m != nil {
		return m.NumBuckets
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsSink)(nil), "envoy.config.metrics.v2.StatsSink")
	proto.RegisterType((*StatsConfig)(nil), "envoy.config.metrics.v2.StatsConfig")
	proto.RegisterType((*StatsMatcher)(nil), "envoy.config.metrics.v2.StatsMatcher")
	proto.RegisterType((*TagSpecifier)(nil), "envoy.config.metrics.v2.TagSpecifier")
	proto.RegisterType((*StatsdSink)(nil), "envoy.config.metrics.v2.StatsdSink")
	proto.RegisterType((*DogStatsdSink)(nil), "envoy.config.metrics.v2.DogStatsdSink")
	proto.RegisterType((*HystrixSink)(nil), "envoy.config.metrics.v2.HystrixSink")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v2/stats.proto", fileDescriptor_51e6d6532209c486)
}

var fileDescriptor_51e6d6532209c486 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x8d, 0xe3, 0x7e, 0xc4, 0xd7, 0x49, 0xd5, 0x67, 0x45, 0x2f, 0x6d, 0xdf, 0xd3, 0x6b, 0x5e,
	0x50, 0xa5, 0x88, 0x85, 0x2d, 0x82, 0x84, 0xd4, 0x65, 0xdc, 0x2e, 0xac, 0x02, 0x55, 0xe5, 0x54,
	0x2c, 0xd8, 0x58, 0x53, 0x7b, 0x62, 0x86, 0x3a, 0x1e, 0x6b, 0x66, 0x1c, 0x12, 0xb1, 0x41, 0xfc,
	0x0e, 0x36, 0xfc, 0x38, 0x76, 0xfc, 0x00, 0xc4, 0x0a, 0xcd, 0x47, 0xaa, 0xd0, 0x92, 0x05, 0x12,
	0xbb, 0x78, 0xce, 0x39, 0xf7, 0x9e, 0x73, 0xef, 0x4c, 0xe0, 0x11, 0x2e, 0xe7, 0x74, 0x19, 0xa4,
	0xb4, 0x9c, 0x92, 0x3c, 0x98, 0x61, 0xc1, 0x48, 0xca, 0x83, 0xf9, 0x28, 0xe0, 0x02, 0x09, 0xee,
	0x57, 0x8c, 0x0a, 0xea, 0xf5, 0x14, 0xc9, 0xd7, 0x24, 0xdf, 0x90, 0xfc, 0xf9, 0xe8, 0xe8, 0x58,
	0xab, 0x51, 0x45, 0xa4, 0x24, 0xa5, 0x0c, 0x07, 0x28, 0xcb, 0x18, 0xe6, 0x46, 0xb9, 0x22, 0x88,
	0x65, 0x85, 0x83, 0x19, 0x12, 0xe9, 0x1b, 0xcc, 0x02, 0x2e, 0x18, 0x29, 0x73, 0x43, 0x38, 0xcc,
	0x29, 0xcd, 0x0b, 0x1c, 0xa8, 0xaf, 0x9b, 0x7a, 0x1a, 0xa0, 0x72, 0x69, 0xa0, 0x7f, 0xef, 0x43,
	0x5c, 0xb0, 0x3a, 0x15, 0x06, 0xfd, 0xef, 0x3e, 0xfa, 0x8e, 0xa1, 0xaa, 0xc2, 0x6c, 0xd5, 0xb9,
	0x37, 0x47, 0x05, 0xc9, 0x90, 0xc0, 0xc1, 0xea, 0x87, 0x06, 0x06, 0x9f, 0x2c, 0x70, 0x26, 0x32,
	0xdc, 0x84, 0x94, 0xb7, 0x9e, 0x07, 0x5b, 0x25, 0x9a, 0xe1, 0x03, 0xab, 0x6f, 0x0d, 0x9d, 0x58,
	0xfd, 0xf6, 0x9e, 0xc0, 0x8e, 0x8e, 0x7a, 0xd0, 0xec, 0x5b, 0x43, 0x77, 0xd4, 0xf3, 0x75, 0x2f,
	0x7f, 0xd5, 0xcb, 0x9f, 0x28, 0x27, 0x51, 0x23, 0x36, 0x44, 0xef, 0x14, 0xda, 0x32, 0x63, 0x96,
	0x18, 0xa1, 0xad, 0x84, 0xdd, 0x07, 0xc2, 0x71, 0xb9, 0x8c, 0x1a, 0xb1, 0xab, 0xb8, 0x67, 0x8a,
	0x1a, 0x76, 0xc0, 0xd5, 0xa2, 0x44, 0x9e, 0x0e, 0xbe, 0x5a, 0xe0, 0x2a, 0x7b, 0x1a, 0xf6, 0xce,
	0x01, 0xd4, 0x2a, 0x12, 0x81, 0x72, 0x7e, 0x60, 0xf5, 0xed, 0xa1, 0x3b, 0x3a, 0xf1, 0x37, 0x2c,
	0xc4, 0xbf, 0x46, 0xf9, 0xa4, 0xc2, 0x29, 0x99, 0x12, 0xcc, 0x62, 0x47, 0x09, 0xaf, 0x51, 0xce,
	0xbd, 0xe7, 0xd0, 0xad, 0x39, 0x4e, 0x50, 0x51, 0x24, 0x19, 0x9e, 0xa2, 0xba, 0x10, 0xba, 0x9e,
	0x0e, 0x78, 0xf4, 0xc0, 0x67, 0x48, 0x69, 0xf1, 0x0a, 0x15, 0x35, 0x8e, 0xff, 0xaa, 0x39, 0x1e,
	0x17, 0xc5, 0xb9, 0x56, 0xa9, 0x62, 0x17, 0xd0, 0xd1, 0x96, 0xcc, 0x46, 0x4d, 0xda, 0xcd, 0xae,
	0x54, 0x9e, 0x97, 0x9a, 0x1c, 0xb7, 0xf9, 0xda, 0xd7, 0xe0, 0x8b, 0x05, 0xed, 0x75, 0xd8, 0x3b,
	0x06, 0x60, 0xf8, 0x2d, 0x4e, 0x85, 0x34, 0xab, 0xd6, 0xd2, 0x8a, 0x1a, 0xb1, 0xa3, 0xcf, 0xc6,
	0x45, 0xe1, 0x5d, 0xc2, 0x1e, 0x5e, 0xa4, 0x45, 0xcd, 0x09, 0x2d, 0x93, 0x82, 0x70, 0x61, 0x42,
	0xac, 0xda, 0xcb, 0x29, 0xfa, 0xc6, 0x99, 0xff, 0x82, 0x70, 0x31, 0x51, 0xf7, 0xcd, 0xd4, 0x8f,
	0x1a, 0x71, 0xe7, 0x4e, 0x2e, 0x51, 0x59, 0x8f, 0x94, 0x3f, 0xd5, 0xb3, 0x7f, 0xb3, 0xde, 0x9d,
	0x5c, 0xa2, 0x61, 0xf7, 0xde, 0x74, 0x3c, 0xfb, 0x5b, 0x68, 0x0d, 0xde, 0x43, 0x7b, 0x7d, 0x37,
	0xde, 0x21, 0xb4, 0x04, 0xca, 0x93, 0xb5, 0xbb, 0xb7, 0x2b, 0x50, 0x7e, 0x29, 0xaf, 0x5f, 0x1f,
	0xb6, 0x19, 0xce, 0xf1, 0x42, 0xe5, 0x72, 0xc2, 0xd6, 0xf7, 0x70, 0x9b, 0xd9, 0xc3, 0x0f, 0x72,
	0x0c, 0x1a, 0xf0, 0xfe, 0x07, 0x77, 0x4a, 0x16, 0x38, 0x4b, 0xe6, 0x72, 0x45, 0xca, 0xaf, 0x13,
	0x35, 0x62, 0x50, 0x87, 0x6a, 0x6d, 0xa1, 0x0b, 0x8e, 0xac, 0xaf, 0x08, 0x83, 0xcf, 0x16, 0x80,
	0x1a, 0x72, 0xa6, 0xee, 0xfc, 0x33, 0xd8, 0x35, 0xaf, 0x54, 0xb5, 0x96, 0xfb, 0xd7, 0x51, 0x51,
	0x45, 0xe4, 0xba, 0xe4, 0x3b, 0xf6, 0xc7, 0x9a, 0x11, 0x35, 0xe2, 0x15, 0xd9, 0x7b, 0x0c, 0xfb,
	0x22, 0xad, 0x12, 0x19, 0x56, 0x60, 0xa6, 0xbd, 0x37, 0x4d, 0xef, 0x3d, 0x91, 0x56, 0x67, 0x1a,
	0x50, 0x21, 0xfe, 0x86, 0x9d, 0x8a, 0xe1, 0x29, 0x59, 0x68, 0x77, 0xb1, 0xf9, 0x0a, 0x7b, 0xb0,
	0xaf, 0xa6, 0x93, 0x25, 0xfc, 0x6e, 0x16, 0x6a, 0x40, 0x1f, 0x2d, 0xe8, 0x9c, 0xd3, 0xfc, 0x0f,
	0xd8, 0xdc, 0xd4, 0xfa, 0x1f, 0xe8, 0x66, 0x34, 0x4f, 0x7e, 0xd9, 0xfe, 0x62, 0xab, 0xd5, 0xdc,
	0xb7, 0x07, 0x3e, 0xb8, 0xd1, 0x52, 0xfe, 0x3f, 0x2d, 0x94, 0x83, 0x63, 0x70, 0xcb, 0x7a, 0x96,
	0xdc, 0xd4, 0xe9, 0x2d, 0x16, 0xda, 0x85, 0x1d, 0x43, 0x59, 0xcf, 0x42, 0x7d, 0x12, 0x9e, 0xc2,
	0x09, 0xa1, 0xda, 0x55, 0xc5, 0xe8, 0x62, 0xb9, 0xe9, 0x05, 0x84, 0x7a, 0xfc, 0x57, 0xf2, 0x79,
	0x5d, 0x59, 0xaf, 0x9b, 0xf3, 0xd1, 0xcd, 0x8e, 0x7a, 0x6b, 0x4f, 0x7f, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x43, 0x6b, 0xad, 0xf2, 0x81, 0x05, 0x00, 0x00,
}
