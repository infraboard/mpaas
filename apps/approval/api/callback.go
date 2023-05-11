package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/apps/approval/api/callback"
)

func init() {
	app.RegistryRESTfulApp(&callbackHandler{})
}

type callbackHandler struct {
	service approval.Service
	log     logger.Logger
	mcenter *rpc.ClientSet
}

func (h *callbackHandler) Config() error {
	h.log = zap.L().Named(approval.AppName)
	h.service = app.GetGrpcApp(approval.AppName).(approval.Service)
	h.mcenter = rpc.C()
	return nil
}

func (h *callbackHandler) Name() string {
	return "callbacks"
}

func (h *callbackHandler) Version() string {
	return "v1"
}

func (h *callbackHandler) Registry(ws *restful.WebService) {
	tags := []string{"审核对接第三方"}
	ws.Route(ws.POST("/feishu").To(h.FeishuCard).
		Doc("飞书卡片处理审核").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(approval.CreateApprovalRequest{}).
		Writes(approval.Approval{}))
}

// 飞书卡片回调, 参考文档: https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
func (h *callbackHandler) FeishuCard(r *restful.Request, w *restful.Response) {
	req := callback.NewFeishuCardCallback()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	// 根据飞书UserId查询用户信息
	descReq := user.NewDescriptUserRequestByFeishuUserId(req.UserId)
	u, err := h.mcenter.User().DescribeUser(r.Request.Context(), descReq)
	if err != nil {
		response.Failed(w, err)
		return
	}

	in, err := req.BuildUpdateApprovalStatusRequest(u.Meta.Id)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins, err := h.service.UpdateApprovalStatus(r.Request.Context(), in)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
