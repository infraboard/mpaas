package admin_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/meta"

	corev1 "k8s.io/api/core/v1"
)

func TestListNamespace(t *testing.T) {
	v, err := impl.ListNamespace(ctx, meta.NewListRequest())
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
	v, err := impl.CreateNamespace(ctx, ns)
	if err != nil {
		t.Log(err)
	}
	t.Log(v.Name)
}
