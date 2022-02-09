package k8s

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (c *Client) ListPod(ctx context.Context, req *ListDeploymentRequest) (*apiv1.PodList, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.CoreV1().Pods(req.Namespace).List(ctx, metav1.ListOptions{})
}

type GetPodRequest struct {
	Namespace string
	Name      string
}

func (c *Client) GetPod(ctx context.Context, req *GetPodRequest) (*apiv1.Pod, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.CoreV1().Pods(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
}

func (c *Client) DeletePod(ctx context.Context) error {
	return c.client.CoreV1().Pods("").Delete(ctx, "", metav1.DeleteOptions{})
}

type LoginContainerRequest struct {
	Namespace     string
	PodName       string
	ContainerName string
	Command       []string
}

// 登录容器
func (c *Client) LoginContainer(req *LoginContainerRequest) {
	restReq := c.client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(req.PodName).
		Namespace(req.Namespace).
		SubResource("exec")

	restReq.VersionedParams(&v1.PodExecOptions{
		Container: req.ContainerName,
		Command:   req.Command,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(c.restconf, "POST", restReq.URL())
	if err != nil {
		return
	}
	fmt.Println(executor)
}

type WatchConainterLogRequest struct {
	ContainerName string
}

// 查看容器日志
func (c *Client) WatchConainterLog(req *WatchConainterLogRequest) {
	count := int64(100)
	opt := &v1.PodLogOptions{
		Container:                    req.ContainerName,
		Follow:                       false,
		TailLines:                    &count,
		InsecureSkipTLSVerifyBackend: true,
	}
	fmt.Println(opt)
}
