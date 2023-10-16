package tools

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"

	// 注册所有服务
	_ "github.com/infraboard/mpaas/apps"
)

func DevelopmentSetup() {
	// 针对http handler的测试需要提前设置默认数据格式
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)

	ioc.DevelopmentSetup()
}
