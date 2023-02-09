package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_build@default.default")
	version := job.NewVersionedRunParam("v1")
	version.Params = job.NewRunParamWithKVPaire(
		"cluster_id", "k8s-test",
		"DB_PASS", "test",
	)
	req.Params = version

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

	t.Log(tools.MustToYaml(set))
}

func TestDescribeJobTask(t *testing.T) {
	req := task.NewDescribeJobTaskRequest("cfhfufua0brh83njg8ag")
	set, err := impl.DescribeJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDeleteJobTask(t *testing.T) {
	req := task.NewDeleteJobTaskRequest("cfhfufua0brh83njg8ag")
	set, err := impl.DeleteJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
