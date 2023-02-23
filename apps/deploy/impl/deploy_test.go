package impl_test

import (
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
	t.Log(ds)
}

func TestCreateDeployConfig(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	k8sConf.ClusterName = "test-k8s"

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
