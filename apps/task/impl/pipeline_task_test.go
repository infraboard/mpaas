package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunPipeline(t *testing.T) {
	req := task.NewRunPipelineRequest("cfkenp5s99bvhd6hgatg")
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
	req := task.NewDescribePipelineTaskRequest("cfkeq7ds99bvnft65p20")
	ins, err := impl.DescribePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeletePipelineTask(t *testing.T) {
	req := task.NewDeletePipelineTaskRequest("cfkeq7ds99bvnft65p20")
	ins, err := impl.DeletePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}
