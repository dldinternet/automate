// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/request/root.proto

package request // import "github.com/chef/automate/api/external/cfgmgmt/request"

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

type Organizations struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Organizations) Reset()         { *m = Organizations{} }
func (m *Organizations) String() string { return proto.CompactTextString(m) }
func (*Organizations) ProtoMessage()    {}
func (*Organizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_2ac5a06cdbec6d6d, []int{0}
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

type SourceFqdns struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceFqdns) Reset()         { *m = SourceFqdns{} }
func (m *SourceFqdns) String() string { return proto.CompactTextString(m) }
func (*SourceFqdns) ProtoMessage()    {}
func (*SourceFqdns) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_2ac5a06cdbec6d6d, []int{1}
}
func (m *SourceFqdns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceFqdns.Unmarshal(m, b)
}
func (m *SourceFqdns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceFqdns.Marshal(b, m, deterministic)
}
func (dst *SourceFqdns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceFqdns.Merge(dst, src)
}
func (m *SourceFqdns) XXX_Size() int {
	return xxx_messageInfo_SourceFqdns.Size(m)
}
func (m *SourceFqdns) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceFqdns.DiscardUnknown(m)
}

var xxx_messageInfo_SourceFqdns proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Organizations)(nil), "chef.automate.api.cfgmgmt.request.Organizations")
	proto.RegisterType((*SourceFqdns)(nil), "chef.automate.api.cfgmgmt.request.SourceFqdns")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/request/root.proto", fileDescriptor_root_2ac5a06cdbec6d6d)
}

var fileDescriptor_root_2ac5a06cdbec6d6d = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xce, 0x31, 0xcb, 0xc2, 0x30,
	0x10, 0x80, 0xe1, 0xed, 0x1b, 0xf2, 0x51, 0x04, 0xff, 0x81, 0x5d, 0xdc, 0xee, 0x06, 0x11, 0x77,
	0x07, 0x57, 0x07, 0x37, 0xb7, 0x6b, 0xbc, 0xa6, 0x01, 0x93, 0x4b, 0x2f, 0x17, 0x10, 0x7f, 0xbd,
	0x28, 0x75, 0x75, 0x7f, 0x79, 0x79, 0xdc, 0x96, 0x4a, 0x44, 0x7e, 0x18, 0x6b, 0xa6, 0x3b, 0xfa,
	0x31, 0xa4, 0x90, 0x0c, 0x95, 0xe7, 0xc6, 0xd5, 0x50, 0x45, 0x0c, 0x8a, 0x8a, 0xc9, 0x7a, 0xe3,
	0x27, 0x1e, 0x81, 0x9a, 0x49, 0x22, 0x63, 0xa0, 0x12, 0x61, 0xa9, 0x61, 0xa9, 0xfb, 0x95, 0xeb,
	0xce, 0x1a, 0x28, 0xc7, 0x27, 0x59, 0x94, 0x5c, 0xfb, 0xce, 0xfd, 0x5f, 0xa4, 0xa9, 0xe7, 0xd3,
	0x7c, 0xcb, 0xf5, 0x78, 0xb8, 0xee, 0x43, 0xb4, 0xa9, 0x0d, 0xe0, 0x25, 0xe1, 0xfb, 0x87, 0xdf,
	0x1f, 0xfe, 0x62, 0x0c, 0x7f, 0x1f, 0xc2, 0xee, 0x15, 0x00, 0x00, 0xff, 0xff, 0x07, 0x80, 0xa6,
	0x92, 0xad, 0x00, 0x00, 0x00,
}