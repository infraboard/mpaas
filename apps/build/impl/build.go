package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/build"
)

func (i *impl) CreateBuildConfig(ctx context.Context, in *build.CreateBuildConfigRequest) (
	*build.BuildConfig, error) {
	ins, err := build.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a build document error, %s", err)
	}
	return ins, nil
}

func (i *impl) QueryBuildConfig(ctx context.Context, in *build.QueryBuildConfigRequest) (
	*build.BuildConfigSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find build error, error is %s", err)
	}

	set := build.NewBuildConfigSet()
	// 循环
	for resp.Next(ctx) {
		ins := build.NewDefaultDeploy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode build error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get build count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *impl) UpdateBuildConfig(ctx context.Context, in *build.UpdateBuildConfigRequest) (
	*build.BuildConfig, error) {
	return nil, nil
}
