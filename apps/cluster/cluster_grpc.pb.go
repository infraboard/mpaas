// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: apps/cluster/pb/cluster.proto

package cluster

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	QueryCluster(ctx context.Context, in *QueryClusterRequest, opts ...grpc.CallOption) (*ClusterSet, error)
	DescribeCluster(ctx context.Context, in *DescribeClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.cluster.Service/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryCluster(ctx context.Context, in *QueryClusterRequest, opts ...grpc.CallOption) (*ClusterSet, error) {
	out := new(ClusterSet)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.cluster.Service/QueryCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeCluster(ctx context.Context, in *DescribeClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.cluster.Service/DescribeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.cluster.Service/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.cluster.Service/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	QueryCluster(context.Context, *QueryClusterRequest) (*ClusterSet, error)
	DescribeCluster(context.Context, *DescribeClusterRequest) (*Cluster, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*Cluster, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedServiceServer) QueryCluster(context.Context, *QueryClusterRequest) (*ClusterSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCluster not implemented")
}
func (UnimplementedServiceServer) DescribeCluster(context.Context, *DescribeClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCluster not implemented")
}
func (UnimplementedServiceServer) UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (UnimplementedServiceServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.cluster.Service/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.cluster.Service/QueryCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryCluster(ctx, req.(*QueryClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.cluster.Service/DescribeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeCluster(ctx, req.(*DescribeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.cluster.Service/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.cluster.Service/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mpaas.cluster.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _Service_CreateCluster_Handler,
		},
		{
			MethodName: "QueryCluster",
			Handler:    _Service_QueryCluster_Handler,
		},
		{
			MethodName: "DescribeCluster",
			Handler:    _Service_DescribeCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _Service_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _Service_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/cluster/pb/cluster.proto",
}
