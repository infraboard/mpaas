package event

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/events/v1"
	"k8s.io/client-go/kubernetes"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
)

func NewEvent(cs *kubernetes.Clientset) *Client {
	return &Client{
		eventsv1: cs.EventsV1(),
	}
}

type Client struct {
	eventsv1 eventsv1.EventsV1Interface
}

func (c *Client) ListEvent(ctx context.Context, req *meta.ListRequest) (*v1.EventList, error) {
	return c.eventsv1.Events(req.Namespace).List(ctx, req.Opts)
}
