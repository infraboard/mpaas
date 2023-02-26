package trigger

import (
	"github.com/infraboard/mcenter/common/validate"
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
