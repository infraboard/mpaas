package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mpaas/apps/deploy"
)

func (h *downloadHandler) Registry() {
	tags := []string{"部署配置下载"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.GET("/deploys/{id}").To(h.DownloadDeployment).
		Doc("下载配置详情").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Writes(deploy.Deployment{}).
		Returns(200, "OK", deploy.Deployment{}).
		Returns(404, "Not Found", nil))
}

func (h *downloadHandler) DownloadDeployment(r *restful.Request, w *restful.Response) {
	req := deploy.NewDescribeDeploymentRequest(r.PathParameter("id"))

	// 查询部署配置
	ins, err := h.service.DescribeDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 校验集群访问Token
	if ins.Spec.AuthEnabled {
		err = ins.ValidateToken(r.QueryParameter("token"))
		if err != nil {
			response.Failed(w, err)
			return
		}
	}

	switch ins.Spec.Type {
	case deploy.TYPE_HOST:
	case deploy.TYPE_KUBERNETES:
		w.Write([]byte(ins.Spec.K8STypeConfig.WorkloadConfig))
	}
}
