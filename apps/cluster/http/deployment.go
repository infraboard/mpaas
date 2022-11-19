package http

import (
	"io/ioutil"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mpaas/provider/k8s"
	"sigs.k8s.io/yaml"

	appsv1 "k8s.io/api/apps/v1"
)

func (h *handler) registryDeploymentHandler(r router.SubRouter) {
	// dr := r.ResourceRouter("deployment")
	// dr.BasePath("clusters/:id/deployments")
	// dr.Handle("GET", "/", h.QueryDeployments).AddLabel(label.List)
	// dr.Handle("POST", "/", h.CreateDeployment).AddLabel(label.Create)
}

func (h *handler) QueryDeployments(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewListRequestFromHttp(r.Request)
	ins, err := client.ListDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateDeployment(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &appsv1.Deployment{}

	data, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		response.Failed(w, err)
		return
	}
	defer r.Request.Body.Close()

	if err := yaml.Unmarshal(data, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
