package k8s

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/task"
)

func (r *K8sRunner) Run(ctx context.Context, in *task.RunJobRequest) (*task.Task, error) {
	runnerParams := in.Params.K8SJobRunnerParams()
	cReq := cluster.NewDescribeClusterRequest(runnerParams.ClusterId)
	c, err := r.cluster.DescribeCluster(ctx, cReq)
	if err != nil {
		return nil, err
	}
	k8sClient, err := c.Client()
	if err != nil {
		return nil, err
	}
	fmt.Println(k8sClient)
	return nil, nil
}
