package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryJob(t *testing.T) {
	req := job.NewQueryJobRequest()
	set, err := impl.QueryJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestCreateBuildJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "docker_build"
	req.CreateBy = "test"
	req.RunnerSpec = tools.MustReadContentFile("test/build.yml")

	ins, err := impl.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateDeployJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "docker_deploy"
	req.CreateBy = "test"
	req.RunnerSpec = tools.MustReadContentFile("test/deployment.yml")
	v1 := job.NewVersionedRunParam("v1")
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "cluster_id",
		Desc:     "job运行时的k8s集群",
	})

	req.AddVersionParams(v1)

	ins, err := impl.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateJob(t *testing.T) {
	req := job.NewPatchJobRequest(os.Getenv("JOB_ID"))
	req.Spec.RunnerSpec = tools.MustReadContentFile("test/build.yml")
	v1 := job.NewVersionedRunParam("v1")
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "cluster_id",
		Desc:     "job运行时的k8s集群",
	})
	// 代码仓库地址, 比如 git@github.com:infraboard/mpaas.git
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "GIT_ADDRESS",
		Desc:     "应用git代码仓库地址",
	})
	// docker push registry.cn-hangzhou.aliyuncs.com/inforboard/mpaas:[镜像版本号]
	// IMAGE_REPOSITORY: egistry.cn-hangzhou.aliyuncs.com/inforboard/mpaas
	// IMAGE_VERSION: [镜像版本号]
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "IMAGE_REPOSITORY",
		Desc:     "镜像推送地址",
	})
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "IMAGE_VERSION",
		Desc:     "镜像版本",
	})

	req.Spec.AddVersionParams(v1)

	ins, err := impl.UpdateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDescribeJob(t *testing.T) {
	req := job.NewDescribeJobRequest("docker_build@default.default")
	ins, err := impl.DescribeJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
