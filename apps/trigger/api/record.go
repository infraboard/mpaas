package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/trigger"
)

func (h *Handler) QueryRecord(r *restful.Request, w *restful.Response) {
	req := trigger.NewQueryRecordRequestFromHTTP(r.Request)

	set, err := h.svc.QueryRecord(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
