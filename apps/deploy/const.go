package deploy

const (
	// 该注解标签 用于部署成功后 回调更新部署状态, 具体功能由operator实现
	ANNOTATION_DEPLOY_ID = "deploy.mpaas.infraboard.io/id"
)
