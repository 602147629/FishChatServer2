// Code generated by protoc-gen-go.
// source: register.proto
// DO NOT EDIT!

package rpc

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

type RGLoginReq struct {
	UID        int64  `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	AccessAddr string `protobuf:"bytes,3,opt,name=accessAddr" json:"accessAddr,omitempty"`
}

func (m *RGLoginReq) Reset()                    { *m = RGLoginReq{} }
func (m *RGLoginReq) String() string            { return proto.CompactTextString(m) }
func (*RGLoginReq) ProtoMessage()               {}
func (*RGLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *RGLoginReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *RGLoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RGLoginReq) GetAccessAddr() string {
	if m != nil {
		return m.AccessAddr
	}
	return ""
}

type RGLoginRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *RGLoginRes) Reset()                    { *m = RGLoginRes{} }
func (m *RGLoginRes) String() string            { return proto.CompactTextString(m) }
func (*RGLoginRes) ProtoMessage()               {}
func (*RGLoginRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RGLoginRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *RGLoginRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *RGLoginRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RGAccessReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *RGAccessReq) Reset()                    { *m = RGAccessReq{} }
func (m *RGAccessReq) String() string            { return proto.CompactTextString(m) }
func (*RGAccessReq) ProtoMessage()               {}
func (*RGAccessReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RGAccessReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type RGAccessRes struct {
	ErrCode    uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr     string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	AccessAddr string `protobuf:"bytes,3,opt,name=accessAddr" json:"accessAddr,omitempty"`
}

func (m *RGAccessRes) Reset()                    { *m = RGAccessRes{} }
func (m *RGAccessRes) String() string            { return proto.CompactTextString(m) }
func (*RGAccessRes) ProtoMessage()               {}
func (*RGAccessRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *RGAccessRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *RGAccessRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *RGAccessRes) GetAccessAddr() string {
	if m != nil {
		return m.AccessAddr
	}
	return ""
}

type RGAuthReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *RGAuthReq) Reset()                    { *m = RGAuthReq{} }
func (m *RGAuthReq) String() string            { return proto.CompactTextString(m) }
func (*RGAuthReq) ProtoMessage()               {}
func (*RGAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *RGAuthReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type RGAuthRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *RGAuthRes) Reset()                    { *m = RGAuthRes{} }
func (m *RGAuthRes) String() string            { return proto.CompactTextString(m) }
func (*RGAuthRes) ProtoMessage()               {}
func (*RGAuthRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *RGAuthRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *RGAuthRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *RGAuthRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RGPingReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *RGPingReq) Reset()                    { *m = RGPingReq{} }
func (m *RGPingReq) String() string            { return proto.CompactTextString(m) }
func (*RGPingReq) ProtoMessage()               {}
func (*RGPingReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *RGPingReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type RGPingRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *RGPingRes) Reset()                    { *m = RGPingRes{} }
func (m *RGPingRes) String() string            { return proto.CompactTextString(m) }
func (*RGPingRes) ProtoMessage()               {}
func (*RGPingRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *RGPingRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *RGPingRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type RGOnlineReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *RGOnlineReq) Reset()                    { *m = RGOnlineReq{} }
func (m *RGOnlineReq) String() string            { return proto.CompactTextString(m) }
func (*RGOnlineReq) ProtoMessage()               {}
func (*RGOnlineReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *RGOnlineReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type RGOnlineRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Online  bool   `protobuf:"varint,3,opt,name=online" json:"online,omitempty"`
}

func (m *RGOnlineRes) Reset()                    { *m = RGOnlineRes{} }
func (m *RGOnlineRes) String() string            { return proto.CompactTextString(m) }
func (*RGOnlineRes) ProtoMessage()               {}
func (*RGOnlineRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *RGOnlineRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *RGOnlineRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *RGOnlineRes) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func init() {
	proto.RegisterType((*RGLoginReq)(nil), "rpc.RGLoginReq")
	proto.RegisterType((*RGLoginRes)(nil), "rpc.RGLoginRes")
	proto.RegisterType((*RGAccessReq)(nil), "rpc.RGAccessReq")
	proto.RegisterType((*RGAccessRes)(nil), "rpc.RGAccessRes")
	proto.RegisterType((*RGAuthReq)(nil), "rpc.RGAuthReq")
	proto.RegisterType((*RGAuthRes)(nil), "rpc.RGAuthRes")
	proto.RegisterType((*RGPingReq)(nil), "rpc.RGPingReq")
	proto.RegisterType((*RGPingRes)(nil), "rpc.RGPingRes")
	proto.RegisterType((*RGOnlineReq)(nil), "rpc.RGOnlineReq")
	proto.RegisterType((*RGOnlineRes)(nil), "rpc.RGOnlineRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RegisterServerRPC service

type RegisterServerRPCClient interface {
	Login(ctx context.Context, in *RGLoginReq, opts ...grpc.CallOption) (*RGLoginRes, error)
	RouterAccess(ctx context.Context, in *RGAccessReq, opts ...grpc.CallOption) (*RGAccessRes, error)
	Auth(ctx context.Context, in *RGAuthReq, opts ...grpc.CallOption) (*RGAuthRes, error)
	Ping(ctx context.Context, in *RGPingReq, opts ...grpc.CallOption) (*RGPingRes, error)
	Online(ctx context.Context, in *RGOnlineReq, opts ...grpc.CallOption) (*RGOnlineRes, error)
}

type registerServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewRegisterServerRPCClient(cc *grpc.ClientConn) RegisterServerRPCClient {
	return &registerServerRPCClient{cc}
}

func (c *registerServerRPCClient) Login(ctx context.Context, in *RGLoginReq, opts ...grpc.CallOption) (*RGLoginRes, error) {
	out := new(RGLoginRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServerRPCClient) RouterAccess(ctx context.Context, in *RGAccessReq, opts ...grpc.CallOption) (*RGAccessRes, error) {
	out := new(RGAccessRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/RouterAccess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServerRPCClient) Auth(ctx context.Context, in *RGAuthReq, opts ...grpc.CallOption) (*RGAuthRes, error) {
	out := new(RGAuthRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServerRPCClient) Ping(ctx context.Context, in *RGPingReq, opts ...grpc.CallOption) (*RGPingRes, error) {
	out := new(RGPingRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServerRPCClient) Online(ctx context.Context, in *RGOnlineReq, opts ...grpc.CallOption) (*RGOnlineRes, error) {
	out := new(RGOnlineRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/Online", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegisterServerRPC service

type RegisterServerRPCServer interface {
	Login(context.Context, *RGLoginReq) (*RGLoginRes, error)
	RouterAccess(context.Context, *RGAccessReq) (*RGAccessRes, error)
	Auth(context.Context, *RGAuthReq) (*RGAuthRes, error)
	Ping(context.Context, *RGPingReq) (*RGPingRes, error)
	Online(context.Context, *RGOnlineReq) (*RGOnlineRes, error)
}

func RegisterRegisterServerRPCServer(s *grpc.Server, srv RegisterServerRPCServer) {
	s.RegisterService(&_RegisterServerRPC_serviceDesc, srv)
}

func _RegisterServerRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).Login(ctx, req.(*RGLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterServerRPC_RouterAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGAccessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).RouterAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/RouterAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).RouterAccess(ctx, req.(*RGAccessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterServerRPC_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).Auth(ctx, req.(*RGAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterServerRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGPingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).Ping(ctx, req.(*RGPingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterServerRPC_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/Online",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).Online(ctx, req.(*RGOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.RegisterServerRPC",
	HandlerType: (*RegisterServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RegisterServerRPC_Login_Handler,
		},
		{
			MethodName: "RouterAccess",
			Handler:    _RegisterServerRPC_RouterAccess_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _RegisterServerRPC_Auth_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _RegisterServerRPC_Ping_Handler,
		},
		{
			MethodName: "Online",
			Handler:    _RegisterServerRPC_Online_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}

func init() { proto.RegisterFile("register.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x4b, 0xfb, 0x40,
	0x10, 0x6d, 0x7e, 0xf9, 0x35, 0xda, 0x51, 0x6b, 0x5d, 0xa4, 0x84, 0x80, 0x5a, 0x72, 0x2a, 0x08,
	0x39, 0xa8, 0x57, 0x0f, 0xa5, 0x42, 0x11, 0x04, 0xcb, 0x56, 0xf1, 0x28, 0x9a, 0x0e, 0x31, 0x28,
	0xd9, 0x38, 0xbb, 0xf5, 0x8f, 0xf7, 0x24, 0xd9, 0x6c, 0xf3, 0x45, 0x40, 0x08, 0xde, 0xf2, 0x66,
	0xde, 0xce, 0x9b, 0x8f, 0x17, 0x18, 0x12, 0x46, 0xb1, 0x54, 0x48, 0x41, 0x4a, 0x42, 0x09, 0x66,
	0x53, 0x1a, 0xfa, 0x0f, 0x00, 0x7c, 0x71, 0x27, 0xa2, 0x38, 0xe1, 0xf8, 0xc9, 0x46, 0x60, 0x3f,
	0xde, 0xde, 0xb8, 0xd6, 0xc4, 0x9a, 0xda, 0x3c, 0xfb, 0x64, 0xc7, 0xd0, 0x57, 0xe2, 0x1d, 0x13,
	0xf7, 0xdf, 0xc4, 0x9a, 0x0e, 0x78, 0x0e, 0xd8, 0x29, 0xc0, 0x4b, 0x18, 0xa2, 0x94, 0xb3, 0xf5,
	0x9a, 0x5c, 0x5b, 0xa7, 0x2a, 0x91, 0x5a, 0x55, 0xc9, 0x5c, 0xd8, 0x41, 0xa2, 0xb9, 0x58, 0xa3,
	0xae, 0x7c, 0xc0, 0xb7, 0x90, 0x8d, 0xc1, 0x41, 0xa2, 0x95, 0x22, 0x53, 0xde, 0xa0, 0x52, 0xd5,
	0xae, 0xa8, 0xfa, 0x67, 0xb0, 0xc7, 0x17, 0x33, 0xad, 0xd2, 0xda, 0xac, 0xff, 0x5c, 0x25, 0x74,
	0xd1, 0xfd, 0x6d, 0xae, 0x13, 0x18, 0xf0, 0xc5, 0x6c, 0xa3, 0xde, 0xda, 0xf5, 0x57, 0x65, 0xfa,
	0xef, 0xa6, 0xd6, 0x9a, 0xcb, 0x38, 0x89, 0xda, 0x35, 0xaf, 0xcb, 0x74, 0x07, 0xcd, 0x7c, 0xa7,
	0xf7, 0xc9, 0x47, 0x9c, 0x60, 0x7b, 0xfd, 0xa7, 0x2a, 0xa1, 0xcb, 0x54, 0x63, 0x70, 0x84, 0x7e,
	0xae, 0xc7, 0xda, 0xe5, 0x06, 0x5d, 0x7c, 0x5b, 0x70, 0xc4, 0x8d, 0x23, 0x57, 0x48, 0x5f, 0x48,
	0x7c, 0x39, 0x67, 0xe7, 0xd0, 0xd7, 0xbe, 0x61, 0x87, 0x01, 0xa5, 0x61, 0x50, 0x7a, 0xd3, 0x6b,
	0x04, 0xa4, 0xdf, 0x63, 0x57, 0xb0, 0xcf, 0xc5, 0x46, 0x21, 0xe5, 0x37, 0x67, 0x23, 0x43, 0x29,
	0x3c, 0xe2, 0x35, 0x23, 0xd9, 0xab, 0x29, 0xfc, 0xcf, 0x6e, 0xc4, 0x86, 0xdb, 0x5c, 0x7e, 0x4f,
	0xaf, 0x8e, 0x0d, 0x33, 0xdb, 0x6c, 0xc1, 0x34, 0x57, 0xf0, 0xea, 0x38, 0x63, 0x06, 0xe0, 0xe4,
	0x3b, 0x2a, 0x7a, 0x28, 0x76, 0xea, 0x35, 0x23, 0xd2, 0xef, 0xbd, 0x3a, 0xfa, 0x17, 0xbc, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x19, 0x03, 0x33, 0x2c, 0x94, 0x03, 0x00, 0x00,
}
