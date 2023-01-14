package config_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func TestListConfigMap(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListConfigMap(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}
