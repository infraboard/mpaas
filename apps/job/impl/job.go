package impl

import (
	"context"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mpaas/apps/job"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		ins.AddExtension()
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
	case job.DESCRIBE_BY_JOB_ID:
		filter["_id"] = in.DescribeValue
	case job.DESCRIBE_BY_JOB_UNIQ_NAME:
		version, name, ns, domain := job.ParseUniqName(in.DescribeValue)
		filter["name"] = name
		filter["namespace"] = ns
		filter["domain"] = domain
		if version != "" {
			if version == job.LATEST_VERSION_NAME {
				filter["status.stage"] = job.JOB_STAGE_PUBLISHED
			} else {
				filter["status.version"] = version
			}
		}
	}

	opt := &options.FindOneOptions{
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
	}
	ins := job.NewDefaultJob()
	if err := i.col.FindOne(ctx, filter, opt).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("job %s not found", in)
		}

		return nil, exception.NewInternalServerError("find job %s error, %s", in.DescribeValue, err)
	}

	ins.AddExtension()
	return ins, nil
}

func (i *impl) UpdateJob(ctx context.Context, in *job.UpdateJobRequest) (
	*job.Job, error) {
	ins, err := i.DescribeJob(ctx, job.NewDescribeJobRequestByName("#"+in.Id))
	if err != nil {
		return nil, err
	}

	switch in.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(in)
	case request.UpdateMode_PATCH:
		err := ins.Patch(in)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	ins.Spec.BuildSearchLabels()
	if err := ins.Spec.Validate(); err != nil {
		return nil, err
	}

	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) UpdateJobStatus(ctx context.Context, in *job.UpdateJobStatusRequest) (
	*job.Job, error) {
	ins, err := i.DescribeJob(ctx, job.NewDescribeJobRequestByName(in.Id))
	if err != nil {
		return nil, err
	}

	if err := in.Validate(); err != nil {
		return nil, err
	}

	if in.Status.Stage < ins.Status.Stage {
		return nil, exception.NewBadRequest("状态不合法")
	}

	switch in.Status.Stage {
	case job.JOB_STAGE_PUBLISHED:
		ins.Status.PublishedAt = time.Now().Unix()
	}

	ins.Status = in.Status
	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) DeleteJob(ctx context.Context, in *job.DeleteJobRequest) (
	*job.Job, error) {
	// 查询删除Job
	ins, err := i.DescribeJob(ctx, job.NewDescribeJobRequestByName("#"+in.Id))
	if err != nil {
		return nil, err
	}

	// 检查是否允许删除
	err = ins.CheckAllowDelete()
	if err != nil {
		return nil, err
	}

	_, err = i.col.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete  job(%s) error, %s", in.Id, err)
	}
	return ins, nil
}
