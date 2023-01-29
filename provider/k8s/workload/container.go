package workload

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
func (c *Workload) LoginContainer(ctx context.Context, req *LoginContainerRequest) error {
	restReq := c.corev1.RESTClient().Post().
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

	return executor.StreamWithContext(ctx, remotecommand.StreamOptions{
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
func (c *Workload) WatchConainterLog(ctx context.Context, req *WatchConainterLogRequest) (io.ReadCloser, error) {
	opt := &v1.PodLogOptions{
		Container:                    req.ContainerName,
		Follow:                       req.Follow,
		TailLines:                    &req.TailLines,
		Previous:                     req.Previous,
		InsecureSkipTLSVerifyBackend: true,
	}

	restReq := c.corev1.Pods(req.Namespace).GetLogs(req.PodName, opt)
	return restReq.Stream(ctx)
}
