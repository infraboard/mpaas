package trigger

import (
	"fmt"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/rs/xid"
)

func NewRecordSet() *RecordSet {
	return &RecordSet{
		Items: []*Record{},
	}
}

func (s *RecordSet) Add(items ...*Record) {
	s.Items = append(s.Items, items...)
}

func (e *Event) Validate() error {
	return validate.Validate(e)
}

func (e *GitlabWebHookEvent) Validate() error {
	return validate.Validate(e)
}

func (e *GitlabWebHookEvent) ParseInfoFromHeader(r *restful.Request) {
	e.EventToken = r.HeaderParameter(GITLAB_HEADER_EVENT_TOKEN)
	e.Instance = r.HeaderParameter(GITLAB_HEADER_INSTANCE)
	e.UserAgent = r.HeaderParameter("User-Agent	")
	e.ParseEventType(r.HeaderParameter(GITLAB_HEADER_EVENT_NAME))
}

func (e *GitlabWebHookEvent) ParseEventType(et string) {
	e.EventDescribe = et
	switch et {
	case "Push Hook":
		e.EventType = EVENT_TYPE_PUSH
	case "Tag Push Hook":
		e.EventType = EVENT_TYPE_TAG
	case "Merge Request Hook":
		e.EventType = EVENT_TYPE_MERGE_REQUEST
	case "Note Hook":
		e.EventType = EVENT_TYPE_COMMENT
	case "Issue Hook":
		e.EventType = EVENT_TYPE_ISSUE
	}
}

// Event产生的事件参数, 作用于Pipeline运行
// 事件通用变量:
// EVENT_PROVIDER: GITLAB
// EVENT_TYPE: PUSH
// EVENT_DESC: Push Hook
// EVENT_INSTANCE: "https://gitlab.com"
// EVENT_USER_AGENT: "GitLab/15.5.0-pre"
// EVENT_TOKEN
//
// PUSH事件变量:
// GIT_SSH_URL: git@github.com:infraboard/mpaas.git
// GIT_BRANCH: master
// GIT_COMMIT_ID: bfacd86c647935aea532f29421fe83c6a6111260
func (e *GitlabWebHookEvent) GitRunParams() *job.VersionedRunParam {
	params := job.NewVersionedRunParam("v1")
	params.Add(
		// 补充gitlab事件相关变量
		job.NewRunParam(VARIABLE_EVENT_PROVIDER, EVENT_PROVIDER_GITLAB.String()),
		job.NewRunParam(VARIABLE_EVENT_TYPE, e.EventType.String()),
		job.NewRunParam(VARIABLE_EVENT_DESC, e.EventDescribe),
		job.NewRunParam(VARIABLE_EVENT_INSTANCE, e.Instance),
		job.NewRunParam(VARIABLE_EVENT_TOKEN, e.EventToken),
		job.NewRunParam(VARIABLE_EVENT_USER_AGENT, e.UserAgent),
		job.NewRunParam(VARIABLE_EVENT_CONTENT, e.EventRaw),
		// 补充项目相关信息
		job.NewRunParam(VARIABLE_GIT_PROJECT_NAME, e.Project.Name),
		job.NewRunParam(VARIABLE_GIT_SSH_URL, e.Project.GitSshUrl),
		job.NewRunParam(VARIABLE_GIT_HTTP_URL, e.Project.GitHttpUrl),
	)

	switch e.EventType {
	case EVENT_TYPE_PUSH:
		params.Add(
			job.NewRunParam(VARIABLE_GIT_BRANCH, e.GetBaseRef()),
		)
		cm := e.GetLatestCommit()
		if cm != nil {
			params.Add(job.NewRunParam(VARIABLE_GIT_COMMIT, cm.Id))
		}
	case EVENT_TYPE_TAG:
		params.Add(
			job.NewRunParam(VARIABLE_GIT_TAG, e.GetBaseRef()),
		)
	case EVENT_TYPE_MERGE_REQUEST:
		oa := e.ObjectAttributes
		params.Add(
			job.NewRunParam(VARIABLE_GIT_MR_ACTION, oa.Action),
			job.NewRunParam(VARIABLE_GIT_MR_STATUS, oa.MergeStatus),
			job.NewRunParam(VARIABLE_GIT_MR_SOURCE_BRANCE, oa.SourceBranch),
			job.NewRunParam(VARIABLE_GIT_MR_TARGET_BRANCE, oa.TargetBranch),
		)
		if e.LastCommit != nil {
			params.Add(job.NewRunParam(VARIABLE_GIT_COMMIT, e.LastCommit.Id))
		}
	case EVENT_TYPE_COMMENT:
	case EVENT_TYPE_ISSUE:
	}

	return params
}

func (e *GitlabWebHookEvent) DateCommitVersion(prefix string) *job.RunParam {
	version := e.GenBuildVersion()
	if !strings.HasPrefix(version, prefix) {
		version = prefix + version
	}
	return job.NewRunParam(job.SYSTEM_VARIABLE_APP_VERSION, version)
}

func (e *GitlabWebHookEvent) TagVersion(prefix string) *job.RunParam {
	version := e.GetBaseRef()
	if !strings.HasPrefix(version, prefix) {
		version = prefix + version
	}
	return job.NewRunParam(job.SYSTEM_VARIABLE_APP_VERSION, version)
}

func (e *GitlabWebHookEvent) GenBuildVersion() string {
	return fmt.Sprintf("%s-%s-%s",
		time.Now().Format("20060102"),
		e.GetBaseRef(),
		e.GetLatestCommitShortId(),
	)
}

func NewRecord(e *Event) *Record {
	if e.Id == "" {
		e.Id = xid.New().String()
	}
	return &Record{
		Event:       e,
		BuildStatus: []*BuildStatus{},
	}
}

func (e *Record) AddBuildStatus(bs *BuildStatus) {
	e.BuildStatus = append(e.BuildStatus, bs)
}

func NewDefaultRecord() *Record {
	return NewRecord(&Event{})
}
