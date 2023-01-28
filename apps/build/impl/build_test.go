package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/build"
)

func TestQueryDeploy(t *testing.T) {
	req := build.NewQueryBuildConfigRequest()
	ds, err := impl.QueryBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
