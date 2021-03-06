// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ui.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	ui.proto

It has these top-level messages:
	Event
	Statistics
	PingRequest
	PingReply
	Connection
	Operator
	Rule
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Event struct {
	Time       string      `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	Connection *Connection `protobuf:"bytes,2,opt,name=connection" json:"connection,omitempty"`
	Rule       *Rule       `protobuf:"bytes,3,opt,name=rule" json:"rule,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Event) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *Event) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type Statistics struct {
	DaemonVersion string            `protobuf:"bytes,1,opt,name=daemon_version,json=daemonVersion" json:"daemon_version,omitempty"`
	Rules         uint64            `protobuf:"varint,2,opt,name=rules" json:"rules,omitempty"`
	Uptime        uint64            `protobuf:"varint,3,opt,name=uptime" json:"uptime,omitempty"`
	DnsResponses  uint64            `protobuf:"varint,4,opt,name=dns_responses,json=dnsResponses" json:"dns_responses,omitempty"`
	Connections   uint64            `protobuf:"varint,5,opt,name=connections" json:"connections,omitempty"`
	Ignored       uint64            `protobuf:"varint,6,opt,name=ignored" json:"ignored,omitempty"`
	Accepted      uint64            `protobuf:"varint,7,opt,name=accepted" json:"accepted,omitempty"`
	Dropped       uint64            `protobuf:"varint,8,opt,name=dropped" json:"dropped,omitempty"`
	RuleHits      uint64            `protobuf:"varint,9,opt,name=rule_hits,json=ruleHits" json:"rule_hits,omitempty"`
	RuleMisses    uint64            `protobuf:"varint,10,opt,name=rule_misses,json=ruleMisses" json:"rule_misses,omitempty"`
	ByProto       map[string]uint64 `protobuf:"bytes,11,rep,name=by_proto,json=byProto" json:"by_proto,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ByAddress     map[string]uint64 `protobuf:"bytes,12,rep,name=by_address,json=byAddress" json:"by_address,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ByHost        map[string]uint64 `protobuf:"bytes,13,rep,name=by_host,json=byHost" json:"by_host,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ByPort        map[string]uint64 `protobuf:"bytes,14,rep,name=by_port,json=byPort" json:"by_port,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ByUid         map[string]uint64 `protobuf:"bytes,15,rep,name=by_uid,json=byUid" json:"by_uid,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	ByExecutable  map[string]uint64 `protobuf:"bytes,16,rep,name=by_executable,json=byExecutable" json:"by_executable,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Events        []*Event          `protobuf:"bytes,17,rep,name=events" json:"events,omitempty"`
}

func (m *Statistics) Reset()                    { *m = Statistics{} }
func (m *Statistics) String() string            { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()               {}
func (*Statistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Statistics) GetDaemonVersion() string {
	if m != nil {
		return m.DaemonVersion
	}
	return ""
}

func (m *Statistics) GetRules() uint64 {
	if m != nil {
		return m.Rules
	}
	return 0
}

func (m *Statistics) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Statistics) GetDnsResponses() uint64 {
	if m != nil {
		return m.DnsResponses
	}
	return 0
}

func (m *Statistics) GetConnections() uint64 {
	if m != nil {
		return m.Connections
	}
	return 0
}

func (m *Statistics) GetIgnored() uint64 {
	if m != nil {
		return m.Ignored
	}
	return 0
}

func (m *Statistics) GetAccepted() uint64 {
	if m != nil {
		return m.Accepted
	}
	return 0
}

func (m *Statistics) GetDropped() uint64 {
	if m != nil {
		return m.Dropped
	}
	return 0
}

func (m *Statistics) GetRuleHits() uint64 {
	if m != nil {
		return m.RuleHits
	}
	return 0
}

func (m *Statistics) GetRuleMisses() uint64 {
	if m != nil {
		return m.RuleMisses
	}
	return 0
}

func (m *Statistics) GetByProto() map[string]uint64 {
	if m != nil {
		return m.ByProto
	}
	return nil
}

