package workload

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCopyPodRunRequest() *CopyPodRunRequest {
	return &CopyPodRunRequest{
		SourcePod:     meta.NewGetRequest(""),
		TargetPodOpts: meta.NewCreateRequest(),
	}
}

type CopyPodRunRequest struct {
	// 需要被Copy的容器
	SourcePod *meta.GetRequest
	// 克隆后的目标容器运行时的其他参数
	TargetPodMeta metav1.ObjectMeta
	// 目标创建容器的创建选项
	TargetPodOpts *meta.CreateRequest
	// 登录目标容器的名称
	ExecContainer string `json:"exec_container"`
	// 用于Hold主Container的命令, 默认 sleep infinity
	ExecHoldCmd []string `json:"exec_hold_cmd"`
	// 登录目标容器的命令
	ExecRunCmd []string `json:"exec_run_cmd"`
	// 是否登录目录容器
	Attach bool `json:"attach"`
	// 当登录终端后,退出终端是否删除容器
	Remove bool `json:"remove"`
	// 目标容器的优雅关闭时间, 默认30秒
	TerminationGracePeriodSeconds int64
	// 登录终端
	Terminal ContainerTerminal `json:"-"`
}

func (r *CopyPodRunRequest) SetAttachTerminal(term ContainerTerminal) {
	r.Attach = true
	r.Terminal = term
}

func (c *Client) CopyPodRun(ctx context.Context, req *CopyPodRunRequest) (*v1.Pod, error) {
	sourcePod, err := c.GetPod(ctx, req.SourcePod)
	if err != nil {
		return nil, err
	}

	targetPod := &v1.Pod{}
	targetPod.Spec = sourcePod.DeepCopy().Spec
	targetPod.ObjectMeta = req.TargetPodMeta
	// 调整Pod关闭操时时长
	if req.TerminationGracePeriodSeconds != 0 {
		targetPod.Spec.TerminationGracePeriodSeconds = &req.TerminationGracePeriodSeconds
	}

	if len(targetPod.Spec.Containers) == 0 {
		return nil, fmt.Errorf("no container found in spec")
	}

	// 需要Debug的容器 Hold住

	execContainer := &targetPod.Spec.Containers[0]
	if req.ExecContainer != "" {
		execContainer = GetContainerFromPodSpec(targetPod.Spec, req.ExecContainer)
		if execContainer == nil {
			return nil, fmt.Errorf("container not found")
		}
	}
	if len(req.ExecHoldCmd) > 0 {
		execContainer.Command = req.ExecHoldCmd
	}

	// 创建目标Pod
	pod, err := c.CreatePod(ctx, targetPod, req.TargetPodOpts)
	if err != nil {
		return nil, err
	}

	if req.Attach {
		// 自动删除Pod
		if req.Remove {
			defer func() {
				delReq := meta.NewDeleteRequest(req.TargetPodMeta.Name).
					WithNamespace(req.TargetPodMeta.Namespace)
				err := c.DeletePod(ctx, delReq)
				if err != nil {
					c.l.Error("delete pod error, %s", err)
				}
			}()
		}

		// 等待目标容器启动
		pod, err = c.WaitForPodCondition(ctx, &WaitForContainerRequest{
			Namespace:     req.TargetPodMeta.Namespace,
			PodName:       req.TargetPodMeta.Name,
			ContainerName: req.ExecContainer,
			ExitCondition: WaitForContainerRunning(req.ExecContainer, req.Terminal),
		})
		if err != nil {
			return nil, err
		}

		// 登录容器
		err = c.LoginContainer(ctx, &LoginContainerRequest{
			Namespace:     req.TargetPodMeta.Namespace,
			PodName:       req.TargetPodMeta.Name,
			ContainerName: req.ExecContainer,
			Command:       shellCmd,
			Excutor:       req.Terminal,
		})
		if err != nil {
			return nil, err
		}
	}

	return pod, nil
}