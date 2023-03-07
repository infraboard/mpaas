package trigger

import (
	"fmt"
	"time"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/common/meta"
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

// Event产生的参数, 作用于Pipeline运行
// GIT_REPOSITORY: git@github.com:infraboard/mpaas.git
// GIT_BRANCH: master
// GIT_COMMIT_ID: bfacd86c647935aea532f29421fe83c6a6111260
func (e *GitlabWebHookEvent) GitRunParams() (params []*job.RunParam) {
	switch e.EventType {
	case EVENT_TYPE_PUSH:
		repo := job.NewRunParam(job.SYSTEM_VARIABLE_GIT_REPOSITORY, e.Project.GitSshUrl)
		branche := job.NewRunParam(job.SYSTEM_VARIABLE_GIT_BRANCH, e.GetBranche())
		params = append(params, repo, branche)

		cm := e.GetLatestCommit()
		if cm != nil {
			commit := job.NewRunParam(job.SYSTEM_VARIABLE_GIT_COMMIT_ID, cm.Id)
			params = append(params, commit)
		}
	case EVENT_TYPE_TAG:
	case EVENT_TYPE_COMMENT:
	case EVENT_TYPE_MERGE_REQUEST:
	}

	return params
}

func (e *GitlabWebHookEvent) VersionRunParam(r build.VERSION_NAMED_RULE) *job.RunParam {
	version := e.GenBuildVersion(r)
	return job.NewRunParam(job.SYSTEM_VARIABLE_IMAGE_VERSION, version)
}

func (e *GitlabWebHookEvent) GenBuildVersion(r build.VERSION_NAMED_RULE) string {
	switch r {
	case build.VERSION_NAMED_RULE_DATE_BRANCH_COMMIT:
		return fmt.Sprintf("%s-%s-%s",
			time.Now().Format("20060102"),
			e.GetBranche(),
			e.GetLatestCommitShortId(),
		)
	}
	return ""
}

func NewRecord(e *Event) *Record {
	return &Record{
		Meta:        meta.NewMeta(),
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
