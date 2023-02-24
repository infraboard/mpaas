package deploy

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infraboard/mpaas/apps/job"
	meta "github.com/infraboard/mpaas/common/meta"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
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

func (d *Deployment) SystemVariable() (items []*job.RunParam, err error) {
	switch d.Spec.Type {
	case TYPE_KUBERNETES:
		// 与k8s部署相关的系统变量
		variables, err := d.Spec.K8STypeConfig.DeploySystemVaraible(d.Spec.ServiceName)
		if err != nil {
			return nil, err
		}
		addr, version := variables.ImageDetail()
		items = append(items,
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_WORKLOAD_KIND,
				strings.ToLower(d.Spec.K8STypeConfig.WorkloadKind.String()),
			),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_WORKLOAD_NAME,
				variables.WorloadName,
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

func NewDeploySystemVaraible() *DeploySystemVaraible {
	return &DeploySystemVaraible{}
}

type DeploySystemVaraible struct {
	WorloadName string `json:"workload_name"`
	Image       string `json:"image"`
}

func (v *DeploySystemVaraible) ImageDetail() (addr, version string) {
	if v.Image == "" {
		return
	}
	av := strings.Split(v.Image, ":")
	addr = av[0]
	if len(av) > 1 {
		version = av[1]
	}
	return
}

func (c *K8STypeConfig) DeploySystemVaraible(serviceName string) (*DeploySystemVaraible, error) {
	m := NewDeploySystemVaraible()
	if c.GetWorkloadConfig() == "" {
		return m, nil
	}

	var container v1.Container
	switch c.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		obj := &appsv1.Deployment{}
		err := yaml.Unmarshal([]byte(c.WorkloadConfig), obj)
		if err != nil {
			return m, err
		}
		m.WorloadName = obj.Name
		container = workload.GetContainerFromPodTemplate(obj.Spec.Template, serviceName)
	case WORKLOAD_KIND_STATEFULSET:
		obj := &appsv1.StatefulSet{}
		err := yaml.Unmarshal([]byte(c.WorkloadConfig), obj)
		if err != nil {
			return nil, err
		}
		m.WorloadName = obj.Name
		container = workload.GetContainerFromPodTemplate(obj.Spec.Template, serviceName)
	case WORKLOAD_KIND_DAEMONSET:
		obj := &appsv1.DaemonSet{}
		err := yaml.Unmarshal([]byte(c.WorkloadConfig), obj)
		if err != nil {
			return nil, err
		}
		m.WorloadName = obj.Name
		container = workload.GetContainerFromPodTemplate(obj.Spec.Template, serviceName)
	case WORKLOAD_KIND_CRONJOB:
		obj := &batchv1.CronJob{}
		err := yaml.Unmarshal([]byte(c.WorkloadConfig), obj)
		if err != nil {
			return nil, err
		}
		m.WorloadName = obj.Name
		container = workload.GetContainerFromPodTemplate(obj.Spec.JobTemplate.Spec.Template, serviceName)
	case WORKLOAD_KIND_JOB:
		obj := &batchv1.Job{}
		err := yaml.Unmarshal([]byte(c.WorkloadConfig), obj)
		if err != nil {
			return nil, err
		}
		m.WorloadName = obj.Name
		container = workload.GetContainerFromPodTemplate(obj.Spec.Template, serviceName)
	}
	m.Image = container.Image
	return m, nil
}
