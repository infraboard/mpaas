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

func NewPutDeployRequest(id string) *UpdateBuildConfigRequest {
	return &UpdateBuildConfigRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		Spec:       NewCreateBuildConfigRequest(),
	}
}

func NewPatchDeployRequest(id string) *UpdateBuildConfigRequest {
	return &UpdateBuildConfigRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Spec:       NewCreateBuildConfigRequest(),
	}
}
