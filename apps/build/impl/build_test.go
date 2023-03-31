package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryBuildConfig(t *testing.T) {
	req := build.NewQueryBuildConfigRequest()
	req.AddService(conf.C.MCENTER_SERVICE_ID)
	req.Event = "push"
	set, err := impl.QueryBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestDescribeBuildConfig(t *testing.T) {
	req := build.NewDescribeBuildConfigRequst(conf.C.BUILD_ID)
	ins, err := impl.DescribeBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateBuildConfig(t *testing.T) {
	req := build.NewPatchBuildConfigRequest(conf.C.BUILD_ID)
	req.Spec.Condition.AddEvent("push")
	req.Spec.Condition.AddBranche("master")
	req.Spec.ImageBuild.PipelineId = conf.C.MPAAS_PIPELINE_ID
	ins, err := impl.UpdateBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteBuildConfig(t *testing.T) {
	req := build.NewDeleteBuildConfigRequest(conf.C.BUILD_ID)
	ins, err := impl.DeleteBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateBuildConfig(t *testing.T) {
	req := build.NewCreateBuildConfigRequest()
	req.Name = "mcenter服务构建"
	req.ServiceId = conf.C.MCENTER_SERVICE_ID
	req.Condition.AddEvent("push")
	req.Condition.AddBranche("master")
	req.ImageBuild.PipelineId = conf.C.CICD_PIPELINE_ID
	ins, err := impl.CreateBuildConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
