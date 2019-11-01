// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/certs.proto

package envoy_admin_v2alpha

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

// Proto representation of certificate details. Admin endpoint uses this wrapper for `/certs` to
// display certificate information. See :ref:`/certs <operations_admin_interface_certs>` for more
// information.
type Certificates struct {
	// List of certificates known to an Envoy.
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{0}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	// Details of CA certificate.
	CaCert []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Details of Certificate Chain
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

// [#next-free-field: 7]
type CertificateDetails struct {
	// Path of the certificate.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Certificate Serial Number.
	SerialNumber string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// List of Subject Alternate names.
	SubjectAltNames []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	// Minimum of days until expiration of certificate and it's chain.
	DaysUntilExpiration uint64 `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	// Indicates the time from which the certificate is valid.
	ValidFrom *types.Timestamp `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	// Indicates the time at which the certificate expires.
	ExpirationTime       *types.Timestamp `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{2}
}

func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateDetails.Unmarshal(m, b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return xxx_messageInfo_CertificateDetails.Size(m)
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *types.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *types.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Subject Alternate Name.
	//
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{3}
}

func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectAlternateName.Unmarshal(m, b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return xxx_messageInfo_SubjectAlternateName.Size(m)
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}

type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
	}
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v2alpha.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v2alpha.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v2alpha.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v2alpha.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/certs.proto", fileDescriptor_c7cd1796e05ff7fa) }

var fileDescriptor_c7cd1796e05ff7fa = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x08, 0xca, 0x24, 0x50, 0xb1, 0x05, 0xc9, 0xea, 0xa5, 0x26, 0x1c, 0x08, 0x17,
	0x5b, 0x84, 0x13, 0x37, 0x48, 0x4a, 0xc4, 0xa9, 0xaa, 0x4c, 0x7b, 0x5e, 0x4d, 0x9c, 0x4d, 0xb3,
	0x68, 0x7f, 0xac, 0xdd, 0x75, 0xd4, 0x3c, 0x09, 0xaf, 0xc6, 0xe3, 0xa0, 0x59, 0xf7, 0x27, 0x88,
	0x48, 0x55, 0x6f, 0x9e, 0xef, 0x67, 0x3c, 0xfe, 0xfc, 0xc1, 0xa9, 0x30, 0x5b, 0xbb, 0x2b, 0x70,
	0xa5, 0xa5, 0x29, 0xb6, 0x53, 0x54, 0xf5, 0x06, 0x8b, 0x4a, 0xb8, 0xe0, 0xf3, 0xda, 0xd9, 0x60,
	0xd9, 0x71, 0x14, 0xe4, 0x51, 0x90, 0xdf, 0x0a, 0x4e, 0x4e, 0xaf, 0xad, 0xbd, 0x56, 0xa2, 0x88,
	0x92, 0x65, 0xb3, 0x2e, 0x82, 0xd4, 0xc2, 0x07, 0xd4, 0x75, 0xeb, 0x1a, 0x5f, 0xc2, 0x68, 0x2e,
	0x5c, 0x90, 0x6b, 0x59, 0x61, 0x10, 0x9e, 0x9d, 0xc1, 0xa8, 0xda, 0x9b, 0xd3, 0x24, 0xeb, 0x4e,
	0x86, 0xd3, 0x2c, 0x3f, 0xb0, 0x3c, 0xdf, 0x33, 0x96, 0xff, 0xb8, 0xc6, 0xbf, 0x13, 0x18, 0xee,
	0xb1, 0xec, 0x2b, 0xbc, 0xa8, 0x90, 0x93, 0xe4, 0x76, 0xe1, 0x87, 0xc7, 0x16, 0x9e, 0x89, 0x80,
	0x52, 0xf9, 0xb2, 0x5f, 0x21, 0xa1, 0x6c, 0x01, 0x40, 0x76, 0x5e, 0x6d, 0x50, 0x9a, 0xb4, 0xf3,
	0xb4, 0x25, 0x03, 0xb2, 0xce, 0xc9, 0x39, 0xfe, 0xd3, 0x01, 0xf6, 0xbf, 0x82, 0x31, 0xe8, 0xd5,
	0x18, 0x36, 0x69, 0x92, 0x25, 0x93, 0x41, 0x19, 0x9f, 0xd9, 0x7b, 0x78, 0xe9, 0x85, 0x93, 0xa8,
	0xb8, 0x69, 0xf4, 0x52, 0xb8, 0xb4, 0x13, 0xc9, 0x51, 0x0b, 0x9e, 0x47, 0x8c, 0x5d, 0xc1, 0x6b,
	0xdf, 0x2c, 0x7f, 0x89, 0x2a, 0x70, 0x54, 0x81, 0x1b, 0xd4, 0xc2, 0xa7, 0xdd, 0x78, 0xde, 0xc7,
	0x83, 0xe7, 0xfd, 0x6c, 0xd5, 0xdf, 0x54, 0x10, 0xce, 0x60, 0x10, 0xe7, 0xa8, 0x45, 0x79, 0xe4,
	0xef, 0x51, 0x9a, 0x3d, 0x9b, 0xc2, 0xdb, 0x15, 0xee, 0x3c, 0x6f, 0x4c, 0x90, 0x8a, 0x8b, 0x9b,
	0x5a, 0x3a, 0x0c, 0xd2, 0x9a, 0xb4, 0x97, 0x25, 0x93, 0x5e, 0x79, 0x4c, 0xe4, 0x15, 0x71, 0xdf,
	0xef, 0x29, 0xf6, 0x05, 0x60, 0x8b, 0x4a, 0xae, 0xf8, 0xda, 0x59, 0x9d, 0x3e, 0xcf, 0x92, 0xc9,
	0x70, 0x7a, 0x92, 0xb7, 0x05, 0xc8, 0xef, 0x0a, 0x90, 0x5f, 0xde, 0x15, 0xa0, 0x1c, 0x44, 0xf5,
	0xc2, 0x59, 0xcd, 0xe6, 0x70, 0xf4, 0xf0, 0x0e, 0x4e, 0x1d, 0x49, 0xfb, 0x8f, 0xfa, 0x5f, 0x3d,
	0x58, 0x08, 0x1c, 0x2f, 0xe0, 0xcd, 0xa1, 0x8f, 0x63, 0x0c, 0xba, 0x2b, 0xe3, 0xdb, 0x68, 0x7f,
	0x3c, 0x2b, 0x69, 0x20, 0xac, 0x71, 0xb2, 0x4d, 0x94, 0xb0, 0xc6, 0xc9, 0x59, 0x1f, 0x7a, 0x14,
	0xdf, 0xec, 0x13, 0xbc, 0x93, 0xb6, 0xcd, 0xae, 0x76, 0xf6, 0x66, 0x77, 0x28, 0xc6, 0x19, 0xd0,
	0x4f, 0xf4, 0x17, 0x74, 0xd5, 0x45, 0xb2, 0xec, 0xc7, 0xf3, 0x3e, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xca, 0x32, 0x06, 0xd0, 0x25, 0x03, 0x00, 0x00,
}
