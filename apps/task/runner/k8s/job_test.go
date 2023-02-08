package k8s_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRun(t *testing.T) {
	jobSpec := tools.MustReadContentFile("test/job.yaml")
	params := job.NewVersionedRunParam("v0.1")
	params.Add(
		&job.RunParam{
			Name:     "cluster_id",
			Required: true,
			Value:    "k8s-test",
		},
		&job.RunParam{
			Name:     "DB_PASS",
			Required: true,
			Value:    "test",
		},
	)

	req := task.NewRunTaskRequest("test-job", jobSpec, params)
	ins, err := impl.Run(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
