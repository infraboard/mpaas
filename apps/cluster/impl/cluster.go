package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"

	"github.com/infraboard/mpaas/apps/cluster"
)

func (s *service) CreateCluster(ctx context.Context, req *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := cluster.NewCluster(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create cluster error, %s", err)
	}

	if err := s.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) Describecluster(ctx context.Context, req *cluster.DescribeClusterRequest) (
	*cluster.Cluster, error) {
	return s.get(ctx, req.Id)
}

func (s *service) Querycluster(ctx context.Context, req *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	query := newQueryclusterRequest(req)
	return s.query(ctx, query)
}

func (s *service) Updatecluster(ctx context.Context, req *cluster.UpdateClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := s.Describecluster(ctx, cluster.NewDescribeClusterRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Data.Validate(); err != nil {
		return nil, err
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) Deletecluster(ctx context.Context, req *cluster.DeleteClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := s.Describecluster(ctx, cluster.NewDescribeClusterRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.deletecluster(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
