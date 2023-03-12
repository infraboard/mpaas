package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/test/conf"
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

func TestDescribeDeployment(t *testing.T) {
	req := deploy.NewDescribeDeploymentRequest(conf.C.DEPLOY_ID)
	ds, err := impl.DescribeDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ds.SystemVariable())

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
	req.ServiceId = conf.C.SERVICE_ID
	req.DeployId = "deploy01"

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestCreateMongoDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/mongodb_workload.yml")
	k8sConf.Service = tools.MustReadContentFile("test/mongodb_service.yml")
	k8sConf.ClusterId = "k8s-test"

	req := deploy.NewCreateDeploymentRequest()
	req.Kind = deploy.KIND_MIDDLEWARE
	req.ServiceName = "mongodb"
	req.K8STypeConfig = k8sConf
	req.Provider = "腾讯云"
	req.Region = "上海"
	req.Environment = "生产"
	req.DeployId = "mongodb"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestUpdateDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	req := deploy.NewPatchDeployRequest(conf.C.DEPLOY_ID)
	req.Spec.K8STypeConfig.ClusterId = "k8s-test"
	ds, err := impl.UpdateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestUpdateDeploymentStatus(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/deployment.yml")
	req := deploy.NewUpdateDeploymentStatusRequest(conf.C.DEPLOY_ID)
	req.UpdatedK8SConfig = k8sConf
	ds, err := impl.UpdateDeploymentStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestDeleteDeployment(t *testing.T) {
	req := deploy.NewDeleteDeploymentRequest(conf.C.DEPLOY_ID)
	ds, err := impl.DeleteDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}
