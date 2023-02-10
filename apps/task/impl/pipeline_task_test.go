package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunPipeline(t *testing.T) {
	req := task.NewRunPipelineRequest("cfj3n4ts99brfo8rb8i0")
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
	t.Log(tools.MustToYaml(set))
}

func TestDescribePipelineTask(t *testing.T) {
	req := task.NewDescribePipelineTaskRequest("cfj3o2ts99brj2n5l230")
	ins, err := impl.DescribePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestDeletePipelineTask(t *testing.T) {
	req := task.NewDeletePipelineTaskRequest("cfj3o2ts99brj2n5l230")
	ins, err := impl.DeletePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}
