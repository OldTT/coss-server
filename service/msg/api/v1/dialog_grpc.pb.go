// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: api/v1/dialog.proto

package v1

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
	DialogService_CreateDialog_FullMethodName      = "/v1.DialogService/CreateDialog"
	DialogService_JoinDialog_FullMethodName        = "/v1.DialogService/JoinDialog"
	DialogService_GetUserDialogList_FullMethodName = "/v1.DialogService/GetUserDialogList"
	DialogService_GetDialogByIds_FullMethodName    = "/v1.DialogService/GetDialogByIds"
)

// DialogServiceClient is the client API for DialogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DialogServiceClient interface {
	CreateDialog(ctx context.Context, in *CreateDialogRequest, opts ...grpc.CallOption) (*CreateDialogResponse, error)
	JoinDialog(ctx context.Context, in *JoinDialogRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUserDialogList(ctx context.Context, in *GetUserDialogListRequest, opts ...grpc.CallOption) (*GetUserDialogListResponse, error)
	GetDialogByIds(ctx context.Context, in *GetDialogByIdsRequest, opts ...grpc.CallOption) (*GetDialogByIdsResponse, error)
}

type dialogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDialogServiceClient(cc grpc.ClientConnInterface) DialogServiceClient {
	return &dialogServiceClient{cc}
}

func (c *dialogServiceClient) CreateDialog(ctx context.Context, in *CreateDialogRequest, opts ...grpc.CallOption) (*CreateDialogResponse, error) {
	out := new(CreateDialogResponse)
	err := c.cc.Invoke(ctx, DialogService_CreateDialog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dialogServiceClient) JoinDialog(ctx context.Context, in *JoinDialogRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, DialogService_JoinDialog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dialogServiceClient) GetUserDialogList(ctx context.Context, in *GetUserDialogListRequest, opts ...grpc.CallOption) (*GetUserDialogListResponse, error) {
	out := new(GetUserDialogListResponse)
	err := c.cc.Invoke(ctx, DialogService_GetUserDialogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dialogServiceClient) GetDialogByIds(ctx context.Context, in *GetDialogByIdsRequest, opts ...grpc.CallOption) (*GetDialogByIdsResponse, error) {
	out := new(GetDialogByIdsResponse)
	err := c.cc.Invoke(ctx, DialogService_GetDialogByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DialogServiceServer is the server API for DialogService service.
// All implementations must embed UnimplementedDialogServiceServer
// for forward compatibility
type DialogServiceServer interface {
	CreateDialog(context.Context, *CreateDialogRequest) (*CreateDialogResponse, error)
	JoinDialog(context.Context, *JoinDialogRequest) (*Empty, error)
	GetUserDialogList(context.Context, *GetUserDialogListRequest) (*GetUserDialogListResponse, error)
	GetDialogByIds(context.Context, *GetDialogByIdsRequest) (*GetDialogByIdsResponse, error)
	mustEmbedUnimplementedDialogServiceServer()
}

// UnimplementedDialogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDialogServiceServer struct {
}

func (UnimplementedDialogServiceServer) CreateDialog(context.Context, *CreateDialogRequest) (*CreateDialogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDialog not implemented")
}
func (UnimplementedDialogServiceServer) JoinDialog(context.Context, *JoinDialogRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinDialog not implemented")
}
func (UnimplementedDialogServiceServer) GetUserDialogList(context.Context, *GetUserDialogListRequest) (*GetUserDialogListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDialogList not implemented")
}
func (UnimplementedDialogServiceServer) GetDialogByIds(context.Context, *GetDialogByIdsRequest) (*GetDialogByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDialogByIds not implemented")
}
func (UnimplementedDialogServiceServer) mustEmbedUnimplementedDialogServiceServer() {}

// UnsafeDialogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DialogServiceServer will
// result in compilation errors.
type UnsafeDialogServiceServer interface {
	mustEmbedUnimplementedDialogServiceServer()
}

func RegisterDialogServiceServer(s grpc.ServiceRegistrar, srv DialogServiceServer) {
	s.RegisterService(&DialogService_ServiceDesc, srv)
}

func _DialogService_CreateDialog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDialogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogServiceServer).CreateDialog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DialogService_CreateDialog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogServiceServer).CreateDialog(ctx, req.(*CreateDialogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DialogService_JoinDialog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinDialogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogServiceServer).JoinDialog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DialogService_JoinDialog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogServiceServer).JoinDialog(ctx, req.(*JoinDialogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DialogService_GetUserDialogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDialogListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogServiceServer).GetUserDialogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DialogService_GetUserDialogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogServiceServer).GetUserDialogList(ctx, req.(*GetUserDialogListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DialogService_GetDialogByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDialogByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogServiceServer).GetDialogByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DialogService_GetDialogByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogServiceServer).GetDialogByIds(ctx, req.(*GetDialogByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DialogService_ServiceDesc is the grpc.ServiceDesc for DialogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DialogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DialogService",
	HandlerType: (*DialogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDialog",
			Handler:    _DialogService_CreateDialog_Handler,
		},
		{
			MethodName: "JoinDialog",
			Handler:    _DialogService_JoinDialog_Handler,
		},
		{
			MethodName: "GetUserDialogList",
			Handler:    _DialogService_GetUserDialogList_Handler,
		},
		{
			MethodName: "GetDialogByIds",
			Handler:    _DialogService_GetDialogByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/dialog.proto",
}
