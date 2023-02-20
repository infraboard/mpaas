package trigger

import (
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mpaas/common/meta"
)

func (e *Event) Validate() error {
	return validate.Validate(e)
}

func (e *GitlabWebHookEvent) Validate() error {
	return validate.Validate(e)
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
