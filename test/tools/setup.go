package tools

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/conf"
)

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}
