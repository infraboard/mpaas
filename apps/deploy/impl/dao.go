package impl

import (
	"context"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/deploy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *deploy.QueryDeploymentRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*deploy.QueryDeploymentRequest
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
	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}
	if len(r.ServiceIds) > 0 {
		filter["service_id"] = bson.M{"$in": r.ServiceIds}
	}
	if len(r.Environments) > 0 {
		filter["environment"] = bson.M{"$in": r.Environments}
	}
	if len(r.Clusters) > 0 {
		filter["cluster"] = bson.M{"$in": r.Clusters}
	}
	return filter
}

func (i *impl) update(ctx context.Context, ins *deploy.Deployment) error {
	ins.Meta.UpdateAt = time.Now().Unix()
	_, err := i.col.UpdateOne(ctx, bson.M{"_id": ins.Meta.Id}, bson.M{"$set": ins})
	if err != nil {
		return exception.NewInternalServerError("update deploy(%s) error, %s", ins.Meta.Id, err)
	}
	return nil
}
