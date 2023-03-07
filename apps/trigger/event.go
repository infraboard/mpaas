package trigger

import (
	"github.com/infraboard/mcenter/common/validate"
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
func (e *GitlabWebHookEvent) PipelineRunParams() (params []*job.RunParam) {
	switch e.EventType {
	case EVENT_TYPE_PUSH:
		repo := job.NewRunParam("GIT_REPOSITORY", e.Project.GitSshUrl)
		branche := job.NewRunParam("GIT_BRANCH", e.GetBranche())
		cm := e.GetLatestCommit()
		if cm != nil {
			commit := job.NewRunParam("GIT_COMMIT_ID", cm.Short())
			params = append(params, commit)
		}

		params = append(params, repo, branche)
	case EVENT_TYPE_TAG:
	case EVENT_TYPE_COMMENT:
	case EVENT_TYPE_MERGE_REQUEST:
	}

	return params
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
