package impl

import (
	"github.com/infraboard/mpaas/apps/build"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *build.QueryBuildConfigRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*build.QueryBuildConfigRequest
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

	if len(r.ServiceIds) > 0 {
		filter["spec.service_id"] = bson.M{"$in": r.ServiceIds}
	}

	return filter
}
