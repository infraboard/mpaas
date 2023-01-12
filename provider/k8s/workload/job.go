package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (b *Workload) ListJob(ctx context.Context, req *meta.ListRequest) (*v1.JobList, error) {
	return b.batchV1.Jobs(req.Namespace).List(ctx, req.Opts)
}

func (b *Workload) GetJob(ctx context.Context, req *meta.GetRequest) (*v1.Job, error) {
	return b.batchV1.Jobs(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (b *Workload) CreateJob(ctx context.Context, job *v1.Job) (*v1.Job, error) {
	return b.batchV1.Jobs(job.Namespace).Create(ctx, job, metav1.CreateOptions{})
}
