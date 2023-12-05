package upstream_test

import (
	"context"

	"github.com/infraboard/mpaas/apps/gateway/provider/apisix"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/upstream"
)

var (
	client *upstream.Client
	ctx    = context.Background()
)

func init() {
	conn := apisix.NewClient(
		"http://127.0.0.1:9180",
		"edd1c9f034335f136f87ad84b625c8f1",
	).Conn()

	client = upstream.NewClient(conn)
}
