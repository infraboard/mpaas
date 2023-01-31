package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
)

func (i *impl) RunJob(ctx context.Context, in *task.RunJobRequest) (
	*task.Task, error) {
	// 1. 查询需要执行的Job
	j, err := i.job.DescribeJob(ctx, nil)
	if err != nil {
		return nil, err
	}
	i.log.Debug(j)

	switch j.Spec.RunnerType {
	case job.RUNNER_TYPE_K8S_JOB:
		runnerParams := in.Params.K8SJobRunnerParams()
		cReq := cluster.NewDescribeClusterRequest(runnerParams.ClusterId)
		c, err := i.cluster.DescribeCluster(ctx, cReq)
		if err != nil {
			return nil, err
		}
		k8sClient, err := c.Client()
		if err != nil {
			return nil, err
		}
		fmt.Println(k8sClient)
	}

	return nil, nil
}
