package admin_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func TestListNode(t *testing.T) {
	v, err := impl.ListNode(ctx, meta.NewListRequest())
	if err != nil {
		t.Fatal(err)
	}
	for i := range v.Items {
		t.Log(v.Items[i].Name)
	}
}
