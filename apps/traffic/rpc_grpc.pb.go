// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: mpaas/apps/traffic/pb/rpc.proto

package traffic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RPC_QueryRule_FullMethodName    = "/infraboard.mpaas.traffic.RPC/QueryRule"
	RPC_DescribeRule_FullMethodName = "/infraboard.mpaas.traffic.RPC/DescribeRule"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RPC 流量管理
type RPCClient interface {
	// 查询流量规则配置列表
	QueryRule(ctx context.Context, in *QueryRuleRequest, opts ...grpc.CallOption) (*RuleSet, error)
	// 查询流量规则配置详情
	DescribeRule(ctx context.Context, in *DescribeRuleRequest, opts ...grpc.CallOption) (*Rule, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) QueryRule(ctx context.Context, in *QueryRuleRequest, opts ...grpc.CallOption) (*RuleSet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RuleSet)
	err := c.cc.Invoke(ctx, RPC_QueryRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeRule(ctx context.Context, in *DescribeRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Rule)
	err := c.cc.Invoke(ctx, RPC_DescribeRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility.
//
// RPC 流量管理
type RPCServer interface {
	// 查询流量规则配置列表
	QueryRule(context.Context, *QueryRuleRequest) (*RuleSet, error)
	// 查询流量规则配置详情
	DescribeRule(context.Context, *DescribeRuleRequest) (*Rule, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCServer struct{}

func (UnimplementedRPCServer) QueryRule(context.Context, *QueryRuleRequest) (*RuleSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRule not implemented")
}
func (UnimplementedRPCServer) DescribeRule(context.Context, *DescribeRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRule not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}
func (UnimplementedRPCServer) testEmbeddedByValue()             {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	// If the following call pancis, it indicates UnimplementedRPCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_QueryRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryRule(ctx, req.(*QueryRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeRule(ctx, req.(*DescribeRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mpaas.traffic.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryRule",
			Handler:    _RPC_QueryRule_Handler,
		},
		{
			MethodName: "DescribeRule",
			Handler:    _RPC_DescribeRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mpaas/apps/traffic/pb/rpc.proto",
}
