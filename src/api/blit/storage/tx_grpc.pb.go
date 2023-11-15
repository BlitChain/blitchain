// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: blit/storage/tx.proto

package storage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateParams_FullMethodName  = "/blit.storage.Msg/UpdateParams"
	Msg_CreateStorage_FullMethodName = "/blit.storage.Msg/CreateStorage"
	Msg_UpdateStorage_FullMethodName = "/blit.storage.Msg/UpdateStorage"
	Msg_DeleteStorage_FullMethodName = "/blit.storage.Msg/DeleteStorage"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateStorage(ctx context.Context, in *MsgCreateStorage, opts ...grpc.CallOption) (*MsgCreateStorageResponse, error)
	UpdateStorage(ctx context.Context, in *MsgUpdateStorage, opts ...grpc.CallOption) (*MsgUpdateStorageResponse, error)
	DeleteStorage(ctx context.Context, in *MsgDeleteStorage, opts ...grpc.CallOption) (*MsgDeleteStorageResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateStorage(ctx context.Context, in *MsgCreateStorage, opts ...grpc.CallOption) (*MsgCreateStorageResponse, error) {
	out := new(MsgCreateStorageResponse)
	err := c.cc.Invoke(ctx, Msg_CreateStorage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateStorage(ctx context.Context, in *MsgUpdateStorage, opts ...grpc.CallOption) (*MsgUpdateStorageResponse, error) {
	out := new(MsgUpdateStorageResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateStorage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteStorage(ctx context.Context, in *MsgDeleteStorage, opts ...grpc.CallOption) (*MsgDeleteStorageResponse, error) {
	out := new(MsgDeleteStorageResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteStorage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateStorage(context.Context, *MsgCreateStorage) (*MsgCreateStorageResponse, error)
	UpdateStorage(context.Context, *MsgUpdateStorage) (*MsgUpdateStorageResponse, error)
	DeleteStorage(context.Context, *MsgDeleteStorage) (*MsgDeleteStorageResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateStorage(context.Context, *MsgCreateStorage) (*MsgCreateStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorage not implemented")
}
func (UnimplementedMsgServer) UpdateStorage(context.Context, *MsgUpdateStorage) (*MsgUpdateStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorage not implemented")
}
func (UnimplementedMsgServer) DeleteStorage(context.Context, *MsgDeleteStorage) (*MsgDeleteStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorage not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStorage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStorage(ctx, req.(*MsgCreateStorage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateStorage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateStorage(ctx, req.(*MsgUpdateStorage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteStorage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteStorage(ctx, req.(*MsgDeleteStorage))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blit.storage.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateStorage",
			Handler:    _Msg_CreateStorage_Handler,
		},
		{
			MethodName: "UpdateStorage",
			Handler:    _Msg_UpdateStorage_Handler,
		},
		{
			MethodName: "DeleteStorage",
			Handler:    _Msg_DeleteStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blit/storage/tx.proto",
}
