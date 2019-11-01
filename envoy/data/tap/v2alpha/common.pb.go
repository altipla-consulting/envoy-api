// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/common.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
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

// Wrapper for tapped body data. This includes HTTP request/response body, transport socket received
// and transmitted data, etc.
type Body struct {
	// Types that are valid to be assigned to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType isBody_BodyType `protobuf_oneof:"body_type"`
	// Specifies whether body data has been truncated to fit within the specified
	// :ref:`max_buffered_rx_bytes
	// <envoy_api_field_service.tap.v2alpha.OutputConfig.max_buffered_rx_bytes>` and
	// :ref:`max_buffered_tx_bytes
	// <envoy_api_field_service.tap.v2alpha.OutputConfig.max_buffered_tx_bytes>` settings.
	Truncated            bool     `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_a560f1e2899ebe7a, []int{0}
}

func (m *Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Body.Unmarshal(m, b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Body.Marshal(b, m, deterministic)
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return xxx_messageInfo_Body.Size(m)
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

type isBody_BodyType interface {
	isBody_BodyType()
}

type Body_AsBytes struct {
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof"`
}

type Body_AsString struct {
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof"`
}

func (*Body_AsBytes) isBody_BodyType() {}

func (*Body_AsString) isBody_BodyType() {}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (m *Body) GetAsBytes() []byte {
	if x, ok := m.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (m *Body) GetAsString() string {
	if x, ok := m.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

func (m *Body) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Body) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
}

func init() {
	proto.RegisterType((*Body)(nil), "envoy.data.tap.v2alpha.Body")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/common.proto", fileDescriptor_a560f1e2899ebe7a)
}

var fileDescriptor_a560f1e2899ebe7a = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0xc6, 0xf1, 0x1b, 0x15, 0x6d, 0x72, 0x9d, 0x32, 0x48, 0x41, 0x85, 0xa2, 0x0e, 0x9d, 0x12,
	0x50, 0x7c, 0x81, 0xb8, 0x74, 0x2c, 0xf5, 0x01, 0xca, 0x69, 0x13, 0xb4, 0x60, 0x73, 0x42, 0x72,
	0x2c, 0xe6, 0xed, 0xa5, 0x41, 0x70, 0xb9, 0xeb, 0xc7, 0x7f, 0xf8, 0x7d, 0xe2, 0xd1, 0xf9, 0x0d,
	0xb3, 0xb6, 0x40, 0xa0, 0x09, 0x82, 0xde, 0x9e, 0xe1, 0x2b, 0x7c, 0x82, 0x9e, 0x71, 0x5d, 0xd1,
	0xab, 0x10, 0x91, 0x50, 0xde, 0x94, 0x48, 0xed, 0x91, 0x22, 0x08, 0xea, 0x2f, 0x7a, 0x58, 0xc5,
	0x85, 0x41, 0x9b, 0xe5, 0xad, 0xa8, 0x20, 0x8d, 0x53, 0x26, 0x97, 0x6a, 0xd6, 0xb0, 0xf6, 0xba,
	0x3b, 0x0c, 0x57, 0x90, 0xcc, 0x3e, 0xc8, 0x7b, 0xc1, 0x21, 0x8d, 0x89, 0xe2, 0xe2, 0x3f, 0xea,
	0xb3, 0x86, 0xb5, 0xbc, 0x3b, 0x0c, 0x15, 0xa4, 0xf7, 0xb2, 0xc8, 0x3b, 0xc1, 0x29, 0x7e, 0xfb,
	0x19, 0xc8, 0xd9, 0xfa, 0xbc, 0x61, 0x6d, 0x35, 0xfc, 0x0f, 0xe6, 0x28, 0xf8, 0x84, 0x36, 0x8f,
	0x94, 0x83, 0x33, 0xaf, 0xe2, 0x69, 0x41, 0x55, 0x2c, 0x21, 0xe2, 0x4f, 0x56, 0xa7, 0x59, 0xe6,
	0xf8, 0x56, 0xf0, 0xfd, 0x6e, 0xef, 0xd9, 0x74, 0x59, 0x4e, 0xbc, 0xfc, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xee, 0x15, 0x32, 0x80, 0xeb, 0x00, 0x00, 0x00,
}
