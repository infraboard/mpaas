package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	k8s_cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/provider/k8s"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 查询集群列表
func (i *impl) QueryCluster(ctx context.Context, in *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	r := newQueryRequest(in)
	i.log.Debugf("filter: %s", r.FindFilter())
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

	// 查询集群关联的部署
	if in.WithDeployment && set.Len() > 0 {
		dquery := deploy.NewQueryDeploymentRequest()
		dquery.Clusters = set.ClusterIds()
		ds, err := i.deploy.QueryDeployment(ctx, dquery)
		if err != nil {
			return nil, err
		}
		set.UpdateDeploymens(ds)
	}

	// 查询关联的服务
	if in.WithService && set.Len() > 0 {
		squery := service.NewQueryServiceRequest()
		squery.Ids = set.ServiceIds()
		services, err := i.mcenter.Service().QueryService(ctx, squery)
		if err != nil {
			return nil, err
		}
		set.ForEatch(func(item *cluster.Cluster) {
			item.Service = services.GetServiceById(item.Spec.ServiceId)
		})
	}

	return set, nil
}

func (i *impl) CreateCluster(ctx context.Context, in *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins, err := cluster.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 获取Service信息
	switch in.Kind {
	case cluster.KIND_WORKLOAD:
		svc, err := i.mcenter.Service().DescribeService(
			ctx,
			service.NewDescribeServiceRequest(in.ServiceId),
		)
		if err != nil {
			return nil, err
		}
		ins.Scope.Domain = svc.Spec.Domain
		ins.Scope.Namespace = svc.Spec.Namespace
	default:
		tk := token.GetTokenFromCtx(ctx)
		if tk != nil {
			ins.Scope = tk.GenScope()
		}
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a cluster document error, %s", err)
	}
	return ins, nil
}

func (i *impl) GetDeployK8sClient(ctx context.Context, k8sClusterId string) (*k8s.Client, error) {
	descReq := k8s_cluster.NewDescribeClusterRequest(k8sClusterId)
	c, err := i.k8s.DescribeCluster(ctx, descReq)
	if err != nil {
		return nil, err
	}
	return c.Client()
}

// 查询集群详情
func (i *impl) DescribeCluster(ctx context.Context, in *cluster.DescribeClusterRequest) (
	*cluster.Cluster, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	ins := cluster.NewDefaultCluster()
	if err := i.col.FindOne(ctx, bson.M{"_id": in.Id}).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("cluster %s not found", in.Id)
		}

		return nil, exception.NewInternalServerError("find cluster %s error, %s", in.Id, err)
	}
	return ins, nil
}

func (i *impl) UpdateCluster(ctx context.Context, in *cluster.UpdateClusterRequest) (
	*cluster.Cluster, error) {
	return nil, nil
}

func (i *impl) DeleteCluster(ctx context.Context, in *cluster.DeleteClusterRequest) (
	*cluster.Cluster, error) {
	req := cluster.NewDescribeClusterRequest(in.Id)
	ins, err := i.DescribeCluster(ctx, req)
	if err != nil {
		return nil, err
	}

	_, err = i.col.DeleteOne(ctx, bson.M{"_id": ins.Meta.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete cluster(%s) error, %s", in.Id, err)
	}
	return ins, nil
}
