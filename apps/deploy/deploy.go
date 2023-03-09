package deploy

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mpaas/apps/job"
	meta "github.com/infraboard/mpaas/common/meta"
	"github.com/infraboard/mpaas/provider/k8s/network"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/core/v1"
)

func NewDeploymentSet() *DeploymentSet {
	return &DeploymentSet{
		Items: []*Deployment{},
	}
}

func (s *DeploymentSet) Add(item *Deployment) {
	s.Items = append(s.Items, item)
}

func NewDefaultDeploy() *Deployment {
	return &Deployment{
		Spec: NewCreateDeploymentRequest(),
	}
}

func (d *Deployment) ValidateToken(token string) error {
	if d.Spec == nil {
		return nil
	}

	if !d.Spec.AuthEnabled {
		return nil
	}

	if d.Credential.Token != token {
		return fmt.Errorf("集群访问Token不合法")
	}

	return nil
}

func (d *Deployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*meta.Scope
		*CreateDeploymentRequest
		Status *Status `json:"status"`
	}{d.Meta, d.Scope, d.Spec, d.Status})
}

func (d *Deployment) GetK8sClusterId() string {
	if d.Spec == nil {
		return ""
	}
	if d.Spec.K8STypeConfig == nil {
		return ""
	}
	return d.Spec.K8STypeConfig.ClusterId
}

// 部署时的系统变量, 在部署任务时注入
func (d *Deployment) SystemVariable() ([]v1.EnvVar, error) {
	items := []*job.RunParam{}
	switch d.Spec.Type {
	case TYPE_KUBERNETES:
		wc := d.Spec.K8STypeConfig
		// 与k8s部署相关的系统变量, 不要注入版本信息, 部署版本由用户自己传人
		wl, err := wc.GetWorkLoad()
		if err != nil {
			return nil, err
		}
		variables := wl.SystemVaraible(d.Spec.ServiceName)
		addr, _ := variables.ImageDetail()
		items = append(items,
			job.NewRunParam(
				job.SYSTEM_VARIABLE_WORKLOAD_KIND,
				strings.ToLower(d.Spec.K8STypeConfig.WorkloadKind),
			).SetReadOnly(true).SetSearchLabel(true),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_WORKLOAD_NAME,
				variables.WorkloadName,
			).SetReadOnly(true).SetSearchLabel(true),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_SERVICE_NAME,
				d.Spec.ServiceName,
			).SetReadOnly(true).SetSearchLabel(true),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_IMAGE_REPOSITORY,
				addr,
			),
		)
	}

	return job.ParamsToEnvVar(items), nil
}

func (c *K8STypeConfig) GetWorkLoad() (*workload.WorkLoad, error) {
	return workload.ParseWorkloadFromYaml(c.WorkloadKind, c.WorkloadConfig)
}

func (c *K8STypeConfig) Merge(target *K8STypeConfig) error {
	if target == nil {
		return nil
	}

	if c.ClusterId != target.ClusterId {
		return fmt.Errorf("k8s cluster id can't update")
	}

	return mergo.MergeWithOverwrite(c, target)
}

func (c *K8STypeConfig) GetServiceObj() (*v1.Service, error) {
	if c.Service == "" {
		return nil, nil
	}
	return network.ParseServiceFromYaml(c.Service)
}

func NewStatus() *Status {
	return &Status{}
}

func (s *Status) MarkCreating() {
	s.Stage = STAGE_CREATING
}

func (s *Status) Update(target *Status) {
	if target == nil {
		return
	}
	target.UpdateAt = time.Now().Unix()
	s = target
}
