package build

import (
	"encoding/json"

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

func NewCreateBuildConfigRequest() *CreateBuildConfigRequest {
	return &CreateBuildConfigRequest{
		Condition:  NewTrigger(),
		ImageBuild: NewImageBuild(),
		PkgBuild:   NewPkgBuildConfig(),
		Labels:     make(map[string]string),
	}
}

func NewTrigger() *Trigger {
	return &Trigger{
		Events:   []string{},
		Branches: []string{},
	}
}

func NewImageBuild() *ImageBuildConfig {
	return &ImageBuildConfig{
		BuildEnvVars: make(map[string]string),
		Extra:        make(map[string]string),
	}
}

func NewPkgBuildConfig() *PkgBuildConfig {
	return &PkgBuildConfig{
		Extra: make(map[string]string),
	}
}

func (t *Trigger) AddEvent(event string) {
	t.Events = append(t.Events, event)
}

func (t *Trigger) AddBranche(branche string) {
	t.Branches = append(t.Branches, branche)
}
