package gateway_test

import (
	"testing"

	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func TestListGateway(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListGateway(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pretty.MustToYaml(v))
}
