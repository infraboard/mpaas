package protocol

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcenter/clients/rpc/middleware"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/infraboard/mcube/grpc/middleware/recovery"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/application"
	"github.com/infraboard/mcube/ioc/config/logger"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		rc.UnaryServerInterceptor(),
		otelgrpc.UnaryServerInterceptor(),
		middleware.GrpcAuthUnaryServerInterceptor(application.App().AppName),
	))

	// 控制Grpc启动其他服务, 比如注册中心
	ctx, cancel := context.WithCancel(context.Background())

	return &GRPCService{
		svr: grpcServer,
		l:   logger.Sub("server.grpc"),
		c:   application.App().GRPC,

		ctx:    ctx,
		cancel: cancel,
		client: rpc.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   *zerolog.Logger
	c   *application.Grpc

	ctx    context.Context
	cancel context.CancelFunc
	ins    *instance.Instance
	client *rpc.ClientSet
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	ioc.LoadGrpcController(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.Addr())
	if err != nil {
		s.l.Error().Msgf("listen grpc tcp conn error, %s", err)
		return
	}

	time.AfterFunc(1*time.Second, s.registry)

	s.l.Info().Msgf("GRPC 服务监听地址: %s", s.c.Addr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info().Msg("service is stopped")
		}

		s.l.Error().Msgf("start grpc service error, %s", err.Error())
		return
	}
}

func (s *GRPCService) registry() {
	req := instance.NewRegistryRequest()
	req.Address = s.c.Addr()
	ins, err := s.client.Instance().RegistryInstance(s.ctx, req)
	if err != nil {
		s.l.Error().Msgf("registry to mcenter error, %s", err)
		return
	}
	s.ins = ins

	s.l.Info().Msgf("registry instance to mcenter success")
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	// 提前 剔除注册中心的地址
	if s.ins != nil {
		req := instance.NewUnregistryRequest(s.ins.Id)
		if _, err := s.client.Instance().UnRegistryInstance(s.ctx, req); err != nil {
			s.l.Error().Msgf("unregistry error, %s", err)
		} else {
			s.l.Info().Msg("unregistry success")
		}
	}

	s.svr.GracefulStop()

	s.cancel()
	return nil
}
