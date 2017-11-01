// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: management_rpc.proto

/*
Package rpcpb is a generated protocol buffer package.

It is generated from these files:
	management_rpc.proto

It has these top-level messages:
	NewAccountRequest
	NewAccountResponse
	UnlockAccountRequest
	UnlockAccountResponse
	LockAccountRequest
	LockAccountResponse
	SignTransactionRequest
	SignTransactionResponse
	SendTransactionPassphraseRequest
	SendTransactionPassphraseResponse
*/
package rpcpb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NewAccountRequest struct {
	Passphrase string `protobuf:"bytes,1,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
}

func (m *NewAccountRequest) Reset()                    { *m = NewAccountRequest{} }
func (m *NewAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*NewAccountRequest) ProtoMessage()               {}
func (*NewAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptorManagementRpc, []int{0} }

func (m *NewAccountRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

type NewAccountResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *NewAccountResponse) Reset()                    { *m = NewAccountResponse{} }
func (m *NewAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*NewAccountResponse) ProtoMessage()               {}
func (*NewAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptorManagementRpc, []int{1} }

func (m *NewAccountResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UnlockAccountRequest struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Passphrase string `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
}

func (m *UnlockAccountRequest) Reset()         { *m = UnlockAccountRequest{} }
func (m *UnlockAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UnlockAccountRequest) ProtoMessage()    {}
func (*UnlockAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorManagementRpc, []int{2}
}

func (m *UnlockAccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UnlockAccountRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

type UnlockAccountResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *UnlockAccountResponse) Reset()         { *m = UnlockAccountResponse{} }
func (m *UnlockAccountResponse) String() string { return proto.CompactTextString(m) }
func (*UnlockAccountResponse) ProtoMessage()    {}
func (*UnlockAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorManagementRpc, []int{3}
}

func (m *UnlockAccountResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type LockAccountRequest struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Passphrase string `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
}

func (m *LockAccountRequest) Reset()                    { *m = LockAccountRequest{} }
func (m *LockAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*LockAccountRequest) ProtoMessage()               {}
func (*LockAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptorManagementRpc, []int{4} }

func (m *LockAccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LockAccountRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

type LockAccountResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *LockAccountResponse) Reset()                    { *m = LockAccountResponse{} }
func (m *LockAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*LockAccountResponse) ProtoMessage()               {}
func (*LockAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptorManagementRpc, []int{5} }

func (m *LockAccountResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type SignTransactionRequest struct {
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// Amount of value sending with this transaction.
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *SignTransactionRequest) Reset()         { *m = SignTransactionRequest{} }
func (m *SignTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SignTransactionRequest) ProtoMessage()    {}
func (*SignTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorManagementRpc, []int{6}
}

func (m *SignTransactionRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SignTransactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SignTransactionRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SignTransactionRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type SignTransactionResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SignTransactionResponse) Reset()         { *m = SignTransactionResponse{} }
func (m *SignTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SignTransactionResponse) ProtoMessage()    {}
func (*SignTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorManagementRpc, []int{7}
}

func (m *SignTransactionResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendTransactionPassphraseRequest struct {
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// Amount of value sending with this transaction.
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// from account passphrase
	Passphrase string `protobuf:"bytes,5,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
}

func (m *SendTransactionPassphraseRequest) Reset()         { *m = SendTransactionPassphraseRequest{} }
func (m *SendTransactionPassphraseRequest) String() string { return proto.CompactTextString(m) }
func (*SendTransactionPassphraseRequest) ProtoMessage()    {}
func (*SendTransactionPassphraseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorManagementRpc, []int{8}
}

