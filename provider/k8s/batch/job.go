package batch

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/batch/v1"
)

func (b *Batch) ListJob(ctx context.Context, req *meta.ListRequest) (*v1.JobList, error) {
	return b.batchV1.Jobs(req.Namespace).List(ctx, req.Opts)
}
