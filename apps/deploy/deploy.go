package deploy

import (
	"encoding/json"
	"fmt"

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

func (d *DeployConfig) ValidateToken(token string) error {
	if d.Spec == nil {
		return nil
	}

	if !d.Spec.AuthEnabled {
		return nil
	}

	if d.Status.Token != token {
		return fmt.Errorf("集群访问Token不合法")
	}

	return nil
}

func (d *DeployConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateDeployConfigRequest
	}{d.Meta, d.Spec})
}
