package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func (c *Client) ListReferenceGrant(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1beta1.ReferenceGrantList, error) {
	d, err := c.dynamic.Resource(c.referenceGrantResource()).Namespace("default").List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1beta1.ReferenceGrantList)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetReferenceGrant(
	ctx context.Context,
	req *meta.GetRequest) (
	*gatewayv1beta1.ReferenceGrant, error) {
	d, err := c.dynamic.Resource(c.referenceGrantResource()).Namespace(req.Namespace).Get(ctx, req.Name, req.Opts)
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1beta1.ReferenceGrant)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) CreateReferenceGrant(
	ctx context.Context,
	req *gatewayv1beta1.ReferenceGrant) (
	*gatewayv1beta1.ReferenceGrant, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.referenceGrantResource()).Namespace(req.Namespace).Create(ctx, us, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1beta1.ReferenceGrant)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) UpdateReferenceGrant(
	ctx context.Context,
	req *gatewayv1beta1.ReferenceGrant) (
	*gatewayv1beta1.ReferenceGrant, error) {
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(req)
	if err != nil {
		return nil, err
	}
	us := &unstructured.Unstructured{Object: m}
	us, err = c.dynamic.Resource(c.referenceGrantResource()).Namespace(req.Namespace).Update(ctx, us, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	obj := new(gatewayv1beta1.ReferenceGrant)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(us.UnstructuredContent(), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) DeleteReferenceGrant(
	ctx context.Context,
	req *meta.DeleteRequest) error {
	return c.dynamic.Resource(c.referenceGrantResource()).Namespace(req.Namespace).Delete(ctx, req.Name, req.Opts)
}
