// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/retry/previous_priorities/previous_priorities_config.proto

package envoy_config_retry_previous_priorities

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A retry host selector that attempts to spread retries between priorities, even if certain
// priorities would not normally be attempted due to higher priorities being available.
//
// As priorities get excluded, load will be distributed amongst the remaining healthy priorities
// based on the relative health of the priorities, matching how load is distributed during regular
// host selection. For example, given priority healths of {100, 50, 50}, the original load will be
// {100, 0, 0} (since P0 has capacity to handle 100% of the traffic). If P0 is excluded, the load
// changes to {0, 50, 50}, because P1 is only able to handle 50% of the traffic, causing the
// remaining to spill over to P2.
//
// Each priority attempted will be excluded until there are no healthy priorities left, at which
// point the list of attempted priorities will be reset, essentially starting from the beginning.
// For example, given three priorities P0, P1, P2 with healthy % of 100, 0 and 50 respectively, the
// following sequence of priorities would be selected (assuming update_frequency = 1):
// Attempt 1: P0 (P0 is 100% healthy)
// Attempt 2: P2 (P0 already attempted, P2 only healthy priority)
// Attempt 3: P0 (no healthy priorities, reset)
// Attempt 4: P2
//
// In the case of all upstream hosts being unhealthy, no adjustments will be made to the original
// priority load, so behavior should be identical to not using this plugin.
//
// Using this PriorityFilter requires rebuilding the priority load, which runs in O(# of
// priorities), which might incur significant overhead for clusters with many priorities.
type PreviousPrioritiesConfig struct {
	// How often the priority load should be updated based on previously attempted priorities. Useful
	// to allow each priorities to receive more than one request before being excluded or to reduce
	// the number of times that the priority load has to be recomputed.
	//
	// For example, by setting this to 2, then the first two attempts (initial attempt and first
	// retry) will use the unmodified priority load. The third and fourth attempt will use priority
	// load which excludes the priorities routed to with the first two attempts, and the fifth and
	// sixth attempt will use the priority load excluding the priorities used for the first four
	// attempts.
	UpdateFrequency      int32    `protobuf:"varint,1,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousPrioritiesConfig) Reset()         { *m = PreviousPrioritiesConfig{} }
func (m *PreviousPrioritiesConfig) String() string { return proto.CompactTextString(m) }
func (*PreviousPrioritiesConfig) ProtoMessage()    {}
func (*PreviousPrioritiesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_previous_priorities_config_458c31d10e9decf3, []int{0}
}
func (m *PreviousPrioritiesConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreviousPrioritiesConfig.Unmarshal(m, b)
}
func (m *PreviousPrioritiesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreviousPrioritiesConfig.Marshal(b, m, deterministic)
}
func (dst *PreviousPrioritiesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousPrioritiesConfig.Merge(dst, src)
}
func (m *PreviousPrioritiesConfig) XXX_Size() int {
	return xxx_messageInfo_PreviousPrioritiesConfig.Size(m)
}
func (m *PreviousPrioritiesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousPrioritiesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousPrioritiesConfig proto.InternalMessageInfo

func (m *PreviousPrioritiesConfig) GetUpdateFrequency() int32 {
	if m != nil {
		return m.UpdateFrequency
	}
	return 0
}

func init() {
	proto.RegisterType((*PreviousPrioritiesConfig)(nil), "envoy.config.retry.previous_priorities.PreviousPrioritiesConfig")
}

func init() {
	proto.RegisterFile("envoy/config/retry/previous_priorities/previous_priorities_config.proto", fileDescriptor_previous_priorities_config_458c31d10e9decf3)
}

var fileDescriptor_previous_priorities_config_458c31d10e9decf3 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0x2f,
	0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0x2f, 0x28, 0xca, 0xcc, 0x2f, 0xca, 0x2c, 0xc9,
	0x4c, 0x2d, 0xc6, 0x26, 0x16, 0x0f, 0xd1, 0xa4, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0xa4, 0x06,
	0x36, 0x48, 0x0f, 0x2a, 0x06, 0x36, 0x48, 0x0f, 0x8b, 0x26, 0x25, 0x57, 0x2e, 0x89, 0x00, 0xa8,
	0x70, 0x00, 0x5c, 0xd4, 0x19, 0xac, 0x4b, 0x48, 0x93, 0x4b, 0xa0, 0xb4, 0x20, 0x25, 0xb1, 0x24,
	0x35, 0x3e, 0xad, 0x28, 0xb5, 0xb0, 0x34, 0x35, 0x2f, 0xb9, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x35, 0x88, 0x1f, 0x22, 0xee, 0x06, 0x13, 0x4e, 0x62, 0x03, 0xdb, 0x6a, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xb5, 0xc6, 0x07, 0x1b, 0xc0, 0x00, 0x00, 0x00,
}