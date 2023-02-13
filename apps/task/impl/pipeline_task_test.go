package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunPipeline(t *testing.T) {
	req := task.NewRunPipelineRequest(os.Getenv("PIPELINE_ID"))
	ins, err := impl.RunPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestQueryPipelineTask(t *testing.T) {
	req := task.NewQueryPipelineTaskRequest()
	set, err := impl.QueryPipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestDescribePipelineTask(t *testing.T) {
	req := task.NewDescribePipelineTaskRequest(os.Getenv("PIPELINE_TASK_ID"))
	ins, err := impl.DescribePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeletePipelineTask(t *testing.T) {
	req := task.NewDeletePipelineTaskRequest(os.Getenv("PIPELINE_TASK_ID"))
	ins, err := impl.DeletePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}
