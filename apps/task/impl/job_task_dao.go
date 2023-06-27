package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *task.QueryJobTaskRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*task.QueryJobTaskRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	token.MakeMongoFilter(filter, r.Scope)
	policy.MakeMongoFilter(filter, "labels", r.Filters)

	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}

	if r.PipelineTaskId != "" {
		filter["pipeline_task"] = r.PipelineTaskId
	}

	if r.Stage != nil {
		filter["status.stage"] = *r.Stage
	}

	if r.HasLabel() {
		for k, v := range r.Labels {
			filter["labels."+k] = v
		}
	}

	return filter
}

func (i *impl) updateJobTask(ctx context.Context, ins *task.JobTask) error {
	// 更新数据库
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.TaskId, bson.M{"$set": ins}); err != nil {
		return exception.NewInternalServerError("update task(%s) document error, %s",
			ins.Spec.TaskId, err)
	}
	return nil
}
