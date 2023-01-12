package batch

import (
	"k8s.io/client-go/kubernetes"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func NewBatch(cs *kubernetes.Clientset) *Batch {
	return &Batch{
		batchV1: cs.BatchV1(),
	}
}

type Batch struct {
	batchV1 batchv1.BatchV1Interface
}
