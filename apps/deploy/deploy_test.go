package deploy_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	"github.com/infraboard/mpaas/test/tools"
)

func TestObjectMeta(t *testing.T) {
	wc := deploy.NewK8STypeConfig()
	wc.WorkloadConfig = tools.MustReadContentFile("impl/test/deployment.yml")
	wc.ClusterId = "k8s-test"

	wl, err := workload.ParseWorkloadFromYaml(wc.WorkloadKind, wc.WorkloadConfig)
	if err != nil {
		t.Fatal(err)
	}
	v := wl.SystemVaraible("nginx")
	t.Log(tools.MustToYaml(v))
}
