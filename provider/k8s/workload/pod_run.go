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
	ExecContainer string
	// 用于Hold主Container的命令, 默认 sleep infinity
	ExecHoldCmd []string
	// 登录目标容器的命令
	ExecRunCmd []string
	// 执行器
	Excutor ContainerTerminal `json:"-"`
}

func (c *Client) CopyPodRun(ctx context.Context, req *CopyPodRunRequest) (*v1.Pod, error) {
	sourcePod, err := c.GetPod(ctx, req.SourcePod)
	if err != nil {
		return nil, err
	}

	targetPod := &v1.Pod{}
	targetPod.Spec = sourcePod.DeepCopy().Spec
	targetPod.ObjectMeta = req.TargetPodMeta

	if len(targetPod.Spec.Containers) == 0 {
		return nil, fmt.Errorf("no container found in spec")
	}

	// 需要Debug的容器 Hold住
	execContainer := &targetPod.Spec.Containers[0]
	if req.ExecContainer == "" {
		execContainer = GetContainerFromPodSpec(targetPod.Spec, req.ExecContainer)
		if execContainer == nil {
			return nil, fmt.Errorf("container not found")
		}
	}
	execContainer.Command = sleepCmd
	if len(req.ExecHoldCmd) > 0 {
		execContainer.Command = req.ExecHoldCmd
	}

	// 创建目标Pod
	pod, err := c.CreatePod(ctx, targetPod, req.TargetPodOpts)
	if err != nil {
		return nil, err
	}

	// 等待Pod运行成功
	// pod, err = c.WaitForPodCondition(ctx, &WaitForContainerRequest{
	// 	Namespace: req.TargetPodMeta.Namespace,
	// 	PodName:   req.TargetPodMeta.Name,
	// 	ExitCondition: func(event watch.Event) (bool, error) {
	// 		switch event.Type {
	// 		case watch.Deleted:
	// 			return false, errors.NewNotFound(schema.GroupResource{Resource: "pods"}, "")
	// 		}
	// 		switch t := event.Object.(type) {
	// 		case *v1.Pod:
	// 			switch t.Status.Phase {
	// 			case v1.PodFailed, v1.PodSucceeded:
	// 				return false, ErrPodCompleted
	// 			case v1.PodRunning:
	// 				conditions := t.Status.Conditions
	// 				if conditions == nil {
	// 					return false, nil
	// 				}
	// 				_, err := req.Excutor.Write([]byte(fmt.Sprintf("%s: %s", t.Status.Phase, t.Status.Message)))
	// 				if err != nil {
	// 					c.l.Errorf("write event error, %s", err)
	// 				}
	// 				for i := range conditions {
	// 					if conditions[i].Type == v1.PodReady &&
	// 						conditions[i].Status == v1.ConditionTrue {
	// 						return true, nil
	// 					}
	// 				}
	// 			}
	// 		}
	// 		return false, nil
	// 	},
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// 登录容器
	// c.LoginContainer(ctx, &LoginContainerRequest{
	// 	Namespace: req.TargetPodMeta.Namespace,
	// 	PodName:   req.TargetPodMeta.Name,
	// })

	return pod, nil
}

// ErrPodCompleted is returned by PodRunning or PodContainerRunning to indicate that
// the pod has already reached completed state.
var ErrPodCompleted = fmt.Errorf("pod ran to completion")
