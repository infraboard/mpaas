package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"构建配置管理"}
	fmt.Println(tags)
}
