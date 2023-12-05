package rpc

import "github.com/infraboard/mcube/v2/ioc"

const (
	MPAAS = "mpaas"
)

func C() *ClientSet {
	return ioc.Config().Get(MPAAS).(*Mpaas).cs
}
