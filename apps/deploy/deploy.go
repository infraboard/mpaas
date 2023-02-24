package deploy

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infraboard/mpaas/apps/job"
	meta "github.com/infraboard/mpaas/common/meta"
	k8smeta "github.com/infraboard/mpaas/provider/k8s/meta"
	"sigs.k8s.io/yaml"
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
		*meta.Scope
		*CreateDeployConfigRequest
	}{d.Meta, d.Scope, d.Spec})
}

func (d *DeployConfig) SystemVariable() (items []*job.RunParam) {
	switch d.Spec.Type {
	case TYPE_KUBERNETES:
		// 与k8s部署相关的系统变量

		items = append(items,
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_WORKLOAD_KIND,
				strings.ToLower(d.Spec.K8STypeConfig.WorkloadKind.String()),
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_WORKLOAD_NAME,
				strings.ToLower(d.Spec.K8STypeConfig.WorkloadKind.String()),
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_SERVICE_NAME,
				d.Spec.ServiceName,
			),
		)
	}
	return
}

func (c *K8STypeConfig) ObjectMeta() (*k8smeta.ObjectMeta, error) {
	m := k8smeta.NewObjectMeta()
	if c.GetWorkloadConfig() == "" {
		return m, nil
	}

	err := yaml.Unmarshal([]byte(c.WorkloadConfig), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
