package impl_test

import (
	"os"
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
	req := approval.NewDescribeApprovalRequest(os.Getenv("APPROVAL_ID"))
	ins, err := impl.DescribeApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestEditApproval(t *testing.T) {
	req := approval.NewEditApprovalRequest(os.Getenv("APPROVAL_ID"))
	ins, err := impl.EditApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateApprovalStatus(t *testing.T) {
	req := approval.NewUpdateApprovalStatusRequest(os.Getenv("APPROVAL_ID"))
	ins, err := impl.UpdateApprovalStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateApproval(t *testing.T) {
	req := approval.NewCreateApprovalRequest()
	req.Version = "1.0.0"
	req.Describe = "发布说明, 支持Markdown语法"
	req.AddProposer("admin")
	req.AddAuditor("admin")
	tools.MustReadYamlFile("test/create.yml", req.PipelineSpec)
	set, err := impl.CreateApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
