package task_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestNewPipelineTask(t *testing.T) {
	p := pipeline.NewDefaultPipeline()
	tools.MustReadYamlFile("impl/test/pipeline.yml", p)
	pt := task.NewPipelineTask(p)
	t.Log(tools.MustToYaml(pt))

	tasks := pt.NextRun()
	t.Log(tools.MustToYaml(tasks))
}
