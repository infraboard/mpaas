package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
)

func TestRunJob(t *testing.T) {
	req := task.NewRunJobRequest()
	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	set, err := impl.QueryTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
