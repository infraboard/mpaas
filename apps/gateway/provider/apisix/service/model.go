package service

import (
	"github.com/infraboard/mcube/tools/pretty"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix"
)

type ServiceList struct {
	Total int        `json:"total"`
	List  []*Service `json:"list"`
}

type Service struct {
	*apisix.Meta
	*CreateServiceRequest
}

type CreateServiceRequest struct {
}

func (r *CreateServiceRequest) ToJSON() string {
	return pretty.ToJSON(r)
}
