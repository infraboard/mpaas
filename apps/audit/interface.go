package audit

import (
	request "github.com/infraboard/mcube/http/request"
	resource "github.com/infraboard/mcube/pb/resource"
)

const (
	AppName = "audits"
)

type Service interface {
	RPCServer
}

func NewQueryRecordRequest() *QueryRecordRequest {
	return &QueryRecordRequest{
		Page:    request.NewDefaultPageRequest(),
		Filters: []*resource.LabelRequirement{},
		Label:   map[string]string{},
		Ids:     []string{},
	}
}
