package apps

import (
	// 内置健康检查
	_ "github.com/infraboard/mcube/app/health/api"

	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/infraboard/mpaas/apps/build/api"
	_ "github.com/infraboard/mpaas/apps/deploy/api"
	_ "github.com/infraboard/mpaas/apps/gateway/api"
	_ "github.com/infraboard/mpaas/apps/job/api"
	_ "github.com/infraboard/mpaas/apps/k8s/api"
	_ "github.com/infraboard/mpaas/apps/pipeline/api"
	_ "github.com/infraboard/mpaas/apps/task/api"
	_ "github.com/infraboard/mpaas/apps/trigger/api"
)
