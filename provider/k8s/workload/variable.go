package workload

import (
	"strings"

	v1 "k8s.io/api/core/v1"
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

	var container v1.Container
	switch w.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		m.WorkloadName = w.Deployment.Name
		container = GetContainerFromPodTemplate(w.Deployment.Spec.Template, serviceName)
	case WORKLOAD_KIND_STATEFULSET:
		m.WorkloadName = w.StatefulSet.Name
		container = GetContainerFromPodTemplate(w.StatefulSet.Spec.Template, serviceName)
	case WORKLOAD_KIND_DAEMONSET:
		m.WorkloadName = w.DaemonSet.Name
		container = GetContainerFromPodTemplate(w.DaemonSet.Spec.Template, serviceName)
	case WORKLOAD_KIND_CRONJOB:
		m.WorkloadName = w.CronJob.Name
		container = GetContainerFromPodTemplate(w.CronJob.Spec.JobTemplate.Spec.Template, serviceName)
	case WORKLOAD_KIND_JOB:
		m.WorkloadName = w.Job.Name
		container = GetContainerFromPodTemplate(w.Job.Spec.Template, serviceName)
	}
	m.Image = container.Image

	return m
}
