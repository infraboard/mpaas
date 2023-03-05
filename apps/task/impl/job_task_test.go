package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunBuildJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_build@default.default")
	version := job.NewVersionedRunParam("v1")
	version.Params = job.NewRunParamWithKVPaire(
		"GIT_REPOSITORY", "git@github.com:infraboard/mpaas.git",
		"GIT_BRANCH", "master",
		"GIT_COMMIT_ID", "57612b40df7fc9619ddc537e3dc117ab335ed294",
		job.SYSTEM_VARIABLE_IMAGE_REPOSITORY, "registry.cn-hangzhou.aliyuncs.com/inforboard/mpaas",
		job.SYSTEM_VARIABLE_IMAGE_VERSION, "v0.0.5",
	)
	req.RunParams = version

	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestRunDeployJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_deploy@default.default")
	version := job.NewVersionedRunParam("v1")
	version.Params = job.NewRunParamWithKVPaire(
		job.SYSTEM_VARIABLE_DEPLOY_ID, conf.C.DEPLOY_ID,
		job.SYSTEM_VARIABLE_IMAGE_VERSION, "1.30",
	)
	req.RunParams = version
	req.DryRun = false

	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins.Status.Detail))
}

func TestQueryJobTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	req.PipelineTaskId = conf.C.PIPELINE_TASK_ID
	set, err := impl.QueryJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tools.MustToJson(set))
}

func TestUpdateJobTaskOutput(t *testing.T) {
	req := task.NewUpdateJobTaskOutputRequest(conf.C.JOB_TASK_ID)
	req.UpdateToken = conf.C.JOB_TASK_TOKEN
	req.AddRuntimeEnv(job.SYSTEM_VARIABLE_IMAGE_VERSION, "v0.0.5")
	req.MarkdownOutput = "构建产物描述信息"
	ins, err := impl.UpdateJobTaskOutput(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestUpdateJobTaskStatus(t *testing.T) {
	req := task.NewUpdateJobTaskStatusRequest(conf.C.JOB_TASK_ID)
	req.Stage = task.STAGE_SUCCEEDED
	req.Message = "执行成功"
	req.Force = true
	req.UpdateToken = conf.C.JOB_TASK_TOKEN
	req.Detail = tools.MustReadContentFile("test/k8s_job.yml")
	ins, err := impl.UpdateJobTaskStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestDescribeJobTask(t *testing.T) {
	req := task.NewDescribeJobTaskRequest(conf.C.JOB_TASK_ID)
	ins, err := impl.DescribeJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteJobTask(t *testing.T) {
	req := task.NewDeleteJobTaskRequest(conf.C.JOB_TASK_ID)
	req.Force = true
	set, err := impl.DeleteJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
