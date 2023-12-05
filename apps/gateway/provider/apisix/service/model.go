package service

import (
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/common"
)

type ServiceList struct {
	Total int        `json:"total"`
	List  []*Service `json:"list"`
}

type Service struct {
	*common.Meta
	*CreateServiceRequest
}

type CreateServiceRequest struct {
}

func (r *CreateServiceRequest) ToJSON() string {
	return pretty.ToJSON(r)
}
