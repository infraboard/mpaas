package trigger

const (
	// GitLab相关事件变量
	// 事件提供商
	SYSTEM_VARIABLE_EVENT_PROVIDER = "EVENT_PROVIDER"
	// 事件提供商
	SYSTEM_VARIABLE_EVENT_TYPE = "EVENT_TYPE"
)

const (
	GITLAB_HEADER_EVENT       = "X-Gitlab-Event"
	GITLAB_HEADER_EVENT_TOKEN = "X-Gitlab-Token"
	GITLAB_HEADER_INSTANCE    = "X-Gitlab-Instance"
	GITLAB_HEADER_EVENT_UUID  = "X-Gitlab-Event-UUID"
)
