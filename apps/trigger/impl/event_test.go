package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/test/tools"
)

func TestHandleEvent(t *testing.T) {
	event := trigger.NewGitlabWebHookEvent()
	err := tools.ReadJsonFile("test/webhook.json", event)
	if err != nil {
		t.Fatal(err)
	}

	req := trigger.NewGitlabEvent(os.Getenv("SERVICE_ID"), event)
	ps, err := impl.HandleEvent(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ps)
}
