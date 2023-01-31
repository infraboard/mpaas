package k8s_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRun(t *testing.T) {
	jobSpec := tools.MustReadContentFile("test/job.yaml")
	req := task.NewRunTaskRequest(jobSpec, nil)
	ins, err := impl.Run(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
