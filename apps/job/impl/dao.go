package impl

import (
	"github.com/infraboard/mpaas/apps/job"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *job.QueryJobRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*job.QueryJobRequest
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
	filter := bson.M{"$or": bson.A{}}

	if r.VisiableMode != nil {
		filter["spec.visiable_mode"] = *r.VisiableMode
	}

	if r.Domain != "" {
		filter["spec.domain"] = r.Domain
	}

	if r.Namespace != "" {
		filter["spec.namespace"] = r.Namespace
	}

	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}

	if len(r.Names) > 0 {
		filter["spec.name"] = bson.M{"$in": r.Names}
	}

	return filter
}
