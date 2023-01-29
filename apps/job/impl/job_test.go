package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryDeploy(t *testing.T) {
	req := job.NewQueryJobRequest()
	ds, err := impl.QueryJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestCreateBuildJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "容器镜像构建"
	req.RunnerSpec = tools.MustReadContentFile("test/build.yaml")
	ds, err := impl.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestCreateDeployJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "容器镜像部署"
	req.RunnerSpec = tools.MustReadContentFile("test/deploy.yaml")
	ds, err := impl.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
