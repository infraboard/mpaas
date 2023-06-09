package trigger

// Event 相关变量
const (
	// 事件提供商, 比如 GITLAB
	VARIABLE_EVENT_PROVIDER = "EVENT_PROVIDER"
	// 事件描述, 比如 Push Hook
	VARIABLE_EVENT_NAME = "EVENT_NAME"
	// 提供事件的实例 比如 "https://gitlab.com"
	VARIABLE_EVENT_INSTANCE = "EVENT_INSTANCE"
	// 事件发生者UA 比如: "GitLab/15.5.0-pre"
	VARIABLE_EVENT_USER_AGENT = "EVENT_USER_AGENT"
	// 事件Token, 这里固定为service id
	VARIABLE_EVENT_TOKEN = "EVENT_TOKEN"
	// 事件的具体内容
	VARIABLE_EVENT_CONTENT = "EVENT_CONTENT"
)

// Git 相关变量
const (
	// GitLab相关事件变量
	VARIABLE_GIT_PROJECT_NAME = "GIT_PROJECT_NAME"
	// Git 项目代码仓库SSH地址
	VARIABLE_GIT_SSH_URL = "GIT_SSH_URL"
	// Git 项目代码仓库HTTP地址
	VARIABLE_GIT_HTTP_URL = "GIT_HTTP_URL"

	// PUSH
	// Git 项目分支
	VARIABLE_GIT_BRANCH = "GIT_BRANCH"
	// Git 项目Commit号
	VARIABLE_GIT_COMMIT = "GIT_COMMIT_ID"

	// TAG
	// Git 项目Tag
	VARIABLE_GIT_TAG = "GIT_TAG"

	// MR
	// Git MR状态
	VARIABLE_GIT_MR_STATUS = "GIT_MR_STATUS"
	// Git MR动作
	VARIABLE_GIT_MR_ACTION = "GIT_MR_ACTION"
	// Git MR source
	VARIABLE_GIT_MR_SOURCE_BRANCE = "GIT_MR_SOURCE_BRANCH"
	// Git MR target
	VARIABLE_GIT_MR_TARGET_BRANCE = "GIT_MR_TARGET_BRANCH"
)

const (
	GITLAB_HEADER_EVENT_NAME     = "X-Gitlab-Event"
	GITLAB_HEADER_EVENT_TOKEN    = "X-Gitlab-Token"
	GITLAB_HEADER_INSTANCE       = "X-Gitlab-Instance"
	GITLAB_HEADER_EVENT_UUID     = "X-Gitlab-Event-UUID"
	GITLAB_HEADER_EVENT_UUID_OLD = "X-Gitlab-Event-Uuid"
)
