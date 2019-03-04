// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/auth/teams/request/teams.proto

package request // import "github.com/chef/automate/components/automate-gateway/api/auth/teams/request"

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

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	UserIds              []string `protobuf:"bytes,4,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{0}
}
func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (dst *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(dst, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Team) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsReq) Reset()         { *m = GetTeamsReq{} }
func (m *GetTeamsReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamsReq) ProtoMessage()    {}
func (*GetTeamsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{1}
}
func (m *GetTeamsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsReq.Unmarshal(m, b)
}
func (m *GetTeamsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsReq.Marshal(b, m, deterministic)
}
func (dst *GetTeamsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsReq.Merge(dst, src)
}
func (m *GetTeamsReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamsReq.Size(m)
}
func (m *GetTeamsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsReq proto.InternalMessageInfo

type GetTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamReq) Reset()         { *m = GetTeamReq{} }
func (m *GetTeamReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamReq) ProtoMessage()    {}
func (*GetTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{2}
}
func (m *GetTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamReq.Unmarshal(m, b)
}
func (m *GetTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamReq.Marshal(b, m, deterministic)
}
func (dst *GetTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamReq.Merge(dst, src)
}
func (m *GetTeamReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamReq.Size(m)
}
func (m *GetTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamReq proto.InternalMessageInfo

func (m *GetTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateTeamReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamReq) Reset()         { *m = CreateTeamReq{} }
func (m *CreateTeamReq) String() string { return proto.CompactTextString(m) }
func (*CreateTeamReq) ProtoMessage()    {}
func (*CreateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{3}
}
func (m *CreateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamReq.Unmarshal(m, b)
}
func (m *CreateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamReq.Marshal(b, m, deterministic)
}
func (dst *CreateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamReq.Merge(dst, src)
}
func (m *CreateTeamReq) XXX_Size() int {
	return xxx_messageInfo_CreateTeamReq.Size(m)
}
func (m *CreateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamReq proto.InternalMessageInfo

func (m *CreateTeamReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTeamReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamReq) Reset()         { *m = UpdateTeamReq{} }
func (m *UpdateTeamReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamReq) ProtoMessage()    {}
func (*UpdateTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{4}
}
func (m *UpdateTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamReq.Unmarshal(m, b)
}
func (m *UpdateTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamReq.Marshal(b, m, deterministic)
}
func (dst *UpdateTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamReq.Merge(dst, src)
}
func (m *UpdateTeamReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamReq.Size(m)
}
func (m *UpdateTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamReq proto.InternalMessageInfo

func (m *UpdateTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTeamReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTeamReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type DeleteTeamReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTeamReq) Reset()         { *m = DeleteTeamReq{} }
func (m *DeleteTeamReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTeamReq) ProtoMessage()    {}
func (*DeleteTeamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{5}
}
func (m *DeleteTeamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTeamReq.Unmarshal(m, b)
}
func (m *DeleteTeamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTeamReq.Marshal(b, m, deterministic)
}
func (dst *DeleteTeamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTeamReq.Merge(dst, src)
}
func (m *DeleteTeamReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTeamReq.Size(m)
}
func (m *DeleteTeamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTeamReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTeamReq proto.InternalMessageInfo

func (m *DeleteTeamReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddUsersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserIds              []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUsersReq) Reset()         { *m = AddUsersReq{} }
func (m *AddUsersReq) String() string { return proto.CompactTextString(m) }
func (*AddUsersReq) ProtoMessage()    {}
func (*AddUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{6}
}
func (m *AddUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUsersReq.Unmarshal(m, b)
}
func (m *AddUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUsersReq.Marshal(b, m, deterministic)
}
func (dst *AddUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUsersReq.Merge(dst, src)
}
func (m *AddUsersReq) XXX_Size() int {
	return xxx_messageInfo_AddUsersReq.Size(m)
}
func (m *AddUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddUsersReq proto.InternalMessageInfo

func (m *AddUsersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddUsersReq) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetUsersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersReq) Reset()         { *m = GetUsersReq{} }
func (m *GetUsersReq) String() string { return proto.CompactTextString(m) }
func (*GetUsersReq) ProtoMessage()    {}
func (*GetUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{7}
}
func (m *GetUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersReq.Unmarshal(m, b)
}
func (m *GetUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersReq.Marshal(b, m, deterministic)
}
func (dst *GetUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersReq.Merge(dst, src)
}
func (m *GetUsersReq) XXX_Size() int {
	return xxx_messageInfo_GetUsersReq.Size(m)
}
func (m *GetUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersReq proto.InternalMessageInfo

func (m *GetUsersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveUsersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserIds              []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveUsersReq) Reset()         { *m = RemoveUsersReq{} }
func (m *RemoveUsersReq) String() string { return proto.CompactTextString(m) }
func (*RemoveUsersReq) ProtoMessage()    {}
func (*RemoveUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{8}
}
func (m *RemoveUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveUsersReq.Unmarshal(m, b)
}
func (m *RemoveUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveUsersReq.Marshal(b, m, deterministic)
}
func (dst *RemoveUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveUsersReq.Merge(dst, src)
}
func (m *RemoveUsersReq) XXX_Size() int {
	return xxx_messageInfo_RemoveUsersReq.Size(m)
}
func (m *RemoveUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveUsersReq proto.InternalMessageInfo

func (m *RemoveUsersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RemoveUsersReq) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetTeamsForUserReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsForUserReq) Reset()         { *m = GetTeamsForUserReq{} }
func (m *GetTeamsForUserReq) String() string { return proto.CompactTextString(m) }
func (*GetTeamsForUserReq) ProtoMessage()    {}
func (*GetTeamsForUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_teams_e519873069c6a2d4, []int{9}
}
func (m *GetTeamsForUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsForUserReq.Unmarshal(m, b)
}
func (m *GetTeamsForUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsForUserReq.Marshal(b, m, deterministic)
}
func (dst *GetTeamsForUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsForUserReq.Merge(dst, src)
}
func (m *GetTeamsForUserReq) XXX_Size() int {
	return xxx_messageInfo_GetTeamsForUserReq.Size(m)
}
func (m *GetTeamsForUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsForUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsForUserReq proto.InternalMessageInfo

func (m *GetTeamsForUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Team)(nil), "chef.automate.api.teams.request.Team")
	proto.RegisterType((*GetTeamsReq)(nil), "chef.automate.api.teams.request.GetTeamsReq")
	proto.RegisterType((*GetTeamReq)(nil), "chef.automate.api.teams.request.GetTeamReq")
	proto.RegisterType((*CreateTeamReq)(nil), "chef.automate.api.teams.request.CreateTeamReq")
	proto.RegisterType((*UpdateTeamReq)(nil), "chef.automate.api.teams.request.UpdateTeamReq")
	proto.RegisterType((*DeleteTeamReq)(nil), "chef.automate.api.teams.request.DeleteTeamReq")
	proto.RegisterType((*AddUsersReq)(nil), "chef.automate.api.teams.request.AddUsersReq")
	proto.RegisterType((*GetUsersReq)(nil), "chef.automate.api.teams.request.GetUsersReq")
	proto.RegisterType((*RemoveUsersReq)(nil), "chef.automate.api.teams.request.RemoveUsersReq")
	proto.RegisterType((*GetTeamsForUserReq)(nil), "chef.automate.api.teams.request.GetTeamsForUserReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/auth/teams/request/teams.proto", fileDescriptor_teams_e519873069c6a2d4)
}

var fileDescriptor_teams_e519873069c6a2d4 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0x65, 0xdd, 0x50, 0xf7, 0x8d, 0xee, 0x90, 0x53, 0x05, 0x65, 0x23, 0x78, 0xd8, 0xc5, 0xe6,
	0xe0, 0x45, 0x10, 0x04, 0x7f, 0x23, 0xe2, 0xa5, 0xb8, 0x8b, 0x17, 0xc9, 0x9a, 0xcf, 0x35, 0x60,
	0x9a, 0x34, 0x49, 0x15, 0xff, 0x7b, 0x69, 0xb7, 0x4a, 0xd5, 0xf5, 0xa0, 0x78, 0xeb, 0xf7, 0xeb,
	0xbd, 0xd7, 0x97, 0x07, 0xa7, 0xa9, 0x56, 0x46, 0xe7, 0x98, 0x7b, 0xc7, 0x78, 0xe9, 0xb5, 0xe2,
	0x1e, 0x0f, 0x97, 0xdc, 0xe3, 0x1b, 0x7f, 0x67, 0xdc, 0xc8, 0xaa, 0x99, 0x31, 0x8f, 0x5c, 0x39,
	0x66, 0xb1, 0x28, 0xd1, 0xf9, 0x55, 0x15, 0x1b, 0xab, 0xbd, 0x26, 0x93, 0x34, 0xc3, 0xe7, 0xb8,
	0xb9, 0x8c, 0xb9, 0x91, 0xf1, 0x6a, 0xbc, 0x5e, 0xa6, 0x4b, 0x18, 0x3c, 0x20, 0x57, 0x64, 0x0c,
	0x81, 0x14, 0x51, 0x6f, 0xda, 0x9b, 0x0d, 0x93, 0x40, 0x0a, 0x42, 0x60, 0x90, 0x73, 0x85, 0x51,
	0x50, 0x77, 0xea, 0x6f, 0x32, 0x85, 0x91, 0x40, 0x97, 0x5a, 0x69, 0xbc, 0xd4, 0x79, 0xd4, 0xaf,
	0x47, 0xed, 0x16, 0xd9, 0x85, 0x9d, 0xd2, 0xa1, 0x7d, 0x92, 0xc2, 0x45, 0x83, 0x69, 0x7f, 0x36,
	0x4c, 0xb6, 0xab, 0xfa, 0x56, 0x38, 0x1a, 0xc2, 0xe8, 0x06, 0x7d, 0xc5, 0xe5, 0x12, 0x2c, 0xe8,
	0x1e, 0xc0, 0xba, 0x4c, 0xb0, 0xf8, 0xce, 0x4e, 0xaf, 0x20, 0xbc, 0xb0, 0xc8, 0x3d, 0x36, 0x0b,
	0x8d, 0x9c, 0x5e, 0xb7, 0x9c, 0xe0, 0x87, 0x1c, 0x3a, 0x87, 0x70, 0x6e, 0x44, 0x0b, 0xe6, 0x5f,
	0xfe, 0x92, 0x4e, 0x20, 0xbc, 0xc4, 0x17, 0xec, 0x84, 0xa5, 0xc7, 0x30, 0x3a, 0x13, 0x62, 0xee,
	0xd0, 0xba, 0x4d, 0xac, 0x6d, 0x97, 0x82, 0xaf, 0x2e, 0xed, 0xd7, 0x2e, 0x75, 0x5d, 0xd2, 0x13,
	0x18, 0x27, 0xa8, 0xf4, 0x2b, 0xfe, 0x05, 0xfb, 0x00, 0x48, 0xf3, 0x02, 0xd7, 0xda, 0x56, 0x08,
	0x1b, 0x00, 0xce, 0xef, 0x1f, 0xef, 0x96, 0xd2, 0x67, 0xe5, 0x22, 0x4e, 0xb5, 0x62, 0x55, 0x7c,
	0x3e, 0x83, 0xc7, 0x7e, 0x1f, 0xc6, 0xc5, 0x56, 0x9d, 0xc3, 0xa3, 0x8f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x83, 0x73, 0x08, 0x3a, 0xc9, 0x02, 0x00, 0x00,
}