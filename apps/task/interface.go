package task

import (
	context "context"
	"fmt"
	"strings"

	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
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
	JobRPCServer
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

func NewRunTaskRequest(name, spec string, params *job.VersionedRunParam) *RunTaskRequest {
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
