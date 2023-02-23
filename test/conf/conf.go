package conf

func newConfig() *Config {
	return &Config{}
}

type Config struct {
	BUILD_CONFIG_ID  string `env:"BUILD_CONFIG_ID"`
	DEPLOY_CONFIG_ID string `env:"DEPLOY_CONFIG_ID"`
	SERVICE_ID       string `env:"SERVICE_ID"`

	FEISHU_BOT_URL   string `env:"FEISHU_BOT_URL"`
	DINGDING_BOT_URL string `env:"DINGDING_BOT_URL"`
	WECHAT_BOT_URL   string `env:"WECHAT_BOT_URL"`
	JOB_TASK_ID      string `env:"JOB_TASK_ID"`
	PIPELINE_TASK_ID string `env:"PIPELINE_TASK_ID"`

	DEPLOY_JOB_ID string `env:"DEPLOY_JOB_ID"`
	BUILD_JOB_ID  string `env:"BUILD_JOB_ID"`
	PIPELINE_ID   string `env:"PIPELINE_ID"`
	APPROVAL_ID   string `env:"APPROVAL_ID"`

	MCENTER_GRPC_ADDRESS  string `env:"MCENTER_GRPC_ADDRESS"`
	MCENTER_CLINET_ID     string `env:"MCENTER_CLINET_ID"`
	MCENTER_CLIENT_SECRET string `env:"MCENTER_CLIENT_SECRET"`
}
