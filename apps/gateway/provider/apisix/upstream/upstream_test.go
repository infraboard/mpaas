package upstream_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/upstream"
)

func TestDescribeUpstream(t *testing.T) {
	in := upstream.NewDescribeUpstreamRequest("1")
	ins, err := client.DescribeUpstream(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
