package deploy

import (
	context "context"
	"fmt"
	"net/http"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	pb_request "github.com/infraboard/mcube/v2/pb/request"
	resource "github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/tools/hash"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

const (
	AppName        = "deploys"
	ResourcePreifx = "dep"
)

type Service interface {
	CreateDeployment(context.Context, *CreateDeploymentRequest) (*Deployment, error)
	UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*Deployment, error)
	DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*Deployment, error)
	RPCServer
}

// New 新建一个部署配置
func New(req *CreateDeploymentRequest) *Deployment {
	m := resource.NewMeta().IdWithPrefix(ResourcePreifx)
	d := &Deployment{
		Meta:             m,
		Spec:             req,
		Status:           NewStatus(),
		Credential:       NewCredential(),
		DynamicInjection: NewDdynamicInjection(),
	}

	return d
}

func (req *CreateDeploymentRequest) Validate() error {
	return validator.Validate(req)
}

func (req *CreateDeploymentRequest) ValidateWorkLoad() error {
	if req.ServiceId == "" {
		return fmt.Errorf("when workload, service_id required")
	}

	return nil
}

func (req *CreateDeploymentRequest) ValidateMiddleware() error {
	if req.ServiceName == "" {
		return fmt.Errorf("when middleware, service_name required")
	}

	return nil
}

func (req *CreateDeploymentRequest) SetDefault() {
	if req.Name == "" {
		req.Name = req.ServiceName
	}
}

func (req *CreateDeploymentRequest) UUID() string {
	return hash.FnvHash(req.Domain, req.Namespace, req.ServiceName, req.Cluster, req.Name)
}

func NewQueryDeploymentRequestFromHttp(r *http.Request) *QueryDeploymentRequest {
	req := NewQueryDeploymentRequest()
	req.Page = request.NewPageRequestFromHTTP(r)
	return req
}

func NewQueryDeploymentRequest() *QueryDeploymentRequest {
	return &QueryDeploymentRequest{
		Page:  request.NewDefaultPageRequest(),
		Scope: resource.NewScope(),
	}
}

func (r *QueryDeploymentRequest) AddServiceId(sid ...string) {
	r.ServiceIds = append(r.ServiceIds, sid...)
}

func (r *QueryDeploymentRequest) AddClusterId(cid ...string) {
	r.Clusters = append(r.Clusters, cid...)
}

func NewCreateDeploymentRequest() *CreateDeploymentRequest {
	return &CreateDeploymentRequest{
		AuthEnabled:    false,
		EventNotify:    NewEventNotify(),
		Labels:         make(map[string]string),
		K8STypeConfig:  NewK8STypeConfig(),
		HostTypeConfig: NewHostTypeConfig(),
		Provider:       instance.DEFAULT_PROVIDER,
		Region:         instance.DEFAULT_REGION,
		Environment:    instance.DEFAULT_ENV,
		Group:          instance.DEFAULT_GROUP,
	}
}

func (r *CreateDeploymentRequest) UpdateOwner(tk *token.Token) {
	r.Domain = tk.Domain
	r.Namespace = tk.Namespace
	r.CreateBy = tk.Username
}

func NewK8STypeConfig() *K8STypeConfig {
	return &K8STypeConfig{
		WorkloadKind: workload.WORKLOAD_KIND_DEPLOYMENT.String(),
		Pods:         map[string]string{},
		Extras:       map[string]string{},
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
		return validator.Validate(req)
	}

	return nil
}

func NewDescribeDeploymentRequest(id string) *DescribeDeploymentRequest {
	return &DescribeDeploymentRequest{
		DescribeValue: id,
	}
}

func (req *DescribeDeploymentRequest) Validate() error {
	return validator.Validate(req)
}

func NewPutDeployRequest(id string) *UpdateDeploymentRequest {
	return &UpdateDeploymentRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		Spec:       NewCreateDeploymentRequest(),
		Sync:       true,
	}
}

func NewPatchDeployRequest(id string) *UpdateDeploymentRequest {
	return &UpdateDeploymentRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Spec:       NewCreateDeploymentRequest(),
		Sync:       true,
	}
}

func NewDeleteDeploymentRequest(id string) *DeleteDeploymentRequest {
	return &DeleteDeploymentRequest{
		Id: id,
	}
}

func NewUpdateDeploymentStatusRequest(id string) *UpdateDeploymentStatusRequest {
	return &UpdateDeploymentStatusRequest{
		UpdatedK8SConfig: NewK8STypeConfig(),
		Id:               id,
	}
}

func NewQueryDeploymentInjectEnvRequest(id string) *QueryDeploymentInjectEnvRequest {
	return &QueryDeploymentInjectEnvRequest{
		Id: id,
	}
}
