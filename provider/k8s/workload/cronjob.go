package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/batch/v1"
)

func (b *Client) ListCronJob(ctx context.Context, req *meta.ListRequest) (*v1.CronJobList, error) {
	return b.batchV1.CronJobs(req.Namespace).List(ctx, req.Opts)
}

func (b *Client) DeleteCronJob(ctx context.Context, req *meta.DeleteRequest) error {
	return b.batchV1.CronJobs(req.Namespace).Delete(ctx, req.Name, req.Opts)
}

func GetCronJobStatus(*v1.CronJob) *WorkloadStatus {
	return NewWorklaodStatus()
}
