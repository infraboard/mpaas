package start

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/infraboard/mcube/ioc"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/infraboard/mcube/ioc/config/logger"

	"github.com/infraboard/mpaas/protocol"

	// 注册所有服务
	_ "github.com/infraboard/mpaas/apps"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "mpaas API服务",
	Long:  "mpaas API服务",
	Run: func(cmd *cobra.Command, args []string) {
		// 启动服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 初始化服务
		svr, err := newService()
		cobra.CheckErr(err)

		// 启动服务
		svr.start(context.Background())
	},
}

func newService() (*service, error) {
	http := protocol.NewHTTPService()
	grpc := protocol.NewGRPCService()

	svr := &service{
		http: http,
		grpc: grpc,
		log:  logger.Sub("cli"),
	}

	return svr, nil
}

type service struct {
	http *protocol.HTTPService
	grpc *protocol.GRPCService
	ch   chan os.Signal
	log  *zerolog.Logger
}

func (s *service) start(ctx context.Context) {
	s.log.Info().Msgf("loaded configs: %s", ioc.Config().List())
	s.log.Info().Msgf("loaded controllers: %s", ioc.Config().List())
	s.log.Info().Msgf("loaded apis: %s", ioc.Api().List())

	go s.grpc.Start()
	go s.http.Start(ctx)
	s.waitSign(ctx)
}

func (s *service) waitSign(ctx context.Context) {
	// 处理信号量
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	for sg := range ch {
		switch v := sg.(type) {
		default:
			s.log.Info().Msgf("receive signal '%v', start graceful shutdown", v.String())

			if err := s.grpc.Stop(); err != nil {
				s.log.Error().Msgf("grpc graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Info().Msg("grpc service stop complete")
			}

			if err := s.http.Stop(ctx); err != nil {
				s.log.Error().Msgf("http graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Info().Msgf("http service stop complete")
			}
			return
		}
	}
}
