package job

const (
	UNIQ_NAME_SPLITER      = "@"
	UNIQ_NAMESPACE_SPLITER = "."
)

// 系统变量统一以_开头, 系统变量由Runner处理并注入
type SYSTEM_VARIABLE string

const (
	// 部署配置ID, runner执行时，会挂载DEPLOY_CONFIG中配置的集群ID相应的kubeconf文件
	SYSTEM_VARIABLE_DEPLOY_CONFIG_ID = "_DEPLOY_CONFIG_ID"
	// 任务运行的Pipeline Task Id, 由Pipeline 允许时创建, runner注入
	SYSTEM_VARIABLE_PIPELINE_TASK_ID = "_PIPELINE_TASK_ID"
)
