package network

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func (c *Client) Run(ctx context.Context, yml string) (*v1.Service, error) {
	obj, err := ParseServiceFromYaml(yml)
	if err != nil {
		return nil, err
	}
	return c.CreateService(ctx, obj)
}

func (c *Client) CreateService(ctx context.Context, req *v1.Service) (*v1.Service, error) {
	return c.corev1.Services(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) UpdateService(ctx context.Context, req *v1.Service) (*v1.Service, error) {
	return c.corev1.Services(req.Namespace).Update(ctx, req, metav1.UpdateOptions{})
}

func (c *Client) ListService(ctx context.Context, req *meta.ListRequest) (*v1.ServiceList, error) {
	return c.corev1.Services(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetService(ctx context.Context, req *meta.GetRequest) (*v1.Service, error) {
	return c.corev1.Services(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Client) DeleteService(ctx context.Context, req *meta.DeleteRequest) error {
	return c.corev1.Services(req.Namespace).Delete(ctx, req.Name, req.Opts)
}

func ParseServiceFromYaml(yml string) (*v1.Service, error) {
	obj := &v1.Service{}
	err := yaml.Unmarshal([]byte(yml), obj)
	if err != nil {
		return nil, err
	}
	if obj.Annotations == nil {
		obj.Annotations = make(map[string]string)
	}
	return obj, nil
}
