package cluster

import (
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/pb/resource"
)

func NewClusterSet() *ClusterSet {
	return &ClusterSet{
		Items: []*Cluster{},
	}
}

func (s *ClusterSet) Add(items ...*Cluster) {
	s.Items = append(s.Items, items...)
}

func New(req *CreateClusterRequest) (*Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Cluster{
		Meta: resource.NewMeta(),
		Spec: req,
	}, nil
}

func (req *CreateClusterRequest) Validate() error {
	return validate.Validate(req)
}

func NewDefaultCluster() *Cluster {
	return &Cluster{}
}
