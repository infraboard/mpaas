package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListGatewayClass(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.GatewayClass, error) {
	d, err := c.dynamic.Resource(c.gatewayClassResource()).List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.GatewayClass)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetGatewayClass(
	ctx context.Context,
	req *meta.GetRequest) (
	*gatewayv1.GatewayClass, error) {
	d, err := c.dynamic.Resource(c.gatewayClassResource()).Namespace(req.Namespace).Get(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.GatewayClass)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) CreateGatewayClass(
	ctx context.Context,
	ins *gatewayv1.GatewayClass,
	req *meta.CreateRequest,
) (
	*gatewayv1.GatewayClass, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(ins)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.gatewayClassResource()).Namespace(ins.Namespace).Create(ctx, us, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.GatewayClass)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) UpdateGatewayClass(
	ctx context.Context,
	req *gatewayv1.GatewayClass) (
	*gatewayv1.GatewayClass, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.gatewayClassResource()).Namespace(req.Namespace).Update(ctx, us, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.GatewayClass)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) DeleteGatewayClass(
	ctx context.Context,
	req *meta.DeleteRequest) (*gatewayv1.GatewayClass, error) {
	gc, err := c.GetGatewayClass(ctx, meta.NewGetRequest(req.Name).WithNamespace(req.Namespace))
	if err != nil {
		return nil, err
	}

	err = c.dynamic.Resource(c.gatewayClassResource()).Namespace(req.Namespace).Delete(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	return gc, nil
}
