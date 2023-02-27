package task_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

func TestParseRuntimeEnvFromBytes(t *testing.T) {
	data := tools.MustReadContentFile("impl/test/pipeline.env")
	envs, err := task.ParseRuntimeEnvFromBytes([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(envs)
}
