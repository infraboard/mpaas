package impl

import (
	"context"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s"
)

func (s *service) CreateCluster(ctx context.Context, req *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := cluster.NewCluster(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create cluster error, %s", err)
	}

	// 连接集群检查状态
	s.checkStatus(ins)
	if err := ins.IsAlive(); err != nil {
		return nil, err
	}

	// 加密
	err = ins.EncryptKubeConf(s.encryptoKey)
	if err != nil {
		return nil, err
	}

	if err := s.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) checkStatus(ins *cluster.Cluster) {
	client, err := k8s.NewClient(ins.Data.KubeConfig)
	if err != nil {
		ins.Status.Message = err.Error()
		return
	}

	if ctx := client.CurrentContext(); ctx != nil {
		ins.Id = ctx.Cluster
		ins.ServerInfo.AuthUser = ctx.AuthInfo
	}

	if k := client.CurrentCluster(); k != nil {
		ins.ServerInfo.Server = k.Server
	}

	ins.Status.CheckAt = time.Now().UnixMilli()
	v, err := client.ServerVersion()
	if err != nil {
		ins.Status.IsAlive = false
		ins.Status.Message = err.Error()
	} else {
		ins.Status.IsAlive = true
		ins.ServerInfo.Version = v
	}
}

func (s *service) DescribeCluster(ctx context.Context, req *cluster.DescribeClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := s.get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if err := ins.DecryptKubeConf(s.encryptoKey); err != nil {
		return nil, err
	}
	return ins, nil
}

func (s *service) QueryCluster(ctx context.Context, req *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	query := newQueryclusterRequest(req)
	set, err := s.query(ctx, query)
	if err != nil {
		return nil, err
	}
	if err := set.DecryptKubeConf(s.encryptoKey); err != nil {
		return nil, err
	}
	return set, nil
}

func (s *service) UpdateCluster(ctx context.Context, req *cluster.UpdateClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := s.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(req.Id))
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

	// 加密
	err = ins.EncryptKubeConf(s.encryptoKey)
	if err != nil {
		return nil, err
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DeleteCluster(ctx context.Context, req *cluster.DeleteClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := s.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.deletecluster(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
