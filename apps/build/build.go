package build

import (
	"encoding/json"
	"regexp"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/common/meta"
)

// New 新建一个domain
func New(req *CreateBuildConfigRequest) (*BuildConfig, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &BuildConfig{
		Meta: meta.NewMeta(),
		Spec: req,
	}

	return d, nil
}

func NewBuildConfigSet() *BuildConfigSet {
	return &BuildConfigSet{
		Items: []*BuildConfig{},
	}
}

func (s *BuildConfigSet) Add(item *BuildConfig) {
	s.Items = append(s.Items, item)
}

// 比如: "foo.*"
func (s *BuildConfigSet) MatchBranch(branchRegExp string) *BuildConfigSet {
	set := NewBuildConfigSet()

	for i := range s.Items {
		item := s.Items[i]
		if item.Spec.Condition.MatchBranch(branchRegExp) {
			set.Add(item)
		}
	}

	return set
}

func NewDefaultBuildConfig() *BuildConfig {
	return &BuildConfig{
		Spec: NewCreateBuildConfigRequest(),
	}
}

func (b *BuildConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateBuildConfigRequest
	}{b.Meta, b.Spec})
}

func (b *BuildConfig) BuildRunParams() *job.VersionedRunParam {
	params := job.NewVersionedRunParam("build")
	// 补充部署Id
	if b.Spec.DeployId != "" {
		params.Add(job.NewRunParam(
			job.SYSTEM_VARIABLE_DEPLOY_ID,
			b.Spec.DeployId,
		))
	}

	// 补充自定义变量
	envs := map[string]string{}
	switch b.Spec.TargetType {
	case TARGET_TYPE_IMAGE:
		envs = b.Spec.ImageBuild.BuildEnvVars
	case TARGET_TYPE_PKG:
	}
	for k, v := range envs {
		params.Add(job.NewRunParam(k, v))
	}

	return params
}

func NewCreateBuildConfigRequest() *CreateBuildConfigRequest {
	return &CreateBuildConfigRequest{
		VersionPrefix: "v",
		Condition:     NewTrigger(),
		ImageBuild:    NewImageBuild(),
		PkgBuild:      NewPkgBuildConfig(),
		Labels:        make(map[string]string),
	}
}

func NewImageBuild() *ImageBuildConfig {
	return &ImageBuildConfig{
		DockerFile:   "Dockerfile",
		BuildEnvVars: make(map[string]string),
		Extra:        make(map[string]string),
	}
}

// 如果没有配置，则使用默认配置
func (c *ImageBuildConfig) GetDockerFileWithDefault(defaulDockerFile string) string {
	if c.DockerFile == "" {
		return defaulDockerFile
	}
	return c.DockerFile
}

// 如果没有配置，则使用默认配置
func (c *ImageBuildConfig) GetImageRepositoryWithDefault(defaultImageRepository string) string {
	if c.ImageRepository == "" {
		return defaultImageRepository
	}
	return c.ImageRepository
}

func NewPkgBuildConfig() *PkgBuildConfig {
	return &PkgBuildConfig{
		Extra: make(map[string]string),
	}
}

func NewTrigger() *Trigger {
	return &Trigger{
		Events:   []string{},
		Branches: []string{},
	}
}

func (t *Trigger) AddEvent(event string) {
	t.Events = append(t.Events, event)
}

func (t *Trigger) AddBranche(branche string) {
	t.Branches = append(t.Branches, branche)
}

func (t *Trigger) MatchBranch(branchRegExp string) bool {
	for _, b := range t.Branches {
		ok, _ := regexp.MatchString(branchRegExp, b)
		if ok {
			return true
		}
	}

	return false
}
