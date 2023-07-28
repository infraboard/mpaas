package impl

import (
	"github.com/infraboard/mpaas/apps/trigger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *trigger.QueryRecordRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*trigger.QueryRecordRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "time", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}
	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ServiceId != "" {
		filter["token"] = r.ServiceId
	}

	if r.PipelineTaskId != "" {
		filter["build_status.pipeline_task_id"] = r.PipelineTaskId
	}

	return filter
}
