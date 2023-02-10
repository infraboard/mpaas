package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunPipeline(t *testing.T) {
	req := task.NewRunPipelineRequest("cfiucuea0brqa1kj3go0")
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
	req := task.NewDescribePipelineTaskRequest("cfiuic6a0brro18jh7kg")
	ins, err := impl.DescribePipelineTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}
