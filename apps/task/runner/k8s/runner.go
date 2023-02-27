package k8s

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/runner"
	"github.com/infraboard/mpaas/provider/k8s"
)

type K8sRunner struct {
	cluster cluster.Service
	deploy  deploy.Service
	task    task.PipelineService
	log     logger.Logger

	k8sClient *k8s.Client
}

func (r *K8sRunner) Init() error {
	r.cluster = app.GetInternalApp(cluster.AppName).(cluster.Service)
	r.deploy = app.GetInternalApp(deploy.AppName).(deploy.Service)
	r.task = app.GetInternalApp(task.AppName).(task.PipelineService)
	r.log = zap.L().Named("runner.k8s")
	return nil
}

func (r *K8sRunner) RunnerType() job.RUNNER_TYPE {
	return job.RUNNER_TYPE_K8S_JOB
}

func init() {
	runner.Registry(&K8sRunner{})
}
