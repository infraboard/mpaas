package workload_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	"sigs.k8s.io/yaml"
)

func TestListPod(t *testing.T) {
	req := meta.NewListRequest()
	req.Namespace = "default"
	req.Opts.LabelSelector = "job-name=test-job"
	pods, err := impl.ListPod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	for _, v := range pods.Items {
		t.Log(v.Namespace, v.Name)
	}
}

func TestGetPod(t *testing.T) {
	req := meta.NewGetRequest("kubernetes-proxy-78d4f87b58-crmlm")

	pods, err := impl.GetPod(ctx, req)
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
