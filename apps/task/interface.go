package task

import (
	context "context"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/tools/pretty"
	job "github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/common/terminal"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/core/v1"
)

const (
	AppName = "tasks"
)

type Service interface {
	JobService
	PipelineService
}

type JobService interface {
	// 执行Job
	RunJob(context.Context, *pipeline.RunJobRequest) (*JobTask, error)
	// 删除任务
	DeleteJobTask(context.Context, *DeleteJobTaskRequest) (*JobTask, error)
	// 任务Debug
	JobTaskDebug(context.Context, *JobTaskDebugRequest)
	JobRPCServer
}

func NewJobTaskDebugRequest(taskId string) *JobTaskDebugRequest {
	return &JobTaskDebugRequest{
		TaskId: taskId,
	}
}

type JobTaskDebugRequest struct {
	// 任务Id
	TaskId string `json:"task_id"`
	// 容器名称
	ContainerName string `json:"container_name"`
	// 终端配置
	Terminal TerminalConfig `json:"terminal"`
	// Debug容器终端
	terminal *terminal.WebSocketTerminal
}

func (r *JobTaskDebugRequest) CopyPodRunRequest(namespace, podName string) *workload.CopyPodRunRequest {
	req := workload.NewCopyPodRunRequest()
	req.SourcePod.Name = podName
	req.SourcePod.Namespace = namespace
	req.TargetPodMeta.Namespace = namespace
	req.TargetPodMeta.Name = fmt.Sprintf("%s-debug", req.SourcePod.Name)
	req.ExecContainer = r.ContainerName
	req.ExecHoldCmd = workload.HoldContaienrCmd(1 * time.Hour)
	return req
}

type TerminalConfig struct {
	// 终端的宽度
	// @gotags: json:"width"
	Width uint16 `json:"width"`
	// 终端的高度
	// @gotags: json:"heigh"
	Heigh uint16 `json:"heigh"`
}

func (r *JobTaskDebugRequest) WebTerminal() *terminal.WebSocketTerminal {
	return r.terminal
}

func (r *JobTaskDebugRequest) SetWebTerminal(term *terminal.WebSocketTerminal) {
	r.terminal = term
	r.terminal.SetSize(r.Terminal.Width, r.Terminal.Heigh)
}

func NewQueryTaskRequest() *QueryJobTaskRequest {
	return &QueryJobTaskRequest{
		Page: request.NewDefaultPageRequest(),
		Ids:  []string{},
	}
}

func (req *QueryJobTaskRequest) HasLabel() bool {
	return req.Labels != nil && len(req.Labels) > 0
}

func NewRunTaskRequest(name, spec string, params *job.RunParamSet) *RunTaskRequest {
	return &RunTaskRequest{
		Name:    name,
		JobSpec: spec,
		Params:  params,
	}
}

// Job运行时需要的注解
func (r *RunTaskRequest) Annotations() map[string]string {
	annotations := map[string]string{}
	if !r.ManualUpdateStatus {
		annotations[ANNOTATION_TASK_ID] = r.Params.GetParamValue(job.SYSTEM_VARIABLE_JOB_TASK_ID)
	}
	return annotations
}

func (r *RunTaskRequest) RenderJobSpec() string {
	renderedSpec := r.JobSpec
	vars := r.Params.TemplateVars()
	for i := range vars {
		t := vars[i]
		renderedSpec = strings.ReplaceAll(renderedSpec, t.RefName(), t.Value)
	}
	return renderedSpec
}

func NewJobTaskEnvConfigMapName(TaskId string) string {
	return fmt.Sprintf("task-%s", TaskId)
}

// Job Task 挂载一个空的config 用于收集运行时的 环境变量
func (s *RunTaskRequest) RuntimeEnvConfigMap(mountPath string) *v1.ConfigMap {
	cm := new(v1.ConfigMap)
	s.Params.GetDeploymentId()
	cm.Name = NewJobTaskEnvConfigMapName(s.Params.GetJobTaskId())
	cm.BinaryData = map[string][]byte{
		CONFIG_MAP_RUNTIME_ENV_KEY: []byte(""),
	}
	cm.Annotations = map[string]string{
		workload.ANNOTATION_CONFIGMAP_MOUNT: mountPath,
	}
	return cm
}

