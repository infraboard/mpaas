package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/common/yaml"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *impl) CreateDeployment(ctx context.Context, in *deploy.CreateDeploymentRequest) (
	*deploy.Deployment, error) {
	// 补充服务相关信息
	ins, err := deploy.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 查询服务
	svc, err := i.mcenter.Service().DescribeService(ctx, service.NewDescribeServiceRequest(in.ServiceId))
	if err != nil {
		return nil, err
	}
	ins.Spec.ServiceName = svc.Spec.Name
	ins.Scope.Domain = svc.Spec.Namespace
	ins.Scope.Namespace = svc.Spec.Namespace

	// 如果服务是k8s服务则直接执行部署
	switch in.Type {
	case deploy.TYPE_KUBERNETES:
		wc := ins.Spec.K8STypeConfig

		wl, err := wc.GetWorkLoad()
		if err != nil {
			return nil, err
		}
		wl.SetDefaultNamespace(ins.Scope.Namespace)
		wl.SetAnnotations(deploy.ANNOTATION_DEPLOY_ID, ins.Meta.Id)

		// 检查主容器是否存在
		serviceContainer := wl.GetServiceContainer(ins.Spec.ServiceName)
		if serviceContainer == nil {
			return nil, fmt.Errorf("部署配置必须包含一个服务名称同名的容器 作为主容器")
		}
		// 从镜像中获取部署的版本信息
		ins.Spec.ServiceVersion = wl.GetServiceContainerVersion(ins.Spec.ServiceName)

		// 查询部署的k8s集群
		k8sClient, err := i.GetDeployK8sClient(ctx, wc.ClusterId)
		if err != nil {
			return nil, err
		}
		// 运行工作负载
		ins.Status.MarkCreating()
		wl, err = k8sClient.WorkLoad().Run(ctx, wl)
		if err != nil {
			return nil, err
		}
		wc.WorkloadConfig = wl.MustToYaml()
		// 创建服务
		if wc.Service != "" {
			service, err := k8sClient.Network().Run(ctx, wc.Service)
			if err != nil {
				return nil, err
			}
			wc.Service = yaml.MustToYaml(service)
		}
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a deploy document error, %s", err)
	}
	return ins, nil
}

func (i *impl) QueryDeployment(ctx context.Context, in *deploy.QueryDeploymentRequest) (
	*deploy.DeploymentSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find deploy error, error is %s", err)
	}

	set := deploy.NewDeploymentSet()
	// 循环
	for resp.Next(ctx) {
		ins := deploy.NewDefaultDeploy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode deploy error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (i *impl) DescribeDeployment(ctx context.Context, req *deploy.DescribeDeploymentRequest) (*deploy.Deployment, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case deploy.DESCRIBE_BY_ID:
		filter["_id"] = req.DescribeValue
	}

	d := deploy.NewDefaultDeploy()
	if err := i.col.FindOne(ctx, filter).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("deploy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find deploy %s error, %s", req.DescribeValue, err)
	}

	return d, nil
}

func (i *impl) UpdateDeployment(ctx context.Context, in *deploy.UpdateDeploymentRequest) (
	*deploy.Deployment, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := deploy.NewDescribeDeploymentRequest(in.Id)
	d, err := i.DescribeDeployment(ctx, req)
	if err != nil {
		return nil, err
	}

	switch in.UpdateMode {
	case request.UpdateMode_PUT:
		d.Spec = in.Spec
	case request.UpdateMode_PATCH:
		if err := mergo.MergeWithOverwrite(d.Spec, in.Spec); err != nil {
			return nil, err
		}
		if err := d.Spec.Validate(); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", in.UpdateMode)
	}

	d.Meta.UpdateAt = time.Now().Unix()
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": d.Meta.Id}, bson.M{"$set": d})
	if err != nil {
		return nil, exception.NewInternalServerError("update deploy(%s) error, %s", d.Meta.Id, err)
	}

	return d, nil
}

func (i *impl) DeleteDeployment(ctx context.Context, in *deploy.DeleteDeploymentRequest) (
	*deploy.Deployment, error) {
	req := deploy.NewDescribeDeploymentRequest(in.Id)
	ins, err := i.DescribeDeployment(ctx, req)
	if err != nil {
		return nil, err
	}

	cid := ins.GetK8sClusterId()
	if cid != "" {
		kc := ins.Spec.K8STypeConfig
		wl, err := kc.GetWorkLoad()
		if err != nil {
			return nil, err
		}

		k8sClient, err := i.GetDeployK8sClient(ctx, cid)
		if err != nil {
			return nil, err
		}
		// 删除工作负载
		if wl != nil {
			_, err = k8sClient.WorkLoad().Delete(ctx, wl)
			if err != nil {
				return nil, err
			}
		}

		// 删除服务
		svc, err := kc.GetServiceObj()
		if err != nil {
			return nil, err
		}
		if svc != nil {
			delReq := meta.NewDeleteRequest(svc.Name).WithNamespace(svc.Namespace)
			err = k8sClient.Network().DeleteService(ctx, delReq)
			if err != nil {
				return nil, err
			}
		}
	}

	_, err = i.col.DeleteOne(ctx, bson.M{"_id": ins.Meta.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete deploy(%s) error, %s", ins.Meta.Id, err)
	}
	return ins, nil
}

func (i *impl) GetDeployK8sClient(ctx context.Context, k8sClusterId string) (*k8s.Client, error) {
	descReq := cluster.NewDescribeClusterRequest(k8sClusterId)
	c, err := i.cluster.DescribeCluster(ctx, descReq)
	if err != nil {
		return nil, err
	}
	return c.Client()
}

func (i *impl) UpdateDeploymentStatus(ctx context.Context, in *deploy.UpdateDeploymentStatusRequest) (
	*deploy.Deployment, error) {
	return nil, nil
}
