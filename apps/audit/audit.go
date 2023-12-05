package audit

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	resource "github.com/infraboard/mcube/v2/pb/resource"
)

func NewRecordSet() *RecordSet {
	return &RecordSet{
		Items: []*Record{},
	}
}

func NewDefaultRecord() *Record {
	return &Record{}
}

func (s *RecordSet) Add(items ...*Record) {
	s.Items = append(s.Items, items...)
}

func NewSaveRecordRequest() *SaveRecordRequest {
	return &SaveRecordRequest{
		Labels: map[string]string{},
		Extra:  map[string]string{},
	}
}

func (r *SaveRecordRequest) Validate() error {
	return validator.Validate(r)
}

func New(req *SaveRecordRequest) (*Record, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	ins := &Record{
		Meta:  resource.NewMeta(),
		Scope: resource.NewScope(),
		Spec:  req,
	}
	return ins, nil
}
