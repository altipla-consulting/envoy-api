// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/quic_config.proto

package envoy_api_v3alpha_listener

import (
	fmt "fmt"
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

// Configuration specific to the QUIC protocol.
// Next id: 4
type QuicProtocolOptions struct {
	// Maximum number of streams that the client can negotiate per connection. 100
	// if not specified.
	MaxConcurrentStreams *types.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	// Maximum number of milliseconds that connection will be alive when there is
	// no network activity. 300000ms if not specified.
	IdleTimeout *types.Duration `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Connection timeout in milliseconds before the crypto handshake is finished.
	// 20000ms if not specified.
	CryptoHandshakeTimeout *types.Duration `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *QuicProtocolOptions) Reset()         { *m = QuicProtocolOptions{} }
func (m *QuicProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*QuicProtocolOptions) ProtoMessage()    {}
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_30110b89b5635b16, []int{0}
}

func (m *QuicProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuicProtocolOptions.Unmarshal(m, b)
}
func (m *QuicProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuicProtocolOptions.Marshal(b, m, deterministic)
}
func (m *QuicProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuicProtocolOptions.Merge(m, src)
}
func (m *QuicProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_QuicProtocolOptions.Size(m)
}
func (m *QuicProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QuicProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QuicProtocolOptions proto.InternalMessageInfo

func (m *QuicProtocolOptions) GetMaxConcurrentStreams() *types.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *QuicProtocolOptions) GetIdleTimeout() *types.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *QuicProtocolOptions) GetCryptoHandshakeTimeout() *types.Duration {
	if m != nil {
		return m.CryptoHandshakeTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*QuicProtocolOptions)(nil), "envoy.api.v3alpha.listener.QuicProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/listener/quic_config.proto", fileDescriptor_30110b89b5635b16)
}

var fileDescriptor_30110b89b5635b16 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x39, 0x05, 0x87, 0x54, 0x10, 0x4e, 0x29, 0xb5, 0x88, 0x88, 0x53, 0x07, 0x49, 0xa0,
	0x5d, 0x75, 0x69, 0x1d, 0x74, 0xb2, 0xb6, 0xea, 0x7a, 0xbc, 0xa6, 0x69, 0x1b, 0xcc, 0xe5, 0xc5,
	0xfc, 0xa8, 0xbd, 0xff, 0xdd, 0x41, 0x2e, 0xb9, 0xd3, 0xa1, 0x88, 0x63, 0x78, 0xdf, 0xcf, 0x27,
	0xc9, 0xf7, 0x91, 0x1b, 0xa1, 0xb7, 0x58, 0x31, 0x30, 0x92, 0x6d, 0x47, 0xa0, 0xcc, 0x06, 0x98,
	0x92, 0xce, 0x0b, 0x2d, 0x2c, 0xfb, 0x08, 0x92, 0x17, 0x1c, 0xf5, 0x4a, 0xae, 0xa9, 0xb1, 0xe8,
	0x31, 0xef, 0xc7, 0x34, 0x05, 0x23, 0x69, 0x93, 0xa6, 0x6d, 0xba, 0x7f, 0xb9, 0x46, 0x5c, 0x2b,
	0xc1, 0x62, 0x72, 0x11, 0x56, 0x6c, 0x19, 0x2c, 0x78, 0x89, 0x3a, 0xb1, 0xfb, 0xf3, 0x4f, 0x0b,
	0xc6, 0x08, 0xeb, 0xd2, 0xfc, 0xfa, 0x2b, 0x23, 0xa7, 0xcf, 0x41, 0xf2, 0x69, 0x7d, 0xe2, 0xa8,
	0x9e, 0x4c, 0x0d, 0xbb, 0x7c, 0x46, 0xba, 0x25, 0xec, 0xea, 0x77, 0xf0, 0x60, 0xad, 0xd0, 0xbe,
	0x70, 0xde, 0x0a, 0x28, 0x5d, 0x2f, 0xbb, 0xca, 0x06, 0x9d, 0xe1, 0x05, 0x4d, 0x62, 0xda, 0x8a,
	0xe9, 0xeb, 0xa3, 0xf6, 0xa3, 0xe1, 0x1b, 0xa8, 0x20, 0x66, 0x67, 0x25, 0xec, 0x26, 0x3f, 0xe8,
	0x3c, 0x91, 0xf9, 0x2d, 0x39, 0x96, 0x4b, 0x25, 0x0a, 0x2f, 0x4b, 0x81, 0xc1, 0xf7, 0x0e, 0xa2,
	0xe9, 0x7c, 0xcf, 0x74, 0xdf, 0x7c, 0x61, 0xd6, 0xa9, 0xe3, 0x2f, 0x29, 0x9d, 0xcf, 0x49, 0x8f,
	0xdb, 0xca, 0x78, 0x2c, 0x36, 0xa0, 0x97, 0x6e, 0x03, 0xef, 0xbf, 0xa6, 0xc3, 0xff, 0x4c, 0xdd,
	0x84, 0x3e, 0xb4, 0x64, 0x23, 0x1d, 0xdf, 0x91, 0x81, 0x44, 0x1a, 0xfb, 0x35, 0x16, 0x77, 0x15,
	0xfd, 0xbb, 0xea, 0xf1, 0x49, 0xdd, 0xd3, 0x24, 0x2e, 0x26, 0xb6, 0x35, 0xcd, 0x16, 0x47, 0xf1,
	0xa6, 0xd1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x37, 0x3c, 0x0f, 0xd0, 0x01, 0x00, 0x00,
}