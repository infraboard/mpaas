package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/mpaas/apps/cluster"
)

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	req := cluster.NewCreateClusterRequest()

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) QueryBook(w http.ResponseWriter, r *http.Request) {
	req := cluster.NewQueryClusterRequestFromHTTP(r)
	set, err := h.service.QueryCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := cluster.NewDescribeClusterRequest(ctx.PS.ByName("id"))
	ins, err := h.service.DescribeCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) PutBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := cluster.NewPutClusterRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := cluster.NewPatchClusterRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := cluster.NewDeleteClusterRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
