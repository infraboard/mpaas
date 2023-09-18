package deploy

const (
	// 该注解标签 用于部署成功后 回调更新部署状态, 具体功能由operator实现
	ANNOTATION_DEPLOY_ID = "deploy.mpaas.infraboard.io/id"

	// k8s 部署时的标签
	LABEL_SERVICE_NAME_KEY = "devcloud.service"
	LABEL_NAMESPACE_KEY    = "devcloud.namespace"
	LABEL_CLUSTER_KEY      = "devcloud.cluster"
	LABEL_DEPLOY_GROUP_KEY = "devcloud.group"
	LABEL_DEPLOY_ID_KEY    = "devcloud.deploy"
)
