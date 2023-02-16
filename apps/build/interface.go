package build

import (
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
	return validate.Validate(req)
}

func NewQueryBuildConfigRequest() *QueryBuildConfigRequest {
	return &QueryBuildConfigRequest{
		Page: request.NewDefaultPageRequest(),
	}
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
