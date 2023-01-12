package k8s_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/provider/k8s"
	corev1 "k8s.io/api/core/v1"
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

func TestListNode(t *testing.T) {
	v, err := client.ListNode(ctx, k8s.NewListRequest())
	if err != nil {
		t.Fatal(err)
	}
	for i := range v.Items {
		t.Log(v.Items[i].Name)
	}
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

func TestCreateNamespace(t *testing.T) {
	ns := &corev1.Namespace{}
	ns.Name = "go8"
	v, err := client.CreateNamespace(ctx, ns)
	if err != nil {
		t.Log(err)
	}
	t.Log(v.Name)
}

func TestListPod(t *testing.T) {
	req := k8s.NewListRequest()
	req.Namespace = "kube-system"
	req.Opts.LabelSelector = "k8s-app=kube-dns"
	pods, err := client.ListPod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	for _, v := range pods.Items {
		t.Log(v.Namespace, v.Name)
	}
}

func TestGetPod(t *testing.T) {
	req := k8s.NewGetRequest("kubernetes-proxy-78d4f87b58-crmlm")

	pods, err := client.GetPod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	yd, err := yaml.Marshal(pods)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(yd))
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
