package impl

import (
	"github.com/infraboard/mpaas/apps/cluster"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *cluster.QueryClusterRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*cluster.QueryClusterRequest
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

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if r.Label != nil && len(r.Label) > 0 {
		for k, v := range r.Label {
			filter["labels."+k] = v
		}
	}

	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}

	if len(r.Names) > 0 {
		filter["name"] = bson.M{"$in": r.Names}
	}

	return filter
}
