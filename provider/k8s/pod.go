package k8s

import (
	"context"
	"io"

	"github.com/go-playground/validator/v10"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

var (
	validate = validator.New()
)

func (c *Client) CreatePod(ctx context.Context, pod *v1.Pod, req *CreateRequest) (*v1.Pod, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().Pods(req.Namespace).Create(ctx, pod, req.Opts)
}

func (c *Client) ListPod(ctx context.Context, req *ListRequest) (*v1.PodList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().Pods(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetPod(ctx context.Context, req *GetRequest) (*v1.Pod, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().Pods(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Client) DeletePod(ctx context.Context, req *DeleteRequest) error {
	return c.client.CoreV1().Pods("").Delete(ctx, "", req.Opts)
}

func NewLoginContainerRequest(cmd []string, ce ContainerExecutor) *LoginContainerRequest {
	return &LoginContainerRequest{
		Command: cmd,
		Excutor: ce,
	}
}

type LoginContainerRequest struct {
	Namespace     string            `json:"namespace" validate:"required"`
	PodName       string            `json:"pod_name" validate:"required"`
	ContainerName string            `json:"container_name"`
	Command       []string          `json:"command"`
	Excutor       ContainerExecutor `json:"-"`
}

func (req *LoginContainerRequest) Validate() error {
	return validate.Struct(req)
}

type ContainerExecutor interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

// 登录容器
func (c *Client) LoginContainer(req *LoginContainerRequest) error {
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
		return err
	}

	return executor.Stream(remotecommand.StreamOptions{
		Stdin:             req.Excutor,
		Stdout:            req.Excutor,
		Stderr:            req.Excutor,
		Tty:               true,
		TerminalSizeQueue: req.Excutor,
	})
}

func NewWatchConainterLogRequest() *WatchConainterLogRequest {
	return &WatchConainterLogRequest{
		TailLines: 100,
		Follow:    false,
		Previous:  false,
	}
}

type WatchConainterLogRequest struct {
	Namespace     string `json:"namespace" validate:"required"`
	PodName       string `json:"pod_name" validate:"required"`
	ContainerName string `json:"container_name"`
	TailLines     int64  `json:"tail_lines"`
	Follow        bool   `json:"follow"`
	Previous      bool   `json:"previous"`
}

func (req *WatchConainterLogRequest) Validate() error {
	return validate.Struct(req)
}

// 查看容器日志
func (c *Client) WatchConainterLog(ctx context.Context, req *WatchConainterLogRequest) (io.ReadCloser, error) {
	opt := &v1.PodLogOptions{
		Container:                    req.ContainerName,
		Follow:                       req.Follow,
		TailLines:                    &req.TailLines,
		Previous:                     req.Previous,
		InsecureSkipTLSVerifyBackend: true,
	}

	restReq := c.client.CoreV1().Pods(req.Namespace).GetLogs(req.PodName, opt)
	return restReq.Stream(ctx)
}
