package job

const (
	UNIQ_NAME_SPLITER      = "@"
	UNIQ_NAMESPACE_SPLITER = "."
)

// 系统变量统一以_开头
type SYSTEM_VARIABLE string

const (
	// 部署配置ID, runner执行时，会挂载DEPLOY_CONFIG中配置的集群ID相应的kubeconf文件
	SYSTEM_VARIABLE_DEPLOY_CONFIG_ID = "_DEPLOY_CONFIG_ID"
)
