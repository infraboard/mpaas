package webhook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/webhook/dingding"
	"github.com/infraboard/mpaas/apps/task/webhook/feishu"
	"github.com/infraboard/mpaas/apps/task/webhook/wechat"
)

const (
	MAX_WEBHOOKS_PER_SEND = 12
)

const (
	feishuBot   = "feishu"
	dingdingBot = "dingding"
	wechatBot   = "wechat"
)

var (
	client = &http.Client{
		Timeout: 3 * time.Second,
	}
)

func newRequest(hook *pipeline.WebHook, t *task.JobTask) *request {
	// 因为AddWebhookStatus是非并非安全的， 因此不能放到Push(Push 是并发的)里面跑
	status := task.NewCallbackStatus(hook.ShowName())
	t.Status.AddWebhookStatus(status)

	return &request{
		hook:   hook,
		task:   t,
		status: status,
	}
}

type request struct {
	hook     *pipeline.WebHook
	task     *task.JobTask
	matchRes string

	status *task.CallbackStatus
}

func (r *request) Push() {
	// 准备请求,适配主流机器人
	var messageObj interface{}
	switch r.BotType() {
	case feishuBot:
		messageObj = r.NewFeishuMessage()
		r.matchRes = `"StatusCode":0,`
	case dingdingBot:
		messageObj = dingding.NewStepCardMessage(r.task)
		r.matchRes = `"errcode":0,`
	case wechatBot:
		messageObj = wechat.NewStepMarkdownMessage(r.task)
		r.matchRes = `"errcode":0,`
	default:
		messageObj = r.task
	}

	body, err := json.Marshal(messageObj)
	if err != nil {
		r.status.SendFailed("marshal step to json error, %s", err)
		return
	}

	req, err := http.NewRequest("POST", r.hook.Url, bytes.NewReader(body))
	if err != nil {
		r.status.SendFailed("new post request error, %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range r.hook.Header {
		req.Header.Add(k, v)
	}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		r.status.SendFailed("send request error, %s", err)
		return
	}
	defer resp.Body.Close()

	// 读取body
	bytesB, err := io.ReadAll(resp.Body)
	if err != nil {
		r.status.SendFailed("read response error, %s", err)
		return
	}
	respString := string(bytesB)

	if (resp.StatusCode / 100) != 2 {
		r.status.SendFailed("status code[%d] is not 200, response %s", resp.StatusCode, respString)
		return
	}

	// 通过返回匹配字符串来判断通知是否成功
	if r.matchRes != "" {
		if !strings.Contains(respString, r.matchRes) {
			r.status.SendFailed("reponse not match string %s, response: %s",
				r.matchRes, respString)
			return
		}
	}

	r.status.SendSuccess(respString)
}

func (r *request) BotType() string {
	if strings.HasPrefix(r.hook.Url, feishu.URL_PREFIX) {
		return feishuBot
	}
	if strings.HasPrefix(r.hook.Url, dingding.URL_PREFIX) {
		return dingdingBot
	}
	if strings.HasPrefix(r.hook.Url, wechat.URL_PREFIX) {
		return wechatBot
	}

	return ""
}

func (r *request) NewFeishuMessage() *feishu.Message {
	s := r.task
	msg := &feishu.NotifyMessage{
		Title:    s.ShowTitle(),
		Content:  s.String(),
		RobotURL: r.hook.Url,
		Note:     []string{"💡 该消息由极乐研发云提供"},
		Color:    feishu.COLOR_PURPLE,
	}
	return feishu.NewCardMessage(msg)
}
