package protocol

import (
	"context"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/apps/apidoc"
	"github.com/infraboard/mcube/ioc/apps/health"
	"github.com/infraboard/mcube/ioc/config/application"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcenter/clients/rpc/middleware"

	"github.com/infraboard/mpaas/swagger"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	r := restful.DefaultContainer
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)

	// CORS中间件
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedDomains: []string{"*"},
		AllowedMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		Container:      r,
	}
	r.Filter(cors.Filter)
	// trace中间件
	filter := otelrestful.OTelFilter(application.App().AppName)
	restful.DefaultContainer.Filter(filter)
	// 认证中间件
	r.Filter(middleware.RestfulServerInterceptor())

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              application.App().HTTP.Addr(),
		Handler:           r,
	}

	return &HTTPService{
		r:          r,
		server:     server,
		l:          logger.Sub("server.http"),
		c:          application.App().HTTP,
		endpoint:   rpc.C().Endpoint(),
		apiDocPath: "/apidocs.json",
	}
}

// HTTPService http服务
type HTTPService struct {
	r      *restful.Container
	l      *zerolog.Logger
	c      *application.Http
	server *http.Server

	endpoint   endpoint.RPCClient
	apiDocPath string
}

func (s *HTTPService) PathPrefix() string {
	return application.App().HTTPPrefix()
}

// Start 启动服务
func (s *HTTPService) Start(ctx context.Context) {
	// 装置子服务路由
	ioc.LoadGoRestfulApi(s.PathPrefix(), s.r)

	// API Doc
	s.r.Add(apidoc.APIDocs(s.apiDocPath, swagger.Docs))
	s.l.Info().Msgf("Get the API Doc using http://%s%s", s.c.Addr(), s.apiDocPath)

	// HealthCheck
	hc := health.NewDefaultHealthChecker()
	s.r.Add(hc.WebService())
	s.l.Info().Msgf("健康检查地址: http://%s%s", s.c.Addr(), hc.HealthCheckPath)

	// 注册路由条目
	s.RegistryEndpoint(ctx)

	// 启动 HTTP服务
	s.l.Info().Msgf("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info().Msg("service is stopped")
			return
		}
		s.l.Error().Msgf("start service error, %s", err.Error())
	}
}

// Stop 停止server
func (s *HTTPService) Stop(ctx context.Context) error {
	s.l.Info().Msg("start graceful shutdown")
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Error().Msgf("graceful shutdown timeout, force exit")
	}
	return nil
}

func (s *HTTPService) RegistryEndpoint(ctx context.Context) {
	// 注册服务权限条目
	s.l.Info().Msg("start registry endpoints ...")

	entries := []*endpoint.Entry{}
	wss := s.r.RegisteredWebServices()
	for i := range wss {
		es := endpoint.TransferRoutesToEntry(wss[i].Routes())
		entries = append(entries, endpoint.GetPRBACEntry(es)...)
	}

	req := endpoint.NewRegistryRequest(application.Short(), entries)
	_, err := s.endpoint.RegistryEndpoint(ctx, req)
	if err != nil {
		s.l.Warn().Msgf("registry endpoints error, %s", err)
	} else {
		s.l.Debug().Msg("service endpoints registry success")
	}
}
