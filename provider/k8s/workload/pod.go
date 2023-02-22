package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
)

func (c *Workload) CreatePod(ctx context.Context, pod *v1.Pod, req *meta.CreateRequest) (*v1.Pod, error) {
	return c.corev1.Pods(req.Namespace).Create(ctx, pod, req.Opts)
}

func (c *Workload) ListPod(ctx context.Context, req *meta.ListRequest) (*v1.PodList, error) {
	return c.corev1.Pods(req.Namespace).List(ctx, req.Opts)
}

func (c *Workload) GetPod(ctx context.Context, req *meta.GetRequest) (*v1.Pod, error) {
	return c.corev1.Pods(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Workload) DeletePod(ctx context.Context, req *meta.DeleteRequest) error {
	return c.corev1.Pods("").Delete(ctx, "", req.Opts)
}

func InjectPodEnvVars(pod *v1.PodSpec, envs []v1.EnvVar) {
	// 给Init容器注入环境变量
	for i, c := range pod.InitContainers {
		InjectContainerEnvVars(&c, envs)
		// 替换掉原来的container的值
		pod.InitContainers[i] = c
	}

	// 给用户容器注入环境变量
	for i, c := range pod.Containers {
		InjectContainerEnvVars(&c, envs)
		// 替换掉原来的container的值
		pod.Containers[i] = c
	}
}
