package batch_test

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/batch"
)

var (
	impl *batch.Batch
	ctx  = context.Background()
)

func init() {
	client, err := k8s.NewClientFromFile("../kube_config.yml")
	if err != nil {
		panic(err)
	}
	impl = client.Batch()
}
