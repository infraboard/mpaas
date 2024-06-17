package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListGateway(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.GatewayList, error) {
	d, err := c.dynamic.Resource(c.gatewayResource()).Namespace(req.Namespace).List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.GatewayList)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetGateway(
	ctx context.Context,
	req *meta.GetRequest) (
	*gatewayv1.Gateway, error) {
	d, err := c.dynamic.Resource(c.gatewayResource()).Namespace(req.Namespace).Get(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.Gateway)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) CreateGateway(
	ctx context.Context,
	req *gatewayv1.Gateway) (
	*gatewayv1.Gateway, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.gatewayResource()).Namespace(req.Namespace).Create(ctx, us, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.Gateway)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) UpdateGateway(
	ctx context.Context,
	req *gatewayv1.Gateway) (
	*gatewayv1.Gateway, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.gatewayResource()).Namespace(req.Namespace).Update(ctx, us, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.Gateway)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) DeleteGateway(
	ctx context.Context,
	req *meta.DeleteRequest) error {
	return c.dynamic.Resource(c.gatewayResource()).Namespace(req.Namespace).Delete(ctx, req.Name, req.Opts)
}
