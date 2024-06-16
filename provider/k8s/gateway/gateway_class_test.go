package gateway_test

import (
	"testing"

	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func TestListGatewayClass(t *testing.T) {
	req := meta.NewListRequest()
	v, err := impl.ListGatewayClass(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pretty.MustToYaml(v))
}
