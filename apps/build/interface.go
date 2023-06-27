package build

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
)

const (
	AppName = "builds"
)

type Service interface {
	RPCServer
}

func (req *CreateBuildConfigRequest) Validate() error {
	if req.VersionPrefix == "" {
		req.VersionPrefix = "v"
	}
	return validate.Validate(req)
}

func NewQueryBuildConfigRequestFromHTTP(r *restful.Request) *QueryBuildConfigRequest {
	req := NewQueryBuildConfigRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	req.Filters = policy.GetScopeFilterFromRequest(r)
	return req
}

func NewQueryBuildConfigRequest() *QueryBuildConfigRequest {
	return &QueryBuildConfigRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryBuildConfigRequest) AddService(serviceId string) {
	req.ServiceIds = append(req.ServiceIds, serviceId)
}

func (req *QueryBuildConfigRequest) SetEnabled(v bool) {
	req.Enabled = &v
}

func NewDescribeBuildConfigRequst(id string) *DescribeBuildConfigRequst {
	return &DescribeBuildConfigRequst{
		Id: id,
	}
}

func (req *DescribeBuildConfigRequst) Validate() error {
	return validate.Validate(req)
}

func NewDeleteBuildConfigRequest(id string) *DeleteBuildConfigRequest {
	return &DeleteBuildConfigRequest{
		Id: id,
	}
}

func NewPutBuildConfigRequest(id string) *UpdateBuildConfigRequest {
	return &UpdateBuildConfigRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		Spec:       NewCreateBuildConfigRequest(),
	}
}

func NewPatchBuildConfigRequest(id string) *UpdateBuildConfigRequest {
	return &UpdateBuildConfigRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Spec:       NewCreateBuildConfigRequest(),
	}
}
