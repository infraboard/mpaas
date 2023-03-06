package job

const (
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
	// 部署工作负载类型
	SYSTEM_VARIABLE_WORKLOAD_KIND = "_WORKLOAD_KIND"
	// 部署名称
	SYSTEM_VARIABLE_WORKLOAD_NAME = "_WORKLOAD_NAME"
	// 部署服务名称
	SYSTEM_VARIABLE_SERVICE_NAME = "_SERVICE_NAME"
	// 服务部署镜像地址
	SYSTEM_VARIABLE_IMAGE_REPOSITORY = "IMAGE_REPOSITORY"
	// 服务部署镜像版本
	SYSTEM_VARIABLE_IMAGE_VERSION = "IMAGE_VERSION"
)
