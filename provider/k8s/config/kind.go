package config

const (
	// 普通配置
	CONFIG_KIND_CONFIG_MAP CONFIG_KIND = 0
	// 加密配置
	CONFIG_KIND_SECRET CONFIG_KIND = 1
)

type CONFIG_KIND int32

func (c CONFIG_KIND) String() string {
	return CONFIG_KIND_NAME[int32(c)]
}

// Enum value maps for WORKLOAD_KIND.
var (
	CONFIG_KIND_NAME = map[int32]string{
		0: "CONFIG_MAP",
		1: "SECRET",
	}
	CONFIG_KIND_VALUE = map[string]int32{
		"CONFIG_MAP": 0,
		"SECRET":     1,
	}
)
