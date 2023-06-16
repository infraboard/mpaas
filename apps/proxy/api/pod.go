package api

import (
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryPodHandler(ws *restful.WebService) {
	tags := []string{"Pod管理"}

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
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Pod{}).
		Returns(200, "OK", corev1.Pod{}))

	ws.Route(ws.GET("/{cluster_id}/pods/{name}/login").To(h.WatchConainterLog).
		Doc("查看Pod日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
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

var (
	defaultCmd = `TERM=xterm-256color; export TERM; [ -x /bin/bash ] && ([ -x /usr/bin/script ] && /usr/bin/script -q -c "/bin/bash" /dev/null || exec /bin/bash) || exec /bin/sh`
)

// Login Container Websocket
func (h *handler) LoginContainer(r *restful.Request, w *restful.Response) {
	// term, err := h.newWebsocketTerminal(w, r.Request)
	// if err != nil {
	// 	h.log.Errorf("new websocket terminal error, %s", err)
	// 	response.Failed(w, err)
	// 	return
	// }
	// defer term.Close()

	// client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	// // 获取参数
	// req := workload.NewLoginContainerRequest([]string{"sh", "-c", defaultCmd}, term)
	// term.ParseParame(req)

	// err = client.WorkLoad().LoginContainer(r.Request.Context(), req)
	// if err != nil {
	// 	// term.WriteMessage(k8s.NewOperatinonParamMessage(err.Error()))
	// 	return
	// }
}

// Watch Container Log Websocket
func (h *handler) WatchConainterLog(r *restful.Request, w *restful.Response) {
	// term, err := h.newWebsocketTerminal(w, r.Request)
	// if err != nil {
	// 	h.log.Errorf("new websocket terminal error, %s", err)
	// 	response.Failed(w, err)
	// 	return
	// }
	// defer term.Close()

	// client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	// // 获取参数
	// req := workload.NewWatchConainterLogRequest()
	// term.ParseParame(req)

	// reader, err := client.WorkLoad().WatchConainterLog(r.Request.Context(), req)
	// if err != nil {
	// 	// term.WriteMessage(k8s.NewOperatinonParamMessage(err.Error()))
	// 	return
	// }

	// // 读取出来的数据流 copy到term
	// _, err = io.Copy(term, reader)
	// if err != nil {
	// 	h.log.Errorf("copy log to weboscket error, %s", err)
	// }
}

// func (h *handler) newWebsocketTerminal(w http.ResponseWriter, r *http.Request) (*k8s.WebsocketTerminal, error) {
// 	// websocket handshake
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	term := k8s.NewWebsocketTerminal(ws)
// 	term.Auth(h.websocketAuth)
// 	return term, nil
// }

func (h *handler) websocketAuth(payload string) error {
	h.log.Debugf("auth payload: %s", payload)
	return nil
}
