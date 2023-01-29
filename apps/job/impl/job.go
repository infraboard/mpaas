package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/job"
)

func (i *impl) CreateJob(ctx context.Context, in *job.CreateJobRequest) (
	*job.Job, error) {
	ins, err := job.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a job document error, %s", err)
	}
	return ins, nil
}

func (i *impl) QueryJob(ctx context.Context, in *job.QueryJobRequest) (
	*job.JobSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find job error, error is %s", err)
	}

	set := job.NewJobSet()
	// 循环
	for resp.Next(ctx) {
		ins := job.NewDefaultJob()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode job error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get job count error, error is %s", err)
	}
	set.Total = count

	// 补充全局的
	if in.WithGlobal {

	}

	return set, nil
}
