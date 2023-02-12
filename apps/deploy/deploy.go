package deploy

import (
	"encoding/json"

	meta "github.com/infraboard/mpaas/common/meta"
)

func NewDeployConfigSet() *DeployConfigSet {
	return &DeployConfigSet{
		Items: []*DeployConfig{},
	}
}

func (s *DeployConfigSet) Add(item *DeployConfig) {
	s.Items = append(s.Items, item)
}

func NewDefaultDeploy() *DeployConfig {
	return &DeployConfig{
		Spec: NewCreateDeployConfigRequest(),
	}
}

func (d *DeployConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateDeployConfigRequest
	}{d.Meta, d.Spec})
}
