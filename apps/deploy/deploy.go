package deploy

import (
	"encoding/json"
	"fmt"
	"time"

	"dario.cat/mergo"
	"github.com/infraboard/mcube/v2/pb/resource"

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

func (s *DeploymentSet) ForEatch(fn func(item *Deployment)) {
	for i := range s.Items {
		fn(s.Items[i])
	}
}

func (s *DeploymentSet) Len() int {
	return len(s.Items)
}

func (s *DeploymentSet) GetK8sClusterIds() (ids []string) {
	m := map[string]string{}
	for i := range s.Items {
		item := s.Items[i]
		m[item.GetK8sClusterId()] = ""
	}

	for k := range m {
		ids = append(ids, k)
	}
	return
}

func NewDefaultDeploy() *Deployment {
	return &Deployment{
		Spec:             NewCreateDeploymentRequest(),
		Credential:       NewCredential(),
		Status:           NewStatus(),
		DynamicInjection: NewDdynamicInjection(),
	}
}

func (d *Deployment) SystemInjectionEnvGroup() *InjectionEnvGroup {
	group := NewInjectionEnvGroup()
	group.Name = "mpaas system env"
	group.AddEnv(
		NewInjectionEnv("MPAAS_DEPLOY_ID", d.Meta.Id),
		NewInjectionEnv("MCENTER_SERVICE_NAME", d.Spec.ServiceName),

		// 服务发现相关变量
		NewInjectionEnv("MCENTER_INSTANCE_PROVIDER", d.Spec.Provider),
		NewInjectionEnv("MCENTER_INSTANCE_REGION", d.Spec.Region),
		NewInjectionEnv("MCENTER_INSTANCE_ENV", d.Spec.Environment),
		NewInjectionEnv("MCENTER_INSTANCE_CLUSTER", d.Spec.Cluster),
		NewInjectionEnv("MCENTER_INSTANCE_GROUP", d.Spec.Group),
	)
	return group
}

func (d *Deployment) SetDefault() {
	if d.Status == nil {
		d.Status = NewStatus()
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
		*resource.Meta
		*CreateDeploymentRequest
		Status           *Status            `json:"status"`
		DynamicInjection *DdynamicInjection `json:"dynamic_injection"`
	}{d.Meta, d.Spec, d.Status, d.DynamicInjection})
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

func (d *Deployment) InjectPodLabel() map[string]string {
	m := map[string]string{
		LABEL_SERVICE_NAME_KEY: d.Spec.ServiceName,
		LABEL_NAMESPACE_KEY:    d.Spec.Namespace,
		LABEL_CLUSTER_KEY:      d.Spec.Cluster,
		LABEL_DEPLOY_GROUP_KEY: d.Spec.Group,
		LABEL_DEPLOY_ID_KEY:    d.Meta.Id,
	}
	return m
}

func (c *K8STypeConfig) GetWorkLoad() (*workload.WorkLoad, error) {
	return workload.ParseWorkloadFromYaml(c.WorkloadKind, c.WorkloadConfig)
}

func (c *K8STypeConfig) Merge(target *K8STypeConfig) error {
	if target == nil {
		return nil
	}

	// 不能更新集群的clusterId
	if target.ClusterId != "" && target.ClusterId != c.ClusterId {
		return fmt.Errorf("状态更新, 禁止修改集群, 修改值: %s", target.ClusterId)
	}

	// 清除删除的Pod
	for k, v := range target.Pods {
		if v == "" {
			delete(c.Pods, k)
		}
	}

	err := mergo.MergeWithOverwrite(c, target)
	if err != nil {
		return err
	}

	return nil
}

func (c *K8STypeConfig) GetServiceObj() (*v1.Service, error) {
	if c.Service == "" {
		return nil, nil
	}
	return network.ParseServiceFromYaml(c.Service)
}

func NewCredential() *Credential {
	return &Credential{}
}

func NewStatus() *Status {
	return &Status{}
}

func (s *Status) MarkFailed(err error) {
	s.Stage = STAGE_ERROR
	s.Message = err.Error()
}

func (s *Status) MarkCreating() {
	s.Stage = STAGE_CREATING
}

func (s *Status) Update(target *Status) error {
	if target == nil {
		return fmt.Errorf("Status为nil")
	}

	if target.Stage <= STAGE_CREATING {
		return fmt.Errorf("更新的状态不合法, 不能更新为PENDDING或者CREATING")
	}

	target.UpdateAt = time.Now().Unix()
	s = target

	return nil
}

func (s *Status) UpdateK8sWorkloadStatus(status *workload.WorkloadStatus) {
	if status == nil {
		return
	}

	switch status.Stage {
	case workload.WORKLOAD_STAGE_PROGERESS:
		s.Stage = STAGE_UPGRADING
	case workload.WORKLOAD_STAGE_ACTIVE:
		s.Stage = STAGE_ACTIVE
	case workload.WORKLOAD_STAGE_ERROR:
		s.Stage = STAGE_ERROR
	}
	s.Reason = status.Reason
	s.Message = status.Message
	s.UpdateAt = time.Now().Unix()
}

func NewEventNotify() *EventNotify {
	return &EventNotify{
		Users: []string{},
	}
}
