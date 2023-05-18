package rest

import (
	"context"

	cluster "github.com/infraboard/mpaas/apps/k8s"
)

func (c *ClientSet) CreateCluster(ctx context.Context, req *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.Group("clusters").
		Post("").
		Body(req).
		Do(ctx).
		Into(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *ClientSet) QueryCluster(ctx context.Context, req *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	set := cluster.NewClusterSet()

	err := c.c.Group("clusters").
		Get("").
		Body(req).
		Do(ctx).
		Into(set)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (c *ClientSet) DescribeCluster(ctx context.Context, req *cluster.DescribeClusterRequest) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.Group("clusters").
		Get(req.Id).
		Body(req).
		Do(ctx).
		Into(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
