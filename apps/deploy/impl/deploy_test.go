package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryDeploy(t *testing.T) {
	req := deploy.NewQueryDeploymentRequest()
	ds, err := impl.QueryDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestCreateDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	k8sConf.ClusterId = "k8s-test"

	req := deploy.NewCreateDeploymentRequest()
	req.K8STypeConfig = k8sConf
	req.Provider = "腾讯云"
	req.Region = "上海"
	req.Environment = "生产"

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestUpdateDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	req := deploy.NewPatchDeployRequest(os.Getenv("DEPLOY_JOB_ID"))
	req.Spec.K8STypeConfig.ClusterId = "k8s-test"
	ds, err := impl.UpdateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestDeleteDeployment(t *testing.T) {
	req := deploy.NewDeleteDeploymentRequest(os.Getenv("DEPLOY_JOB_ID"))
	ds, err := impl.DeleteDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}