func (m *SendTransactionPassphraseRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendTransactionPassphraseRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendTransactionPassphraseRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SendTransactionPassphraseRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SendTransactionPassphraseRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

type SendTransactionPassphraseResponse struct {
	// Hex string of transaction hash.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *SendTransactionPassphraseResponse) Reset()         { *m = SendTransactionPassphraseResponse{} }
func (m *SendTransactionPassphraseResponse) String() string { return proto.CompactTextString(m) }
func (*SendTransactionPassphraseResponse) ProtoMessage()    {}
func (*SendTransactionPassphraseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorManagementRpc, []int{9}
}

func (m *SendTransactionPassphraseResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*NewAccountRequest)(nil), "rpcpb.NewAccountRequest")
	proto.RegisterType((*NewAccountResponse)(nil), "rpcpb.NewAccountResponse")
	proto.RegisterType((*UnlockAccountRequest)(nil), "rpcpb.UnlockAccountRequest")
	proto.RegisterType((*UnlockAccountResponse)(nil), "rpcpb.UnlockAccountResponse")
	proto.RegisterType((*LockAccountRequest)(nil), "rpcpb.LockAccountRequest")
	proto.RegisterType((*LockAccountResponse)(nil), "rpcpb.LockAccountResponse")
	proto.RegisterType((*SignTransactionRequest)(nil), "rpcpb.SignTransactionRequest")
	proto.RegisterType((*SignTransactionResponse)(nil), "rpcpb.SignTransactionResponse")
	proto.RegisterType((*SendTransactionPassphraseRequest)(nil), "rpcpb.SendTransactionPassphraseRequest")
	proto.RegisterType((*SendTransactionPassphraseResponse)(nil), "rpcpb.SendTransactionPassphraseResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ManagementService service

type ManagementServiceClient interface {
	// NewAccount create a new account with passphrase
	NewAccount(ctx context.Context, in *NewAccountRequest, opts ...grpc.CallOption) (*NewAccountResponse, error)
	// UnlockAccount unlock account with passphrase
	UnlockAccount(ctx context.Context, in *UnlockAccountRequest, opts ...grpc.CallOption) (*UnlockAccountResponse, error)
	// LockAccount lock account
	LockAccount(ctx context.Context, in *LockAccountRequest, opts ...grpc.CallOption) (*LockAccountResponse, error)
	// SignTransaction sign transaction
	SignTransaction(ctx context.Context, in *SignTransactionRequest, opts ...grpc.CallOption) (*SignTransactionResponse, error)
	// sendTransactionWithPassphrase send transaction with passphrase
	SendTransactionWithPassphrase(ctx context.Context, in *SendTransactionPassphraseRequest, opts ...grpc.CallOption) (*SendTransactionPassphraseResponse, error)
}

type managementServiceClient struct {
	cc *grpc.ClientConn
}

func NewManagementServiceClient(cc *grpc.ClientConn) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) NewAccount(ctx context.Context, in *NewAccountRequest, opts ...grpc.CallOption) (*NewAccountResponse, error) {
	out := new(NewAccountResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ManagementService/NewAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) UnlockAccount(ctx context.Context, in *UnlockAccountRequest, opts ...grpc.CallOption) (*UnlockAccountResponse, error) {
	out := new(UnlockAccountResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ManagementService/UnlockAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) LockAccount(ctx context.Context, in *LockAccountRequest, opts ...grpc.CallOption) (*LockAccountResponse, error) {
	out := new(LockAccountResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ManagementService/LockAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) SignTransaction(ctx context.Context, in *SignTransactionRequest, opts ...grpc.CallOption) (*SignTransactionResponse, error) {
	out := new(SignTransactionResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ManagementService/SignTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) SendTransactionWithPassphrase(ctx context.Context, in *SendTransactionPassphraseRequest, opts ...grpc.CallOption) (*SendTransactionPassphraseResponse, error) {
	out := new(SendTransactionPassphraseResponse)
	err := grpc.Invoke(ctx, "/rpcpb.ManagementService/sendTransactionWithPassphrase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManagementService service

type ManagementServiceServer interface {
	// NewAccount create a new account with passphrase
	NewAccount(context.Context, *NewAccountRequest) (*NewAccountResponse, error)
	// UnlockAccount unlock account with passphrase
	UnlockAccount(context.Context, *UnlockAccountRequest) (*UnlockAccountResponse, error)
	// LockAccount lock account
	LockAccount(context.Context, *LockAccountRequest) (*LockAccountResponse, error)
	// SignTransaction sign transaction
	SignTransaction(context.Context, *SignTransactionRequest) (*SignTransactionResponse, error)
	// sendTransactionWithPassphrase send transaction with passphrase
	SendTransactionWithPassphrase(context.Context, *SendTransactionPassphraseRequest) (*SendTransactionPassphraseResponse, error)
}

func RegisterManagementServiceServer(s *grpc.Server, srv ManagementServiceServer) {
	s.RegisterService(&_ManagementService_serviceDesc, srv)
}

func _ManagementService_NewAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ManagementService/NewAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).NewAccount(ctx, req.(*NewAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_UnlockAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).UnlockAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ManagementService/UnlockAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).UnlockAccount(ctx, req.(*UnlockAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_LockAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).LockAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ManagementService/LockAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).LockAccount(ctx, req.(*LockAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_SignTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).SignTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ManagementService/SignTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).SignTransaction(ctx, req.(*SignTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_SendTransactionWithPassphrase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionPassphraseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).SendTransactionWithPassphrase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ManagementService/SendTransactionWithPassphrase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).SendTransactionWithPassphrase(ctx, req.(*SendTransactionPassphraseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAccount",
			Handler:    _ManagementService_NewAccount_Handler,
		},
		{
			MethodName: "UnlockAccount",
			Handler:    _ManagementService_UnlockAccount_Handler,
		},
		{
			MethodName: "LockAccount",
			Handler:    _ManagementService_LockAccount_Handler,
		},
		{
			MethodName: "SignTransaction",
			Handler:    _ManagementService_SignTransaction_Handler,
		},
		{
			MethodName: "sendTransactionWithPassphrase",
			Handler:    _ManagementService_SendTransactionWithPassphrase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management_rpc.proto",
}

func init() { proto.RegisterFile("management_rpc.proto", fileDescriptorManagementRpc) }

var fileDescriptorManagementRpc = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x25, 0xdd, 0x74, 0x81, 0x61, 0x01, 0xed, 0x50, 0x96, 0x10, 0x68, 0x15, 0x72, 0x21, 0x97,
	0x06, 0x89, 0x1e, 0x38, 0x23, 0x24, 0x4e, 0xa5, 0xaa, 0x52, 0x10, 0x47, 0xe4, 0x3a, 0xa6, 0x89,
	0x68, 0xed, 0x60, 0x3b, 0xe5, 0x3f, 0xf8, 0x1a, 0x3e, 0x0f, 0x35, 0x75, 0xd2, 0x34, 0x49, 0xa9,
	0x90, 0xe0, 0x36, 0xe3, 0x99, 0x79, 0xef, 0x79, 0xfc, 0x64, 0x18, 0x6c, 0x08, 0x27, 0x2b, 0xb6,
	0x61, 0x5c, 0x7f, 0x91, 0x19, 0x0d, 0x33, 0x29, 0xb4, 0xc0, 0xbe, 0xcc, 0x68, 0xb6, 0xf4, 0x27,
	0x70, 0x3d, 0x63, 0x3f, 0xde, 0x52, 0x2a, 0x72, 0xae, 0x23, 0xf6, 0x3d, 0x67, 0x4a, 0xe3, 0x08,
	0x20, 0x23, 0x4a, 0x65, 0x89, 0x24, 0x8a, 0x39, 0x96, 0x67, 0x05, 0x77, 0xa3, 0xda, 0x89, 0x1f,
	0x02, 0xd6, 0x87, 0x54, 0x26, 0xb8, 0x62, 0xe8, 0xc0, 0x6d, 0x12, 0xc7, 0x92, 0x29, 0x65, 0x46,
	0xca, 0xd4, 0x9f, 0xc3, 0xe0, 0x13, 0x5f, 0x0b, 0xfa, 0xad, 0xc1, 0x73, 0x72, 0xa2, 0xa1, 0xa0,
	0xd7, 0x52, 0xf0, 0x0a, 0x1e, 0x37, 0x10, 0x8d, 0x88, 0x1b, 0xb8, 0x94, 0x4c, 0xe5, 0x6b, 0x5d,
	0x20, 0xde, 0x89, 0x4c, 0xe6, 0xcf, 0x00, 0xa7, 0xff, 0x52, 0xc0, 0x18, 0x1e, 0x4d, 0xff, 0x82,
	0x3e, 0x81, 0x9b, 0x45, 0xba, 0xe2, 0x1f, 0x25, 0xe1, 0x8a, 0x50, 0x9d, 0x0a, 0x5e, 0x4a, 0x40,
	0xb0, 0xbf, 0x4a, 0xb1, 0x31, 0xfc, 0x45, 0x8c, 0x0f, 0xa0, 0xa7, 0x85, 0x21, 0xed, 0x69, 0x81,
	0x03, 0xe8, 0x6f, 0xc9, 0x3a, 0x67, 0xce, 0x85, 0x67, 0x05, 0x57, 0xd1, 0x3e, 0xd9, 0x9d, 0x72,
	0xc1, 0x29, 0x73, 0x6c, 0xcf, 0x0a, 0xec, 0x68, 0x9f, 0xf8, 0x63, 0x78, 0xd2, 0x62, 0x32, 0xe2,
	0x10, 0xec, 0x98, 0x68, 0x52, 0x50, 0x5d, 0x45, 0x45, 0xec, 0xff, 0xb4, 0xc0, 0x5b, 0x30, 0x1e,
	0xd7, 0xfa, 0xe7, 0xd5, 0x2d, 0xff, 0x93, 0xc6, 0xc6, 0x72, 0xfb, 0xad, 0xe5, 0xbe, 0x81, 0x17,
	0x7f, 0xd0, 0x74, 0xb8, 0x4d, 0x42, 0x54, 0x52, 0x8a, 0xda, 0xc5, 0xaf, 0x7f, 0x5d, 0xc0, 0xf5,
	0x87, 0xca, 0xed, 0x0b, 0x26, 0xb7, 0x29, 0x65, 0xf8, 0x0e, 0xe0, 0x60, 0x57, 0x74, 0xc2, 0xc2,
	0xf9, 0x61, 0xcb, 0xf6, 0xee, 0xd3, 0x8e, 0xca, 0x9e, 0xcc, 0xbf, 0x85, 0x53, 0xb8, 0x7f, 0xe4,
	0x38, 0x7c, 0x66, 0xba, 0xbb, 0x9c, 0xed, 0x3e, 0xef, 0x2e, 0x56, 0x68, 0xef, 0xe1, 0x5e, 0xcd,
	0x3e, 0x58, 0x32, 0xb7, 0x2d, 0xea, 0xba, 0x5d, 0xa5, 0x0a, 0x27, 0x82, 0x87, 0x8d, 0xd7, 0xc6,
	0xa1, 0x19, 0xe8, 0xf6, 0x9b, 0x3b, 0x3a, 0x55, 0xae, 0x30, 0x35, 0x0c, 0xd5, 0xf1, 0xf6, 0x3f,
	0xa7, 0x3a, 0x39, 0xbc, 0x00, 0xbe, 0x2c, 0x21, 0xce, 0xf8, 0xc6, 0x0d, 0xce, 0x37, 0x96, 0xac,
	0xcb, 0xcb, 0xe2, 0x5b, 0x9a, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x10, 0xad, 0xd1, 0xae,
	0x04, 0x00, 0x00,
}