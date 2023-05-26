package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/cluster"
)

// 查询集群列表
func (i *impl) QueryCluster(ctx context.Context, in *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find cluster error, error is %s", err)
	}

	set := cluster.NewClusterSet()
	// 循环
	for resp.Next(ctx) {
		ins := cluster.NewDefaultCluster()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode cluster error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get cluster count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (i *impl) CreateCluster(ctx context.Context, in *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := cluster.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a cluster document error, %s", err)
	}
	return ins, nil
}

// 查询集群详情
func (i *impl) DescribeCluster(ctx context.Context, in *cluster.DescribeClusterRequest) (
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
