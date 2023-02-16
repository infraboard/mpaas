package impl_test

import (
	"os"
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
	req := build.NewDescribeBuildConfigRequst(os.Getenv("BUILD_CONFIG_ID"))
	ins, err := impl.DescribeBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateBuildConfig(t *testing.T) {
	req := build.NewPatchDeployRequest(os.Getenv("BUILD_CONFIG_ID"))
	req.Spec.Condition.AddEvent("put")
	req.Spec.Condition.AddBranche("master")
	ins, err := impl.UpdateBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteBuildConfig(t *testing.T) {
	req := build.NewDeleteBuildConfigRequest(os.Getenv("BUILD_CONFIG_ID"))
	ins, err := impl.DeleteBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateBuildConfig(t *testing.T) {
	req := build.NewCreateBuildConfigRequest()
	req.Name = "测试构建"
	req.ServiceId = os.Getenv("SERVICE_ID")
	ins, err := impl.CreateBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
