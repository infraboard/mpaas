package impl

import (
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/pb/resource"
	"github.com/infraboard/mpaas/apps/pipeline"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *pipeline.QueryPipelineRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*pipeline.QueryPipelineRequest
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
	if r.IsTemplate != nil {
		filter["is_template"] = *r.IsTemplate
	}

	return bson.M{"$or": bson.A{
		filter,
		bson.M{"visiable_mode": resource.VISIABLE_GLOBAL},
		bson.M{"visiable_mode": resource.VISIABLE_DOMAIN, "domain": r.Scope.Domain},
	}}
}
