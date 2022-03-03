package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/binding"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mpaas/provider/k8s"

	appsv1 "k8s.io/api/apps/v1"
)

func (h *handler) registryDeploymentHandler(r router.SubRouter) {
	dr := r.ResourceRouter("deployment")
	dr.BasePath("clusters/:id/deployments")
	dr.Handle("GET", "/", h.QueryDeployments).AddLabel(label.List)
	dr.Handle("POST", "/", h.CreateDeployment).AddLabel(label.Create)
}

func (h *handler) QueryDeployments(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewListDeploymentRequestFromHttp(r)
	ins, err := client.ListDeployment(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateDeployment(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &appsv1.Deployment{}
	if err := binding.Bind(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateDeployment(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
