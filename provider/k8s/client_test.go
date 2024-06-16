package k8s_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	client *k8s.Client
)

func TestServerVersion(t *testing.T) {
	v, err := client.ServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestServerResources(t *testing.T) {
	rs := client.ServerResources()

	for i := range rs.Items {
		item := rs.Items[i]
		t.Log(item.GroupVersion, item.APIVersion)
		for _, r := range item.APIResources {
			t.Log(r)
		}
	}
}

func init() {
	kubeConf := tools.MustReadContentFile("kube_config.yml")
	c, err := k8s.NewClient(kubeConf)
	if err != nil {
		panic(err)
	}
	client = c
}
