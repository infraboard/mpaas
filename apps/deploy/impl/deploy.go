package impl

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	deploy_cluster "github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/common/yaml"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/provider/k8s/network"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *impl) CreateDeployment(ctx context.Context, in *deploy.CreateDeploymentRequest) (
	*deploy.Deployment, error) {
	// 检查Cluster是否存在
	c, err := i.cluster.DescribeCluster(ctx, deploy_cluster.NewDescribeClusterRequest(in.Cluster))
	if err != nil {
		return nil, err
	}
	in.ServiceId = c.Spec.ServiceId

	err = i.validate(ctx, c.Spec.Kind, in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := deploy.New(in)
	ins.Spec.SetDefault()

	// 因为有WebHook 需要提前保存好集群信息
	ins.Meta.Id = ins.Spec.UUID()
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a deploy document error, %s", err)
	}

	switch in.Type {
	case deploy.TYPE_KUBERNETES:
		// 如果创建成功, 等待回调更新状态, 如果失败则直接更新状态
		err := i.RunK8sDeploy(ctx, ins)
		if err != nil {
			ins.Status.MarkFailed(err)
		}
		i.update(ctx, ins)
	}

	return ins, nil
}

func (i *impl) validate(ctx context.Context, kind deploy_cluster.KIND, in *deploy.CreateDeploymentRequest) error {
	if err := in.Validate(); err != nil {
		return err
	}

	wc := in.K8STypeConfig
	wl, err := wc.GetWorkLoad()
	if err != nil {
		return err
	}

	// 补充服务相关信息
	switch kind {
	case deploy_cluster.KIND_WORKLOAD:
		err := in.ValidateWorkLoad()
		if err != nil {
			return exception.NewBadRequest(err.Error())
		}

		svc, err := i.mcenter.Service().DescribeService(ctx, service.NewDescribeServiceRequest(in.ServiceId))
		if err != nil {
			return err
		}
		in.ServiceName = svc.Spec.Name
		in.Domain = svc.Spec.Namespace
		in.Namespace = svc.Spec.Namespace
	case deploy_cluster.KIND_MIDDLEWARE:
		err := in.ValidateMiddleware()
		if err != nil {
			return err
		}
	}

	// 获取部署名称
	in.Name = wl.GetObjectMeta().Name
	// 检查主容器是否存在
	serviceContainer := wl.GetMainContainer()
	if serviceContainer == nil {
		return fmt.Errorf("无主容器[第一个容器]配置")
	}
	// 检查主容器名称
	if serviceContainer.Name != in.ServiceName {
		return fmt.Errorf("主容器[第一个容器]名称(%s)与服务名称(%s)不相同, 请修改主容器名称为服务名称",
			serviceContainer.Name,
			in.ServiceName,
		)
	}

	// 从镜像中获取部署的版本信息
	in.ServiceVersion = wl.GetMainContainerVersion()
	return nil
}

