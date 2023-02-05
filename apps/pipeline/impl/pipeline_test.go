package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/pipeline"
)

func TestQueryDeploy(t *testing.T) {
	req := pipeline.NewQueryPipelineRequest()
	set, err := impl.QueryPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
