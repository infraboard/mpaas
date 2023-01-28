package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
)

func TestQueryDeploy(t *testing.T) {
	req := task.NewRunJobRequest()
	ds, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
