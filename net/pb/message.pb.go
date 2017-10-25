// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

/*
Package netpb is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Hello
	Peers
	PeerInfo
*/
package netpb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TODO: change field name to lowercase with _
type Hello struct {
	NodeId        string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ClientVersion string `protobuf:"bytes,2,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
}

func (m *Hello) Reset()                    { *m = Hello{} }
func (m *Hello) String() string            { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()               {}
func (*Hello) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

func (m *Hello) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Hello) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type Peers struct {
	Peers []*PeerInfo `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *Peers) Reset()                    { *m = Peers{} }
func (m *Peers) String() string            { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()               {}
func (*Peers) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func (m *Peers) GetPeers() []*PeerInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PeerInfo struct {
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addrs string `protobuf:"bytes,2,opt,name=addrs,proto3" json:"addrs,omitempty"`
}

func (m *PeerInfo) Reset()                    { *m = PeerInfo{} }
func (m *PeerInfo) String() string            { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()               {}
func (*PeerInfo) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{2} }

func (m *PeerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerInfo) GetAddrs() string {
	if m != nil {
		return m.Addrs
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "netpb.Hello")
	proto.RegisterType((*Peers)(nil), "netpb.Peers")
	proto.RegisterType((*PeerInfo)(nil), "netpb.PeerInfo")
}

func init() { proto.RegisterFile("message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x41, 0x0b, 0x82, 0x40,
	0x10, 0x46, 0xd1, 0x58, 0xab, 0x09, 0x0d, 0x96, 0x20, 0x8f, 0x22, 0x08, 0x9e, 0x96, 0xa8, 0x1f,
	0x51, 0xde, 0xc2, 0x43, 0x57, 0xd1, 0x66, 0x8a, 0x05, 0xdb, 0x91, 0x5d, 0xe9, 0xf7, 0x87, 0x6e,
	0x75, 0x9b, 0xef, 0xbd, 0xc3, 0x1b, 0x88, 0x5f, 0xe4, 0x5c, 0xfb, 0x24, 0x35, 0x58, 0x1e, 0x59,
	0x0a, 0x43, 0xe3, 0xd0, 0xe5, 0x67, 0x10, 0x17, 0xea, 0x7b, 0x96, 0x7b, 0x58, 0x1a, 0x46, 0x6a,
	0x34, 0xa6, 0x41, 0x16, 0x94, 0xeb, 0x3a, 0x9a, 0x66, 0x85, 0xb2, 0x80, 0xe4, 0xde, 0x6b, 0x32,
	0x63, 0xf3, 0x26, 0xeb, 0x34, 0x9b, 0x34, 0x9c, 0x7d, 0xec, 0xe9, 0xcd, 0xc3, 0x5c, 0x81, 0xb8,
	0x12, 0x59, 0x27, 0x0b, 0x10, 0xc3, 0x74, 0xa4, 0x41, 0xb6, 0x28, 0x37, 0xc7, 0xad, 0x9a, 0x43,
	0x6a, 0x92, 0x95, 0x79, 0x70, 0xed, 0x6d, 0x7e, 0x80, 0xd5, 0x0f, 0xc9, 0x04, 0xc2, 0x7f, 0x36,
	0xd4, 0x28, 0x77, 0x20, 0x5a, 0x44, 0xeb, 0xbe, 0x25, 0x3f, 0xba, 0x68, 0x7e, 0xfc, 0xf4, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x04, 0xe0, 0xa7, 0xc9, 0xc9, 0x00, 0x00, 0x00,
}