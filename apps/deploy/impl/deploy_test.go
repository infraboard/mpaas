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
	req.Scope.Domain = domain.DEFAULT_DOMAIN
	ds, err := impl.QueryDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}

func TestDescribeDeployment(t *testing.T) {
	req := deploy.NewDescribeDeploymentRequest(conf.C.MCENTER_DEPLOY_ID)
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
	k8sConf.ClusterId = "docker-desktop"

	req := deploy.NewCreateDeploymentRequest()
	req.Cluster = "2a90c4eec422c171"
	req.ServiceName = "mongodb"
	req.K8STypeConfig = k8sConf
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}

func TestCreateMcenterDeployment(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("test/mcenter_workload.yml")
	k8sConf.ClusterId = "docker-desktop"

	req := deploy.NewCreateDeploymentRequest()
	req.K8STypeConfig = k8sConf
	req.Cluster = "6d6586cacdc2aa95"

	ds, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
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
	req := deploy.NewUpdateDeploymentStatusRequest(conf.C.MCENTER_DEPLOY_ID)
	req.UpdatedK8SConfig = k8sConf
	ds, err := impl.UpdateDeploymentStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}

func TestQueryDeploymentInjectEnv(t *testing.T) {
	req := deploy.NewQueryDeploymentInjectEnvRequest(conf.C.MCENTER_DEPLOY_ID)
	env, err := impl.QueryDeploymentInjectEnv(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(env))

	// 测试Group匹配
	for i := range env.EnvGroups {
		group := env.EnvGroups[i]
		t.Log(group.IsLabelMatched(map[string]string{}))
		t.Log(tools.MustToJson(group.ToConfigMap()))
	}
}

func TestDeleteDeployment(t *testing.T) {
	req := deploy.NewDeleteDeploymentRequest(conf.C.MCENTER_DEPLOY_ID)
	ds, err := impl.DeleteDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ds))
}
