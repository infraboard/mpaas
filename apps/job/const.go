package job

const (
	UNIQ_VERSION_SPLITER   = ":"
	UNIQ_NAME_SPLITER      = "@"
	UNIQ_NAMESPACE_SPLITER = "."
)

// 系统变量统一以_开头, 系统变量由Runner处理并注入
type SYSTEM_VARIABLE string

const (
	// 部署配置ID, runner执行时，会挂载Deploy中配置的集群ID相应的kubeconf文件
	SYSTEM_VARIABLE_DEPLOY_ID = "_DEPLOY_ID"
	// 任务运行的Pipeline Task Id, 由Pipeline 运行时创建, runner注入
	SYSTEM_VARIABLE_PIPELINE_TASK_ID = "_PIPELINE_TASK_ID"
	// 任务运行的Job Task Id, 由Job 运行时创建, runner注入
	SYSTEM_VARIABLE_JOB_TASK_ID = "_JOB_TASK_ID"
	// 任务运行的Job Id, 由Job 运行时创建, runner注入
	SYSTEM_VARIABLE_JOB_ID = "_JOB_ID"
	// 用于Task内部使用Update Token 回写OUTPUT任务的输出信息
	SYSTEM_VARIABLE_JOB_TASK_UPDATE_TOKEN = "_JOB_TASK_UPDATE_TOKEN"
	// 部署工作负载类型
	SYSTEM_VARIABLE_WORKLOAD_KIND = "_WORKLOAD_KIND"
	// 部署名称
	SYSTEM_VARIABLE_WORKLOAD_NAME = "_WORKLOAD_NAME"
	// 部署服务名称
	SYSTEM_VARIABLE_SERVICE_NAME = "_SERVICE_NAME"
)

const (
	// 获取最新发布的版本
	LATEST_VERSION_NAME = "latest"
)
