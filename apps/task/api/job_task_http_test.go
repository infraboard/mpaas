package api_test

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/task/api"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl *api.JobTaskHandler
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetApi("job_tasks").(*api.JobTaskHandler)
}
