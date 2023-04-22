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
	t.Log(tools.MustToJson(ds))
}

func TestDescribeDeployment(t *testing.T) {
	req := deploy.NewDescribeDeploymentRequest(conf.C.DEPLOY_ID)
	ds, err := impl.DescribeDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ds.SystemVariable())
	t.Log(tools.MustToJson(ds))
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
	req.DeployId = "mongodb"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestCreateMcenterDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/mcenter_workload.yml")
	k8sConf.ClusterId = "k8s-test"

	req := deploy.NewCreateDeploymentRequest()
	req.K8STypeConfig = k8sConf
	req.ServiceId = conf.C.MCENTER_SERVICE_ID
	req.DeployId = "mcenter_v1"

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func TestUpdateDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/mcenter_workload.yml")
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

func TestQueryDeploymentInjectEnv(t *testing.T) {
	req := deploy.NewQueryDeploymentInjectEnvRequest(conf.C.DEPLOY_ID)
	env, err := impl.QueryDeploymentInjectEnv(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(env))
}

func TestDeleteDeployment(t *testing.T) {
	req := deploy.NewDeleteDeploymentRequest(conf.C.DEPLOY_ID)
	ds, err := impl.DeleteDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}
