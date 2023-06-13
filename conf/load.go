package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"

	"github.com/infraboard/mcenter/clients/rpc"
)

var (
	conf *Config
)

// C 全局配置对象
func C() *Config {
	if conf == nil {
		panic("Load Config first")
	}

	// 提前加载好 mcenter客户端
	err := rpc.LoadClientFromConfig(conf.Mcenter)
	if err != nil {
		panic("load mcenter client from config error: " + err.Error())
	}
	return conf
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	conf = newConfig()
	if _, err := toml.DecodeFile(filePath, conf); err != nil {
		return err
	}

	return conf.InitGlobal()
}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	conf = newConfig()
	if err := env.Parse(conf); err != nil {
		return err
	}
	return conf.InitGlobal()
}
