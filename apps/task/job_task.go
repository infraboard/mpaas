package task

import (
	"fmt"
	"time"

	"github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	"github.com/rs/xid"
)

func NewJobTaskSet() *JobTaskSet {
	return &JobTaskSet{
		Items: []*JobTask{},
	}
}

func (s *JobTaskSet) Add(tasks ...*JobTask) {
	s.Items = append(s.Items, tasks...)
}

func (s *JobTaskSet) Len() int {
	return len(s.Items)
}

func NewDefaultJobTask() *JobTask {
	req := pipeline.NewRunJobRequest("")
	return NewJobTask(req)
}

func NewJobTask(req *pipeline.RunJobRequest) *JobTask {
	if req.Id == "" {
		req.Id = xid.New().String()
	}

	return &JobTask{
		Id:       req.Id,
		CreateAt: time.Now().Unix(),
		Spec:     req,
		Job:      nil,
		Status:   NewJobTaskStatus(),
	}
}

func (t *JobTask) GetStatusDetail() string {
	if t.Status != nil {
		return t.Status.Detail
	}

	return ""
}

func (t *JobTask) Update(job *job.Job, status *JobTaskStatus) {
	t.Job = job
	t.Status = status
}

func (s *JobTask) ShowTitle() string {
	return fmt.Sprintf("任务[%s]当前状态: %s", s.Spec.Job, s.Status.Stage.String())
}

func NewJobTaskStatus() *JobTaskStatus {
	return &JobTaskStatus{
		StartAt: time.Now().Unix(),
	}
}

func (t *JobTaskStatus) IsComplete() bool {
	return t.Stage > 10
}

func (t *JobTaskStatus) Update(req *UpdateJobTaskStatusRequest) {
	t.Stage = req.Stage
	t.Message = req.Message

	// 取消的任务 不需要更新detail详情
	if !t.Stage.Equal(STAGE_CANCELED) {
		t.Detail = req.Detail
	}

	// 结束时标记结束时间
	if t.IsComplete() {
		t.EndAt = time.Now().Unix()
	}
}
