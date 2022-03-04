package k8s

import (
	"context"

	v1 "k8s.io/api/events/v1"
)

func (c *Client) ListEvent(ctx context.Context, req *ListRequest) (*v1.EventList, error) {
	return c.client.EventsV1().Events(req.Namespace).List(ctx, req.Opts)
}
