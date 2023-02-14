package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryApproval(t *testing.T) {
	req := approval.NewQueryApprovalRequest()
	set, err := impl.QueryApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestDescribeApproval(t *testing.T) {
	req := approval.NewDescribeApprovalRequest("xx")
	ins, err := impl.DescribeApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateApproval(t *testing.T) {
	req := approval.NewCreateApprovalRequest()
	tools.MustReadYamlFile("test/create.yml", req.DeployPipelineSpec)
	set, err := impl.CreateApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
