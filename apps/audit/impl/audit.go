package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/audit"
)

// 保存审计日志
func (i *impl) SaveRecord(ctx context.Context, in *audit.SaveRecordRequest) (
	*audit.Record, error) {
	ins, err := audit.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a audit record document error, %s", err)
	}
	return ins, nil
}

// 查询升级日志
func (i *impl) QueryRecord(ctx context.Context, in *audit.QueryRecordRequest) (
	*audit.RecordSet, error) {
	r := newQueryRequest(in)
	i.log.Debugf("filter: %s", r.FindFilter())
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find record error, error is %s", err)
	}

	set := audit.NewRecordSet()
	// 循环
	for resp.Next(ctx) {
		ins := audit.NewDefaultRecord()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode record error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get record count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
