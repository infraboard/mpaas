package workload

import (
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

func NewWorkload(cs *kubernetes.Clientset, restconf *rest.Config) *Client {
	return &Client{
		appsv1:   cs.AppsV1(),
		batchV1:  cs.BatchV1(),
		corev1:   cs.CoreV1(),
		restconf: restconf,
		l:        log.Sub("workload"),
	}
}

type Client struct {
	appsv1   appsv1.AppsV1Interface
	batchV1  batchv1.BatchV1Interface
	corev1   corev1.CoreV1Interface
	restconf *rest.Config
	l        *zerolog.Logger
}
