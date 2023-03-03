package workload_test

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	test "github.com/infraboard/mpaas/test/conf"
)

var (
	impl *workload.Client
	ctx  = context.Background()
)

func init() {
	client, err := k8s.NewClientFromFile("../kube_config.yml")
	if err != nil {
		panic(err)
	}
	// 加载单元测试的变量
	test.LoadConfigFromEnv()
	impl = client.WorkLoad()
}
