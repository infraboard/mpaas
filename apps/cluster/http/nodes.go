package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"
)

func (h *handler) registryNodeHandler(r router.SubRouter) {
	nr := r.ResourceRouter("node")
	nr.BasePath("clusters/:id/nodes")
	nr.Handle("GET", "/", h.QueryNodes).AddLabel(label.List)
}

func (h *handler) QueryNodes(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.ListNode(r.Context())
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
