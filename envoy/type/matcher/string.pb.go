// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/string.proto

package matcher

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Specifies the way to match a string.
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_Regex
	MatchPattern         isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc62c75a0f154e3, []int{0}
}

func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMatcher.Unmarshal(m, b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return xxx_messageInfo_StringMatcher.Size(m)
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatcher_Regex struct {
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Prefix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Suffix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Regex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatcher) GetRegex() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Regex); ok {
		return x.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_Regex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc62c75a0f154e3, []int{1}
}

func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStringMatcher.Unmarshal(m, b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return xxx_messageInfo_ListStringMatcher.Size(m)
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.ListStringMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/string.proto", fileDescriptor_1dc62c75a0f154e3) }

var fileDescriptor_1dc62c75a0f154e3 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xeb, 0xa4, 0x69, 0x83, 0xab, 0x0e, 0x58, 0x08, 0x22, 0x16, 0xdc, 0x4e, 0x99, 0x1c,
	0x09, 0x36, 0x46, 0x2f, 0x74, 0x00, 0xa9, 0x0a, 0x1b, 0x42, 0x42, 0xa6, 0xb8, 0xc5, 0x12, 0xc4,
	0x96, 0x73, 0x54, 0xc9, 0xc6, 0x33, 0xf0, 0x0e, 0xbc, 0x23, 0xea, 0x84, 0x6c, 0x07, 0xa4, 0x2a,
	0xdd, 0x6c, 0xff, 0xdf, 0x77, 0x77, 0x3e, 0x7c, 0x21, 0xab, 0xad, 0x6e, 0x0b, 0x68, 0x8d, 0x2c,
	0xde, 0x05, 0xac, 0x5e, 0xa5, 0x2d, 0x6a, 0xb0, 0xaa, 0xda, 0x30, 0x63, 0x35, 0x68, 0x42, 0x3c,
	0xc0, 0x1c, 0xc0, 0x3a, 0xe0, 0xfc, 0x6c, 0x2b, 0xde, 0xd4, 0x8b, 0x00, 0x59, 0xfc, 0x1d, 0x02,
	0x3c, 0xff, 0x46, 0x78, 0x7a, 0xef, 0xed, 0xbb, 0x80, 0x92, 0x53, 0x9c, 0xc8, 0x46, 0xac, 0x20,
	0x43, 0x14, 0xe5, 0x47, 0x8b, 0x41, 0x19, 0xae, 0x64, 0x86, 0x47, 0xc6, 0xca, 0xb5, 0x6a, 0xb2,
	0xc8, 0x05, 0x7c, 0xbc, 0xe3, 0x43, 0x1b, 0x51, 0xb4, 0x18, 0x94, 0x5d, 0xe0, 0x90, 0xfa, 0x63,
	0xed, 0x90, 0xb8, 0x87, 0x84, 0x80, 0x50, 0x9c, 0x58, 0xb9, 0x91, 0x4d, 0x36, 0xf4, 0x44, 0xba,
	0xe3, 0x89, 0x8d, 0xf3, 0xcf, 0xd4, 0xf5, 0xf1, 0x01, 0x3f, 0xc1, 0x53, 0x3f, 0xf5, 0x93, 0x11,
	0x00, 0xd2, 0x56, 0x24, 0xfe, 0xe1, 0x68, 0xfe, 0x88, 0x8f, 0x6f, 0x55, 0x0d, 0xfb, 0xa3, 0xde,
	0xe0, 0xb4, 0x83, 0xea, 0x0c, 0xd1, 0x38, 0x9f, 0x5c, 0xce, 0x58, 0xff, 0xf3, 0x6c, 0x4f, 0xf2,
	0x2d, 0xbf, 0x50, 0x94, 0xa2, 0xf2, 0x5f, 0xe6, 0xd7, 0x98, 0x2a, 0x1d, 0x54, 0x63, 0x75, 0xd3,
	0x1e, 0xa8, 0xc2, 0x27, 0xa1, 0xcc, 0xd2, 0xad, 0x6d, 0x89, 0x1e, 0xc6, 0xdd, 0xfb, 0xf3, 0xc8,
	0x2f, 0xf2, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x17, 0x53, 0x56, 0x98, 0x01, 0x00, 0x00,
}
