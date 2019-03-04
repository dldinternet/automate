// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/request/tokens.proto

package request // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"

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

type CreateTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenReq) Reset()         { *m = CreateTokenReq{} }
func (m *CreateTokenReq) String() string { return proto.CompactTextString(m) }
func (*CreateTokenReq) ProtoMessage()    {}
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_e6ac700d926d65fa, []int{0}
}
func (m *CreateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenReq.Unmarshal(m, b)
}
func (m *CreateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenReq.Marshal(b, m, deterministic)
}
func (dst *CreateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenReq.Merge(dst, src)
}
func (m *CreateTokenReq) XXX_Size() int {
	return xxx_messageInfo_CreateTokenReq.Size(m)
}
func (m *CreateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenReq proto.InternalMessageInfo

func (m *CreateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTokenReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTokenReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *CreateTokenReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type GetTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenReq) Reset()         { *m = GetTokenReq{} }
func (m *GetTokenReq) String() string { return proto.CompactTextString(m) }
func (*GetTokenReq) ProtoMessage()    {}
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_e6ac700d926d65fa, []int{1}
}
func (m *GetTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenReq.Unmarshal(m, b)
}
func (m *GetTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenReq.Marshal(b, m, deterministic)
}
func (dst *GetTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenReq.Merge(dst, src)
}
func (m *GetTokenReq) XXX_Size() int {
	return xxx_messageInfo_GetTokenReq.Size(m)
}
func (m *GetTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenReq proto.InternalMessageInfo

func (m *GetTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateTokenReq struct {
	// ID can't be changed; ID used to discover token
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTokenReq) Reset()         { *m = UpdateTokenReq{} }
func (m *UpdateTokenReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenReq) ProtoMessage()    {}
func (*UpdateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_e6ac700d926d65fa, []int{2}
}
func (m *UpdateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenReq.Unmarshal(m, b)
}
func (m *UpdateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenReq.Marshal(b, m, deterministic)
}
func (dst *UpdateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenReq.Merge(dst, src)
}
func (m *UpdateTokenReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenReq.Size(m)
}
func (m *UpdateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenReq proto.InternalMessageInfo

func (m *UpdateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTokenReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTokenReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *UpdateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenReq) Reset()         { *m = DeleteTokenReq{} }
func (m *DeleteTokenReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenReq) ProtoMessage()    {}
func (*DeleteTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_e6ac700d926d65fa, []int{3}
}
func (m *DeleteTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenReq.Unmarshal(m, b)
}
func (m *DeleteTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenReq.Marshal(b, m, deterministic)
}
func (dst *DeleteTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenReq.Merge(dst, src)
}
func (m *DeleteTokenReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenReq.Size(m)
}
func (m *DeleteTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenReq proto.InternalMessageInfo

func (m *DeleteTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListTokensReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTokensReq) Reset()         { *m = ListTokensReq{} }
func (m *ListTokensReq) String() string { return proto.CompactTextString(m) }
func (*ListTokensReq) ProtoMessage()    {}
func (*ListTokensReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_e6ac700d926d65fa, []int{4}
}
func (m *ListTokensReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTokensReq.Unmarshal(m, b)
}
func (m *ListTokensReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTokensReq.Marshal(b, m, deterministic)
}
func (dst *ListTokensReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTokensReq.Merge(dst, src)
}
func (m *ListTokensReq) XXX_Size() int {
	return xxx_messageInfo_ListTokensReq.Size(m)
}
func (m *ListTokensReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTokensReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTokensReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTokenReq)(nil), "chef.automate.api.iam.v2beta.CreateTokenReq")
	proto.RegisterType((*GetTokenReq)(nil), "chef.automate.api.iam.v2beta.GetTokenReq")
	proto.RegisterType((*UpdateTokenReq)(nil), "chef.automate.api.iam.v2beta.UpdateTokenReq")
	proto.RegisterType((*DeleteTokenReq)(nil), "chef.automate.api.iam.v2beta.DeleteTokenReq")
	proto.RegisterType((*ListTokensReq)(nil), "chef.automate.api.iam.v2beta.ListTokensReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/request/tokens.proto", fileDescriptor_tokens_e6ac700d926d65fa)
}

var fileDescriptor_tokens_e6ac700d926d65fa = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0xb9, 0xbf, 0xdc, 0x8d, 0x18, 0x61, 0x11, 0x09, 0xa2, 0x10, 0x52, 0xa5, 0x71, 0x17,
	0xf4, 0x03, 0x08, 0x2a, 0x58, 0xa8, 0x4d, 0xd0, 0xc6, 0x6e, 0x92, 0x8c, 0x97, 0xd5, 0xdb, 0xec,
	0x26, 0x3b, 0x89, 0xd8, 0xf8, 0xd9, 0xe5, 0x12, 0xbd, 0xee, 0x04, 0xc1, 0x6e, 0xde, 0x30, 0x8f,
	0xf7, 0x98, 0x1f, 0x5c, 0xe6, 0xd6, 0x38, 0x5b, 0x51, 0xc5, 0x5e, 0x61, 0xcb, 0xd6, 0x20, 0xd3,
	0xd9, 0x0a, 0x99, 0xde, 0xf1, 0x43, 0xa1, 0xd3, 0x4a, 0xa3, 0x51, 0xdd, 0x79, 0x46, 0x8c, 0xaa,
	0xa1, 0xba, 0x25, 0xcf, 0x8a, 0xed, 0x1b, 0x55, 0x5e, 0xba, 0xc6, 0xb2, 0x15, 0x27, 0x79, 0x49,
	0x2f, 0xf2, 0xc7, 0x2a, 0xd1, 0x69, 0xa9, 0xd1, 0xc8, 0xc1, 0x12, 0x7f, 0x42, 0x70, 0xdd, 0x10,
	0x32, 0x3d, 0x6e, 0x3c, 0x29, 0xd5, 0x22, 0x80, 0xb1, 0x2e, 0xc2, 0x51, 0x34, 0x4a, 0x96, 0xe9,
	0x58, 0x17, 0x42, 0xc0, 0xb4, 0x42, 0x43, 0xe1, 0xb8, 0xdf, 0xf4, 0xb3, 0x38, 0x82, 0x39, 0xe6,
	0xac, 0x3b, 0x0a, 0x27, 0xd1, 0x28, 0x59, 0xa4, 0xdf, 0x4a, 0x1c, 0xc2, 0xac, 0xc3, 0x75, 0x4b,
	0xe1, 0xb4, 0x3f, 0x1e, 0x84, 0x38, 0x86, 0x85, 0x6b, 0xec, 0x2b, 0xe5, 0xec, 0xc3, 0x59, 0x34,
	0x49, 0x96, 0xe9, 0x56, 0xc7, 0xa7, 0xb0, 0x77, 0x4b, 0xbc, 0x2b, 0x3c, 0x2e, 0x21, 0x78, 0x72,
	0xc5, 0x7f, 0xd5, 0xfb, 0xad, 0x48, 0x04, 0xc1, 0x0d, 0xad, 0x69, 0x77, 0x52, 0x7c, 0x00, 0xfb,
	0xf7, 0xda, 0x0f, 0x5d, 0x7d, 0x4a, 0xf5, 0xd5, 0xc3, 0xf3, 0xdd, 0x4a, 0x73, 0xd9, 0x66, 0x32,
	0xb7, 0x46, 0x6d, 0xde, 0xbc, 0x25, 0xa4, 0xfe, 0x4e, 0x2d, 0x9b, 0xf7, 0xbc, 0x2e, 0xbe, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0x79, 0x73, 0x3c, 0xf2, 0x01, 0x00, 0x00,
}