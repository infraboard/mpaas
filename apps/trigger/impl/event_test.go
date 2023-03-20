package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestHandleEvent(t *testing.T) {
	event := trigger.NewGitlabWebHookEvent()
	err := tools.ReadJsonFile("test/gitlab_push.json", event)
	if err != nil {
		t.Fatal(err)
	}

	req := trigger.NewGitlabEvent(conf.C.SERVICE_ID, event)
	req.SkipRunPipeline = false
	ps, err := impl.HandleEvent(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ps))
}

func TestQueryRecord(t *testing.T) {
	req := trigger.NewQueryRecordRequest()
	set, err := impl.QueryRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
