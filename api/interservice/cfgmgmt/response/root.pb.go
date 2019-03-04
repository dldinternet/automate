// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/response/root.proto

package response // import "github.com/chef/automate/api/interservice/cfgmgmt/response"

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

type VersionInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	SHA                  string   `protobuf:"bytes,3,opt,name=SHA,proto3" json:"SHA,omitempty" toml:"SHA,omitempty" mapstructure:"SHA,omitempty"`
	Built                string   `protobuf:"bytes,4,opt,name=built,proto3" json:"built,omitempty" toml:"built,omitempty" mapstructure:"built,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_9f8ce96ff5a7ad31, []int{0}
}
func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionInfo.Unmarshal(m, b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
}
func (dst *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(dst, src)
}
func (m *VersionInfo) XXX_Size() int {
	return xxx_messageInfo_VersionInfo.Size(m)
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

func (m *VersionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionInfo) GetSHA() string {
	if m != nil {
		return m.SHA
	}
	return ""
}

func (m *VersionInfo) GetBuilt() string {
	if m != nil {
		return m.Built
	}
	return ""
}

// Health message
//
// The config-mgmt-service health is constructed with:
// * Status:
//            => ok:             Everything is alright
//            => initialization: The service is in its initialization process
//            => warning:        Something might be wrong?
//            => critical:       Something is wrong!
//
// @afiune: Here we can add more health information to the response
type Health struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_9f8ce96ff5a7ad31, []int{1}
}
func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (dst *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(dst, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Organizations struct {
	Organizations        []string `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty" toml:"organizations,omitempty" mapstructure:"organizations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Organizations) Reset()         { *m = Organizations{} }
func (m *Organizations) String() string { return proto.CompactTextString(m) }
func (*Organizations) ProtoMessage()    {}
func (*Organizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_9f8ce96ff5a7ad31, []int{2}
}
func (m *Organizations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organizations.Unmarshal(m, b)
}
func (m *Organizations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organizations.Marshal(b, m, deterministic)
}
func (dst *Organizations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organizations.Merge(dst, src)
}
func (m *Organizations) XXX_Size() int {
	return xxx_messageInfo_Organizations.Size(m)
}
func (m *Organizations) XXX_DiscardUnknown() {
	xxx_messageInfo_Organizations.DiscardUnknown(m)
}

var xxx_messageInfo_Organizations proto.InternalMessageInfo

func (m *Organizations) GetOrganizations() []string {
	if m != nil {
		return m.Organizations
	}
	return nil
}

