package rpc

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func init() {
	ioc.Config().Registry(&Mpaas{})
}

type Mpaas struct {
	ioc.ObjectImpl
	cs *ClientSet
}

func (m *Mpaas) String() string {
	return pretty.ToJSON(m)
}

func (m *Mpaas) Name() string {
	return MPAAS
}

func (m *Mpaas) Init() error {
	cs, err := NewClient()
	if err != nil {
		return err
	}
	m.cs = cs
	return nil
}

func (c *Mpaas) Close(ctx context.Context) error {
	if c.cs != nil {
		c.cs.conn.Close()
	}

	return nil
}
