package deploy

import (
	context "context"
	"fmt"
	"net/http"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mpaas/common/meta"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

const (
	AppName = "deploys"
)

type Service interface {
	CreateDeployment(context.Context, *CreateDeploymentRequest) (*Deployment, error)
	DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*Deployment, error)
	RPCServer
}

// New 新建一个部署配置
func New(req *CreateDeploymentRequest) (*Deployment, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	m := meta.NewMeta()
	if req.DeployId != "" {
		m.Id = req.DeployId
	}

	d := &Deployment{
		Meta:  m,
		Scope: meta.NewScope(),
		Spec:  req,
	}

	return d, nil
}

func (req *CreateDeploymentRequest) Validate() error {
	return validate.Validate(req)
}

func NewQueryDeploymentRequestFromHttp(r *http.Request) *QueryDeploymentRequest {
	req := NewQueryDeploymentRequest()
	req.Page = request.NewPageRequestFromHTTP(r)
	return req
}

func NewQueryDeploymentRequest() *QueryDeploymentRequest {
	return &QueryDeploymentRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewCreateDeploymentRequest() *CreateDeploymentRequest {
	return &CreateDeploymentRequest{
		AuthEnabled:    false,
		Labels:         make(map[string]string),
		K8STypeConfig:  NewK8STypeConfig(),
		HostTypeConfig: NewHostTypeConfig(),
	}
}

func NewK8STypeConfig() *K8STypeConfig {
	return &K8STypeConfig{
		WorkloadKind: workload.WORKLOAD_KIND_DEPLOYMENT.String(),
	}
}

func NewHostTypeConfig() *HostTypeConfig {
	return &HostTypeConfig{}
}

// Validate 校验请求是否合法
func (req *UpdateDeploymentRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("id required")
	}
	if req.UpdateMode.Equal(pb_request.UpdateMode_PUT) {
		return validate.Validate(req)
	}

	return nil
}

func NewDescribeDeploymentRequest(id string) *DescribeDeploymentRequest {
	return &DescribeDeploymentRequest{
		DescribeValue: id,
	}
}

func (req *DescribeDeploymentRequest) Validate() error {
	return validate.Validate(req)
}

func NewPutDeployRequest(id string) *UpdateDeploymentRequest {
	return &UpdateDeploymentRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		Spec:       NewCreateDeploymentRequest(),
	}
}

func NewPatchDeployRequest(id string) *UpdateDeploymentRequest {
	return &UpdateDeploymentRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Spec:       NewCreateDeploymentRequest(),
	}
}

func NewDeleteDeploymentRequest(id string) *DeleteDeploymentRequest {
	return &DeleteDeploymentRequest{
		Id: id,
	}
}
