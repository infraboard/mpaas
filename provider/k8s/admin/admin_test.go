package admin_test

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/admin"
)

var (
	impl *admin.Admin
	ctx  = context.Background()
)

func init() {
	client, err := k8s.NewClientFromFile("../kube_config.yml")
	if err != nil {
		panic(err)
	}
	impl = client.Admin()
}
