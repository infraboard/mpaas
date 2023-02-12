package gateway

import (
	context "context"

	"github.com/infraboard/mcenter/common/validate"
)

const (
	AppName = "gateways"
)

type Service interface {
	CreateGateway(context.Context, *CreateGatewayRequest) (*Gateway, error)
	RPCServer
}

func NewDefaultTraefikConfig() *TraefikConfig {
	return &TraefikConfig{
		Endpoints: []string{"127.0.0.1:2379"},
	}
}

func (req *CreateGatewayRequest) Validate() error {
	return validate.Validate(req)
}

func NewCreateGatewayRequest() *CreateGatewayRequest {
	return &CreateGatewayRequest{}
}

func (req *QueryGatewayRequest) Validate() error {
	return validate.Validate(req)
}

func (req *DescribeGatewayRequest) Validate() error {
	return validate.Validate(req)
}
