package webhook_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/apps/task/webhook"
)

func TestFeishuWebHook(t *testing.T) {
	should := assert.New(t)

	hooks := testPipelineWebHook(os.Getenv("FEISHU_BOT_URL"))
	sender := webhook.NewWebHook()
	err := sender.Send(
		context.Background(),
		hooks,
		testPipelineStep(),
	)
	should.NoError(err)
	t.Log(hooks[0])
}

func TestDingDingWebHook(t *testing.T) {
	should := assert.New(t)

	hooks := testPipelineWebHook(os.Getenv("DINGDING_BOT_URL"))
	sender := webhook.NewWebHook()
	err := sender.Send(
		context.Background(),
		hooks,
		testPipelineStep(),
	)
	should.NoError(err)

	t.Log(hooks[0])
}

func TestWechatWebHook(t *testing.T) {
	should := assert.New(t)

	hooks := testPipelineWebHook(os.Getenv("WECHAT_BOT_URL"))
	sender := webhook.NewWebHook()
	err := sender.Send(
		context.Background(),
		hooks,
		testPipelineStep(),
	)
	should.NoError(err)
	t.Log(hooks[0])
}

func testPipelineWebHook(url string) []*pipeline.WebHook {
	h1 := &pipeline.WebHook{
		Url:         url,
		Events:      []string{task.STAGE_SUCCEEDED.String()},
		Description: "测试",
	}
	return []*pipeline.WebHook{h1}
}

func testPipelineStep() *task.JobTask {
	t := task.NewJobTask(pipeline.NewRunJobRequest("test"))
	return t
}

func init() {
	zap.DevelopmentSetup()
}
