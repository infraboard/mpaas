package trigger

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/logger/zap"
	build "github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/common/format"
	"github.com/infraboard/mpaas/conf"
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

func (e *Event) ToJson() string {
	return format.Prettify(e)
}

func (e *Event) Validate() error {
	return validate.Validate(e)
}

func (e *Event) UUID() string {
	return fmt.Sprintf("event-%s", e.Id)
}

func ParseGitLabEventFromRequest(r *restful.Request) (*Event, error) {
	// 读取body数据
	body, err := io.ReadAll(r.Request.Body)
	defer r.Request.Body.Close()
	if err != nil {
		return nil, err
	}
	e := NewGitlabEvent(string(body))

	// 读取Header中的数据
	e.Name = r.HeaderParameter(GITLAB_HEADER_EVENT_NAME)
	e.Id = r.PathParameter(GITLAB_HEADER_EVENT_UUID)
	e.Token = r.PathParameter(GITLAB_HEADER_EVENT_TOKEN)
	e.From = r.PathParameter(GITLAB_HEADER_INSTANCE)
	e.UserAgent = r.Request.UserAgent()

	// 读取URL参数
	e.SkipRunPipeline = r.QueryParameter("skip_run_pipeline") == "true"
	return e, nil
}

func (e *GitlabWebHookEvent) DefaultRepository() string {
	return fmt.Sprintf("%s/%s",
		conf.C().Image.DefaultRegistry,
		e.Project.NamespacePath,
	)
}

func (e *GitlabWebHookEvent) Validate() error {
	return validate.Validate(e)
}

func (e *GitlabWebHookEvent) ToJson() string {
	return format.Prettify(e)
}

func (e *GitlabWebHookEvent) ParseEventType(eventHeaderName string) {
	switch eventHeaderName {
	case "Push Hook":
		e.EventType = GITLAB_EVENT_TYPE_PUSH
	case "Tag Push Hook":
		e.EventType = GITLAB_EVENT_TYPE_TAG
	case "Merge Request Hook":
		e.EventType = GITLAB_EVENT_TYPE_MERGE_REQUEST
	case "Note Hook":
		e.EventType = GITLAB_EVENT_TYPE_COMMENT
	case "Issue Hook":
		e.EventType = GITLAB_EVENT_TYPE_ISSUE
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
func (e *Event) GitRunParams() *job.VersionedRunParam {
	params := job.NewVersionedRunParam("v1")
	params.Add(
		// 补充gitlab事件相关变量
		job.NewRunParam(VARIABLE_EVENT_PROVIDER, EVENT_PROVIDER_GITLAB.String()),
		job.NewRunParam(VARIABLE_EVENT_NAME, e.Name),
		job.NewRunParam(VARIABLE_EVENT_INSTANCE, e.From),
		job.NewRunParam(VARIABLE_EVENT_TOKEN, e.Token),
		job.NewRunParam(VARIABLE_EVENT_USER_AGENT, e.UserAgent),
		job.NewRunParam(VARIABLE_EVENT_CONTENT, e.Raw),
	)

	// 补充GITLAB事件相关变量
	switch e.Provider {
	case EVENT_PROVIDER_GITLAB:
		ge, err := e.GetGitlabEvent()
		if err != nil {
			zap.L().Errorf("parse gitlab event error, %s", err)
		} else {
			ge.GitRunParams(params)
		}
	}

	return params
}

func (e *Event) GetGitlabEvent() (*GitlabWebHookEvent, error) {
	if !e.Provider.Equal(EVENT_PROVIDER_GITLAB) {
		return nil, fmt.Errorf("not gitlab")
	}

	if e.Raw == "" {
		return nil, fmt.Errorf("event raw data not found")
	}

	event := NewGitlabWebHookEvent()
	err := json.Unmarshal([]byte(e.Raw), event)
	if err != nil {
		return nil, err
	}
	event.ParseEventType(e.From)
	return event, nil
}

func (e *GitlabWebHookEvent) GitRunParams(params *job.VersionedRunParam) {
	// 补充项目相关信息
	params.Add(
		job.NewRunParam(VARIABLE_GIT_PROJECT_NAME, e.Project.Name).SetSearchLabel(true),
		job.NewRunParam(VARIABLE_GIT_SSH_URL, e.Project.GitSshUrl).SetSearchLabel(true),
		job.NewRunParam(VARIABLE_GIT_HTTP_URL, e.Project.GitHttpUrl).SetSearchLabel(true),
	)

	switch e.EventType {
	case GITLAB_EVENT_TYPE_PUSH:
		params.Add(
			job.NewRunParam(VARIABLE_GIT_BRANCH, e.GetBranch()),
		)
		cm := e.GetLatestCommit()
		if cm != nil {
			params.Add(job.NewRunParam(VARIABLE_GIT_COMMIT, cm.Id))
		}
	case GITLAB_EVENT_TYPE_TAG:
		params.Add(
			job.NewRunParam(VARIABLE_GIT_TAG, e.GetTag()),
		)
	case GITLAB_EVENT_TYPE_MERGE_REQUEST:
		oa := e.ObjectAttributes
		params.Add(
			job.NewRunParam(VARIABLE_GIT_MR_ACTION, oa.Action),
			job.NewRunParam(VARIABLE_GIT_MR_STATUS, oa.MergeStatus),
			job.NewRunParam(VARIABLE_GIT_MR_SOURCE_BRANCE, oa.SourceBranch),
			job.NewRunParam(VARIABLE_GIT_MR_TARGET_BRANCE, oa.TargetBranch),
		)
		if oa.LastCommit != nil {
			params.Add(job.NewRunParam(VARIABLE_GIT_COMMIT, oa.LastCommit.Id))
		}
	case GITLAB_EVENT_TYPE_COMMENT:
	case GITLAB_EVENT_TYPE_ISSUE:
	}
}

func (e *GitlabWebHookEvent) DateCommitVersion(prefix string) *job.RunParam {
	version := e.GenBuildVersion()
	if !strings.HasPrefix(version, prefix) {
		version = prefix + version
	}
	return job.NewRunParam(build.SYSTEM_VARIABLE_APP_VERSION, version)
}

func (e *GitlabWebHookEvent) TagVersion(prefix string) *job.RunParam {
	version := e.GetTag()
	if !strings.HasPrefix(version, prefix) {
		version = prefix + version
	}
	return job.NewRunParam(build.SYSTEM_VARIABLE_APP_VERSION, version)
}

func (e *GitlabWebHookEvent) GenBuildVersion() string {
	return fmt.Sprintf("%s-%s-%s",
		time.Now().Format("20060102"),
		e.GetBranch(),
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
