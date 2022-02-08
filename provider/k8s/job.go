package k8s

import (
	"context"

	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListJob(ctx context.Context, req *ListDeploymentRequest) (*v1.JobList, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.BatchV1().Jobs(req.Namespace).List(ctx, metav1.ListOptions{})
}
