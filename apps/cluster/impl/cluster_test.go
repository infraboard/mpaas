package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryCluster(t *testing.T) {
	req := cluster.NewQueryClusterRequest()
	ds, err := impl.QueryCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}
