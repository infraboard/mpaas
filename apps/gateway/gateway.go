package gateway

import (
	"encoding/json"

	resource "github.com/infraboard/mcube/v2/pb/resource"
)

func New(req *CreateGatewayRequest) (*Gateway, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Gateway{
		Meta: resource.NewMeta(),
		Spec: req,
	}, nil
}

func NewDefaultGateway() *Gateway {
	return &Gateway{
		Spec: NewCreateGatewayRequest(),
	}
}

func (d *Gateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateGatewayRequest
	}{d.Meta, d.Spec})
}

func NewGatewaySet() *GatewaySet {
	return &GatewaySet{
		Items: []*Gateway{},
	}
}

func (s *GatewaySet) Add(item *Gateway) {
	s.Items = append(s.Items, item)
}
