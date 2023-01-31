package k8s

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/task"
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"
)

func (r *K8sRunner) Run(ctx context.Context, in *task.RunTaskRequest) (*task.Task, error) {
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

	// 执行Job
	obj := new(v1.Job)
	if err := yaml.Unmarshal([]byte(in.JobSpec), obj); err != nil {
		return nil, err
	}
	obj, err = k8sClient.WorkLoad().CreateJob(ctx, obj)
	if err != nil {
		return nil, err
	}

	fmt.Println(obj)
	return nil, nil
}
