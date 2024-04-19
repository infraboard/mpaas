package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/common/yaml"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/provider/k8s/network"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

func NewSyncK8sDeployRequest() *SyncK8sDeployRequest {
	return &SyncK8sDeployRequest{
		SyncDeployment: true,
		SyncPods:       true,
		SyncService:    false,
	}
}

type SyncK8sDeployRequest struct {
	SyncDeployment bool
	SyncPods       bool
	SyncService    bool
}

func (i *impl) SyncK8sDeploy(ctx context.Context, req *SyncK8sDeployRequest, ins *deploy.Deployment) error {
	wc := ins.Spec.K8STypeConfig

	wl, err := wc.GetWorkLoad()
	if err != nil {
		return err
	}
	wl.SetDefaultNamespace(ins.Spec.Namespace)
	// 设置Match Lable
	wl.SetMatchLabel(deploy.LABEL_DEPLOY_ID_KEY, ins.Meta.Id)

	// 查询部署的k8s集群
	k8sClient, err := i.GetDeployK8sClient(ctx, wc.ClusterId)
	if err != nil {
		return err
	}

	// 运行工作负载
	if req.SyncDeployment {
		ins.Status.MarkCreating()
		dReq := meta.NewGetRequest(ins.Spec.Name).WithNamespace(ins.Spec.Namespace)
		err = k8sClient.WorkLoad().Get(ctx, wl, dReq)
		if err != nil {
			return err
		}
		wc.WorkloadConfig = wl.MustToYaml()
	}

	// pod同步
	if req.SyncPods {
		pReq := meta.NewListRequest().
			WithNamespace(ins.Spec.Namespace).
			WithLabelSelector(meta.NewLabelSelector().Add(deploy.LABEL_DEPLOY_ID_KEY, ins.Meta.Id))
		pods, err := k8sClient.WorkLoad().ListPod(ctx, pReq)
		if err != nil {
			return err
		}
		newPods := map[string]string{}
		for i := range pods.Items {
			p := pods.Items[i]
			newPods[p.Name] = yaml.MustToYaml(p)
		}
		wc.Pods = newPods
	}

	// 同步服务
	if wc.Service != "" && req.SyncService {
		svc, err := network.ParseServiceFromYaml(wc.Service)
		if err != nil {
			return err
		}
		svcReq := meta.NewGetRequest(svc.Name).WithNamespace(svc.Namespace)
		service, err := k8sClient.Network().GetService(ctx, svcReq)
		if err != nil {
			return err
		}
		wc.Service = yaml.MustToYaml(service)
	}

	return nil
}

func (i *impl) UpdateK8sDeploy(ctx context.Context, ins *deploy.Deployment) error {
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

	// 查询部署的k8s集群
	k8sClient, err := i.GetDeployK8sClient(ctx, wc.ClusterId)
	if err != nil {
		return err
	}

	// 运行工作负载
	ins.Status.MarkCreating()
	wl, err = k8sClient.WorkLoad().Update(ctx, wl)
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
		service, err := k8sClient.Network().UpdateService(ctx, svc)
		if err != nil {
			return err
		}
		wc.Service = yaml.MustToYaml(service)
	}

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
	}

	return nil
}
