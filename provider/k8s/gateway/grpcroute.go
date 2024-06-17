package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListGRPCRouteList(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.GRPCRouteList, error) {
	d, err := c.dynamic.Resource(c.grpcRouteResource()).Namespace("default").List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.GRPCRouteList)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetGRPCRoute(
	ctx context.Context,
	req *meta.GetRequest) (
	*gatewayv1.GRPCRoute, error) {
	d, err := c.dynamic.Resource(c.grpcRouteResource()).Namespace(req.Namespace).Get(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.GRPCRoute)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) CreateGRPCRoute(
	ctx context.Context,
	req *gatewayv1.HTTPRoute) (
	*gatewayv1.GRPCRoute, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.grpcRouteResource()).Namespace(req.Namespace).Create(ctx, us, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.GRPCRoute)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) UpdateGRPCRoute(
	ctx context.Context,
	req *gatewayv1.GRPCRoute) (
	*gatewayv1.GRPCRoute, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.grpcRouteResource()).Namespace(req.Namespace).Update(ctx, us, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.GRPCRoute)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) DeleteGRPCRoute(
	ctx context.Context,
	req *meta.DeleteRequest) error {
	return c.dynamic.Resource(c.grpcRouteResource()).Namespace(req.Namespace).Delete(ctx, req.Name, req.Opts)
}
