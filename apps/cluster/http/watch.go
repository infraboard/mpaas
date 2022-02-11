package http

import (
	"io"
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/provider/k8s"
)

// Watch 资源变化
func (h *handler) Watch(w http.ResponseWriter, r *http.Request) {
	term, err := h.newWebsocketTerminal(w, r)
	if err != nil {
		h.log.Errorf("new websocket terminal error, %s", err)
		response.Failed(w, err)
		return
	}
	defer term.Close()

	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		term.WriteMessage(k8s.NewOperatinonParamMessage(err.Error()))
		return
	}

	// 获取参数
	req := k8s.NewWatchRequest()
	term.ParseParame(req)

	wi, err := client.Watch(r.Context(), req)
	if err != nil {
		term.WriteMessage(k8s.NewOperatinonParamMessage(err.Error()))
		return
	}

	reader := k8s.NewWatchReader(wi)
	// 读取出来的数据流 copy到term
	_, err = io.Copy(term, reader)
	if err != nil {
		h.log.Errorf("copy log to weboscket error, %s", err)
	}
}
