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

func TestDescribeBuildConfig(t *testing.T) {
	req := build.NewDescribeBuildConfigRequst("cfmtca6a0brmve8mlgu0")
	ins, err := impl.DescribeBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteBuildConfig(t *testing.T) {
	req := build.NewDeleteBuildConfigRequest("cfmtca6a0brmve8mlgu0")
	ins, err := impl.DeleteBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateBuildConfig(t *testing.T) {
	req := build.NewCreateBuildConfigRequest()
	ins, err := impl.CreateBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
