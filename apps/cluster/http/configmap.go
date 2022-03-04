package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/binding"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mpaas/provider/k8s"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryConfigMapHandler(r router.SubRouter) {
	dr := r.ResourceRouter("configmap")
	dr.BasePath("clusters/:id/configmaps")
	dr.Handle("GET", "/", h.QueryConfigMap).AddLabel(label.List)
	dr.Handle("POST", "/", h.CreateConfigMap).AddLabel(label.Create)
}

func (h *handler) QueryConfigMap(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewListRequestFromHttp(r)
	ins, err := client.ListConfigMap(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateConfigMap(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &corev1.ConfigMap{}
	if err := binding.Bind(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateConfigMap(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
