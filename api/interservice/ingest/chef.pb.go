// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/ingest/chef.proto

package ingest // import "github.com/chef/automate/api/interservice/ingest"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import request "github.com/chef/automate/api/external/ingest/request"
import response "github.com/chef/automate/api/external/ingest/response"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Version message
//
// The ingest-service version constructed with:
// * Service name
// * Built time
// * Semantic version
// * Git SHA
type Version struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	Built                string   `protobuf:"bytes,1,opt,name=built,proto3" json:"built,omitempty" toml:"built,omitempty" mapstructure:"built,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Sha                  string   `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty" toml:"sha,omitempty" mapstructure:"sha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_4bc9d64afcc939d3, []int{0}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (dst *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(dst, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetBuilt() string {
	if m != nil {
		return m.Built
	}
	return ""
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_4bc9d64afcc939d3, []int{1}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (dst *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(dst, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Version)(nil), "chef.automate.domain.ingest.Version")
	proto.RegisterType((*VersionRequest)(nil), "chef.automate.domain.ingest.VersionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChefIngesterClient is the client API for ChefIngester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefIngesterClient interface {
	ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error)
	ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error)
	ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error)
}

type chefIngesterClient struct {
	cc *grpc.ClientConn
}

func NewChefIngesterClient(cc *grpc.ClientConn) ChefIngesterClient {
	return &chefIngesterClient{cc}
}

