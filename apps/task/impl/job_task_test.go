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
	t.Log(tools.MustToJson(ins))
}

func TestQueryJobTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	req.PipelineTaskId = "cfkeq7ds99bvnft65p20"
	set, err := impl.QueryJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tools.MustToJson(set))
}

func TestDescribeJobTask(t *testing.T) {
	req := task.NewDescribeJobTaskRequest("cfk6mdls99bopbk780ig")
	ins, err := impl.DescribeJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateJobTaskStatus(t *testing.T) {
	req := task.NewUpdateJobTaskStatusRequest("cfkeq7ds99bvnft65p30")
	req.Stage = task.STAGE_SUCCEEDED
	req.Message = "执行成功"
	req.Detail = tools.MustReadContentFile("test/k8s_job.yml")
	ins, err := impl.UpdateJobTaskStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestDeleteJobTask(t *testing.T) {
	req := task.NewDeleteJobTaskRequest("cfk6mdls99bopbk780ig")
	set, err := impl.DeleteJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
