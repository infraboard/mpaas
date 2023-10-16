package rpc

import "github.com/infraboard/mcube/ioc"

const (
	MPAAS = "mpaas_client"
)

func C() *ClientSet {
	return ioc.Config().Get(MPAAS).(*Mpaas).cs
}
