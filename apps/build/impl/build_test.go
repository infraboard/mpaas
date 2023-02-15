package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryBuildConfig(t *testing.T) {
	req := build.NewQueryBuildConfigRequest()
	set, err := impl.QueryBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestCreateBuildConfig(t *testing.T) {
	req := build.NewCreateBuildConfigRequest()
	ins, err := impl.CreateBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
