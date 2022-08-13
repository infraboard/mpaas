package http

import (
	"net/http"

	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/mcube/http/binding"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/mpaas/apps/cluster"
)

func (h *handler) registryClusterHandler(r router.SubRouter) {
	rr := r.ResourceRouter("cluster")
	rr.BasePath("clusters")
	rr.Handle("POST", "/", h.CreateCluster).AddLabel(label.Create)
	rr.Handle("GET", "/", h.QueryCluster).AddLabel(label.List)
	rr.Handle("GET", "/:id", h.DescribeCluster).AddLabel(label.Get)
	rr.Handle("PUT", "/:id", h.PutCluster).AddLabel(label.Update)
	rr.Handle("PATCH", "/:id", h.PatchCluster).AddLabel(label.Update)
	rr.Handle("DELETE", "/:id", h.DeleteCluster).AddLabel(label.Delete)
}

func (h *handler) CreateCluster(w http.ResponseWriter, r *http.Request) {
	req := cluster.NewCreateClusterRequest()

	if err := binding.Bind(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner()
	set, err := h.service.CreateCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryCluster(w http.ResponseWriter, r *http.Request) {
	req := cluster.NewQueryClusterRequestFromHTTP(r)
	req.UpdateNamespace()

	set, err := h.service.QueryCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	set.Desense()
	response.Success(w, set)
}

func (h *handler) DescribeCluster(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := cluster.NewDescribeClusterRequest(ctx.PS.ByName("id"))
	ins, err := h.service.DescribeCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desense()
	response.Success(w, ins)
}

func (h *handler) PutCluster(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)
	req := cluster.NewPutClusterRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Account

	set, err := h.service.UpdateCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchCluster(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)
	req := cluster.NewPatchClusterRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Account

	set, err := h.service.UpdateCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteCluster(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := cluster.NewDeleteClusterRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteCluster(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
