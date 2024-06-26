// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/grpc/v1/push.proto

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
	PushService_Push_FullMethodName = "/push_v1.PushService/Push"
)

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushServiceClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type pushServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushServiceClient(cc grpc.ClientConnInterface) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, PushService_Push_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServiceServer is the server API for PushService service.
// All implementations must embed UnimplementedPushServiceServer
// for forward compatibility
type PushServiceServer interface {
	Push(context.Context, *PushRequest) (*PushResponse, error)
	mustEmbedUnimplementedPushServiceServer()
}

// UnimplementedPushServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushServiceServer struct {
}

func (UnimplementedPushServiceServer) Push(context.Context, *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedPushServiceServer) mustEmbedUnimplementedPushServiceServer() {}

// UnsafePushServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServiceServer will
// result in compilation errors.
type UnsafePushServiceServer interface {
	mustEmbedUnimplementedPushServiceServer()
}

func RegisterPushServiceServer(s grpc.ServiceRegistrar, srv PushServiceServer) {
	s.RegisterService(&PushService_ServiceDesc, srv)
}

func _PushService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushService_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PushService_ServiceDesc is the grpc.ServiceDesc for PushService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "push_v1.PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _PushService_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/v1/push.proto",
}
