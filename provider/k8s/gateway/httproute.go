package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListHttpRoute(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.HTTPRouteList, error) {
	d, err := c.dynamic.Resource(c.httpRouteResource()).Namespace(req.Namespace).List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.HTTPRouteList)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetHttpRoute(
	ctx context.Context,
	req *meta.GetRequest) (
	*gatewayv1.HTTPRoute, error) {
	d, err := c.dynamic.Resource(c.httpRouteResource()).Namespace(req.Namespace).Get(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.HTTPRoute)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) CreateHttpRoute(
	ctx context.Context,
	ins *gatewayv1.HTTPRoute,
	req *meta.CreateRequest) (
	*gatewayv1.HTTPRoute, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(ins)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.httpRouteResource()).Namespace(ins.Namespace).Create(ctx, us, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.HTTPRoute)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) UpdateHttpRoute(
	ctx context.Context,
	req *gatewayv1.HTTPRoute) (
	*gatewayv1.HTTPRoute, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.httpRouteResource()).Namespace(req.Namespace).Update(ctx, us, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1.HTTPRoute)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) DeleteHttpRoute(
	ctx context.Context,
	req *meta.DeleteRequest) (
	*gatewayv1.HTTPRoute, error) {
	route, err := c.GetHttpRoute(ctx, meta.NewGetRequest(req.Name).WithNamespace(req.Namespace))
	if err != nil {
		return nil, err
	}

	err = c.dynamic.Resource(c.httpRouteResource()).Namespace(req.Namespace).Delete(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	return route, nil
}