type SourceFQDNS struct {
	SourceFqdns          []string `protobuf:"bytes,1,rep,name=source_fqdns,json=sourceFqdns,proto3" json:"source_fqdns,omitempty" toml:"source_fqdns,omitempty" mapstructure:"source_fqdns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *SourceFQDNS) Reset()         { *m = SourceFQDNS{} }
func (m *SourceFQDNS) String() string { return proto.CompactTextString(m) }
func (*SourceFQDNS) ProtoMessage()    {}
func (*SourceFQDNS) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_9f8ce96ff5a7ad31, []int{3}
}
func (m *SourceFQDNS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceFQDNS.Unmarshal(m, b)
}
func (m *SourceFQDNS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceFQDNS.Marshal(b, m, deterministic)
}
func (dst *SourceFQDNS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceFQDNS.Merge(dst, src)
}
func (m *SourceFQDNS) XXX_Size() int {
	return xxx_messageInfo_SourceFQDNS.Size(m)
}
func (m *SourceFQDNS) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceFQDNS.DiscardUnknown(m)
}

var xxx_messageInfo_SourceFQDNS proto.InternalMessageInfo

func (m *SourceFQDNS) GetSourceFqdns() []string {
	if m != nil {
		return m.SourceFqdns
	}
	return nil
}

type ExportData struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty" toml:"content,omitempty" mapstructure:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ExportData) Reset()         { *m = ExportData{} }
func (m *ExportData) String() string { return proto.CompactTextString(m) }
func (*ExportData) ProtoMessage()    {}
func (*ExportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_9f8ce96ff5a7ad31, []int{4}
}
func (m *ExportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportData.Unmarshal(m, b)
}
func (m *ExportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportData.Marshal(b, m, deterministic)
}
func (dst *ExportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportData.Merge(dst, src)
}
func (m *ExportData) XXX_Size() int {
	return xxx_messageInfo_ExportData.Size(m)
}
func (m *ExportData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportData.DiscardUnknown(m)
}

var xxx_messageInfo_ExportData proto.InternalMessageInfo

func (m *ExportData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "chef.automate.domain.cfgmgmt.response.VersionInfo")
	proto.RegisterType((*Health)(nil), "chef.automate.domain.cfgmgmt.response.Health")
	proto.RegisterType((*Organizations)(nil), "chef.automate.domain.cfgmgmt.response.Organizations")
	proto.RegisterType((*SourceFQDNS)(nil), "chef.automate.domain.cfgmgmt.response.SourceFQDNS")
	proto.RegisterType((*ExportData)(nil), "chef.automate.domain.cfgmgmt.response.ExportData")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/response/root.proto", fileDescriptor_root_9f8ce96ff5a7ad31)
}

var fileDescriptor_root_9f8ce96ff5a7ad31 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x99, 0x9b, 0x93, 0xbd, 0x6d, 0x20, 0x41, 0xa4, 0xc7, 0x59, 0x54, 0x76, 0x90, 0x46,
	0x10, 0x2f, 0xe2, 0x45, 0x99, 0x63, 0x5e, 0x14, 0x37, 0xf0, 0xe0, 0x45, 0xb2, 0xec, 0xb5, 0x0b,
	0xac, 0x79, 0x35, 0x79, 0x1d, 0xe2, 0x5f, 0x2f, 0xcd, 0x5a, 0xd1, 0x93, 0xb7, 0xf7, 0xfb, 0x91,
	0x8f, 0x24, 0xef, 0x83, 0x0b, 0x55, 0x18, 0x69, 0x2c, 0xa3, 0xf3, 0xe8, 0xb6, 0x46, 0xa3, 0xd4,
	0x69, 0x96, 0x67, 0x39, 0x4b, 0x87, 0xbe, 0x20, 0xeb, 0x51, 0x3a, 0x22, 0x4e, 0x0a, 0x47, 0x4c,
	0xe2, 0x4c, 0xaf, 0x31, 0x4d, 0x54, 0xc9, 0x94, 0x2b, 0xc6, 0x64, 0x45, 0xb9, 0x32, 0x36, 0xa9,
	0x13, 0x49, 0x93, 0x88, 0x35, 0xf4, 0x5f, 0xd1, 0x79, 0x43, 0xf6, 0xd1, 0xa6, 0x24, 0x04, 0x74,
	0xac, 0xca, 0x31, 0x6a, 0x8d, 0x5a, 0xe3, 0xde, 0x3c, 0xcc, 0x22, 0x82, 0x83, 0xed, 0xee, 0x48,
	0xb4, 0x17, 0x74, 0x83, 0xe2, 0x10, 0xda, 0x8b, 0xd9, 0x5d, 0xd4, 0x0e, 0xb6, 0x1a, 0xc5, 0x11,
	0xec, 0x2f, 0x4b, 0xb3, 0xe1, 0xa8, 0x13, 0xdc, 0x0e, 0xe2, 0x11, 0x74, 0x67, 0xa8, 0x36, 0xbc,
	0x16, 0xc7, 0xd0, 0xf5, 0xac, 0xb8, 0xf4, 0xf5, 0x0d, 0x35, 0xc5, 0xd7, 0x30, 0x7c, 0x76, 0x99,
	0xb2, 0xe6, 0x4b, 0xb1, 0x21, 0xeb, 0xc5, 0x29, 0x0c, 0xe9, 0xb7, 0x88, 0x5a, 0xa3, 0xf6, 0xb8,
	0x37, 0xff, 0x2b, 0xe3, 0x4b, 0xe8, 0x2f, 0xa8, 0x74, 0x1a, 0xa7, 0x2f, 0x93, 0xa7, 0x85, 0x38,
	0x81, 0x81, 0x0f, 0xf8, 0x9e, 0x7e, 0xac, 0x7e, 0x32, 0xfd, 0x9d, 0x9b, 0x56, 0x2a, 0x3e, 0x07,
	0x78, 0xf8, 0x2c, 0xc8, 0xf1, 0x44, 0xb1, 0xaa, 0xbe, 0xa6, 0xc9, 0x32, 0x5a, 0x0e, 0xef, 0x19,
	0xcc, 0x1b, 0xbc, 0xbf, 0x7d, 0xbb, 0xc9, 0x0c, 0xaf, 0xcb, 0x65, 0xa2, 0x29, 0x97, 0xd5, 0x2e,
	0x65, 0xb3, 0x4b, 0xf9, 0x6f, 0x0f, 0xcb, 0x6e, 0xe8, 0xe0, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0x70, 0x37, 0xa4, 0xb3, 0x01, 0x00, 0x00,
}