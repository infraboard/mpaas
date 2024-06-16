package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func (c *Client) ListReferenceGrant(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1beta1.ReferenceGrant, error) {
	d, err := c.dynamic.Resource(c.referenceGrantResource()).Namespace("default").List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1beta1.ReferenceGrant)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
