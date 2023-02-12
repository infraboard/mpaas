package gateway

import (
	"encoding/json"

	meta "github.com/infraboard/mpaas/common/meta"
)

func New(req *CreateGatewayRequest) (*Gateway, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Gateway{
		Meta: meta.NewMeta(),
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
		*meta.Meta
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
