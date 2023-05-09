// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: mpaas/apps/pipeline/pb/rpc.proto

package pipeline

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
	RPC_QueryPipeline_FullMethodName    = "/infraboard.mpaas.pipeline.RPC/QueryPipeline"
	RPC_DescribePipeline_FullMethodName = "/infraboard.mpaas.pipeline.RPC/DescribePipeline"
	RPC_CreatePipeline_FullMethodName   = "/infraboard.mpaas.pipeline.RPC/CreatePipeline"
	RPC_UpdatePipeline_FullMethodName   = "/infraboard.mpaas.pipeline.RPC/UpdatePipeline"
	RPC_DeletePipeline_FullMethodName   = "/infraboard.mpaas.pipeline.RPC/DeletePipeline"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 查询Pipeline列表
	QueryPipeline(ctx context.Context, in *QueryPipelineRequest, opts ...grpc.CallOption) (*PipelineSet, error)
	// 查询Pipeline详情
	DescribePipeline(ctx context.Context, in *DescribePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	// 创建Pipeline
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	// 更新Pipeline
	UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	// 删除Pipeline
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) QueryPipeline(ctx context.Context, in *QueryPipelineRequest, opts ...grpc.CallOption) (*PipelineSet, error) {
	out := new(PipelineSet)
	err := c.cc.Invoke(ctx, RPC_QueryPipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribePipeline(ctx context.Context, in *DescribePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, RPC_DescribePipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, RPC_CreatePipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, RPC_UpdatePipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, RPC_DeletePipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 查询Pipeline列表
	QueryPipeline(context.Context, *QueryPipelineRequest) (*PipelineSet, error)
	// 查询Pipeline详情
	DescribePipeline(context.Context, *DescribePipelineRequest) (*Pipeline, error)
	// 创建Pipeline
	CreatePipeline(context.Context, *CreatePipelineRequest) (*Pipeline, error)
	// 更新Pipeline
	UpdatePipeline(context.Context, *UpdatePipelineRequest) (*Pipeline, error)
	// 删除Pipeline
	DeletePipeline(context.Context, *DeletePipelineRequest) (*Pipeline, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) QueryPipeline(context.Context, *QueryPipelineRequest) (*PipelineSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPipeline not implemented")
}
func (UnimplementedRPCServer) DescribePipeline(context.Context, *DescribePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePipeline not implemented")
}
func (UnimplementedRPCServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedRPCServer) UpdatePipeline(context.Context, *UpdatePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePipeline not implemented")
}
func (UnimplementedRPCServer) DeletePipeline(context.Context, *DeletePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_QueryPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryPipeline(ctx, req.(*QueryPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribePipeline(ctx, req.(*DescribePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreatePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_UpdatePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdatePipeline(ctx, req.(*UpdatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeletePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mpaas.pipeline.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryPipeline",
			Handler:    _RPC_QueryPipeline_Handler,
		},
		{
			MethodName: "DescribePipeline",
			Handler:    _RPC_DescribePipeline_Handler,
		},
		{
			MethodName: "CreatePipeline",
			Handler:    _RPC_CreatePipeline_Handler,
		},
		{
			MethodName: "UpdatePipeline",
			Handler:    _RPC_UpdatePipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _RPC_DeletePipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mpaas/apps/pipeline/pb/rpc.proto",
}
