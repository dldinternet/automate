// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/request/nodes.proto

package request // import "github.com/chef/automate/api/external/cfgmgmt/request"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import query "github.com/chef/automate/api/external/common/query"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Nodes struct {
	Filter               []string          `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	Pagination           *query.Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Sorting              *query.Sorting    `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_8db8c7d15de50264, []int{0}
}
func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (dst *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(dst, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Nodes) GetPagination() *query.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Nodes) GetSorting() *query.Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

type Node struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_8db8c7d15de50264, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type NodeRun struct {
	NodeId               string               `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	RunId                string               `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NodeRun) Reset()         { *m = NodeRun{} }
func (m *NodeRun) String() string { return proto.CompactTextString(m) }
func (*NodeRun) ProtoMessage()    {}
func (*NodeRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_8db8c7d15de50264, []int{2}
}
func (m *NodeRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRun.Unmarshal(m, b)
}
func (m *NodeRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRun.Marshal(b, m, deterministic)
}
func (dst *NodeRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRun.Merge(dst, src)
}
func (m *NodeRun) XXX_Size() int {
	return xxx_messageInfo_NodeRun.Size(m)
}
func (m *NodeRun) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRun.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRun proto.InternalMessageInfo

func (m *NodeRun) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRun) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *NodeRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type Runs struct {
	NodeId     string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Filter     []string          `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	Pagination *query.Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// TODO: (@afiune) Should we standardize these parameters as well?
	Start                string   `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Runs) Reset()         { *m = Runs{} }
func (m *Runs) String() string { return proto.CompactTextString(m) }
func (*Runs) ProtoMessage()    {}
func (*Runs) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_8db8c7d15de50264, []int{3}
}
func (m *Runs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runs.Unmarshal(m, b)
}
func (m *Runs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runs.Marshal(b, m, deterministic)
}
func (dst *Runs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runs.Merge(dst, src)
}
func (m *Runs) XXX_Size() int {
	return xxx_messageInfo_Runs.Size(m)
}
func (m *Runs) XXX_DiscardUnknown() {
	xxx_messageInfo_Runs.DiscardUnknown(m)
}

var xxx_messageInfo_Runs proto.InternalMessageInfo

func (m *Runs) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Runs) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Runs) GetPagination() *query.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Runs) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Runs) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func init() {
	proto.RegisterType((*Nodes)(nil), "chef.automate.api.cfgmgmt.request.Nodes")
	proto.RegisterType((*Node)(nil), "chef.automate.api.cfgmgmt.request.Node")
	proto.RegisterType((*NodeRun)(nil), "chef.automate.api.cfgmgmt.request.NodeRun")
	proto.RegisterType((*Runs)(nil), "chef.automate.api.cfgmgmt.request.Runs")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/request/nodes.proto", fileDescriptor_nodes_8db8c7d15de50264)
}

var fileDescriptor_nodes_8db8c7d15de50264 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0xc7, 0x95, 0xa6, 0x69, 0x9e, 0xc7, 0x2c, 0xc8, 0xe2, 0x25, 0xea, 0xd2, 0x92, 0x85, 0xaa,
	0x83, 0x2d, 0x81, 0x2a, 0x66, 0xd8, 0xca, 0x80, 0x50, 0x60, 0x62, 0xa9, 0xdc, 0xe6, 0x92, 0x5a,
	0xaa, 0xed, 0xd4, 0xb1, 0x25, 0xf8, 0x4c, 0x48, 0x7c, 0x46, 0x64, 0x27, 0xe1, 0x45, 0xa2, 0x65,
	0x60, 0xcb, 0x9d, 0xfe, 0x77, 0xbf, 0xcb, 0x4f, 0x46, 0x13, 0x56, 0x71, 0x0a, 0xcf, 0x06, 0xb4,
	0x64, 0x1b, 0xba, 0x2a, 0x4a, 0x51, 0x0a, 0x43, 0x35, 0x6c, 0x2d, 0xd4, 0x86, 0x4a, 0x95, 0x43,
	0x4d, 0x2a, 0xad, 0x8c, 0xc2, 0x67, 0xab, 0x35, 0x14, 0x84, 0x59, 0xa3, 0x04, 0x33, 0x40, 0x58,
	0xc5, 0x49, 0x1b, 0x27, 0x6d, 0x7c, 0x38, 0x2a, 0x95, 0x2a, 0x37, 0x40, 0xfd, 0xc0, 0xd2, 0x16,
	0xd4, 0x70, 0x01, 0xb5, 0x61, 0xa2, 0x6a, 0x76, 0x0c, 0xa7, 0xdf, 0x69, 0x4a, 0x08, 0x25, 0xe9,
	0xd6, 0x82, 0x7e, 0xa1, 0x15, 0xd3, 0x4c, 0x80, 0x01, 0xdd, 0xf2, 0xd2, 0xb7, 0x00, 0x45, 0x77,
	0x8e, 0x8f, 0x4f, 0xd0, 0xa0, 0xe0, 0x1b, 0x03, 0x3a, 0x09, 0xc6, 0xe1, 0xe4, 0x7f, 0xd6, 0x56,
	0xf8, 0x16, 0xa1, 0x8a, 0x95, 0x5c, 0x32, 0xc3, 0x95, 0x4c, 0x7a, 0xe3, 0x60, 0x72, 0x70, 0x31,
	0x25, 0x3f, 0x9c, 0xe9, 0x39, 0xc4, 0x73, 0xc8, 0xfd, 0xc7, 0x44, 0xf6, 0x65, 0x1a, 0x5f, 0xa3,
	0xb8, 0x56, 0xda, 0x70, 0x59, 0x26, 0xa1, 0x5f, 0x74, 0xfe, 0xdb, 0xa2, 0x87, 0x26, 0x9e, 0x75,
	0x73, 0xe9, 0x08, 0xf5, 0xdd, 0xbd, 0xf8, 0x14, 0xc5, 0xce, 0xdb, 0x82, 0xe7, 0x49, 0x30, 0x0e,
	0xdc, 0xbd, 0xae, 0x9c, 0xe7, 0x69, 0x85, 0x62, 0x17, 0xc8, 0xac, 0xdc, 0x99, 0xc1, 0xc7, 0x68,
	0xa0, 0xad, 0x74, 0xfd, 0x9e, 0xef, 0x47, 0xda, 0xca, 0x79, 0x8e, 0x67, 0xe8, 0x1f, 0xc8, 0x7c,
	0xe1, 0x7c, 0xb6, 0xf7, 0x0d, 0x49, 0x23, 0x9b, 0x74, 0xb2, 0xc9, 0x63, 0x27, 0x3b, 0x8b, 0x41,
	0xe6, 0xae, 0x4a, 0x5f, 0x03, 0xd4, 0xcf, 0xac, 0xac, 0x77, 0xf3, 0x3e, 0xdd, 0xf6, 0xf6, 0xb8,
	0x0d, 0xff, 0xe4, 0xf6, 0x08, 0x45, 0xb5, 0x61, 0xda, 0x24, 0xfd, 0xe6, 0x97, 0x7c, 0x81, 0x0f,
	0x51, 0x08, 0x32, 0x4f, 0x22, 0xdf, 0x73, 0x9f, 0x37, 0x57, 0x4f, 0xb3, 0x92, 0x9b, 0xb5, 0x5d,
	0xba, 0xad, 0xd4, 0xb1, 0x68, 0xc7, 0xa2, 0xfb, 0x9e, 0xe9, 0x72, 0xe0, 0x1d, 0x5c, 0xbe, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0xc1, 0x6d, 0x8f, 0xcd, 0x02, 0x00, 0x00,
}