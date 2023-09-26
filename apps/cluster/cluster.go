package cluster

import (
	"encoding/json"
	"fmt"
	"strings"

	service "github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/pb/resource"
	"github.com/infraboard/mcube/tools/hash"
	deploy "github.com/infraboard/mpaas/apps/deploy"
	v1 "k8s.io/api/core/v1"
)

func NewClusterSet() *ClusterSet {
	return &ClusterSet{
		Items: []*Cluster{},
	}
}

func (s *ClusterSet) Add(items ...*Cluster) {
	s.Items = append(s.Items, items...)
}

func (s *ClusterSet) Len() int {
	return len(s.Items)
}

func (s *ClusterSet) GetClusterByID(clusterId string) *Cluster {
	for i := range s.Items {
		if s.Items[i].Meta.Id == clusterId {
			return s.Items[i]
		}
	}
	return nil
}

func (s *ClusterSet) UpdateDeploymens(ds *deploy.DeploymentSet) {
	for i := range ds.Items {
		item := ds.Items[i]
		c := s.GetClusterByID(item.Spec.Cluster)
		if c != nil {
			c.Deployments.Add(item)
		}
	}
}

func (s *ClusterSet) ForEatch(fn func(*Cluster)) {
	for i := range s.Items {
		fn(s.Items[i])
	}
}

func (s *ClusterSet) ClusterIds() (ids []string) {
	for i := range s.Items {
		item := s.Items[i]
		ids = append(ids, item.Meta.Id)
	}
	return
}

func (s *ClusterSet) ServiceIds() (ids []string) {
	for i := range s.Items {
		item := s.Items[i]
		ids = append(ids, item.Spec.ServiceId)
	}
	return
}

func New(req *CreateClusterRequest) (*Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	ins := &Cluster{
		Meta:  resource.NewMeta(),
		Scope: resource.NewScope(),
		Spec:  req,
	}

	// 生成唯一键
	ins.Meta.Id = hash.FnvHash(ins.FullName())
	return ins, nil
}

func (req *CreateClusterRequest) Validate() error {
	return validate.Validate(req)
}

func NewDefaultCluster() *Cluster {
	return &Cluster{
		Deployments: deploy.NewDeploymentSet(),
	}
}

func (c *Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*resource.Scope
		*CreateClusterRequest
		Deployments *deploy.DeploymentSet `json:"deployments"`
		Service     *service.Service      `json:"service"`
	}{c.Meta, c.Scope, c.Spec, c.Deployments, c.Service})
}

func (c *Cluster) FullName() string {
	return fmt.Sprintf("%s.%s.%s.%s",
		c.Scope.Domain,
		c.Scope.Namespace,
		c.Spec.ServiceId,
		c.Spec.Name,
	)
}

func NewAccessAddressFromK8sService(svc *v1.Service) *AccessAddress {
	address := NewAccessAddress()
	for i := range svc.Spec.Ports {
		p := svc.Spec.Ports[i]
		name := fmt.Sprintf(
			"%s_PORT_%d_%s",
			strings.ToUpper(svc.Name),
			p.Port,
			strings.ToUpper(string(p.Protocol)),
		)
		address.AddServiceEnv(name, "REDIS_SERVICE_NAME_PORT_6379_TCP")
	}

	return &AccessAddress{}
}

func NewServiceEnv(name, example string) *AccessEnv {
	return &AccessEnv{
		Name:    name,
		Example: example,
	}
}

func NewAccessAddress() *AccessAddress {
	return &AccessAddress{
		AccessEnvs: []*AccessEnv{},
	}
}

func (a *AccessAddress) AddServiceEnv(name, example string) {
	a.AccessEnvs = append(a.AccessEnvs, NewServiceEnv(name, example))
}
