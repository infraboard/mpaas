package job

import (
	"reflect"
	"time"
	"unicode"

	"github.com/rs/xid"
	corev1 "k8s.io/api/core/v1"
)

// New 新建一个部署配置
func New(req *CreateJobRequest) (*Job, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Job{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
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

func NewVersionedRunParam(version string) *VersionedRunParam {
	return &VersionedRunParam{
		Version: version,
		Params:  []*RunParam{},
	}
}

func (r *VersionedRunParam) Add(items ...*RunParam) {
	r.Params = append(r.Params, items...)
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

	if field, ok := pt.FieldByName("ClusterId"); ok {
		tagValue := field.Tag.Get("param")
		params.ClusterId = r.GetParamValue(tagValue)
	}

	return params
}

// 获取需要注入容器的环境变量参数
// 注意: 只有大写的变量才会被导出, 因为一般环境变量都是大写的, 比如 DB_PASS,
// 小写的变量用于系统内部使用, 比如 K8SJobRunnerParams 中的cluster_id
func (r *VersionedRunParam) EnvVars() (envs []corev1.EnvVar) {
	for i := range r.Params {
		item := r.Params[i]
		if item.Name != "" && unicode.IsUpper(rune(item.Name[0])) {
			envs = append(envs, corev1.EnvVar{
				Name:  item.Name,
				Value: item.Value,
			})
		}
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
