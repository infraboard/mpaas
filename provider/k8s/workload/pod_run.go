package workload

import (
	"context"

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
}

func (c *Client) CopyPodRun(ctx context.Context, req *CopyPodRunRequest) (*v1.Pod, error) {
	sourcePod, err := c.GetPod(ctx, req.SourcePod)
	if err != nil {
		return nil, err
	}

	targetPod := &v1.Pod{}
	targetPod.Spec = sourcePod.DeepCopy().Spec
	targetPod.ObjectMeta = req.TargetPodMeta

	return c.CreatePod(ctx, targetPod, req.TargetPodOpts)
}
