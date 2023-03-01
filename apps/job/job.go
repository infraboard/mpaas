package job

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/imdario/mergo"
	"github.com/infraboard/mpaas/common/meta"
	v1 "k8s.io/api/core/v1"
)

// New 新建一个部署配置
func New(req *CreateJobRequest) (*Job, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Job{
		Meta: meta.NewMeta(),
		Spec: req,
	}

	return d, nil
}

func NewJobSet() *JobSet {
	return &JobSet{
		Items: []*Job{},
	}
}

func (s *JobSet) Add(item *Job) {
	s.Items = append(s.Items, item)
}

func NewDefaultJob() *Job {
	return &Job{
		Spec: NewCreateJobRequest(),
	}
}

func (i *Job) Update(req *UpdateJobRequest) {
	i.Meta.UpdateAt = time.Now().Unix()
	i.Meta.UpdateBy = req.UpdateBy
	i.Spec = req.Spec
}

func (i *Job) Patch(req *UpdateJobRequest) error {
	i.Meta.UpdateAt = time.Now().Unix()
	i.Meta.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Spec, req.Spec)
}

func (i *Job) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateJobRequest
	}{i.Meta, i.Spec})
}

func (i *Job) GetVersionedRunParam(version string) *VersionedRunParam {
	for m := range i.Spec.RunParams {
		v := i.Spec.RunParams[m]
		if v.Version == version {
			return v
		}
	}

	return nil
}

func NewVersionedRunParam(version string) *VersionedRunParam {
	return &VersionedRunParam{
		Version: version,
		Params:  []*RunParam{},
	}
}

func (r *VersionedRunParam) Add(items ...*RunParam) {
	r.Params = append(r.Params, items...)
}

func (r *VersionedRunParam) Validate() error {
	for i := range r.Params {
		p := r.Params[i]
		if p.Required && p.Value == "" {
			return fmt.Errorf("参数: %s 不能为空", p.Name)
		}
	}

	return nil
}

// 从参数中提取k8s job执行器(runner)需要的参数
// 这里采用反射来获取Struc Tag, 然后根据Struct Tag 获取参数的具体值
// 关于反射 可以参考: https://blog.csdn.net/bocai_xiaodaidai/article/details/123668047
func (r *VersionedRunParam) K8SJobRunnerParams() *K8SJobRunnerParams {
	params := NewK8SJobRunnerParams()

	// params是一个Pointer Value, 如果需要获取值的类型需要这样处理:
	//	reflect.Indirect(reflect.ValueOf(params)).Type()
	// 因此这里直接采用K8SJobRunnerParams{}获取类型
	pt := reflect.TypeOf(K8SJobRunnerParams{})

	// go语言所有函数传的都是值，所以要想修改原来的值就需要传指
	// 通过Elem()返回指针指向的对象
	v := reflect.ValueOf(params).Elem()

	for i := 0; i < pt.NumField(); i++ {
		field := pt.Field(i)
		if field.IsExported() {
			tagValue := field.Tag.Get("param")
			v.Field(i).SetString(r.GetParamValue(tagValue))
		}
	}

	return params
}

func (r *VersionedRunParam) GetDeploymentId() string {
	return r.GetParamValue(SYSTEM_VARIABLE_DEPLOY_ID)
}

func (r *VersionedRunParam) GetJobTaskId() string {
	return r.GetParamValue(SYSTEM_VARIABLE_JOB_TASK_ID)
}

func (r *VersionedRunParam) GetPipelineTaskId() string {
	return r.GetParamValue(SYSTEM_VARIABLE_PIPELINE_TASK_ID)
}

// 获取需要注入容器的环境变量参数:
//
//	用户变量: 大写开头的变量, 因为一般环境变量都是大写的比如 DB_PASS,
//	系统变量: _开头为系统变量, 由Runner处理并注入, 比如 _DEPLOY_ID
//	Runner变量: 小写的变量, 用于系统内部使用, 不会注入, 比如 K8SJobRunnerParams 中的cluster_id
func (r *VersionedRunParam) EnvVars() (envs []v1.EnvVar) {
	for i := range r.Params {
		item := r.Params[i]
		// 只导出环境变量
		if !item.UsageType.Equal(PARAM_USAGE_TYPE_ENV) {
			continue
		}
		if item.Name != "" && (unicode.IsUpper(rune(item.Name[0])) || strings.HasPrefix(item.Name, "_")) {
			envs = append(envs, v1.EnvVar{
				Name:  item.Name,
				Value: item.Value,
			})
		}
	}
	return
}

func (r *VersionedRunParam) TemplateVars() (vars []*RunParam) {
	for i := range r.Params {
		item := r.Params[i]
		// 只导出模版变量
		if item.UsageType.Equal(PARAM_USAGE_TYPE_TEMPLATE) {
			vars = append(vars, item)
		}
	}
	return
}

func ParamsToEnvVar(params []*RunParam) (envs []v1.EnvVar) {
	for i := range params {
		item := params[i]
		envs = append(envs, v1.EnvVar{
			Name:  item.Name,
			Value: item.Value,
		})
	}
	return
}

// 获取参数的值
func (r *VersionedRunParam) GetParamValue(key string) string {
	for i := range r.Params {
		item := r.Params[i]
		if item.Name == key {
			return item.Value
		}
	}
	return ""
}

// 设置参数的值
func (r *VersionedRunParam) SetParamValue(key, value string) {
	for i := range r.Params {
		item := r.Params[i]
		if item.Name == key {
			item.Value = value
			return
		}
	}
}

func (r *VersionedRunParam) Merge(target *VersionedRunParam) {
	for i := range target.Params {
		t := target.Params[i]
		r.SetParamValue(t.Name, t.Value)
	}
}

func (r *VersionedRunParam) UpdateFromEnvs(targets []v1.EnvVar) {
	for i := range targets {
		t := targets[i]
		r.SetParamValue(t.Name, t.Value)
	}
}

func NewK8SJobRunnerParams() *K8SJobRunnerParams {
	return &K8SJobRunnerParams{}
}

func NewRunParam(name, value string) *RunParam {
	return &RunParam{
		Name:  name,
		Value: value,
	}
}

func NewRunParamWithKVPaire(kvs ...string) (params []*RunParam) {
	if len(kvs)%2 != 0 {
		panic("kvs must paire")
	}

	if len(kvs) == 0 {
		return
	}

	kv := []string{}
	for i, v := range kvs {
		kv = append(kv, v)
		if i%2 != 0 {
			params = append(params, NewRunParam(kv[0], kv[1]))
			kv = []string{}
		}
	}

	return
}

// 引用名称
func (p *RunParam) RefName() string {
	return fmt.Sprintf("${%s}", p.Name)
}
