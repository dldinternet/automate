// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/response/actions.proto

package response // import "github.com/chef/automate/api/external/cfgmgmt/response"

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

type PolicyCookbooks struct {
	PolicyName           string          `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	CookbookLocks        []*CookbookLock `protobuf:"bytes,2,rep,name=cookbook_locks,json=cookbookLocks,proto3" json:"cookbook_locks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PolicyCookbooks) Reset()         { *m = PolicyCookbooks{} }
func (m *PolicyCookbooks) String() string { return proto.CompactTextString(m) }
func (*PolicyCookbooks) ProtoMessage()    {}
func (*PolicyCookbooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_actions_974f3bcfbb65b831, []int{0}
}
func (m *PolicyCookbooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyCookbooks.Unmarshal(m, b)
}
func (m *PolicyCookbooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyCookbooks.Marshal(b, m, deterministic)
}
func (dst *PolicyCookbooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyCookbooks.Merge(dst, src)
}
func (m *PolicyCookbooks) XXX_Size() int {
	return xxx_messageInfo_PolicyCookbooks.Size(m)
}
func (m *PolicyCookbooks) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyCookbooks.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyCookbooks proto.InternalMessageInfo

func (m *PolicyCookbooks) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *PolicyCookbooks) GetCookbookLocks() []*CookbookLock {
	if m != nil {
		return m.CookbookLocks
	}
	return nil
}

type CookbookLock struct {
	Cookbook             string   `protobuf:"bytes,1,opt,name=cookbook,proto3" json:"cookbook,omitempty"`
	PolicyIdentifier     string   `protobuf:"bytes,2,opt,name=policy_identifier,json=policyIdentifier,proto3" json:"policy_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookLock) Reset()         { *m = CookbookLock{} }
func (m *CookbookLock) String() string { return proto.CompactTextString(m) }
func (*CookbookLock) ProtoMessage()    {}
func (*CookbookLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_actions_974f3bcfbb65b831, []int{1}
}
func (m *CookbookLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookLock.Unmarshal(m, b)
}
func (m *CookbookLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookLock.Marshal(b, m, deterministic)
}
func (dst *CookbookLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookLock.Merge(dst, src)
}
func (m *CookbookLock) XXX_Size() int {
	return xxx_messageInfo_CookbookLock.Size(m)
}
func (m *CookbookLock) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookLock.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookLock proto.InternalMessageInfo

func (m *CookbookLock) GetCookbook() string {
	if m != nil {
		return m.Cookbook
	}
	return ""
}

func (m *CookbookLock) GetPolicyIdentifier() string {
	if m != nil {
		return m.PolicyIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*PolicyCookbooks)(nil), "chef.automate.api.cfgmgmt.response.PolicyCookbooks")
	proto.RegisterType((*CookbookLock)(nil), "chef.automate.api.cfgmgmt.response.CookbookLock")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/response/actions.proto", fileDescriptor_actions_974f3bcfbb65b831)
}

var fileDescriptor_actions_974f3bcfbb65b831 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x05, 0xd1, 0xad, 0x7f, 0xf7, 0x14, 0xbc, 0x58, 0x72, 0x2a, 0x14, 0x76, 0x45,
	0x41, 0x3c, 0xeb, 0x49, 0x10, 0x91, 0x5e, 0x0a, 0x5e, 0xca, 0x66, 0x9c, 0xa4, 0x4b, 0xb2, 0x3b,
	0x4b, 0x76, 0x0a, 0xfa, 0x19, 0xfc, 0xd2, 0x62, 0x9a, 0x2d, 0x39, 0xe9, 0x71, 0xde, 0xbc, 0xc7,
	0xfb, 0xf1, 0xc4, 0xc2, 0x04, 0xab, 0xf1, 0x93, 0xb1, 0xf3, 0xa6, 0xd5, 0x50, 0xd5, 0xae, 0x76,
	0xac, 0x3b, 0x8c, 0x81, 0x7c, 0x44, 0x6d, 0x80, 0x2d, 0xf9, 0xa8, 0x42, 0x47, 0x4c, 0xb2, 0x80,
	0x0d, 0x56, 0xca, 0x6c, 0x99, 0x9c, 0x61, 0x54, 0x26, 0x58, 0x35, 0x24, 0x54, 0x4a, 0x14, 0xdf,
	0x99, 0x38, 0x7f, 0xa3, 0xd6, 0xc2, 0xd7, 0x13, 0x51, 0x53, 0x12, 0x35, 0x51, 0x5e, 0x8b, 0x69,
	0xe8, 0xa5, 0xb5, 0x37, 0x0e, 0xf3, 0x6c, 0x96, 0xcd, 0x8f, 0x97, 0x62, 0x27, 0xbd, 0x1a, 0x87,
	0x72, 0x25, 0xce, 0x60, 0x70, 0xaf, 0x5b, 0x82, 0x26, 0xe6, 0x93, 0xd9, 0xc1, 0x7c, 0x7a, 0x7b,
	0xa3, 0xfe, 0x6f, 0x54, 0xa9, 0xe7, 0x85, 0xa0, 0x59, 0x9e, 0xc2, 0xe8, 0x8a, 0xc5, 0x4a, 0x9c,
	0x8c, 0xdf, 0xf2, 0x4a, 0x1c, 0x25, 0xc3, 0x80, 0xb1, 0xbf, 0xe5, 0x42, 0x5c, 0x0e, 0x94, 0xf6,
	0x03, 0x3d, 0xdb, 0xca, 0x62, 0x97, 0x4f, 0x7a, 0xd3, 0xc5, 0xee, 0xf1, 0xbc, 0xd7, 0x1f, 0x1f,
	0xde, 0xef, 0x6b, 0xcb, 0x9b, 0x6d, 0xa9, 0x80, 0x9c, 0xfe, 0xa5, 0xd4, 0x89, 0x52, 0xff, 0x39,
	0x69, 0x79, 0xd8, 0x6f, 0x79, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x98, 0x74, 0x1a, 0x7a,
	0x01, 0x00, 0x00,
}