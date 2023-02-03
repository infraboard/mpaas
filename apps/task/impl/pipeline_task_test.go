package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
)

func TestRunPipeline(t *testing.T) {
	req := task.NewRunPipelineRequest()
	ins, err := impl.RunPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
