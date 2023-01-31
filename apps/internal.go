package apps

import (
	// 监控检查实现
	_ "github.com/infraboard/mpaas/apps/health/impl"

	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/infraboard/mpaas/apps/build/impl"
	_ "github.com/infraboard/mpaas/apps/cluster/impl"
	_ "github.com/infraboard/mpaas/apps/deploy/impl"
	_ "github.com/infraboard/mpaas/apps/gateway/impl"
	_ "github.com/infraboard/mpaas/apps/job/impl"
	_ "github.com/infraboard/mpaas/apps/storage/impl"
	_ "github.com/infraboard/mpaas/apps/trigger/impl"
)
