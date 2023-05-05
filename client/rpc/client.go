package rpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/resolver"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
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
	log := zap.L().Named("sdk.mpaas")

	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout())
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s", resolver.Scheme, "mpaas"),
		grpc.WithPerRPCCredentials(rpc.NewAuthenticationFromEnv()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
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
	log  logger.Logger
}

// Job Task 管理接口
func (s *ClientSet) JobTask() task.JobRPCClient {
	return task.NewJobRPCClient(s.conn)
}

// Deploy 管理接口
func (s *ClientSet) Deploy() deploy.RPCClient {
	return deploy.NewRPCClient(s.conn)
}
