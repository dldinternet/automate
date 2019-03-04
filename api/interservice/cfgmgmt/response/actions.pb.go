// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/response/actions.proto

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

type PolicyCookbooks struct {
	PolicyName           string          `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	CookbookLocks        []*CookbookLock `protobuf:"bytes,2,rep,name=cookbook_locks,json=cookbookLocks,proto3" json:"cookbook_locks,omitempty" toml:"cookbook_locks,omitempty" mapstructure:"cookbook_locks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32           `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PolicyCookbooks) Reset()         { *m = PolicyCookbooks{} }
func (m *PolicyCookbooks) String() string { return proto.CompactTextString(m) }
func (*PolicyCookbooks) ProtoMessage()    {}
func (*PolicyCookbooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_actions_df7e901cdc1664ff, []int{0}
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
	Cookbook             string   `protobuf:"bytes,1,opt,name=cookbook,proto3" json:"cookbook,omitempty" toml:"cookbook,omitempty" mapstructure:"cookbook,omitempty"`
	PolicyIdentifier     string   `protobuf:"bytes,2,opt,name=policy_identifier,json=policyIdentifier,proto3" json:"policy_identifier,omitempty" toml:"policy_identifier,omitempty" mapstructure:"policy_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CookbookLock) Reset()         { *m = CookbookLock{} }
func (m *CookbookLock) String() string { return proto.CompactTextString(m) }
func (*CookbookLock) ProtoMessage()    {}
func (*CookbookLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_actions_df7e901cdc1664ff, []int{1}
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
	proto.RegisterType((*PolicyCookbooks)(nil), "chef.automate.domain.cfgmgmt.response.PolicyCookbooks")
	proto.RegisterType((*CookbookLock)(nil), "chef.automate.domain.cfgmgmt.response.CookbookLock")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/response/actions.proto", fileDescriptor_actions_df7e901cdc1664ff)
}

var fileDescriptor_actions_df7e901cdc1664ff = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x05, 0xd1, 0xad, 0x7f, 0xf7, 0x14, 0xbc, 0x58, 0x0a, 0x42, 0x41, 0xd8, 0x05,
	0x7b, 0x13, 0x4f, 0x7a, 0x12, 0x44, 0xa4, 0x17, 0xa1, 0x97, 0xb2, 0x99, 0x4e, 0xd2, 0x21, 0xdd,
	0x9d, 0xb0, 0x3b, 0x15, 0xfc, 0x12, 0x7e, 0x66, 0x31, 0x4d, 0x4a, 0x6f, 0xf6, 0x38, 0x6f, 0xde,
	0xe3, 0xfd, 0x78, 0xca, 0xba, 0x86, 0x2c, 0x05, 0xc1, 0x98, 0x30, 0x7e, 0x11, 0xa0, 0x85, 0xb2,
	0xf2, 0x95, 0x17, 0x1b, 0x31, 0x35, 0x1c, 0x12, 0x5a, 0x07, 0x42, 0x1c, 0x92, 0x69, 0x22, 0x0b,
	0xeb, 0x3b, 0x58, 0x61, 0x69, 0xdc, 0x46, 0xd8, 0x3b, 0x41, 0xb3, 0x64, 0xef, 0x28, 0x98, 0x2e,
	0x64, 0xfa, 0xd0, 0xf8, 0x27, 0x53, 0x97, 0x1f, 0xbc, 0x26, 0xf8, 0x7e, 0x61, 0xae, 0x0b, 0xe6,
	0x3a, 0xe9, 0x5b, 0x35, 0x6c, 0x5a, 0x69, 0x11, 0x9c, 0xc7, 0x3c, 0x1b, 0x65, 0x93, 0xd3, 0x99,
	0xda, 0x4a, 0xef, 0xce, 0xa3, 0x9e, 0xab, 0x0b, 0xe8, 0xdc, 0x8b, 0x35, 0x43, 0x9d, 0xf2, 0xc1,
	0xe8, 0x68, 0x32, 0x7c, 0x98, 0x9a, 0x83, 0x4a, 0x4d, 0x5f, 0xf5, 0xc6, 0x50, 0xcf, 0xce, 0x61,
	0xef, 0x4a, 0xe3, 0x4f, 0x75, 0xb6, 0xff, 0xd6, 0x37, 0xea, 0xa4, 0x37, 0x74, 0x24, 0xbb, 0x5b,
	0xdf, 0xab, 0xeb, 0x0e, 0x94, 0x96, 0x18, 0x84, 0x4a, 0xc2, 0x98, 0x0f, 0x5a, 0xd3, 0xd5, 0xf6,
	0xf1, 0xba, 0xd3, 0x9f, 0x9f, 0xe6, 0x8f, 0x15, 0xc9, 0x6a, 0x53, 0x18, 0x60, 0x6f, 0xff, 0x40,
	0x6d, 0x0f, 0xfa, 0xff, 0xb8, 0xc5, 0x71, 0xbb, 0xea, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x38,
	0xe8, 0x93, 0x82, 0x88, 0x01, 0x00, 0x00,
}