package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
)

func TestQueryDeploy(t *testing.T) {
	req := job.NewQueryJobRequest()
	ds, err := impl.QueryJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