func (m *Statistics) GetByAddress() map[string]uint64 {
	if m != nil {
		return m.ByAddress
	}
	return nil
}

func (m *Statistics) GetByHost() map[string]uint64 {
	if m != nil {
		return m.ByHost
	}
	return nil
}

func (m *Statistics) GetByPort() map[string]uint64 {
	if m != nil {
		return m.ByPort
	}
	return nil
}

func (m *Statistics) GetByUid() map[string]uint64 {
	if m != nil {
		return m.ByUid
	}
	return nil
}

func (m *Statistics) GetByExecutable() map[string]uint64 {
	if m != nil {
		return m.ByExecutable
	}
	return nil
}

func (m *Statistics) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type PingRequest struct {
	Id    uint64      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Stats *Statistics `protobuf:"bytes,2,opt,name=stats" json:"stats,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PingRequest) GetStats() *Statistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

type PingReply struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PingReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Connection struct {
	Protocol    string   `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	SrcIp       string   `protobuf:"bytes,2,opt,name=src_ip,json=srcIp" json:"src_ip,omitempty"`
	SrcPort     uint32   `protobuf:"varint,3,opt,name=src_port,json=srcPort" json:"src_port,omitempty"`
	DstIp       string   `protobuf:"bytes,4,opt,name=dst_ip,json=dstIp" json:"dst_ip,omitempty"`
	DstHost     string   `protobuf:"bytes,5,opt,name=dst_host,json=dstHost" json:"dst_host,omitempty"`
	DstPort     uint32   `protobuf:"varint,6,opt,name=dst_port,json=dstPort" json:"dst_port,omitempty"`
	UserId      uint32   `protobuf:"varint,7,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ProcessId   uint32   `protobuf:"varint,8,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	ProcessPath string   `protobuf:"bytes,9,opt,name=process_path,json=processPath" json:"process_path,omitempty"`
	ProcessArgs []string `protobuf:"bytes,10,rep,name=process_args,json=processArgs" json:"process_args,omitempty"`
}

func (m *Connection) Reset()                    { *m = Connection{} }
func (m *Connection) String() string            { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()               {}
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Connection) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Connection) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

func (m *Connection) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *Connection) GetDstIp() string {
	if m != nil {
		return m.DstIp
	}
	return ""
}

func (m *Connection) GetDstHost() string {
	if m != nil {
		return m.DstHost
	}
	return ""
}

func (m *Connection) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *Connection) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Connection) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *Connection) GetProcessPath() string {
	if m != nil {
		return m.ProcessPath
	}
	return ""
}

func (m *Connection) GetProcessArgs() []string {
	if m != nil {
		return m.ProcessArgs
	}
	return nil
}

type Operator struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Operand string `protobuf:"bytes,2,opt,name=operand" json:"operand,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *Operator) Reset()                    { *m = Operator{} }
func (m *Operator) String() string            { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()               {}
func (*Operator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Operator) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Operator) GetOperand() string {
	if m != nil {
		return m.Operand
	}
	return ""
}

func (m *Operator) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Rule struct {
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Action   string    `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Duration string    `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	Operator *Operator `protobuf:"bytes,4,opt,name=operator" json:"operator,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Rule) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Rule) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "protocol.Event")
	proto.RegisterType((*Statistics)(nil), "protocol.Statistics")
	proto.RegisterType((*PingRequest)(nil), "protocol.PingRequest")
	proto.RegisterType((*PingReply)(nil), "protocol.PingReply")
	proto.RegisterType((*Connection)(nil), "protocol.Connection")
	proto.RegisterType((*Operator)(nil), "protocol.Operator")
	proto.RegisterType((*Rule)(nil), "protocol.Rule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UI service

type UIClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	AskRule(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Rule, error)
}

type uIClient struct {
	cc *grpc.ClientConn
}

