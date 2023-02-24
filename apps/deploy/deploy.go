package deploy

import (
	"encoding/json"
	"fmt"
	"strings"

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

	if d.Status.Token != token {
		return fmt.Errorf("集群访问Token不合法")
	}

	return nil
}

func (d *Deployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*meta.Scope
		*CreateDeploymentRequest
	}{d.Meta, d.Scope, d.Spec})
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
func (d *Deployment) SystemVariable() (items []*job.RunParam, err error) {
	switch d.Spec.Type {
	case TYPE_KUBERNETES:
		wc := d.Spec.K8STypeConfig
		// 与k8s部署相关的系统变量
		wl, err := wc.GetWorkLoad()
		if err != nil {
			return nil, err
		}
		variables := wl.SystemVaraible(d.Spec.ServiceName)
		addr, version := variables.ImageDetail()
		items = append(items,
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_WORKLOAD_KIND,
				strings.ToLower(d.Spec.K8STypeConfig.WorkloadKind),
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_WORKLOAD_NAME,
				variables.WorkloadName,
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_SERVICE_NAME,
				d.Spec.ServiceName,
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_SERVICE_IMAGE_ADDR,
				addr,
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_SERVICE_IMAGE_VERSION,
				version,
			),
		)
	}
	return
}

func (c *K8STypeConfig) GetWorkLoad() (*workload.WorkLoad, error) {
	return workload.ParseWorkloadFromYaml(c.WorkloadKind, c.WorkloadConfig)
}

func (c *K8STypeConfig) GetServiceObj() (*v1.Service, error) {
	if c.Service == "" {
		return nil, nil
	}
	return network.ParseServiceFromYaml(c.Service)
}
