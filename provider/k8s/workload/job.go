package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// job使用文档: https://kubernetes.io/docs/concepts/workloads/controllers/job/

func (b *Client) ListJob(ctx context.Context, req *meta.ListRequest) (*v1.JobList, error) {
	return b.batchV1.Jobs(req.Namespace).List(ctx, req.Opts)
}

func (b *Client) GetJob(ctx context.Context, req *meta.GetRequest) (*v1.Job, error) {
	return b.batchV1.Jobs(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (b *Client) CreateJob(ctx context.Context, job *v1.Job) (*v1.Job, error) {
	return b.batchV1.Jobs(job.Namespace).Create(ctx, job, metav1.CreateOptions{})
}

func (b *Client) CreateCronJob(ctx context.Context, job *v1.CronJob) (*v1.CronJob, error) {
	return b.batchV1.CronJobs(job.Namespace).Create(ctx, job, metav1.CreateOptions{})
}

func (c *Client) DeleteJob(ctx context.Context, req *meta.DeleteRequest) error {
	return c.batchV1.Jobs(req.Namespace).Delete(ctx, req.Name, req.Opts)
}

// 注入Job标签
func InjectJobLabels(pod *v1.Job, labels map[string]string) {
	for k, v := range labels {
		pod.Labels[k] = v
	}
}
