package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/test/tools"
)

func TestHandleEvent(t *testing.T) {
	req := trigger.NewDefaultWebHookEvent()
	err := tools.ReadJsonFile("test/webhook.json", req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req)

	ps, err := impl.HandleGitlabEvent(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ps)
}