func (i *impl) RunK8sDeploy(ctx context.Context, ins *deploy.Deployment) error {
	wc := ins.Spec.K8STypeConfig

	wl, err := wc.GetWorkLoad()
	if err != nil {
		return err
	}
	wl.SetDefaultNamespace(ins.Spec.Namespace)

	// 补充Pod需要注入的信息
	pts := wl.GetPodTemplateSpec()
	workload.InjectPodTemplateSpecAnnotations(pts, deploy.ANNOTATION_DEPLOY_ID, ins.Meta.Id)
	wl.SetAnnotations(deploy.ANNOTATION_DEPLOY_ID, ins.Meta.Id)
	for k, v := range ins.InjectPodLabel() {
		workload.InjectPodTemplateSpecLabel(pts, k, v)
	}
	// 设置Match Lable
	wl.SetMatchLabel(deploy.LABEL_DEPLOY_ID_KEY, ins.Meta.Id)

	// 补充需要注入的系统变量
	// ins.DynamicInjection.EnvGroups[0].ToContainerEnvVars()

	// 查询部署的k8s集群
	k8sClient, err := i.GetDeployK8sClient(ctx, wc.ClusterId)
	if err != nil {
		return err
	}

	// 运行工作负载
	ins.Status.MarkCreating()
	wl, err = k8sClient.WorkLoad().Run(ctx, wl)
	if err != nil {
		return err
	}
	wc.WorkloadConfig = wl.MustToYaml()
	// 创建服务
	if wc.Service != "" {
		svc, err := network.ParseServiceFromYaml(wc.Service)
		if err != nil {
			return err
		}
		svc.Namespace = ins.Spec.Namespace
		svc.Annotations[deploy.ANNOTATION_DEPLOY_ID] = ins.Meta.Id
		service, err := k8sClient.Network().CreateService(ctx, svc)
		if err != nil {
			return err
		}
		wc.Service = yaml.MustToYaml(service)
		ins.Spec.AccessInfo.Private = deploy.NewAccessAddressFromK8sService(service)
	}

	return nil
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
	case deploy.DESCRIBE_BY_NAME:
		filter["domain"] = req.Domain
		filter["namespace"] = req.Namespace
		filter["name"] = req.DescribeValue
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

	if err := i.update(ctx, d); err != nil {
		return nil, err
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
	c, err := i.k8s.DescribeCluster(ctx, descReq)
	if err != nil {
		return nil, err
	}
	return c.Client()
}

func (i *impl) UpdateDeploymentStatus(ctx context.Context, in *deploy.UpdateDeploymentStatusRequest) (
	*deploy.Deployment, error) {
	req := deploy.NewDescribeDeploymentRequest(in.Id)
	ins, err := i.DescribeDeployment(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := ins.ValidateToken(in.UpdateToken); err != nil {
		return nil, err
	}

	switch ins.Spec.Type {
	case deploy.TYPE_KUBERNETES:
		err := i.UpdateK8sDeployStatus(ctx, ins, in.UpdatedK8SConfig)
		if err != nil {
			return nil, err
		}
	}

	// 更新
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": ins.Meta.Id}, bson.M{"$set": ins})
	if err != nil {
		return nil, exception.NewInternalServerError("update deploy status(%s) error, %s", ins.Meta.Id, err)
	}

	return ins, nil
}

// k8s类型的服务
func (i *impl) UpdateK8sDeployStatus(ctx context.Context, ins *deploy.Deployment, in *deploy.K8STypeConfig) error {
	if in == nil {
		return fmt.Errorf("k8s config 不能为nil")
	}

	ins.SetDefault()

	wc := ins.Spec.K8STypeConfig
	err := wc.Merge(in)
	if err != nil {
		return err
	}

	// 根据workload信息 补充更新 部署的版本和状态
	if in.WorkloadConfig != "" {
		wl, err := in.GetWorkLoad()
		if err != nil {
			return err
		}
		// 从镜像中获取部署的版本信息
		ins.Spec.ServiceVersion = wl.GetMainContainerVersion()
		// 更新部署状态
		ins.Status.UpdateK8sWorkloadStatus(wl.Status())
	}

	return nil
}

// 查询部署是需要动态注入的环境变量, 通过该接口拉取Env进行动态注入
func (i *impl) QueryDeploymentInjectEnv(ctx context.Context, in *deploy.QueryDeploymentInjectEnvRequest) (
	*deploy.InjectionEnvGroupSet, error) {
	req := deploy.NewDescribeDeploymentRequest(in.Id)
	ins, err := i.DescribeDeployment(ctx, req)
	if err != nil {
		return nil, err
	}

	return i.queryDeploymentInjectEnv(ctx, ins)
}

func (i *impl) queryDeploymentInjectEnv(ctx context.Context, ins *deploy.Deployment) (
	*deploy.InjectionEnvGroupSet, error) {
	set := deploy.NewInjectionEnvGroupSet()
	if ins.DynamicInjection == nil {
		return set, nil
	}

	// 注入已经启用的组变量
	ins.DynamicInjection.AddEnabledGroupTo(set)

	// 注入部署相关系统变量
	systemGroup := ins.SystemInjectionEnvGroup()
	if ins.DynamicInjection.SystemEnv {
		set.Add(systemGroup)
	}

	// 注入服务相关信息
	if ins.Spec.ServiceId != "" {
		app, err := i.mcenter.Service().DescribeService(ctx, service.NewDescribeServiceRequest(ins.Spec.ServiceId))
		if err != nil {
			return nil, err
		}
		systemGroup.AddEnv(
			deploy.NewInjectionEnv("MCENTER_CLINET_ID", app.Credential.ClientId),
			deploy.NewInjectionEnv("MCENTER_CLIENT_SECRET", app.Credential.ClientSecret),
		)

		// 使用服务的加密key对需要加密的环境变量加密
		encryptKey := app.Security.EncryptKey
		set.Encrypt(encryptKey)
	}
	return set, nil
}
