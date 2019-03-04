// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/response/teams.proto

package response // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetTeamResp) Reset()         { *m = GetTeamResp{} }
func (m *GetTeamResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamResp) ProtoMessage()    {}
func (*GetTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{0}
}
func (m *GetTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamResp.Unmarshal(m, b)
}
func (m *GetTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamResp.Marshal(b, m, deterministic)
}
func (dst *GetTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamResp.Merge(dst, src)
}
func (m *GetTeamResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamResp.Size(m)
}
func (m *GetTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamResp proto.InternalMessageInfo

func (m *GetTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type GetTeamsResp struct {
	Teams                []*common.Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTeamsResp) Reset()         { *m = GetTeamsResp{} }
func (m *GetTeamsResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamsResp) ProtoMessage()    {}
func (*GetTeamsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{1}
}
func (m *GetTeamsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsResp.Unmarshal(m, b)
}
func (m *GetTeamsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsResp.Marshal(b, m, deterministic)
}
func (dst *GetTeamsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsResp.Merge(dst, src)
}
func (m *GetTeamsResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamsResp.Size(m)
}
func (m *GetTeamsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsResp proto.InternalMessageInfo

func (m *GetTeamsResp) GetTeams() []*common.Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type CreateTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateTeamResp) Reset()         { *m = CreateTeamResp{} }
func (m *CreateTeamResp) String() string { return proto.CompactTextString(m) }
func (*CreateTeamResp) ProtoMessage()    {}
func (*CreateTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{2}
}
func (m *CreateTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamResp.Unmarshal(m, b)
}
func (m *CreateTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamResp.Marshal(b, m, deterministic)
}
func (dst *CreateTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamResp.Merge(dst, src)
}
func (m *CreateTeamResp) XXX_Size() int {
	return xxx_messageInfo_CreateTeamResp.Size(m)
}
func (m *CreateTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamResp proto.InternalMessageInfo

func (m *CreateTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type UpdateTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateTeamResp) Reset()         { *m = UpdateTeamResp{} }
func (m *UpdateTeamResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamResp) ProtoMessage()    {}
func (*UpdateTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{3}
}
func (m *UpdateTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamResp.Unmarshal(m, b)
}
func (m *UpdateTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamResp.Marshal(b, m, deterministic)
}
func (dst *UpdateTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamResp.Merge(dst, src)
}
func (m *UpdateTeamResp) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamResp.Size(m)
}
func (m *UpdateTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamResp proto.InternalMessageInfo

func (m *UpdateTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type DeleteTeamResp struct {
	Team                 *common.Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteTeamResp) Reset()         { *m = DeleteTeamResp{} }
func (m *DeleteTeamResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTeamResp) ProtoMessage()    {}
func (*DeleteTeamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{4}
}
func (m *DeleteTeamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTeamResp.Unmarshal(m, b)
}
func (m *DeleteTeamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTeamResp.Marshal(b, m, deterministic)
}
func (dst *DeleteTeamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTeamResp.Merge(dst, src)
}
func (m *DeleteTeamResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTeamResp.Size(m)
}
func (m *DeleteTeamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTeamResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTeamResp proto.InternalMessageInfo

func (m *DeleteTeamResp) GetTeam() *common.Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type AddTeamMembersResp struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTeamMembersResp) Reset()         { *m = AddTeamMembersResp{} }
func (m *AddTeamMembersResp) String() string { return proto.CompactTextString(m) }
func (*AddTeamMembersResp) ProtoMessage()    {}
func (*AddTeamMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{5}
}
func (m *AddTeamMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTeamMembersResp.Unmarshal(m, b)
}
func (m *AddTeamMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTeamMembersResp.Marshal(b, m, deterministic)
}
func (dst *AddTeamMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTeamMembersResp.Merge(dst, src)
}
func (m *AddTeamMembersResp) XXX_Size() int {
	return xxx_messageInfo_AddTeamMembersResp.Size(m)
}
func (m *AddTeamMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTeamMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddTeamMembersResp proto.InternalMessageInfo

func (m *AddTeamMembersResp) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamMembershipResp struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamMembershipResp) Reset()         { *m = GetTeamMembershipResp{} }
func (m *GetTeamMembershipResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamMembershipResp) ProtoMessage()    {}
func (*GetTeamMembershipResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{6}
}
func (m *GetTeamMembershipResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamMembershipResp.Unmarshal(m, b)
}
func (m *GetTeamMembershipResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamMembershipResp.Marshal(b, m, deterministic)
}
func (dst *GetTeamMembershipResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamMembershipResp.Merge(dst, src)
}
func (m *GetTeamMembershipResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamMembershipResp.Size(m)
}
func (m *GetTeamMembershipResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamMembershipResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamMembershipResp proto.InternalMessageInfo

func (m *GetTeamMembershipResp) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type RemoveTeamMembersResp struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamMembersResp) Reset()         { *m = RemoveTeamMembersResp{} }
func (m *RemoveTeamMembersResp) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamMembersResp) ProtoMessage()    {}
func (*RemoveTeamMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{7}
}
func (m *RemoveTeamMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamMembersResp.Unmarshal(m, b)
}
func (m *RemoveTeamMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamMembersResp.Marshal(b, m, deterministic)
}
func (dst *RemoveTeamMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamMembersResp.Merge(dst, src)
}
func (m *RemoveTeamMembersResp) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamMembersResp.Size(m)
}
func (m *RemoveTeamMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamMembersResp proto.InternalMessageInfo

func (m *RemoveTeamMembersResp) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamsForMemberResp struct {
	Teams                []*common.Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTeamsForMemberResp) Reset()         { *m = GetTeamsForMemberResp{} }
func (m *GetTeamsForMemberResp) String() string { return proto.CompactTextString(m) }
func (*GetTeamsForMemberResp) ProtoMessage()    {}
func (*GetTeamsForMemberResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{8}
}
func (m *GetTeamsForMemberResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsForMemberResp.Unmarshal(m, b)
}
func (m *GetTeamsForMemberResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsForMemberResp.Marshal(b, m, deterministic)
}
func (dst *GetTeamsForMemberResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsForMemberResp.Merge(dst, src)
}
func (m *GetTeamsForMemberResp) XXX_Size() int {
	return xxx_messageInfo_GetTeamsForMemberResp.Size(m)
}
func (m *GetTeamsForMemberResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsForMemberResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsForMemberResp proto.InternalMessageInfo

func (m *GetTeamsForMemberResp) GetTeams() []*common.Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type ApplyV2DataMigrationsResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyV2DataMigrationsResp) Reset()         { *m = ApplyV2DataMigrationsResp{} }
func (m *ApplyV2DataMigrationsResp) String() string { return proto.CompactTextString(m) }
func (*ApplyV2DataMigrationsResp) ProtoMessage()    {}
func (*ApplyV2DataMigrationsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_b23c3125391c9628, []int{9}
}
func (m *ApplyV2DataMigrationsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyV2DataMigrationsResp.Unmarshal(m, b)
}
func (m *ApplyV2DataMigrationsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyV2DataMigrationsResp.Marshal(b, m, deterministic)
}
func (dst *ApplyV2DataMigrationsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyV2DataMigrationsResp.Merge(dst, src)
}
func (m *ApplyV2DataMigrationsResp) XXX_Size() int {
	return xxx_messageInfo_ApplyV2DataMigrationsResp.Size(m)
}
func (m *ApplyV2DataMigrationsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyV2DataMigrationsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyV2DataMigrationsResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTeamResp)(nil), "chef.automate.api.iam.v2beta.GetTeamResp")
	proto.RegisterType((*GetTeamsResp)(nil), "chef.automate.api.iam.v2beta.GetTeamsResp")
	proto.RegisterType((*CreateTeamResp)(nil), "chef.automate.api.iam.v2beta.CreateTeamResp")
	proto.RegisterType((*UpdateTeamResp)(nil), "chef.automate.api.iam.v2beta.UpdateTeamResp")
	proto.RegisterType((*DeleteTeamResp)(nil), "chef.automate.api.iam.v2beta.DeleteTeamResp")
	proto.RegisterType((*AddTeamMembersResp)(nil), "chef.automate.api.iam.v2beta.AddTeamMembersResp")
	proto.RegisterType((*GetTeamMembershipResp)(nil), "chef.automate.api.iam.v2beta.GetTeamMembershipResp")
	proto.RegisterType((*RemoveTeamMembersResp)(nil), "chef.automate.api.iam.v2beta.RemoveTeamMembersResp")
	proto.RegisterType((*GetTeamsForMemberResp)(nil), "chef.automate.api.iam.v2beta.GetTeamsForMemberResp")
	proto.RegisterType((*ApplyV2DataMigrationsResp)(nil), "chef.automate.api.iam.v2beta.ApplyV2DataMigrationsResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/response/teams.proto", fileDescriptor_teams_b23c3125391c9628)
}

var fileDescriptor_teams_b23c3125391c9628 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x4b, 0x33, 0x31,
	0x10, 0x87, 0x29, 0xef, 0xeb, 0x57, 0x2a, 0x3d, 0x14, 0x84, 0x56, 0x3d, 0x94, 0x3d, 0x79, 0x31,
	0x81, 0x15, 0xc4, 0x8b, 0x48, 0xb5, 0x6a, 0x05, 0x2b, 0x58, 0xd4, 0x83, 0x17, 0x99, 0xdd, 0x1d,
	0xdb, 0x40, 0xb3, 0x13, 0x92, 0x69, 0xa5, 0xff, 0xbd, 0x6c, 0xb6, 0x16, 0x4f, 0xd2, 0xea, 0x1e,
	0xf3, 0xf1, 0x3c, 0xbf, 0x61, 0x32, 0x11, 0x17, 0x29, 0x19, 0x4b, 0x39, 0xe6, 0xec, 0x15, 0x4c,
	0x99, 0x0c, 0x30, 0x1e, 0x8f, 0x80, 0xf1, 0x03, 0xe6, 0x0a, 0xac, 0x56, 0x1a, 0x8c, 0x9a, 0xc5,
	0x09, 0x32, 0x28, 0x87, 0xde, 0x52, 0xee, 0x51, 0x31, 0x82, 0xf1, 0xd2, 0x3a, 0x62, 0x6a, 0x1e,
	0xa6, 0x63, 0x7c, 0x97, 0x5f, 0xa8, 0x04, 0xab, 0xa5, 0x06, 0x23, 0x4b, 0x64, 0xff, 0x7c, 0x0d,
	0x7d, 0x4a, 0xc6, 0x50, 0xfe, 0x5d, 0x1e, 0x5d, 0x8b, 0xfa, 0x2d, 0xf2, 0x13, 0x82, 0x19, 0xa2,
	0xb7, 0xcd, 0x53, 0xf1, 0xbf, 0x38, 0x6d, 0xd5, 0x3a, 0xb5, 0xa3, 0x7a, 0x1c, 0xc9, 0x9f, 0xa2,
	0x65, 0xa0, 0xc2, 0xfd, 0xa8, 0x2f, 0x76, 0x17, 0x1a, 0x1f, 0x3c, 0x67, 0x62, 0x23, 0xa4, 0xb4,
	0x6a, 0x9d, 0x7f, 0x2b, 0x8a, 0x4a, 0x20, 0xea, 0x8b, 0xc6, 0x95, 0x43, 0x60, 0xac, 0xa0, 0xa6,
	0xc6, 0xb3, 0xcd, 0x2a, 0x32, 0xf5, 0x70, 0x82, 0x15, 0x98, 0x94, 0x68, 0x76, 0xb3, 0xac, 0xd8,
	0x18, 0xa0, 0x49, 0xd0, 0x95, 0xdd, 0x6a, 0x8b, 0xed, 0xa9, 0x47, 0xf7, 0xa6, 0xb3, 0xb2, 0x61,
	0x3b, 0xc3, 0xad, 0x62, 0x7d, 0x97, 0xf9, 0x28, 0x16, 0x7b, 0x8b, 0xc6, 0x2e, 0x80, 0xb1, 0xb6,
	0x2b, 0x30, 0x43, 0x34, 0x34, 0xc3, 0x35, 0x72, 0x1e, 0x97, 0x39, 0xfe, 0x86, 0x5c, 0x09, 0xfd,
	0xf1, 0x25, 0x0f, 0x44, 0xbb, 0x6b, 0xed, 0x64, 0xfe, 0x12, 0xf7, 0x80, 0x61, 0xa0, 0x47, 0x0e,
	0x58, 0x53, 0x1e, 0x4a, 0xb9, 0x7c, 0x78, 0xbd, 0x1f, 0x69, 0x1e, 0x4f, 0x13, 0x99, 0x92, 0x51,
	0x85, 0x73, 0x39, 0xbd, 0xea, 0x17, 0x1f, 0x26, 0xd9, 0x0c, 0xe3, 0x7c, 0xf2, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x28, 0xb2, 0x40, 0x5d, 0x6e, 0x03, 0x00, 0x00,
}