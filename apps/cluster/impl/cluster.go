package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/cluster"
)

func (i *impl) CreateCluster(ctx context.Context, in *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	return nil, nil
}

func (i *impl) UpdateCluster(ctx context.Context, in *cluster.UpdateClusterRequest) (
	*cluster.Cluster, error) {
	return nil, nil
}

func (i *impl) DeleteCluster(ctx context.Context, in *cluster.DeleteClusterRequest) (
	*cluster.Cluster, error) {
	return nil, nil
}

// 查询集群列表
func (i *impl) QueryCluster(ctx context.Context, in *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	return nil, nil
}

// 查询集群详情
func (i *impl) DescribeCluster(ctx context.Context, in *cluster.DescribeClusterRequest) (
	*cluster.Cluster, error) {
	return nil, nil
}
