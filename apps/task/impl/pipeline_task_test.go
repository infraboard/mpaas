package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryPipelineTask(t *testing.T) {
	req := task.NewQueryPipelineTaskRequest()
	set, err := impl.QueryPipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestRunTestPipeline(t *testing.T) {
	req := pipeline.NewRunPipelineRequest(conf.C.PIPELINE_ID)
	req.RunBy = "test"
	ins, err := impl.RunPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestRunMpaasPipeline(t *testing.T) {
	req := pipeline.NewRunPipelineRequest(conf.C.MPAAS_PIPELINE_ID)
	req.RunBy = "test"
	req.RunParams = job.NewRunParamWithKVPaire(
		"GIT_REPOSITORY", "git@github.com:infraboard/mpaas.git",
		"GIT_BRANCH", "master",
		"GIT_COMMIT_ID", "57953a59e0ff5c93d0596696fbf6ffef6a90b446",
		job.SYSTEM_VARIABLE_IMAGE_VERSION, "v0.0.10",
	)

	ins, err := impl.RunPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestDescribePipelineTask(t *testing.T) {
	req := task.NewDescribePipelineTaskRequest(conf.C.PIPELINE_TASK_ID)
	ins, err := impl.DescribePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeletePipelineTask(t *testing.T) {
	req := task.NewDeletePipelineTaskRequest(conf.C.PIPELINE_TASK_ID)
	ins, err := impl.DeletePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}
