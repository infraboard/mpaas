package rpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcenter/clients/rpc/resolver"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcube/grpc/middleware/exception"
	"github.com/infraboard/mcube/ioc/config/logger"
)

func NewClientSetFromEnv() (*ClientSet, error) {
	// 从环境变量中获取mcenter配置
	mc, err := rpc.NewConfigFromEnv()
	if err != nil {
		return nil, err
	}

	return NewClientSetFromConfig(mc)
}

// NewClient todo
func NewClientSetFromConfig(conf *rpc.Config) (*ClientSet, error) {
	log := logger.Sub("sdk.mpaas")

	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout())
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s?%s", resolver.Scheme, "mpaas", conf.Resolver.ToQueryString()),
		grpc.WithPerRPCCredentials(conf.Credentials()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),

		// 将异常转化为 API Exception
		grpc.WithChainUnaryInterceptor(exception.NewUnaryClientInterceptor()),

		// Grpc Trace
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conf: conf,
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type ClientSet struct {
	conf *rpc.Config
	conn *grpc.ClientConn
	log  *zerolog.Logger
}

// 关闭GRPC连接
func (c *ClientSet) Stop() error {
	return c.conn.Close()
}

// Deploy 管理接口
func (s *ClientSet) Deploy() deploy.RPCClient {
	return deploy.NewRPCClient(s.conn)
}
