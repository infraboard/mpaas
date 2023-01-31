package k8s_test

import (
	"context"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task/runner"
)

var (
	impl runner.Runner
	ctx  = context.Background()
)

func init() {
	impl = runner.GetRunner(job.RUNNER_TYPE_K8S_JOB)
}
