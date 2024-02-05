// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: blit/blit/tx.proto

package blit

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
	Msg_UpdateParams_FullMethodName       = "/blit.blit.Msg/UpdateParams"
	Msg_MintCoins_FullMethodName          = "/blit.blit.Msg/MintCoins"
	Msg_BurnCoins_FullMethodName          = "/blit.blit.Msg/BurnCoins"
	Msg_ForceTransferCoins_FullMethodName = "/blit.blit.Msg/ForceTransferCoins"
	Msg_SetDenomMetadata_FullMethodName   = "/blit.blit.Msg/SetDenomMetadata"
	Msg_CreateTask_FullMethodName         = "/blit.blit.Msg/CreateTask"
	Msg_DeleteTask_FullMethodName         = "/blit.blit.Msg/DeleteTask"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	MintCoins(ctx context.Context, in *MsgMintCoins, opts ...grpc.CallOption) (*MsgMintCoinsResponse, error)
	BurnCoins(ctx context.Context, in *MsgBurnCoins, opts ...grpc.CallOption) (*MsgBurnCoinsResponse, error)
	ForceTransferCoins(ctx context.Context, in *MsgForceTransferCoins, opts ...grpc.CallOption) (*MsgForceTransferCoinsResponse, error)
	SetDenomMetadata(ctx context.Context, in *MsgSetDenomMetadata, opts ...grpc.CallOption) (*MsgSetDenomMetadataResponse, error)
	CreateTask(ctx context.Context, in *MsgCreateTask, opts ...grpc.CallOption) (*MsgCreateTaskResponse, error)
	DeleteTask(ctx context.Context, in *MsgDeleteTask, opts ...grpc.CallOption) (*MsgDeleteTaskResponse, error)
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

func (c *msgClient) MintCoins(ctx context.Context, in *MsgMintCoins, opts ...grpc.CallOption) (*MsgMintCoinsResponse, error) {
	out := new(MsgMintCoinsResponse)
	err := c.cc.Invoke(ctx, Msg_MintCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnCoins(ctx context.Context, in *MsgBurnCoins, opts ...grpc.CallOption) (*MsgBurnCoinsResponse, error) {
	out := new(MsgBurnCoinsResponse)
	err := c.cc.Invoke(ctx, Msg_BurnCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ForceTransferCoins(ctx context.Context, in *MsgForceTransferCoins, opts ...grpc.CallOption) (*MsgForceTransferCoinsResponse, error) {
	out := new(MsgForceTransferCoinsResponse)
	err := c.cc.Invoke(ctx, Msg_ForceTransferCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetDenomMetadata(ctx context.Context, in *MsgSetDenomMetadata, opts ...grpc.CallOption) (*MsgSetDenomMetadataResponse, error) {
	out := new(MsgSetDenomMetadataResponse)
	err := c.cc.Invoke(ctx, Msg_SetDenomMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateTask(ctx context.Context, in *MsgCreateTask, opts ...grpc.CallOption) (*MsgCreateTaskResponse, error) {
	out := new(MsgCreateTaskResponse)
	err := c.cc.Invoke(ctx, Msg_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteTask(ctx context.Context, in *MsgDeleteTask, opts ...grpc.CallOption) (*MsgDeleteTaskResponse, error) {
	out := new(MsgDeleteTaskResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteTask_FullMethodName, in, out, opts...)
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
	MintCoins(context.Context, *MsgMintCoins) (*MsgMintCoinsResponse, error)
	BurnCoins(context.Context, *MsgBurnCoins) (*MsgBurnCoinsResponse, error)
	ForceTransferCoins(context.Context, *MsgForceTransferCoins) (*MsgForceTransferCoinsResponse, error)
	SetDenomMetadata(context.Context, *MsgSetDenomMetadata) (*MsgSetDenomMetadataResponse, error)
	CreateTask(context.Context, *MsgCreateTask) (*MsgCreateTaskResponse, error)
	DeleteTask(context.Context, *MsgDeleteTask) (*MsgDeleteTaskResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) MintCoins(context.Context, *MsgMintCoins) (*MsgMintCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintCoins not implemented")
}
func (UnimplementedMsgServer) BurnCoins(context.Context, *MsgBurnCoins) (*MsgBurnCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnCoins not implemented")
}
func (UnimplementedMsgServer) ForceTransferCoins(context.Context, *MsgForceTransferCoins) (*MsgForceTransferCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceTransferCoins not implemented")
}
func (UnimplementedMsgServer) SetDenomMetadata(context.Context, *MsgSetDenomMetadata) (*MsgSetDenomMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDenomMetadata not implemented")
}
func (UnimplementedMsgServer) CreateTask(context.Context, *MsgCreateTask) (*MsgCreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedMsgServer) DeleteTask(context.Context, *MsgDeleteTask) (*MsgDeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
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

func _Msg_MintCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintCoins)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintCoins(ctx, req.(*MsgMintCoins))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnCoins)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnCoins(ctx, req.(*MsgBurnCoins))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ForceTransferCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgForceTransferCoins)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ForceTransferCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ForceTransferCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ForceTransferCoins(ctx, req.(*MsgForceTransferCoins))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetDenomMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetDenomMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetDenomMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetDenomMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDenomMetadata(ctx, req.(*MsgSetDenomMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTask(ctx, req.(*MsgCreateTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteTask(ctx, req.(*MsgDeleteTask))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blit.blit.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "MintCoins",
			Handler:    _Msg_MintCoins_Handler,
		},
		{
			MethodName: "BurnCoins",
			Handler:    _Msg_BurnCoins_Handler,
		},
		{
			MethodName: "ForceTransferCoins",
			Handler:    _Msg_ForceTransferCoins_Handler,
		},
		{
			MethodName: "SetDenomMetadata",
			Handler:    _Msg_SetDenomMetadata_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Msg_CreateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Msg_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blit/blit/tx.proto",
}
