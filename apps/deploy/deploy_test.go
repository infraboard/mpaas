package deploy_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/test/tools"
)

func TestObjectMeta(t *testing.T) {
	k8sConf := deploy.NewK8STypeConfig()
	k8sConf.WorkloadConfig = tools.MustReadContentFile("impl/test/deployment.yml")
	k8sConf.ClusterId = "k8s-test"
	m, err := k8sConf.DeploySystemVaraible("nginx")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(m))
}
