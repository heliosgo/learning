// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: task.proto

package task

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

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	UpdateProgress(ctx context.Context, in *UpdateProgressReq, opts ...grpc.CallOption) (*UpdateProgressRsp, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) UpdateProgress(ctx context.Context, in *UpdateProgressReq, opts ...grpc.CallOption) (*UpdateProgressRsp, error) {
	out := new(UpdateProgressRsp)
	err := c.cc.Invoke(ctx, "/task.Task/UpdateProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	UpdateProgress(context.Context, *UpdateProgressReq) (*UpdateProgressRsp, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) UpdateProgress(context.Context, *UpdateProgressReq) (*UpdateProgressRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProgress not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_UpdateProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProgressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).UpdateProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/UpdateProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).UpdateProgress(ctx, req.(*UpdateProgressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateProgress",
			Handler:    _Task_UpdateProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}