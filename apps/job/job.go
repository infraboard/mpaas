package job

import (
	"reflect"
	"time"

	"github.com/rs/xid"
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
		Items:   []*RunParam{},
	}
}

func (r *VersionedRunParam) Add(item *RunParam) {
	r.Items = append(r.Items, item)
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

// 获取参数的值
func (r *VersionedRunParam) GetParamValue(key string) string {
	for i := range r.Items {
		item := r.Items[i]
		if item.Name == key {
			return item.Value
		}
	}
	return ""
}

func NewK8SJobRunnerParams() *K8SJobRunnerParams {
	return &K8SJobRunnerParams{}
}
