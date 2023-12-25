package traefik

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/infraboard/mpaas/apps/gateway"
	"github.com/rs/zerolog"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdOperator(config *gateway.TraefikConfig) (*EtcdOperator, error) {
	etcdConfig := clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: time.Duration(5) * time.Second,
		Username:    config.Username,
		Password:    config.Password,
	}
	if config.EnableTls {
		etcdConfig.TLS = &tls.Config{}
	}
	client, err := clientv3.New(etcdConfig)
	if err != nil {
		return nil, err
	}
	op := &EtcdOperator{
		rootKey: config.RootKey,
		client:  client,
		log:     log.Sub("traefik.etcd"),
	}
	return op, nil
}

type EtcdOperator struct {
	client  *clientv3.Client
	log     *zerolog.Logger
	rootKey string
}

func (o *EtcdOperator) ListKeys(ctx context.Context, key string) ([]string, error) {
	keys := []string{}
	resp, err := o.client.Get(ctx, o.rootKey+key, clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}
	for i := range resp.Kvs {
		keys = append(keys, string(resp.Kvs[i].Key))
	}

	return keys, nil
}
