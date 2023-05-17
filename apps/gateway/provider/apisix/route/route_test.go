package route_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/route"
)

func TestQueryRoute(t *testing.T) {
	in := route.NewQueryRouteRequest()
	list, err := client.QueryRoute(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}
