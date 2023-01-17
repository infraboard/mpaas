package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/deploy"
)

func TestQueryDeploy(t *testing.T) {
	req := deploy.NewQueryDeployRequest()
	ds, err := impl.QueryDeploy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
