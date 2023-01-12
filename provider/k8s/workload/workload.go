package workload

import (
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func NewWorkload(cs *kubernetes.Clientset) *Workload {
	return &Workload{
		appsv1: cs.AppsV1(),
	}
}

type Workload struct {
	appsv1 appsv1.AppsV1Interface
}
