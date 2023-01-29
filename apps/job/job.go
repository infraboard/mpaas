package job

import (
	"time"

	"github.com/rs/xid"
)

// New 新建一个部署配置
func New(req *CreateJobRequest) (*Job, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Job{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return d, nil
}

func NewJobSet() *JobSet {
	return &JobSet{
		Items: []*Job{},
	}
}

func (s *JobSet) Add(item *Job) {
	s.Items = append(s.Items, item)
}

func NewDefaultJob() *Job {
	return &Job{
		Spec: NewCreateJobRequest(),
	}
}
