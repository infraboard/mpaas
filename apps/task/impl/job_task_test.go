package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
)

func TestRunJob(t *testing.T) {

	req := task.NewRunJobRequest("docker_build")
	version := job.NewVersionedRunParam("v1")
	version.Params = job.NewRunParamWithKVPaire(
		"cluster_id", "k8s-test",
		"DB_PASS", "test",
	)

	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	set, err := impl.QueryJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
