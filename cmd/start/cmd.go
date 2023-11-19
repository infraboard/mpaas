package start

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/apps/instance"
	mcenter "github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcenter/clients/rpc/middleware"
	"github.com/infraboard/mcenter/clients/rpc/tools"
	"github.com/infraboard/mcube/ioc/config/application"
	"github.com/infraboard/mcube/ioc/config/logger"

	// 注册所有服务
	_ "github.com/infraboard/mpaas/apps"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "mpaas API服务",
	Long:  "mpaas API服务",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(newService().Start())
	},
}

func newService() *service {
	return &service{
		log: logger.Sub("cmd"),
	}
}

type service struct {
	log *zerolog.Logger

	ins *instance.Instance
}

func (s *service) Start() error {
	application.App().UseGoRestful()

	// 补充权限与注册
	application.App().HTTP.RouterBuildConfig.BeforeLoad = s.HttpBeforeLoad
	application.App().HTTP.RouterBuildConfig.AfterLoad = s.HttpAfterLoad

	// 补充Grpc认证
	application.App().GRPC.AddInterceptors(middleware.GrpcAuthUnaryServerInterceptor())
	application.App().GRPC.PostStart = s.GrpcPostStart
	application.App().GRPC.PreStop = s.GrpcPreStop

	return application.App().Start(context.Background())
}

func (s *service) HttpBeforeLoad(r http.Handler) {
	if router, ok := r.(*restful.Container); ok {
		// 认证中间件
		router.Filter(middleware.RestfulServerInterceptor())
	}
}

func (s *service) HttpAfterLoad(r http.Handler) {
	if router, ok := r.(*restful.Container); ok {
		// 注册服务权限条目
		s.log.Info().Msg("start registry endpoints ...")

		register := tools.NewEndpointRegister()
		err := register.Registry(context.Background(), router, application.Short())
		if err != nil {
			s.log.Warn().Msgf("registry endpoints error, %s", err)
		} else {
			s.log.Debug().Msg("service endpoints registry success")
		}
	}
}

func (s *service) GrpcPostStart(ctx context.Context) error {
	mcenter := mcenter.C()

	req := instance.NewRegistryRequest()
	req.Address = application.App().GRPC.Addr()
	ins, err := mcenter.Instance().RegistryInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("registry to mcenter error, %s", err)
	}
	s.ins = ins

	s.log.Info().Msgf("registry instance to mcenter success")
	return nil
}

func (s *service) GrpcPreStop(ctx context.Context) error {
	mcenter := mcenter.C()

	// 提前 剔除注册中心的地址
	if s.ins != nil {
		req := instance.NewUnregistryRequest(s.ins.Id)
		if _, err := mcenter.Instance().UnRegistryInstance(ctx, req); err != nil {
			s.log.Error().Msgf("unregistry error, %s", err)
		} else {
			s.log.Info().Msg("unregistry success")
		}
	}
	return nil
}
