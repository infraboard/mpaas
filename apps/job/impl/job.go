package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/job"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	return set, nil
}

func (i *impl) DescribeJob(ctx context.Context, in *job.DescribeJobRequest) (
	*job.Job, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch in.DescribeBy {
	case job.DESCRIBE_BY_ID:
		filter["_id"] = in.DescribeValue
	}

	ins := job.NewDefaultJob()
	if err := i.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("job %s not found", in)
		}

		return nil, exception.NewInternalServerError("find job %s error, %s", in.DescribeValue, err)
	}

	return ins, nil
}
