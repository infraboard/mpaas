package workload_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func TestListJob(t *testing.T) {
	req := meta.NewListRequest()
	req.Namespace = "kube-system"
	req.Opts.LabelSelector = "k8s-app=kube-dns"
	list, err := impl.ListJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	for _, v := range list.Items {
		t.Log(v.Namespace, v.Name)
	}
}
