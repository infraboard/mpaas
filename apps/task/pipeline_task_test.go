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

	tasks := pt.JobTasks()
	for i := range tasks.Items {
		task := tasks.Items[i]
		t.Log(task.Id, task.Spec.Job, task.Status.Stage)
	}

	// 即将运行的tasks
	nexts := pt.NextRun()
	for nexts.Len() > 0 {
		for i := range nexts.Items {
			next := nexts.Items[i]
			next.Status.MarkedSuccess()
			t.Log(next.Id, next.Spec.Job, next.Status.Stage)
		}

		nexts = pt.NextRun()
	}
}
