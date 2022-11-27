package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/cluster"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) save(ctx context.Context, ins *cluster.Cluster) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted cluster(%s) document error, %s",
			ins.Data.Name, err)
	}
	return nil
}

func (s *service) get(ctx context.Context, id string) (*cluster.Cluster, error) {
	filter := bson.M{"_id": id}

	ins := cluster.NewDefaultCluster()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("cluster %s not found", id)
		}

		return nil, exception.NewInternalServerError("find cluster %s error, %s", id, err)
	}

	return ins, nil
}

func newQueryclusterRequest(r *cluster.QueryClusterRequest) *queryclusterRequest {
	return &queryclusterRequest{
		r,
	}
}

type queryclusterRequest struct {
	*cluster.QueryClusterRequest
}

func (r *queryclusterRequest) FindOptions() *options.FindOptions {
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

func (r *queryclusterRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Domain != "" {
		filter["data.domain"] = r.Domain
	}
	if r.Namespace != "" {
		filter["data.namespace"] = r.Namespace
	}
	if r.Vendor != "" {
		filter["data.vendor"] = r.Vendor
	}
	if r.Region != "" {
		filter["data.region"] = r.Region
	}

	if r.Keywords != "" {
		filter["$or"] = bson.A{
			bson.M{"data.name": bson.M{"$regex": r.Keywords, "$options": "im"}},
		}
	}
	return filter
}

func (s *service) query(ctx context.Context, req *queryclusterRequest) (*cluster.ClusterSet, error) {
	s.log.Debugf("find filter: %s, options: %s", req.FindFilter())

	resp, err := s.col.Find(ctx, req.FindFilter(), req.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find cluster error, error is %s", err)
	}

	set := cluster.NewClusterSet()
	// 循环
	for resp.Next(ctx) {
		ins := cluster.NewDefaultCluster()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode cluster error, error is %s", err)
		}

		ins.Desense()
		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get cluster count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) update(ctx context.Context, ins *cluster.Cluster) error {
	if _, err := s.col.UpdateByID(ctx, ins.Id, ins); err != nil {
		return exception.NewInternalServerError("inserted cluster(%s) document error, %s",
			ins.Data.Name, err)
	}

	return nil
}

func (s *service) deletecluster(ctx context.Context, ins *cluster.Cluster) error {
	if ins == nil || ins.Id == "" {
		return fmt.Errorf("cluster is nil")
	}

	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
	if err != nil {
		return exception.NewInternalServerError("delete cluster(%s) error, %s", ins.Id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("cluster %s not found", ins.Id)
	}

	return nil
}
