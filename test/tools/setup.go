package tools

import (
	"encoding/json"
	"io"
	"os"

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

func ReadJsonFile(filepath string, v any) error {
	fd, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fd.Close()

	payload, err := io.ReadAll(fd)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, v)
}
