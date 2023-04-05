package conf

func newConfig() *Config {
	return &Config{}
}

type Config struct {
	DEPLOY_ID        string `env:"DEPLOY_ID"`
	BUILD_ID         string `env:"BUILD_ID"`
	MCENTER_BUILD_ID string `env:"MCENTER_BUILD_ID"`
	SERVICE_ID       string `env:"SERVICE_ID"`
	FEISHU_BOT_URL   string `env:"FEISHU_BOT_URL"`
	DINGDING_BOT_URL string `env:"DINGDING_BOT_URL"`
	WECHAT_BOT_URL   string `env:"WECHAT_BOT_URL"`
	JOB_TASK_ID      string `env:"JOB_TASK_ID"`
	JOB_TASK_TOKEN   string `env:"JOB_TASK_TOKEN"`
	PIPELINE_TASK_ID string `env:"PIPELINE_TASK_ID"`

	DEPLOY_JOB_ID     string `env:"DEPLOY_JOB_ID"`
	BUILD_JOB_ID      string `env:"BUILD_JOB_ID"`
	CICD_PIPELINE_ID  string `env:"CICD_PIPELINE_ID"`
	MPAAS_PIPELINE_ID string `env:"MPAAS_PIPELINE_ID"`
	APPROVAL_ID       string `env:"APPROVAL_ID"`

	MCENTER_SERVICE_ID    string `env:"MCENTER_SERVICE_ID"`
	MCENTER_GRPC_ADDRESS  string `env:"MCENTER_GRPC_ADDRESS"`
	MCENTER_CLINET_ID     string `env:"MCENTER_CLINET_ID"`
	MCENTER_CLIENT_SECRET string `env:"MCENTER_CLIENT_SECRET"`
}
