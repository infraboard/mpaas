package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/clients/rpc/middleware"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/common/terminal"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/provider/k8s/workload"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryPodHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] Pod管理"}

	ws.Route(ws.POST("/{cluster_id}/pods").To(h.CreatePod).
		Doc("创建Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))

	ws.Route(ws.GET("/{cluster_id}/pods").To(h.QueryPods).
		Doc("查询Pod列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.PodList{}).
		Returns(200, "OK", corev1.PodList{}))

	ws.Route(ws.GET("/{cluster_id}/pods/{name}").To(h.GetPod).
		Doc("查询Pod详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))

	ws.Route(ws.GET("/{cluster_id}/pods/{name}/login").To(h.LoginContainer).
		Doc("登陆Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))

	ws.Route(ws.GET("/{cluster_id}/pods/{name}/log").To(h.WatchConainterLog).
		Doc("查看Pod日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))
}

func (h *handler) CreatePod(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	pod := &corev1.Pod{}
	if err := r.ReadEntity(pod); err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewCreateRequest()
	ins, err := client.WorkLoad().CreatePod(r.Request.Context(), pod, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryPods(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.WorkLoad().ListPod(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetPod(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.WorkLoad().GetPod(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
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
func (h *handler) LoginContainer(r *restful.Request, w *restful.Response) {
	ws, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		response.Failed(w, err)
		return
	}

	term := terminal.NewWebSocketTerminal(ws)
	term.SetAuditor(os.Stdout)

	// 开启认证与鉴权
	entry := endpoint.NewEntryFromRestRequest(r).
		SetAuthEnable(true).
		SetPermissionEnable(true)
	err = middleware.GetHttpAuther().PermissionCheck(r, w, entry)
	if err != nil {
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
	fmt.Println(req)
	err = client.WorkLoad().LoginContainer(r.Request.Context(), req)
	if err != nil {
		term.Failed(err)
		return
	}

	term.Success("ok")
}

// Watch Container Log Websocket
func (h *handler) WatchConainterLog(r *restful.Request, w *restful.Response) {
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
	err = middleware.GetHttpAuther().PermissionCheck(r, w, entry)
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
