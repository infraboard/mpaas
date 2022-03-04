package k8s

import (
	"context"

	v1 "k8s.io/api/batch/v1"
)

func (c *Client) ListCronJob(ctx context.Context, req *ListRequest) (*v1.CronJobList, error) {
	return c.client.BatchV1().CronJobs(req.Namespace).List(ctx, req.Opts)
}
