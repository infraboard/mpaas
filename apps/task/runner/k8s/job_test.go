package k8s_test

import (
	"testing"
	"unicode"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRun(t *testing.T) {
	jobSpec := tools.MustReadContentFile("test/job.yaml")
	params := job.NewVersionedRunParam("v0.1")
	params.Add(
		&job.RunParam{
			Name:  "cluster_id",
			Value: "k8s-test",
		},
		&job.RunParam{
			Name:  "namespace",
			Value: "default",
		},
		&job.RunParam{
			Name:  "DB",
			Value: "xxx",
		},
		&job.RunParam{
			Name:  job.SYSTEM_VARIABLE_DEPLOY_CONFIG_ID,
			Value: conf.C.DEPLOY_CONFIG_ID,
		},
	)

	req := task.NewRunTaskRequest("test-job", jobSpec, params)
	ins, err := impl.Run(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestXxx(t *testing.T) {
	t.Log(unicode.IsUpper('_'))
}
