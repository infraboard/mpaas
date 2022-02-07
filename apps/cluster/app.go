package cluster

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
)

const (
	AppName = "cluster"
)

var (
	validate = validator.New()
)

func NewCreateClusterRequest() *CreateClusterRequest {
	return &CreateClusterRequest{}
}

func NewCluster(req *CreateClusterRequest) (*Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Cluster{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMicro(),
		Data:     req,
	}, nil
}

func (req *CreateClusterRequest) Validate() error {
	return validate.Struct(req)
}

func NewClusterSet() *ClusterSet {
	return &ClusterSet{
		Items: []*Cluster{},
	}
}

func (s *ClusterSet) Add(item *Cluster) {
	s.Items = append(s.Items, item)
}

func NewDefaultCluster() *Cluster {
	return &Cluster{
		Data: &CreateClusterRequest{},
	}
}

func (i *Cluster) Update(req *UpdateClusterRequest) {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	i.Data = req.Data
}

func (i *Cluster) Patch(req *UpdateClusterRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Data, req.Data)
}

func NewDescribeClusterRequest(id string) *DescribeClusterRequest {
	return &DescribeClusterRequest{
		Id: id,
	}
}

func NewQueryClusterRequest() *QueryClusterRequest {
	return &QueryClusterRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewQueryClusterRequestFromHTTP(r *http.Request) *QueryClusterRequest {
	qs := r.URL.Query()

	return &QueryClusterRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewPutClusterRequest(id string) *UpdateClusterRequest {
	return &UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateClusterRequest(),
	}
}

func NewPatchClusterRequest(id string) *UpdateClusterRequest {
	return &UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateClusterRequest(),
	}
}

func NewDeleteClusterRequestWithID(id string) *DeleteClusterRequest {
	return &DeleteClusterRequest{
		Id: id,
	}
}