func NewUIClient(cc *grpc.ClientConn) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/protocol.UI/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) AskRule(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := grpc.Invoke(ctx, "/protocol.UI/AskRule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UI service

type UIServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	AskRule(context.Context, *Connection) (*Rule, error)
}

func RegisterUIServer(s *grpc.Server, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_AskRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).AskRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UI/AskRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).AskRule(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UI_Ping_Handler,
		},
		{
			MethodName: "AskRule",
			Handler:    _UI_AskRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ui.proto",
}

func init() { proto.RegisterFile("ui.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x6d, 0x6f, 0xdc, 0x44,
	0x10, 0x6e, 0xee, 0xc5, 0x67, 0xcf, 0xbd, 0xb4, 0x5d, 0x1a, 0x58, 0xae, 0x42, 0xbd, 0xba, 0x02,
	0x22, 0x3e, 0x9c, 0x44, 0xa8, 0x50, 0x5b, 0x55, 0x42, 0x29, 0x8a, 0xd4, 0x13, 0x20, 0xa2, 0x45,
	0xe5, 0xab, 0x65, 0x7b, 0x57, 0x89, 0xd5, 0x8b, 0xd7, 0xec, 0xac, 0x23, 0xfc, 0x85, 0x7f, 0xc3,
	0x6f, 0xe1, 0x6f, 0xa1, 0x9d, 0xb5, 0xcf, 0x47, 0x4a, 0x2a, 0xdd, 0xa7, 0xdb, 0xe7, 0x79, 0xe6,
	0x99, 0x1b, 0x8d, 0x67, 0x67, 0x21, 0xac, 0x8b, 0x75, 0x65, 0xb4, 0xd5, 0x2c, 0xa4, 0x9f, 0x5c,
	0x6f, 0xe3, 0x1a, 0xc6, 0xe7, 0x37, 0xaa, 0xb4, 0x8c, 0xc1, 0xc8, 0x16, 0xd7, 0x8a, 0x1f, 0xad,
	0x8e, 0x4e, 0x22, 0x41, 0x67, 0xf6, 0x1c, 0x20, 0xd7, 0x65, 0xa9, 0x72, 0x5b, 0xe8, 0x92, 0x0f,
	0x56, 0x47, 0x27, 0xd3, 0xd3, 0x47, 0xeb, 0xce, 0xbb, 0xfe, 0x71, 0xa7, 0x89, 0xbd, 0x38, 0x16,
	0xc3, 0xc8, 0xd4, 0x5b, 0xc5, 0x87, 0x14, 0xbf, 0xe8, 0xe3, 0x45, 0xbd, 0x55, 0x82, 0xb4, 0xf8,
	0x9f, 0x10, 0xe0, 0x37, 0x9b, 0xda, 0x02, 0x6d, 0x91, 0x23, 0xfb, 0x12, 0x16, 0x32, 0x55, 0xd7,
	0xba, 0x4c, 0x6e, 0x94, 0x41, 0xf7, 0x67, 0xbe, 0x8c, 0xb9, 0x67, 0x7f, 0xf7, 0x24, 0x7b, 0x04,
	0x63, 0xe7, 0x46, 0x2a, 0x65, 0x24, 0x3c, 0x60, 0x9f, 0x42, 0x50, 0x57, 0x54, 0xfb, 0x90, 0xe8,
	0x16, 0xb1, 0x67, 0x30, 0x97, 0x25, 0x26, 0x46, 0x61, 0xa5, 0x4b, 0x54, 0xc8, 0x47, 0x24, 0xcf,
	0x64, 0x89, 0xa2, 0xe3, 0xd8, 0x0a, 0xa6, 0x7d, 0xe9, 0xc8, 0xc7, 0x14, 0xb2, 0x4f, 0x31, 0x0e,
	0x93, 0xe2, 0xb2, 0xd4, 0x46, 0x49, 0x1e, 0x90, 0xda, 0x41, 0xb6, 0x84, 0x30, 0xcd, 0x73, 0x55,
	0x59, 0x25, 0xf9, 0x84, 0xa4, 0x1d, 0x76, 0x2e, 0x69, 0x74, 0x55, 0x29, 0xc9, 0x43, 0xef, 0x6a,
	0x21, 0x7b, 0x0c, 0x91, 0xab, 0x3b, 0xb9, 0x2a, 0x2c, 0xf2, 0xc8, 0xdb, 0x1c, 0xf1, 0xb6, 0xb0,
	0xc8, 0x9e, 0xc0, 0x94, 0xc4, 0xeb, 0x02, 0x5d, 0xc5, 0x40, 0x32, 0x38, 0xea, 0x17, 0x62, 0xd8,
	0x6b, 0x08, 0xb3, 0x26, 0xa1, 0x96, 0xf2, 0xe9, 0x6a, 0x78, 0x32, 0x3d, 0x7d, 0xda, 0x37, 0xb8,
	0xef, 0xe8, 0xfa, 0x4d, 0x73, 0xe1, 0xd8, 0xf3, 0xd2, 0x9a, 0x46, 0x4c, 0x32, 0x8f, 0xd8, 0x1b,
	0x80, 0xac, 0x49, 0x52, 0x29, 0x8d, 0x42, 0xe4, 0x33, 0xf2, 0x3f, 0xbb, 0xc3, 0x7f, 0xe6, 0xa3,
	0x7c, 0x86, 0x28, 0xeb, 0x30, 0x7b, 0x09, 0x93, 0xac, 0x49, 0xae, 0x34, 0x5a, 0x3e, 0xa7, 0x04,
	0xab, 0x3b, 0x12, 0xbc, 0xd5, 0x68, 0xbd, 0x3b, 0xc8, 0x08, 0xb4, 0xd6, 0x4a, 0x1b, 0xcb, 0x17,
	0x1f, 0xb5, 0x5e, 0x68, 0xd3, 0x5b, 0x1d, 0x60, 0xdf, 0x43, 0x90, 0x35, 0x49, 0x5d, 0x48, 0x7e,
	0x9f, 0x9c, 0x4f, 0xee, 0x70, 0xbe, 0x2b, 0xa4, 0x37, 0x8e, 0x33, 0x77, 0x66, 0x3f, 0xc1, 0x3c,
	0x6b, 0x12, 0xf5, 0xa7, 0xca, 0x6b, 0x9b, 0x66, 0x5b, 0xc5, 0x1f, 0x90, 0xfd, 0xab, 0x3b, 0xec,
	0xe7, 0xbb, 0x40, 0x9f, 0x65, 0x96, 0xed, 0x51, 0xec, 0x6b, 0x08, 0x94, 0xbb, 0x2c, 0xc8, 0x1f,
	0x52, 0x96, 0xfb, 0x7d, 0x16, 0xba, 0x44, 0xa2, 0x95, 0x97, 0xaf, 0x60, 0xb6, 0xff, 0x01, 0xd8,
	0x03, 0x18, 0xbe, 0x57, 0x4d, 0x3b, 0xd4, 0xee, 0xe8, 0x46, 0xf9, 0x26, 0xdd, 0xd6, 0xaa, 0x1b,
	0x65, 0x02, 0xaf, 0x06, 0x2f, 0x8e, 0x96, 0xaf, 0x61, 0xf1, 0xdf, 0xe6, 0x1f, 0xe4, 0x7e, 0x09,
	0xd3, 0xbd, 0xce, 0x1f, 0x6e, 0xdd, 0x75, 0xfe, 0x20, 0xeb, 0x0b, 0x80, 0xbe, 0xf5, 0x07, 0x39,
	0x7f, 0x80, 0x87, 0x1f, 0x74, 0xfd, 0x90, 0x04, 0xf1, 0x06, 0xa6, 0x17, 0x45, 0x79, 0x29, 0xd4,
	0x1f, 0xb5, 0x42, 0xcb, 0x16, 0x30, 0x28, 0x24, 0x39, 0x47, 0x62, 0x50, 0x48, 0xf6, 0x0d, 0x8c,
	0xd1, 0xa6, 0x16, 0x3f, 0xdc, 0x5e, 0xfd, 0x77, 0x17, 0x3e, 0x24, 0x7e, 0x0c, 0x91, 0x4f, 0x55,
	0x6d, 0x9b, 0xdb, 0x89, 0xe2, 0xbf, 0x07, 0x00, 0xfd, 0xc2, 0x73, 0x77, 0xbf, 0xcb, 0xd4, 0xd6,
	0xb9, 0xc3, 0xec, 0x18, 0x02, 0x34, 0x79, 0x52, 0x54, 0xf4, 0xa7, 0x91, 0x18, 0xa3, 0xc9, 0x37,
	0x15, 0xfb, 0x1c, 0x42, 0x47, 0xd3, 0xf8, 0xbb, 0x4d, 0x35, 0x17, 0x13, 0x34, 0x39, 0x4d, 0xf7,
	0x31, 0x04, 0x12, 0xad, 0x73, 0x8c, 0xbc, 0x43, 0xa2, 0xf5, 0x0e, 0x47, 0xd3, 0x5d, 0x1b, 0x93,
	0x30, 0x91, 0x68, 0xe9, 0x2a, 0xb5, 0x12, 0x25, 0x0b, 0x7c, 0x32, 0x89, 0x96, 0x92, 0x7d, 0x06,
	0x93, 0x1a, 0x95, 0x49, 0x0a, 0xbf, 0x95, 0xe6, 0x22, 0x70, 0x70, 0x23, 0xd9, 0x17, 0x00, 0x95,
	0xd1, 0xb9, 0x42, 0x74, 0x5a, 0x48, 0x5a, 0xd4, 0x32, 0x1b, 0xc9, 0x9e, 0xc2, 0xac, 0x93, 0xab,
	0xd4, 0x5e, 0xd1, 0x6e, 0x8a, 0xc4, 0xb4, 0xe5, 0x2e, 0x52, 0x7b, 0xb5, 0x1f, 0x92, 0x9a, 0x4b,
	0xb7, 0x9f, 0x86, 0x7b, 0x21, 0x67, 0xe6, 0x12, 0xe3, 0x9f, 0x21, 0xfc, 0xb5, 0x52, 0x26, 0xb5,
	0xda, 0xd0, 0x9b, 0xd2, 0x54, 0xfd, 0x9b, 0xd2, 0x54, 0xca, 0x2d, 0x46, 0xed, 0xf4, 0x52, 0xb6,
	0xdd, 0xe9, 0xa0, 0x8b, 0x96, 0xa9, 0x4d, 0xa9, 0x37, 0x91, 0xa0, 0x73, 0xfc, 0x17, 0x8c, 0xdc,
	0xab, 0xe1, 0xb4, 0x32, 0xed, 0x5f, 0x27, 0x77, 0x76, 0x7b, 0x3f, 0xed, 0x5f, 0xa6, 0x48, 0xb4,
	0xc8, 0x7d, 0x1a, 0x59, 0x9b, 0x94, 0x14, 0x9f, 0x6b, 0x87, 0xd9, 0x1a, 0x42, 0xdd, 0x56, 0x47,
	0xad, 0x9e, 0x9e, 0xb2, 0x7e, 0x22, 0xba, 0xba, 0xc5, 0x2e, 0xe6, 0xf4, 0x1a, 0x06, 0xef, 0x36,
	0xec, 0x39, 0x8c, 0xdc, 0x60, 0xb0, 0xe3, 0x3e, 0x76, 0x6f, 0xe6, 0x96, 0x9f, 0xdc, 0xa6, 0xab,
	0x6d, 0x13, 0xdf, 0x63, 0xdf, 0xc2, 0xe4, 0x0c, 0xdf, 0x53, 0xf9, 0xff, 0xfb, 0x68, 0x2e, 0x6f,
	0x3d, 0x8d, 0xf1, 0xbd, 0x2c, 0x20, 0xe2, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x94,
	0x16, 0x93, 0xaa, 0x07, 0x00, 0x00,
}
