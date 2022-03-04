package k8s

import (
	"context"

	v1 "k8s.io/api/batch/v1"
)

func (c *Client) ListJob(ctx context.Context, req *ListRequest) (*v1.JobList, error) {
	return c.client.BatchV1().Jobs(req.Namespace).List(ctx, req.Opts)
}
