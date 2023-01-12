package k8s_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/provider/k8s"
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
	rs, err := client.ServerResources()
	if err != nil {
		t.Log(err)
	}
	for i := range rs {
		t.Log(rs[i].GroupVersion, rs[i].APIVersion)
		for _, r := range rs[i].APIResources {
			t.Log(r.Name)
		}
	}
}

func init() {
	zap.DevelopmentSetup()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	kc, err := os.ReadFile(filepath.Join(wd, "kube_config.yml"))
	if err != nil {
		panic(err)
	}

	client, err = k8s.NewClient(string(kc))
	if err != nil {
		panic(err)
	}

}
