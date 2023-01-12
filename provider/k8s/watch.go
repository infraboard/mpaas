package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
)

var (
	validate = validator.New()
)

type ObjectKind string

func (t ObjectKind) String() string {
	return string(t)
}

const (
	OBJECT_DEPLOY       ObjectKind = "deployment"
	OBJECT_STATEFUL_SET ObjectKind = "statefulset"
	OBJECT_DAEMON_SET   ObjectKind = "daemonset"
	OBJECT_JOB          ObjectKind = "job"
	OBJECT_CRONJOB      ObjectKind = "cronjob"
	OBJECT_POD          ObjectKind = "pod"
)

func NewWatchRequest() *WatchRequest {
	return &WatchRequest{}
}

type WatchRequest struct {
	Namespace  string     `json:"namespace"`
	ObjectKind ObjectKind `json:"kind"`
}

func (req *WatchRequest) WatchOptions() metav1.ListOptions {
	return metav1.ListOptions{}
}

func (req *WatchRequest) Validate() error {
	return validate.Struct(req)
}

func (c *Client) Watch(ctx context.Context, req *WatchRequest) (watch.Interface, error) {
	switch req.ObjectKind {
	case OBJECT_DEPLOY:
		return c.client.AppsV1().Deployments(req.Namespace).Watch(ctx, req.WatchOptions())
	case OBJECT_STATEFUL_SET:
		return c.client.AppsV1().StatefulSets(req.Namespace).Watch(ctx, req.WatchOptions())
	case OBJECT_DAEMON_SET:
		return c.client.AppsV1().DaemonSets(req.Namespace).Watch(ctx, req.WatchOptions())
	case OBJECT_JOB:
		return c.client.BatchV1().Jobs(req.Namespace).Watch(ctx, req.WatchOptions())
	case OBJECT_CRONJOB:
		return c.client.BatchV1().CronJobs(req.Namespace).Watch(ctx, req.WatchOptions())
	case OBJECT_POD:
		return c.client.CoreV1().Pods(req.Namespace).Watch(ctx, req.WatchOptions())
	default:
		return nil, fmt.Errorf("unknown Object Kind %s", req.ObjectKind)
	}
}

func NewWatchReader(w watch.Interface) *WatchReader {
	return &WatchReader{
		w:  w,
		ch: w.ResultChan(),
	}
}

type WatchReader struct {
	w  watch.Interface
	ch <-chan watch.Event
}

func (r *WatchReader) Close() error {
	r.w.Stop()
	return nil
}

func (r *WatchReader) Read(p []byte) (int, error) {
	e := <-r.ch
	jb, err := json.Marshal(e)
	if err != nil {
		return 0, err
	}
	copy(p, jb)
	return len(jb), nil
}
