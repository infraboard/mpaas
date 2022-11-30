package rest

import (
	"context"

	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/cluster"
	appsv1 "k8s.io/api/apps/v1"
)

func (c *ClientSet) CreateDeployment(ctx context.Context, req *appsv1.Deployment) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.Group("clusters").
		Post("").
		Body(req).
		Do(ctx).
		Into(response.NewData(ins))
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *ClientSet) QueryDeployment(ctx context.Context, req *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.
		Post("clusters").
		Body(req).
		Do(ctx).
		Into(response.NewData(ins))
	if err != nil {
		return nil, err
	}

	return ins, nil
}
