package rest

import (
	"context"

	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
)

func (c *ClientSet) CreateDeployment(ctx context.Context, req *cluster.CreateClusterRequest) (
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
