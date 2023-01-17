package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/deploy"
)

func TestQueryDeploy(t *testing.T) {
	req := deploy.NewQueryDeployConfigRequest()
	ds, err := impl.QueryDeployConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
