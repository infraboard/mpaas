package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	conf *Config
)

// C 全局配置对象
func C() *Config {
	if conf == nil {
		panic("Load Config first")
	}
	return conf
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	conf = newConfig()
	if _, err := toml.DecodeFile(filePath, conf); err != nil {
		return err
	}
	return nil
}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	conf = newConfig()
	if err := env.Parse(conf); err != nil {
		return err
	}
	return nil
}
