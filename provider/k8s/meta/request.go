package meta

import (
	"net/http"
	"net/url"

	v1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewListRequestFromHttp(r *http.Request) *ListRequest {
	qs := r.URL.Query()

	req := &ListRequest{
		Namespace:         qs.Get("namespace"),
		SkipManagedFields: qs.Get("skip_managed_fields") == "true",
		Opts: metav1.ListOptions{
			LabelSelector: qs.Get("label"),
		},
	}

	return req
}

func NewGetRequestFromHttp(r *http.Request) *GetRequest {
	qs := r.URL.Query()

	req := &GetRequest{
		Namespace: qs.Get("namespace"),
		Name:      qs.Get("name"),
	}

	return req
}

func NewGetRequest(name string) *GetRequest {
	return &GetRequest{
		Namespace: DEFAULT_NAMESPACE,
		Name:      name,
	}
}

type GetRequest struct {
	Namespace string
	Name      string
	Opts      metav1.GetOptions
}

func (r *GetRequest) WithNamespace(namespace string) *GetRequest {
	r.Namespace = namespace
	return r
}

func NewDeleteRequest(name string) *DeleteRequest {
	req := &DeleteRequest{
		Namespace: DEFAULT_NAMESPACE,
		Name:      name,
		Opts:      metav1.DeleteOptions{},
	}
	req.SetPropagationPolicy(metav1.DeletePropagationBackground)
	return req
}

type DeleteRequest struct {
	Namespace string
	Name      string
	Opts      metav1.DeleteOptions
}

func (r *DeleteRequest) WithNamespace(namespace string) *DeleteRequest {
	r.Namespace = namespace
	return r
}

func (req *DeleteRequest) SetPropagationPolicy(dp metav1.DeletionPropagation) {
	req.Opts.PropagationPolicy = &dp
}

func NewListRequest() *ListRequest {
	return &ListRequest{}
}

type ListRequest struct {
	Namespace         string
	SkipManagedFields bool
	Opts              metav1.ListOptions
}

func (r *ListRequest) WithNamespace(ns string) *ListRequest {
	r.Namespace = ns
	return r
}

func NewLabelSelector() *LabelSelector {
	return &LabelSelector{
		values: url.Values{},
	}
}

type LabelSelector struct {
	values url.Values
}

func (l *LabelSelector) Add(key, value string) *LabelSelector {
	l.values.Add(key, value)
	return l
}

func (r *ListRequest) WithLabelSelector(l *LabelSelector) *ListRequest {
	r.Opts.LabelSelector = l.values.Encode()
	return r
}

func NewCreateRequest() *CreateRequest {
	return &CreateRequest{
		Opts: metav1.CreateOptions{},
	}
}

type CreateRequest struct {
	Opts metav1.CreateOptions
}

func NewScaleRequest() *ScaleRequest {
	return &ScaleRequest{
		Scale:   &v1.Scale{},
		Options: metav1.UpdateOptions{},
	}
}

type ScaleRequest struct {
	Scale   *v1.Scale
	Options metav1.UpdateOptions
}
