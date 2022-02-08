package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryNodes(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.ListNodes(r.Context())
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
