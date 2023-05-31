package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/test/conf"
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

func TestCreateTestJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	req.Name = "test"
	req.CreateBy = "test"
	req.RunnerSpec = tools.MustReadContentFile("test/test.yml")
	param := job.NewRunParamSet()
	param.Add(&job.RunParam{
		Required: true,
		Name:     "cluster_id",
		NameDesc: "job运行时的k8s集群",
		Value:    "k8s-test",
	})
	param.Add(&job.RunParam{
		Required: true,
		Name:     "namespace",
		NameDesc: "job运行时的namespace",
		Value:    "default",
	})
	req.RunParam = param

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
	req.Spec.Labels["Language"] = "*"
	param := job.NewRunParamSet()
	param.Add(&job.RunParam{
		Required:    true,
		Name:        "cluster_id",
		NameDesc:    "job运行时的k8s集群",
		Value:       "k8s-test",
		SearchLabel: true,
	})
	param.Add(&job.RunParam{
		Required:    true,
		Name:        "namespace",
		NameDesc:    "job运行时的namespace",
		Value:       "default",
		SearchLabel: true,
	})

	// 部署运行时变量
	param.Add(&job.RunParam{
		Required:    true,
		Name:        job.SYSTEM_VARIABLE_DEPLOY_ID,
		NameDesc:    "部署id, 部署时由系统传人",
		Example:     "deploy01",
		SearchLabel: true,
	})
	param.Add(&job.RunParam{
		Required:    true,
		Name:        build.SYSTEM_VARIABLE_IMAGE_REPOSITORY,
		NameDesc:    "应用部署的镜像仓库地址",
		Example:     "registry.cn-hangzhou.aliyuncs.com/infraboard/mcenter",
		SearchLabel: true,
	})
	param.Add(&job.RunParam{
		Required:    true,
		Name:        build.SYSTEM_VARIABLE_APP_VERSION,
		NameDesc:    "应用部署时的版本",
		Example:     "v0.0.1",
		SearchLabel: true,
	})
	req.Spec.RunParam = param

	ins, err := impl.UpdateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateBuildJob(t *testing.T) {
	req := job.NewPatchJobRequest(conf.C.BUILD_JOB_ID)
	req.Spec.RunnerSpec = tools.MustReadContentFile("test/container_build.yml")
	req.Spec.Labels["Language"] = "*"
	param := job.NewRunParamSet()
	param.Add(&job.RunParam{
		Required:    true,
		Name:        "cluster_id",
		NameDesc:    "job运行时的k8s集群",
		Value:       "k8s-test",
		SearchLabel: true,
	})
	param.Add(&job.RunParam{
		Required:    true,
		Name:        "namespace",
		NameDesc:    "job运行时的namespace",
		Value:       "default",
		SearchLabel: true,
	})

	// 需要构建的代码信息
	param.Add(&job.RunParam{
		Required:    true,
		Name:        "GIT_SSH_URL",
		NameDesc:    "应用git代码仓库地址",
		Example:     "git@github.com:infraboard/mpaas.git",
		SearchLabel: true,
	})
	param.Add(&job.RunParam{
		Required: false,
		Name:     "APP_DOCKERFILE",
		NameDesc: "应用git代码仓库中用于构建镜像的Dockerfile路径",
		Value:    "Dockerfile",
		Example:  "Dockerfile",
	})
	param.Add(&job.RunParam{
		Required: true,
		Name:     "GIT_BRANCH",
		NameDesc: "需要拉去的代码分支",
		Example:  "master",
	})
	param.Add(&job.RunParam{
		Required: true,
		Name:     "GIT_COMMIT_ID",
		NameDesc: "应用git代码仓库地址",
		Example:  "32d63566098f7e0b0ac3a3d8ddffe71cc6cad7b0",
	})
	param.Add(&job.RunParam{
		UsageType: job.PARAM_USAGE_TYPE_TEMPLATE,
		Name:      "GIT_SSH_SECRET",
		NameDesc:  "用于拉取git仓库代码的secret名称, kubectl create secret generic git-ssh-key --from-file=id_rsa=${HOME}/.ssh/id_rsa",
		Example:   "git-ssh-key",
		Value:     "git-ssh-key",
	})
	// docker push registry.cn-hangzhou.aliyuncs.com/infraboard/mpaas:[镜像版本号]
	param.Add(&job.RunParam{
		Required: true,
		Name:     build.SYSTEM_VARIABLE_IMAGE_REPOSITORY,
		NameDesc: "镜像推送地址",
		Example:  "registry.cn-hangzhou.aliyuncs.com/infraboard/mpaas",
	})
	param.Add(&job.RunParam{
		Required: true,
		Name:     build.SYSTEM_VARIABLE_APP_VERSION,
		NameDesc: "镜像版本",
		Example:  "v0.0.2",
	})
	param.Add(&job.RunParam{
		UsageType: job.PARAM_USAGE_TYPE_TEMPLATE,
		Name:      "IMAGE_PUSH_SECRET",
		NameDesc:  "用于推送镜像的secret名称, 具体文档参考: https://github.com/GoogleContainerTools/kaniko#pushing-to-docker-hub",
		Example:   "kaniko-secret",
		Value:     "kaniko-secret",
	})
	req.Spec.RunParam = param

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

func TestUpdateJobStatus(t *testing.T) {
	req := job.NewUpdateJobStatusRequest("docker_build@default.default")
	req.Status.Stage = job.JOB_STAGE_PUBLISHED
	req.Status.Version = "v1"
	ins, err := impl.UpdateJobStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
