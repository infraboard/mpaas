package webhook_test

import (
	"context"
	"testing"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/apps/task/webhook"
)

var (
	feishuBotURL = "https://open.feishu.cn/open-apis/bot/v2/hook/461ead7b-d856-472c-babc-2d3d0ec9fabb"
	dingBotURL   = "https://oapi.dingtalk.com/robot/send?access_token=xxxx"
	wechatBotURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
)

func TestFeishuWebHook(t *testing.T) {
	should := assert.New(t)

	hooks := testPipelineWebHook(feishuBotURL)
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

	hooks := testPipelineWebHook(dingBotURL)
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

	hooks := testPipelineWebHook(wechatBotURL)
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
	return &task.JobTask{}
}

func init() {
	zap.DevelopmentSetup()
}
