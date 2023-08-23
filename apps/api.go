package apps

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/infraboard/mpaas/apps/cluster/api"
	_ "github.com/infraboard/mpaas/apps/deploy/api"
	_ "github.com/infraboard/mpaas/apps/gateway/api"
	_ "github.com/infraboard/mpaas/apps/k8s/api"
	_ "github.com/infraboard/mpaas/apps/proxy/api"
)
