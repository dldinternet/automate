// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/authz/authz.proto

package authz // import "github.com/chef/automate/components/automate-gateway/api/authz"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import version "github.com/chef/automate/api/external/common/version"
import request "github.com/chef/automate/components/automate-gateway/api/authz/request"
import response "github.com/chef/automate/components/automate-gateway/api/authz/response"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
	CreatePolicy(ctx context.Context, in *request.CreatePolicyReq, opts ...grpc.CallOption) (*response.CreatePolicyResp, error)
	ListPolicies(ctx context.Context, in *request.ListPoliciesReq, opts ...grpc.CallOption) (*response.ListPoliciesResp, error)
	DeletePolicy(ctx context.Context, in *request.DeletePolicyReq, opts ...grpc.CallOption) (*response.DeletePolicyResp, error)
	IntrospectAll(ctx context.Context, in *request.IntrospectAllReq, opts ...grpc.CallOption) (*response.IntrospectResp, error)
	IntrospectSome(ctx context.Context, in *request.IntrospectSomeReq, opts ...grpc.CallOption) (*response.IntrospectResp, error)
	Introspect(ctx context.Context, in *request.IntrospectReq, opts ...grpc.CallOption) (*response.IntrospectResp, error)
	IntrospectAllProjects(ctx context.Context, in *request.IntrospectAllProjectsReq, opts ...grpc.CallOption) (*response.IntrospectProjectsResp, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) CreatePolicy(ctx context.Context, in *request.CreatePolicyReq, opts ...grpc.CallOption) (*response.CreatePolicyResp, error) {
	out := new(response.CreatePolicyResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/CreatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) ListPolicies(ctx context.Context, in *request.ListPoliciesReq, opts ...grpc.CallOption) (*response.ListPoliciesResp, error) {
	out := new(response.ListPoliciesResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) DeletePolicy(ctx context.Context, in *request.DeletePolicyReq, opts ...grpc.CallOption) (*response.DeletePolicyResp, error) {
	out := new(response.DeletePolicyResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IntrospectAll(ctx context.Context, in *request.IntrospectAllReq, opts ...grpc.CallOption) (*response.IntrospectResp, error) {
	out := new(response.IntrospectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/IntrospectAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IntrospectSome(ctx context.Context, in *request.IntrospectSomeReq, opts ...grpc.CallOption) (*response.IntrospectResp, error) {
	out := new(response.IntrospectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/IntrospectSome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) Introspect(ctx context.Context, in *request.IntrospectReq, opts ...grpc.CallOption) (*response.IntrospectResp, error) {
	out := new(response.IntrospectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/Introspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IntrospectAllProjects(ctx context.Context, in *request.IntrospectAllProjectsReq, opts ...grpc.CallOption) (*response.IntrospectProjectsResp, error) {
	out := new(response.IntrospectProjectsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.authz.Authorization/IntrospectAllProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	CreatePolicy(context.Context, *request.CreatePolicyReq) (*response.CreatePolicyResp, error)
	ListPolicies(context.Context, *request.ListPoliciesReq) (*response.ListPoliciesResp, error)
	DeletePolicy(context.Context, *request.DeletePolicyReq) (*response.DeletePolicyResp, error)
	IntrospectAll(context.Context, *request.IntrospectAllReq) (*response.IntrospectResp, error)
	IntrospectSome(context.Context, *request.IntrospectSomeReq) (*response.IntrospectResp, error)
	Introspect(context.Context, *request.IntrospectReq) (*response.IntrospectResp, error)
	IntrospectAllProjects(context.Context, *request.IntrospectAllProjectsReq) (*response.IntrospectProjectsResp, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreatePolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/CreatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).CreatePolicy(ctx, req.(*request.CreatePolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListPoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).ListPolicies(ctx, req.(*request.ListPoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeletePolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).DeletePolicy(ctx, req.(*request.DeletePolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IntrospectAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IntrospectAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/IntrospectAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IntrospectAll(ctx, req.(*request.IntrospectAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IntrospectSome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectSomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IntrospectSome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/IntrospectSome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IntrospectSome(ctx, req.(*request.IntrospectSomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_Introspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Introspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/Introspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Introspect(ctx, req.(*request.IntrospectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IntrospectAllProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectAllProjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IntrospectAllProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.authz.Authorization/IntrospectAllProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IntrospectAllProjects(ctx, req.(*request.IntrospectAllProjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.authz.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Authorization_GetVersion_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _Authorization_CreatePolicy_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _Authorization_ListPolicies_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _Authorization_DeletePolicy_Handler,
		},
		{
			MethodName: "IntrospectAll",
			Handler:    _Authorization_IntrospectAll_Handler,
		},
		{
			MethodName: "IntrospectSome",
			Handler:    _Authorization_IntrospectSome_Handler,
		},
		{
			MethodName: "Introspect",
			Handler:    _Authorization_Introspect_Handler,
		},
		{
			MethodName: "IntrospectAllProjects",
			Handler:    _Authorization_IntrospectAllProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/authz/authz.proto",
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/authz/authz.proto", fileDescriptor_authz_53acb10f4e7fe39c)
}

var fileDescriptor_authz_53acb10f4e7fe39c = []byte{
	// 761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x86, 0xe5, 0xde, 0xab, 0xea, 0x6a, 0x94, 0xa6, 0xb9, 0xa7, 0x4d, 0x3a, 0x4d, 0x6f, 0x55,
	0x5d, 0xc3, 0x86, 0x8a, 0xd8, 0x6d, 0x61, 0x41, 0xbd, 0x40, 0x2d, 0x20, 0xa1, 0x4a, 0x2c, 0xaa,
	0x22, 0xb1, 0x40, 0x48, 0x91, 0xeb, 0x4e, 0x93, 0x41, 0xb6, 0xc7, 0xf5, 0x4c, 0x4a, 0x5b, 0xd4,
	0x05, 0x5e, 0xa1, 0x2c, 0xe1, 0x09, 0x58, 0xf1, 0x04, 0x7e, 0x11, 0x78, 0x00, 0x36, 0xb0, 0x61,
	0x81, 0x78, 0x04, 0x34, 0x63, 0x3b, 0x8e, 0x93, 0xb4, 0x49, 0xba, 0x49, 0x94, 0x39, 0xff, 0x6f,
	0xcf, 0xf7, 0x9f, 0x33, 0x9a, 0xa0, 0x4d, 0x87, 0x79, 0x01, 0xf3, 0x89, 0x2f, 0xb8, 0x69, 0x77,
	0x04, 0xf3, 0x6c, 0x41, 0x1a, 0x2d, 0x5b, 0x90, 0x37, 0xf6, 0xb9, 0x69, 0x07, 0x54, 0x2e, 0xb6,
	0x2f, 0x92, 0x4f, 0x23, 0x08, 0x99, 0x60, 0xb0, 0xe4, 0xb4, 0xc9, 0xb1, 0x91, 0x89, 0x0d, 0x3b,
	0xa0, 0x86, 0x2a, 0xd7, 0xff, 0x6b, 0x31, 0xd6, 0x72, 0x49, 0x62, 0xf3, 0x7d, 0x26, 0x6c, 0x41,
	0x99, 0xcf, 0x13, 0x5b, 0xfd, 0x8e, 0x5c, 0x26, 0x67, 0x82, 0x84, 0xbe, 0xed, 0x9a, 0x0e, 0xf3,
	0x3c, 0xe6, 0x9b, 0xa7, 0x24, 0xe4, 0x34, 0xff, 0x4e, 0xa5, 0xdb, 0x93, 0x6d, 0x2a, 0x24, 0x27,
	0x1d, 0xc2, 0x45, 0xff, 0xe6, 0xea, 0xd6, 0xa4, 0x56, 0x1e, 0x30, 0x9f, 0x93, 0x82, 0x77, 0x67,
	0xa4, 0x37, 0x0c, 0x1c, 0x53, 0xd5, 0x9d, 0x46, 0x8b, 0xf8, 0x8d, 0x80, 0xb9, 0xd4, 0x39, 0xbf,
	0x82, 0x71, 0x9a, 0x27, 0x50, 0xdb, 0x1b, 0x7e, 0xc2, 0xd6, 0xe7, 0x79, 0x34, 0xb7, 0xdb, 0x11,
	0x6d, 0x16, 0xd2, 0x0b, 0x55, 0x80, 0xdf, 0x1a, 0x42, 0x4f, 0x89, 0x78, 0x91, 0x24, 0x04, 0xf7,
	0x8d, 0xe1, 0xf8, 0x93, 0x30, 0x8d, 0x2c, 0xc4, 0x54, 0xba, 0xe7, 0x1f, 0xb3, 0x83, 0x24, 0xa2,
	0x7a, 0x63, 0x2a, 0x97, 0x7e, 0x19, 0xc5, 0xb8, 0x86, 0x16, 0x39, 0x09, 0x4f, 0xa9, 0x43, 0x9a,
	0xd4, 0x3f, 0x66, 0x56, 0xaa, 0x8b, 0x62, 0x3c, 0x0b, 0x7f, 0x87, 0xc4, 0x3e, 0xea, 0xc6, 0x18,
	0xa3, 0x1a, 0x3f, 0xe7, 0x82, 0x78, 0x56, 0x2a, 0xcd, 0x54, 0xdd, 0x18, 0xaf, 0xc0, 0x72, 0xb1,
	0x96, 0xbe, 0xc0, 0x6a, 0x11, 0x11, 0x7d, 0xfd, 0xfe, 0x71, 0x06, 0x43, 0x4d, 0xc5, 0x6f, 0xaa,
	0x40, 0x28, 0xe1, 0xd9, 0x14, 0xc0, 0x37, 0x0d, 0x95, 0x1e, 0x87, 0xc4, 0x16, 0x64, 0x5f, 0x65,
	0x05, 0x1b, 0xc6, 0x15, 0x33, 0x67, 0xa4, 0x33, 0x60, 0xf4, 0xcb, 0x0f, 0xc8, 0x49, 0x7d, 0xeb,
	0x1a, 0x47, 0xd2, 0xfa, 0x01, 0x0b, 0x0f, 0x74, 0x1a, 0xc5, 0x78, 0x1e, 0xcd, 0x49, 0x99, 0x95,
	0xed, 0x29, 0x8a, 0xf1, 0x3f, 0x30, 0xeb, 0x28, 0x69, 0x37, 0xc6, 0x65, 0x54, 0xa2, 0xb6, 0xd7,
	0xab, 0x76, 0x63, 0x5c, 0x85, 0x85, 0xfe, 0x15, 0x2b, 0x91, 0x2a, 0xc0, 0x8a, 0x5e, 0x2e, 0x02,
	0xaa, 0xd5, 0xbf, 0x2c, 0x6d, 0x1d, 0xbe, 0x68, 0xa8, 0xf4, 0x8c, 0x72, 0xb1, 0x9f, 0x56, 0x26,
	0x20, 0xec, 0x97, 0x4f, 0x48, 0x58, 0xb4, 0xf0, 0x40, 0x7f, 0x35, 0x9a, 0x30, 0x6f, 0xe8, 0x30,
	0xdf, 0x02, 0xfc, 0x5b, 0xe0, 0x73, 0x29, 0x4f, 0xda, 0x57, 0x81, 0x01, 0x3a, 0xf8, 0xa1, 0xa1,
	0xd2, 0x13, 0xe2, 0x92, 0x29, 0xda, 0xd6, 0x2f, 0x9f, 0x10, 0xaa, 0x68, 0xe1, 0x81, 0xce, 0xa2,
	0x18, 0x2f, 0x22, 0x28, 0x40, 0x59, 0x6f, 0xe9, 0xd1, 0x65, 0xd2, 0xbb, 0x23, 0xa5, 0x97, 0x24,
	0xa8, 0x48, 0x22, 0x25, 0x23, 0x1a, 0x98, 0xe8, 0x15, 0x62, 0x75, 0x7d, 0x61, 0x60, 0x42, 0xa5,
	0x07, 0x7e, 0x69, 0x68, 0x6e, 0xcf, 0x17, 0x21, 0xe3, 0x01, 0x71, 0xc4, 0xae, 0xeb, 0xc2, 0xe6,
	0x58, 0xd0, 0x82, 0x5e, 0x92, 0x6e, 0x8c, 0x27, 0xcd, 0x3d, 0x8a, 0xf3, 0x2c, 0x8a, 0xf1, 0x2d,
	0xf4, 0xbf, 0x14, 0x35, 0x69, 0xaf, 0x24, 0x4f, 0x54, 0xfe, 0xab, 0x69, 0xbb, 0x6e, 0xa1, 0xa1,
	0x15, 0x54, 0x96, 0x74, 0xb9, 0xa2, 0x1b, 0xe3, 0x25, 0xa8, 0x16, 0xd7, 0xe4, 0x91, 0xdc, 0x75,
	0x5d, 0xc5, 0x0c, 0x50, 0x49, 0x98, 0xf3, 0x32, 0xbc, 0x9b, 0x41, 0xe5, 0x7c, 0x33, 0xcf, 0x99,
	0x47, 0x60, 0x6b, 0x0a, 0x62, 0x69, 0xb8, 0x19, 0xf2, 0x7b, 0x2d, 0x8a, 0xf1, 0x6d, 0xa4, 0x5f,
	0xcf, 0xcc, 0x99, 0x47, 0xc6, 0x42, 0x63, 0xa8, 0x0d, 0x43, 0xcb, 0x9d, 0x29, 0xea, 0x25, 0xbd,
	0x3a, 0x48, 0x9d, 0x3c, 0xb6, 0x77, 0x62, 0x7f, 0x6a, 0x08, 0xe5, 0xbb, 0x03, 0x63, 0x0a, 0xfe,
	0x9b, 0xb1, 0x5f, 0x44, 0x31, 0x5e, 0x43, 0xab, 0xd7, 0xa2, 0x8f, 0xa5, 0x5e, 0x04, 0x18, 0xa6,
	0x4e, 0xfa, 0xac, 0x0f, 0xf5, 0x39, 0x87, 0xfd, 0x34, 0x83, 0xaa, 0x85, 0x89, 0xdd, 0x0f, 0xd9,
	0x6b, 0xe2, 0x08, 0x0e, 0xdb, 0xd3, 0x4d, 0x7a, 0xe6, 0x93, 0x11, 0x3c, 0x98, 0x26, 0x82, 0xdc,
	0xc8, 0x03, 0xfd, 0x83, 0x1c, 0x83, 0xbb, 0x68, 0x7d, 0xec, 0xe8, 0x37, 0x83, 0xd4, 0x36, 0x36,
	0x98, 0x35, 0x58, 0x1d, 0x79, 0x06, 0xb2, 0xf7, 0xaa, 0x2c, 0x56, 0x60, 0x79, 0x68, 0x2a, 0xb2,
	0x37, 0x3c, 0xda, 0x79, 0xf9, 0xb0, 0x45, 0x45, 0xbb, 0x73, 0x28, 0x2f, 0x52, 0x53, 0xa2, 0xf5,
	0xae, 0x7c, 0x73, 0xa2, 0x3f, 0x21, 0x87, 0xb3, 0xea, 0xca, 0xbf, 0xf7, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x75, 0x4d, 0xa3, 0x84, 0x09, 0x00, 0x00,
}