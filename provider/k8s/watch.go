package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
)

type ObjectKind string

const (
	ObjectKindPod     ObjectKind = "pod"
	ObjectKindDeploy  ObjectKind = "deployment"
	ObjectStatefulset ObjectKind = "statefulset"
	ObjectDaemonset   ObjectKind = "daemonset"
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
	case ObjectKindPod:
		return c.client.CoreV1().Pods(req.Namespace).Watch(ctx, req.WatchOptions())
	case ObjectKindDeploy:
		return c.client.AppsV1().Deployments(req.Namespace).Watch(ctx, req.WatchOptions())
	case ObjectStatefulset:
		return c.client.AppsV1().StatefulSets(req.Namespace).Watch(ctx, req.WatchOptions())
	case ObjectDaemonset:
		return c.client.AppsV1().DaemonSets(req.Namespace).Watch(ctx, req.WatchOptions())
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
