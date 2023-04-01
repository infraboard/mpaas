package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryJob(t *testing.T) {
	req := job.NewQueryJobRequest()
	req.Label["cluster_id"] = "k8s-test"
	set, err := impl.QueryJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestCreateTestJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "test"
	req.CreateBy = "test"
	req.RunnerSpec = tools.MustReadContentFile("test/test.yml")
	v1 := job.NewVersionedRunParam("v1")
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "cluster_id",
		Desc:     "job运行时的k8s集群",
		Value:    "k8s-test",
	})
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "namespace",
		Desc:     "job运行时的namespace",
		Value:    "default",
	})
	req.AddVersionParams(v1)

	ins, err := impl.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateBuildJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "docker_build"
	req.CreateBy = "test"
	req.RunnerSpec = tools.MustReadContentFile("test/container_build.yml")

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
	req.RunnerSpec = tools.MustReadContentFile("test/container_deploy.yml")

	ins, err := impl.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateDeployJob(t *testing.T) {
	req := job.NewPatchJobRequest(conf.C.DEPLOY_JOB_ID)
	req.Spec.RunnerSpec = tools.MustReadContentFile("test/container_deploy.yml")

	v1 := job.NewVersionedRunParam("v1")
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        "cluster_id",
		Desc:        "job运行时的k8s集群",
		Value:       "k8s-test",
		SearchLabel: true,
	})
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        "namespace",
		Desc:        "job运行时的namespace",
		Value:       "default",
		SearchLabel: true,
	})

	// 部署运行时变量
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        job.SYSTEM_VARIABLE_DEPLOY_ID,
		Desc:        "部署id, 部署时由系统传人",
		Example:     "deploy01",
		SearchLabel: true,
	})
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        job.SYSTEM_VARIABLE_APP_VERSION,
		Desc:        "应用部署时的版本",
		Example:     "v0.0.1",
		SearchLabel: true,
	})
	req.Spec.AddVersionParams(v1)

	ins, err := impl.UpdateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateBuildJob(t *testing.T) {
	req := job.NewPatchJobRequest(conf.C.BUILD_JOB_ID)
	req.Spec.RunnerSpec = tools.MustReadContentFile("test/container_build.yml")
	v1 := job.NewVersionedRunParam("v1")
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        "cluster_id",
		Desc:        "job运行时的k8s集群",
		Value:       "k8s-test",
		SearchLabel: true,
	})
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        "namespace",
		Desc:        "job运行时的namespace",
		Value:       "default",
		SearchLabel: true,
	})

	// 需要构建的代码信息
	v1.Add(&job.RunParam{
		Required:    true,
		Name:        "GIT_SSH_URL",
		Desc:        "应用git代码仓库地址",
		Example:     "git@github.com:infraboard/mpaas.git",
		SearchLabel: true,
	})
	v1.Add(&job.RunParam{
		Required: false,
		Name:     "APP_DOCKERFILE",
		Desc:     "应用git代码仓库中用于构建镜像的Dockerfile路径",
		Value:    "Dockerfile",
		Example:  "Dockerfile",
	})
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "GIT_BRANCH",
		Desc:     "需要拉去的代码分支",
		Example:  "master",
	})
	v1.Add(&job.RunParam{
		Required: true,
		Name:     "GIT_COMMIT_ID",
		Desc:     "应用git代码仓库地址",
		Example:  "32d63566098f7e0b0ac3a3d8ddffe71cc6cad7b0",
	})
	v1.Add(&job.RunParam{
		UsageType: job.PARAM_USAGE_TYPE_TEMPLATE,
		Name:      "GIT_SSH_SECRET",
		Desc:      "用于拉取git仓库代码的secret名称, kubectl create secret generic git-ssh-key --from-file=id_rsa=${HOME}/.ssh/id_rsa",
		Example:   "git-ssh-key",
		Value:     "git-ssh-key",
	})
	// docker push registry.cn-hangzhou.aliyuncs.com/inforboard/mpaas:[镜像版本号]
	v1.Add(&job.RunParam{
		Required: true,
		Name:     job.SYSTEM_VARIABLE_IMAGE_REPOSITORY,
		Desc:     "镜像推送地址",
		Example:  "registry.cn-hangzhou.aliyuncs.com/inforboard/mpaas",
	})
	v1.Add(&job.RunParam{
		Required: true,
		Name:     job.SYSTEM_VARIABLE_APP_VERSION,
		Desc:     "镜像版本",
		Example:  "v0.0.2",
	})
	v1.Add(&job.RunParam{
		UsageType: job.PARAM_USAGE_TYPE_TEMPLATE,
		Name:      "IMAGE_PUSH_SECRET",
		Desc:      "用于推送镜像的secret名称, 具体文档参考: https://github.com/GoogleContainerTools/kaniko#pushing-to-docker-hub",
		Example:   "kaniko-secret",
		Value:     "kaniko-secret",
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