func (c *chefIngesterClient) ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error) {
	out := new(response.ProcessChefRunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessChefRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error) {
	out := new(response.ProcessChefActionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessChefAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error) {
	out := new(response.ProcessLivenessResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessLivenessPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error) {
	out := new(response.ProcessMultipleNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessMultipleNodeDeletes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error) {
	out := new(response.ProcessNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessNodeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefIngesterServer is the server API for ChefIngester service.
type ChefIngesterServer interface {
	ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error)
	ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error)
	ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error)
	GetVersion(context.Context, *VersionRequest) (*Version, error)
}

func RegisterChefIngesterServer(s *grpc.Server, srv ChefIngesterServer) {
	s.RegisterService(&_ChefIngester_serviceDesc, srv)
}

func _ChefIngester_ProcessChefRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Run)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessChefRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, req.(*request.Run))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessChefAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessChefAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, req.(*request.Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessLivenessPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Liveness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessLivenessPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, req.(*request.Liveness))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessMultipleNodeDeletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.MultipleNodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessMultipleNodeDeletes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, req.(*request.MultipleNodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessNodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Delete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessNodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, req.(*request.Delete))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChefIngester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.ingest.ChefIngester",
	HandlerType: (*ChefIngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessChefRun",
			Handler:    _ChefIngester_ProcessChefRun_Handler,
		},
		{
			MethodName: "ProcessChefAction",
			Handler:    _ChefIngester_ProcessChefAction_Handler,
		},
		{
			MethodName: "ProcessLivenessPing",
			Handler:    _ChefIngester_ProcessLivenessPing_Handler,
		},
		{
			MethodName: "ProcessMultipleNodeDeletes",
			Handler:    _ChefIngester_ProcessMultipleNodeDeletes_Handler,
		},
		{
			MethodName: "ProcessNodeDelete",
			Handler:    _ChefIngester_ProcessNodeDelete_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ChefIngester_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/ingest/chef.proto",
}

func init() {
	proto.RegisterFile("api/interservice/ingest/chef.proto", fileDescriptor_chef_4bc9d64afcc939d3)
}

var fileDescriptor_chef_4bc9d64afcc939d3 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0x95, 0xb6, 0x50, 0x61, 0xa1, 0xb2, 0x75, 0xa9, 0x1a, 0xb9, 0x15, 0xaa, 0x2c, 0x3e,
	0x96, 0xad, 0xe4, 0xa0, 0x72, 0xa2, 0x70, 0x01, 0x2a, 0x21, 0x24, 0x40, 0x55, 0x0e, 0x1c, 0xb8,
	0x20, 0x6f, 0x76, 0x9a, 0xb5, 0x94, 0xb5, 0x43, 0xec, 0xac, 0x7a, 0xe6, 0x05, 0x38, 0x70, 0x44,
	0x82, 0x87, 0xe2, 0x15, 0x78, 0x10, 0x14, 0x7f, 0x94, 0x5a, 0x2d, 0xd9, 0xed, 0x6d, 0x3c, 0xf9,
	0xcf, 0xcc, 0x2f, 0x33, 0xf6, 0x20, 0xca, 0x6b, 0x91, 0x09, 0x69, 0xa0, 0xd1, 0xd0, 0xcc, 0x45,
	0x01, 0x99, 0x90, 0x25, 0x68, 0x93, 0x15, 0x53, 0x38, 0x65, 0x75, 0xa3, 0x8c, 0xc2, 0xbb, 0xd6,
	0xe6, 0xad, 0x51, 0x33, 0x6e, 0x80, 0x4d, 0xd4, 0x8c, 0x0b, 0xc9, 0x9c, 0x8e, 0xec, 0x95, 0x4a,
	0x95, 0x15, 0x64, 0x5d, 0x1e, 0x2e, 0xa5, 0x32, 0xdc, 0x08, 0x25, 0xb5, 0x0b, 0x25, 0x0f, 0x3b,
	0x37, 0x9c, 0x19, 0x68, 0x24, 0xaf, 0x42, 0xea, 0x06, 0xbe, 0xb4, 0x71, 0x09, 0x32, 0xec, 0xd3,
	0xf1, 0xa2, 0x4b, 0xe9, 0x95, 0xa3, 0x3e, 0x65, 0x25, 0xe6, 0x20, 0x41, 0x87, 0xea, 0x8f, 0xae,
	0xd6, 0xea, 0x5a, 0x49, 0x0d, 0x17, 0xcb, 0x3f, 0xee, 0x15, 0x46, 0xf5, 0x0f, 0x7a, 0xa5, 0x31,
	0x00, 0xfd, 0x8c, 0xd6, 0x3f, 0x42, 0xa3, 0x85, 0x92, 0x38, 0x45, 0xeb, 0x73, 0x67, 0xa6, 0x2b,
	0xfb, 0xc9, 0xf0, 0x56, 0x1e, 0x8e, 0xf8, 0x2e, 0xba, 0x31, 0x6e, 0x45, 0x65, 0xd2, 0xc4, 0xfa,
	0xdd, 0x01, 0x63, 0xb4, 0x26, 0xf9, 0x0c, 0xd2, 0x55, 0xeb, 0xb4, 0x36, 0x1e, 0xa0, 0x55, 0x3d,
	0xe5, 0xe9, 0x9a, 0x75, 0x75, 0x26, 0x1d, 0xa0, 0x0d, 0x5f, 0x20, 0x77, 0x2d, 0x38, 0xfc, 0xb9,
	0x8e, 0x6e, 0xbf, 0x9e, 0xc2, 0xe9, 0x5b, 0x4b, 0x06, 0x0d, 0xfe, 0x96, 0xa0, 0x8d, 0x93, 0x46,
	0x15, 0xa0, 0x75, 0xe7, 0xcf, 0x5b, 0x89, 0x1f, 0xb0, 0x78, 0xa2, 0xbc, 0x16, 0x7e, 0x9c, 0xcc,
	0x77, 0x92, 0xe5, 0xad, 0x24, 0xcf, 0xfa, 0x64, 0xee, 0x87, 0x59, 0x9c, 0x39, 0xf7, 0x6e, 0xba,
	0xfb, 0xf5, 0xf7, 0x9f, 0xef, 0x2b, 0xdb, 0x74, 0x90, 0xc1, 0x1c, 0xa4, 0xd1, 0xb6, 0xd9, 0x59,
	0xd3, 0xca, 0xa3, 0x64, 0x84, 0x7f, 0x24, 0x68, 0xf3, 0x42, 0xdc, 0x4b, 0xdb, 0x5e, 0x3c, 0x5c,
	0x0c, 0xe5, 0x94, 0xe4, 0xc5, 0xf5, 0xb8, 0x5c, 0xd4, 0x39, 0xda, 0x3d, 0x8b, 0x96, 0xd2, 0xad,
	0x08, 0xcd, 0xcd, 0xb8, 0xa3, 0xfb, 0x95, 0xa0, 0x2d, 0x1f, 0xfd, 0xce, 0x4f, 0xf3, 0x44, 0xc8,
	0x12, 0x8f, 0x16, 0xf3, 0x05, 0x3d, 0x39, 0x5a, 0x9e, 0x30, 0xc4, 0x9c, 0xf3, 0xed, 0x5b, 0x3e,
	0x42, 0xb7, 0x23, 0xbe, 0x70, 0xb1, 0x3c, 0x21, 0xf1, 0xd1, 0xef, 0xdb, 0xca, 0x88, 0xba, 0x82,
	0x0f, 0x6a, 0x02, 0xc7, 0x50, 0x81, 0x01, 0x8d, 0x9f, 0x2f, 0x06, 0xbd, 0x1c, 0xe6, 0xaf, 0x0f,
	0x39, 0x5e, 0x9e, 0xfc, 0xaa, 0x24, 0x4e, 0xd1, 0x11, 0x86, 0x09, 0xff, 0xfb, 0xba, 0xcc, 0x84,
	0x9d, 0xf2, 0x3a, 0x13, 0xbe, 0x5c, 0x9d, 0x52, 0xdb, 0xc1, 0x3d, 0xba, 0x13, 0x75, 0x50, 0xaa,
	0x09, 0x4c, 0xac, 0xb0, 0xeb, 0xe1, 0x19, 0x42, 0x6f, 0xc0, 0x84, 0xc7, 0x79, 0xc0, 0x7a, 0x56,
	0x1c, 0x8b, 0x5f, 0x18, 0xb9, 0xbf, 0x8c, 0x98, 0xee, 0x58, 0x88, 0x4d, 0x7c, 0x27, 0xec, 0x08,
	0xff, 0xdc, 0x5f, 0x1d, 0x7e, 0x7a, 0x52, 0x0a, 0x33, 0x6d, 0xc7, 0xac, 0x50, 0x33, 0x7f, 0xf9,
	0x7c, 0xaa, 0xec, 0x3f, 0xcb, 0x78, 0x7c, 0xd3, 0xae, 0x93, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x38, 0x56, 0x59, 0xae, 0x05, 0x00, 0x00,
}