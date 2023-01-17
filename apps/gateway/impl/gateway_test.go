package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/gateway"
)

func TestCreateGateway(t *testing.T) {
	req := gateway.NewCreateGatewayRequest()
	ins, err := impl.CreateGateway(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
