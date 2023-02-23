package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryDeploy(t *testing.T) {
	req := deploy.NewQueryDeployConfigRequest()
	ds, err := impl.QueryDeployConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestCreateDeployConfig(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	k8sConf.ClusterId = "test-k8s"

	req := deploy.NewCreateDeployConfigRequest()
	req.K8STypeConfig = k8sConf
	req.Provider = "腾讯云"
	req.Region = "上海"
	req.Environment = "生产"

	ds, err := impl.CreateDeployConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestUpdateDeployConfig(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	req := deploy.NewPatchDeployRequest(os.Getenv("DEPLOY_JOB_ID"))
	req.Spec.K8STypeConfig.ClusterId = "k8s-test"
	ds, err := impl.UpdateDeployConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestDeleteDeployConfig(t *testing.T) {
	req := deploy.NewDeleteDeployConfigRequest(os.Getenv("DEPLOY_JOB_ID"))
	ds, err := impl.DeleteDeployConfig(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}
