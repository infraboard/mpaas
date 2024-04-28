package api

import (
	"io"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/infraboard/mcenter/apps/endpoint"
	middleware "github.com/infraboard/mcenter/clients/rpc/middleware/auth/gorestful"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	corev1 "k8s.io/api/core/v1"

	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/common/terminal"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

func init() {
	ioc.Api().Registry(&websocketHandler{})
}

type websocketHandler struct {
	service cluster.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *websocketHandler) Init() error {
	h.log = log.Sub(cluster.AppName)
	h.service = ioc.Controller().Get(cluster.AppName).(cluster.Service)
	h.Registry()
	return nil
}

// /prifix/proxy/
func (h *websocketHandler) Name() string {
	return "ws/proxy"
}

func (h *websocketHandler) Version() string {
	return "v1"
}

func (h *websocketHandler) Registry() {
	r := gorestful.ObjectRouter(h)
	r.Filter(ClusterMiddleware())
	h.registryPodHandler(r)
}

func (h *websocketHandler) registryPodHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] Pod管理"}

	ws.Route(ws.GET("/{cluster_id}/pods/{name}/login").To(h.LoginContainer).
		Doc("登陆Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "pod_terminal").
		Metadata(label.Action, label.Create.Value()).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))

	ws.Route(ws.GET("/{cluster_id}/pods/{name}/log").To(h.WatchConainterLog).
		Doc("查看Pod日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "pod_terminal").
		Metadata(label.Action, label.Get.Value()).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))
}

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: 60 * time.Second,
		ReadBufferSize:   8192,
		WriteBufferSize:  8192,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Login Container Websocket
func (h *websocketHandler) LoginContainer(r *restful.Request, w *restful.Response) {
	ws, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		response.Failed(w, err)
		return
	}

	term := terminal.NewWebSocketTerminal(ws)
	// term.SetAuditor(os.Stdout)

	// 开启认证与鉴权
	entry := endpoint.NewEntryFromRestRequest(r).
		SetAuthEnable(true).
		SetPermissionEnable(true)

	err = middleware.Get().PermissionCheck(r, w, entry)
	if err != nil {
		h.log.Debug().Msgf("login container error, %s", err)
		term.Failed(err)
		return
	}

	// 获取参数
	req := workload.NewLoginContainerRequest(term)
	if err = term.ReadReq(req); err != nil {
		term.Failed(err)
		return
	}

	// 登录容器
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)
	err = client.WorkLoad().LoginContainer(r.Request.Context(), req)
	if err != nil {
		term.Failed(err)
		return
	}

	term.Success("ok")
}

// Watch Container Log Websocket
func (h *websocketHandler) WatchConainterLog(r *restful.Request, w *restful.Response) {
	ws, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		response.Failed(w, err)
		return
	}

	websocket.Subprotocols(r.Request)
	term := terminal.NewWebSocketTerminal(ws)

	// 开启认证与鉴权
	entry := endpoint.NewEntryFromRestRequest(r).
		SetAuthEnable(true).
		SetPermissionEnable(true)
	err = middleware.Get().PermissionCheck(r, w, entry)
	if err != nil {
		term.Failed(err)
		return
	}

	// 获取参数
	req := workload.NewWatchConainterLogRequest()
	if err = term.ReadReq(req); err != nil {
		term.Failed(err)
		return
	}
	req.PodName = r.PathParameter("name")

	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)
	reader, err := client.WorkLoad().WatchConainterLog(r.Request.Context(), req)
	if err != nil {
		term.Failed(err)
		return
	}

	// 读取出来的数据流 copy到term
	_, err = io.Copy(term, reader)
	if err != nil {
		term.Failed(err)
		return
	}

	term.Success("ok")
}
