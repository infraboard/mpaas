package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/audit"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryRecord(t *testing.T) {
	req := audit.NewQueryRecordRequest()
	ds, err := impl.QueryRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}

func TestSaveRecord(t *testing.T) {
	req := audit.NewSaveRecordRequest()
	ds, err := impl.SaveRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}