type PipelineService interface {
	PipelineRPCServer
}

func NewDescribeJobTaskRequest(id string) *DescribeJobTaskRequest {
	return &DescribeJobTaskRequest{
		Id: id,
	}
}

func (r *DescribeJobTaskRequest) Validate() error {
	if r.Id == "" {
		return fmt.Errorf("job id required")
	}
	return nil
}

func NewDeleteJobTaskRequest(id string) *DeleteJobTaskRequest {
	return &DeleteJobTaskRequest{
		Id: id,
	}
}

func NewUpdateJobTaskOutputRequest(id string) *UpdateJobTaskOutputRequest {
	return &UpdateJobTaskOutputRequest{
		Id:          id,
		RuntimeEnvs: []*RuntimeEnv{},
	}
}

func (req *UpdateJobTaskOutputRequest) AddRuntimeEnv(name, value string) {
	req.RuntimeEnvs = append(req.RuntimeEnvs, NewRuntimeEnv(name, value))
}

func NewUpdateJobTaskStatusRequest(id string) *UpdateJobTaskStatusRequest {
	return &UpdateJobTaskStatusRequest{
		Id: id,
	}
}

func (r *UpdateJobTaskStatusRequest) MarkError(err error) {
	r.Stage = STAGE_FAILED
	r.Message = err.Error()
}

func NewQueryPipelineTaskRequest() *QueryPipelineTaskRequest {
	return &QueryPipelineTaskRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewDeletePipelineTaskRequest(id string) *DeletePipelineTaskRequest {
	return &DeletePipelineTaskRequest{
		Id: id,
	}
}

func NewDescribePipelineTaskRequest(id string) *DescribePipelineTaskRequest {
	return &DescribePipelineTaskRequest{
		Id: id,
	}
}

// IM 消息通知, 独立适配
type TaskMessage interface {
	// 消息来自的源
	GetDomain() string
	// 消息来源的空间
	GetNamespace() string
	// IM消息事件的Title
	ShowTitle() string
	// Markdown格式的消息内容
	MarkdownContent() string
	// HTML格式的消息内容
	HTMLContent() string
	// 消息事件状态
	GetStatusStage() STAGE
	// 通知过程中的事件
	AddErrorEvent(format string, a ...any)
}

type WebHookMessage interface {
	TaskMessage
	// Web事件回调
	AddWebhookStatus(items ...*CallbackStatus)
}

// Task状态变更用户通知
type MentionUserMessage interface {

	// Task状态变化通知消息
	TaskMessage
	// 通知回调, 是否通知成功
	AddNotifyStatus(items ...*CallbackStatus)
}

func NewJobTaskStreamReponse() *JobTaskStreamReponse {
	return &JobTaskStreamReponse{
		Data: make([]byte, 0, 512),
	}
}

func (r *JobTaskStreamReponse) ReSet() {
	r.Data = r.Data[:0]
}

func NewWatchJobTaskLogRequest(taskId string) *WatchJobTaskLogRequest {
	return &WatchJobTaskLogRequest{
		TaskId: taskId,
	}
}

func (req *WatchJobTaskLogRequest) ToJSON() string {
	return pretty.ToJSON(req)
}

func NewTaskLogWebsocketTerminal(conn *websocket.Conn) *TaskLogWebsocketTerminal {
	return &TaskLogWebsocketTerminal{
		terminal.NewWebSocketWriter(conn),
	}
}

type TaskLogWebsocketTerminal struct {
	*terminal.WebSocketWriter
}

func (r *TaskLogWebsocketTerminal) Send(in *JobTaskStreamReponse) (err error) {
	_, err = r.Write(in.Data)
	return
}
