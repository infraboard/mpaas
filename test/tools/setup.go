package tools

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/conf"
	test "github.com/infraboard/mpaas/test/conf"

	// 注册所有服务
	_ "github.com/infraboard/mpaas/apps"
)

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 加载单元测试的变量
	test.LoadConfigFromEnv()

	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}
