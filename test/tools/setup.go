package tools

import (
	"encoding/json"
	"io"
	"os"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/conf"
	"sigs.k8s.io/yaml"
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

func ReadContentFile(filepath string) ([]byte, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	payload, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func MustReadContentFile(filepath string) string {
	content, err := ReadContentFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func ReadJsonFile(filepath string, v any) error {
	content, err := ReadContentFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, v)
}

func ReadYamlFile(filepath string, v any) error {
	content, err := ReadContentFile(filepath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, v)
}

func MustReadYamlFile(filepath string, v any) {
	err := ReadYamlFile(filepath, v)
	if err != nil {
		panic(err)
	}
}
