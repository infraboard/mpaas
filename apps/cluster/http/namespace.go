package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/binding"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mpaas/provider/k8s"

	v1 "k8s.io/api/core/v1"
)

func (h *handler) registryNamespaceHandler(r router.SubRouter) {
	ns := r.ResourceRouter("namespace")
	ns.BasePath("clusters/:id/namespaces")
	ns.Handle("GET", "/", h.QueryNamespaces).AddLabel(label.List)
	ns.Handle("POST", "/", h.CreateNamespaces).AddLabel(label.Create)
}

func (h *handler) QueryNamespaces(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.ListNamespace(r.Context(), k8s.NewListRequest())
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateNamespaces(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &v1.Namespace{}
	if err := binding.Bind(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateNamespace(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
