package workload

import (
	"context"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mpaas/common/format"
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
		PodLogOptions: &v1.PodLogOptions{
			Follow:                       false,
			Previous:                     false,
			InsecureSkipTLSVerifyBackend: true,
		},
	}
}

type WatchConainterLogRequest struct {
	Namespace string `json:"namespace" validate:"required"`
	PodName   string `json:"pod_name" validate:"required"`
	*v1.PodLogOptions
}

func (req *WatchConainterLogRequest) Validate() error {
	return validate.Struct(req)
}

// 查看容器日志
func (c *Workload) WatchConainterLog(ctx context.Context, req *WatchConainterLogRequest) (io.ReadCloser, error) {
	restReq := c.corev1.Pods(req.Namespace).GetLogs(req.PodName, req.PodLogOptions)
	return restReq.Stream(ctx)
}

func InjectEnvVars(c *v1.Container, envs []v1.EnvVar) {
	set := NewEnvVarSet(c.Env)
	for _, env := range envs {
		e := set.GetOrNewEnv(env.Name)
		e.Value = env.Value
		e.ValueFrom = nil
	}
	fmt.Println(set.Items)
	c.Env = set.EnvVars()
}

func NewEnvVarSet(envs []v1.EnvVar) *EnvVarSet {
	set := &EnvVarSet{
		Items: []*v1.EnvVar{},
	}
	for _, env := range envs {
		set.Add(&env)
	}
	return set
}

type EnvVarSet struct {
	Items []*v1.EnvVar
}

func (s *EnvVarSet) String() string {
	return format.Prettify(s)
}

func (s *EnvVarSet) Add(item *v1.EnvVar) {
	s.Items = append(s.Items, item)
}

func (s *EnvVarSet) EnvVars() (envs []v1.EnvVar) {
	for i := range s.Items {
		item := s.Items[i]
		envs = append(envs, *item)
	}
	return
}

// 如果有就返回已有的Env, 如果没有则创建新的Env
func (s *EnvVarSet) GetOrNewEnv(name string) *v1.EnvVar {
	for i := range s.Items {
		item := s.Items[i]
		if item.Name == name {
			return item
		}
	}

	newEnv := &v1.EnvVar{
		Name: name,
	}
	s.Add(newEnv)

	return newEnv
}
