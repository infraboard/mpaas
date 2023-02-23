package conf

import "github.com/caarlos0/env/v6"

var (
	C *Config
)

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() {
	C = newConfig()
	if err := env.Parse(C); err != nil {
		panic(err)
	}
}
