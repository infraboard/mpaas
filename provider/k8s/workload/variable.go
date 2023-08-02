package workload

import (
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewSystemVaraible() *SystemVaraible {
	return &SystemVaraible{}
}

type SystemVaraible struct {
	WorkloadName string `json:"workload_name"`
	Image        string `json:"image"`
}

func (v *SystemVaraible) ImageDetail() (addr, version string) {
	if v.Image == "" {
		return
	}
	av := strings.Split(v.Image, ":")
	addr = av[0]
	if len(av) > 1 {
		version = av[1]
	}
	return
}

func (w *WorkLoad) SystemVaraible(serviceName string) *SystemVaraible {
	m := NewSystemVaraible()

	meta := w.GetObjectMeta()
	m.WorkloadName = meta.Name

	container := w.GetServiceContainer(serviceName)
	if container != nil {
		m.Image = container.Image
	}
	return m
}

// 获取主容器的镜像版本
func (w *WorkLoad) GetMainContainerVersion() string {
	c := w.GetMainContainer()
	if c != nil && c.Image != "" {
		image := strings.Split(c.Image, ":")
		count := len(image)
		if count > 1 {
			return image[count-1]
		}
	}
	return ""
}

func (w *WorkLoad) GetServiceContainer(serviceName string) *v1.Container {
	var container *v1.Container
	switch w.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		container = GetContainerFromPodTemplate(w.Deployment.Spec.Template, serviceName)
	case WORKLOAD_KIND_STATEFULSET:
		container = GetContainerFromPodTemplate(w.StatefulSet.Spec.Template, serviceName)
	case WORKLOAD_KIND_DAEMONSET:
		container = GetContainerFromPodTemplate(w.DaemonSet.Spec.Template, serviceName)
	case WORKLOAD_KIND_CRONJOB:
		container = GetContainerFromPodTemplate(w.CronJob.Spec.JobTemplate.Spec.Template, serviceName)
	case WORKLOAD_KIND_JOB:
		container = GetContainerFromPodTemplate(w.Job.Spec.Template, serviceName)
	}
	return container
}

// 约定第一个容器为主容器, 其他容器为辅助容器
func (w *WorkLoad) GetMainContainer() *v1.Container {
	spec := w.GetPodTemplateSpec()
	if len(spec.Spec.Containers) == 0 {
		return nil
	}
	return &spec.Spec.Containers[0]
}

func (w *WorkLoad) GetObjectMeta() *metav1.ObjectMeta {
	switch w.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		return &w.Deployment.ObjectMeta
	case WORKLOAD_KIND_STATEFULSET:
		return &w.StatefulSet.ObjectMeta
	case WORKLOAD_KIND_DAEMONSET:
		return &w.DaemonSet.ObjectMeta
	case WORKLOAD_KIND_CRONJOB:
		return &w.CronJob.ObjectMeta
	case WORKLOAD_KIND_JOB:
		return &w.Job.ObjectMeta
	}
	return nil
}
