package build

import (
	"net/http"

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

func (req *CreateBuildConfigRequest) PipielineId() string {
	switch req.TargetType {
	case TARGET_TYPE_IMAGE:
		return req.ImageBuild.PipelineId
	case TARGET_TYPE_PKG:
		return req.PkgBuild.PipelineId
	}
	return ""
}

func NewQueryBuildConfigRequestFromHTTP(r *http.Request) *QueryBuildConfigRequest {
	return &QueryBuildConfigRequest{
		Page: request.NewPageRequestFromHTTP(r),
	}
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
