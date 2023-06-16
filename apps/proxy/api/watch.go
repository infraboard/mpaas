package api

import (
	"github.com/emicklei/go-restful/v3"
)

func (h *handler) registryWatchHandler(ws *restful.WebService) {
}

// Watch 资源变化
func (h *handler) Watch(r *restful.Request, w *restful.Response) {
	// term, err := h.newWebsocketTerminal(w, r.Request)
	// if err != nil {
	// 	h.log.Errorf("new websocket terminal error, %s", err)
	// 	response.Failed(w, err)
	// 	return
	// }
	// defer term.Close()

	// client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	// // 获取参数
	// req := k8s.NewWatchRequest()
	// term.ParseParame(req)

	// wi, err := client.Watch(r.Request.Context(), req)
	// if err != nil {
	// 	term.WriteMessage(k8s.NewOperatinonParamMessage(err.Error()))
	// 	return
	// }

	// reader := k8s.NewWatchReader(wi)
	// // 读取出来的数据流 copy到term
	// _, err = io.Copy(term, reader)
	// if err != nil {
	// 	h.log.Errorf("copy log to weboscket error, %s", err)
	// }
}
