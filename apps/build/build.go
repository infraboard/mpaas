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

func NewDefaultDeploy() *BuildConfig {
	return &BuildConfig{
		Spec: NewCreateDeployConfigRequest(),
	}
}

func (b *BuildConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateBuildConfigRequest
	}{b.Meta, b.Spec})
}

func NewCreateDeployConfigRequest() *CreateBuildConfigRequest {
	return &CreateBuildConfigRequest{}
}
