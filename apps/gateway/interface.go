package gateway

import (
	context "context"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mpaas/common/meta"
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

func NewDefaultGateway() *Gateway {
	return &Gateway{
		Spec: NewCreateGatewayRequest(),
	}
}

func NewCreateGatewayRequest() *CreateGatewayRequest {
	return &CreateGatewayRequest{}
}

func New(req *CreateGatewayRequest) (*Gateway, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Gateway{
		Meta: meta.NewMeta(),
		Spec: req,
	}, nil
}

func NewGatewaySet() *GatewaySet {
	return &GatewaySet{
		Items: []*Gateway{},
	}
}

func (s *GatewaySet) Add(item *Gateway) {
	s.Items = append(s.Items, item)
}

func (req *QueryGatewayRequest) Validate() error {
	return validate.Validate(req)
}

func (req *DescribeGatewayRequest) Validate() error {
	return validate.Validate(req)
}
