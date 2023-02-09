package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/pipeline"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 查询Pipeline列表
func (i *impl) QueryPipeline(ctx context.Context, in *pipeline.QueryPipelineRequest) (
	*pipeline.PipelineSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find pipeline error, error is %s", err)
	}

	set := pipeline.NewPipelineSet()
	// 循环
	for resp.Next(ctx) {
		ins := pipeline.NewDefaultPipeline()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode pipeline error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

// 查询Pipeline详情
func (i *impl) DescribePipeline(ctx context.Context, in *pipeline.DescribePipelineRequest) (
	*pipeline.Pipeline, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := pipeline.NewDefaultPipeline()
	if err := i.col.FindOne(ctx, bson.M{"_id": in.Id}).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("pipeline %s not found", in.Id)
		}

		return nil, exception.NewInternalServerError("find pipeline %s error, %s", in.Id, err)
	}

	return ins, nil
}

// 创建Pipeline
func (i *impl) CreatePipeline(ctx context.Context, in *pipeline.CreatePipelineRequest) (
	*pipeline.Pipeline, error) {
	ins, err := pipeline.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a pipeline document error, %s", err)
	}
	return ins, nil
}

// 更新Pipeline
func (i *impl) UpdatePipeline(ctx context.Context, in *pipeline.UpdatePipelineRequest) (
	*pipeline.Pipeline, error) {
	return nil, nil
}

// 删除Pipeline
func (i *impl) DeletePipeline(ctx context.Context, in *pipeline.DeletePipelineRequest) (
	*pipeline.Pipeline, error) {
	req := pipeline.NewDescribePipelineRequest(in.Id)
	ins, err := i.DescribePipeline(ctx, req)
	if err != nil {
		return nil, err
	}

	_, err = i.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete pipeline(%s) error, %s", in.Id, err)
	}
	return ins, nil
}
