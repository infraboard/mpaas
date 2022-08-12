package k8s_test

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/infraboard/mpaas/provider/k8s"
	"sigs.k8s.io/yaml"
)

var (
	client *k8s.Client
	ctx    = context.Background()
)

func TestServerVersion(t *testing.T) {
	v, err := client.ServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)

	rs, err := client.ServerResources()
	if err != nil {
		t.Log(err)
	}
	t.Log(rs)
}

func TestListNamespace(t *testing.T) {
	v, err := client.ListNamespace(ctx, k8s.NewListRequest())
	if err != nil {
		t.Log(err)
	}
	for i := range v.Items {
		t.Log(v.Items[i].Name)
	}
}

func TestListDeployment(t *testing.T) {
	v, err := client.ListDeployment(ctx, k8s.NewListRequest())
	if err != nil {
		t.Log(err)
	}
	for i := range v.Items {
		t.Log(v.Items[i].Name)
	}
}

func TestGetDeployment(t *testing.T) {
	v, err := client.GetDeployment(ctx, k8s.NewGetRequest("nginx"))
	if err != nil {
		t.Log(err)
	}

	yd, err := yaml.Marshal(v)
	if err != nil {
		t.Log(err)
	}

	t.Log(string(yd))
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	kc, err := ioutil.ReadFile(filepath.Join(wd, "kube_config.yml"))
	if err != nil {
		panic(err)
	}

	client, err = k8s.NewClient(string(kc))
	if err != nil {
		panic(err)
	}

}
