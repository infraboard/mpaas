package workload_test

import (
	"testing"
	"time"

	"github.com/infraboard/mpaas/provider/k8s/workload"
	"sigs.k8s.io/yaml"
)

func TestCopyPodRun(t *testing.T) {
	req := workload.NewCopyPodRunRequest()
	req.SourcePod.Namespace = "default"
	req.SourcePod.Name = "task-ci4hr3ro99m7irvib5jg-rjnk5"
	req.TargetPodMeta.Namespace = "default"
	req.TargetPodMeta.Name = "task-ci4hr3ro99m7irvib5jg-rjnk5-debug01"
	req.ExecHoldCmd = workload.HoldContaienrCmd(1 * time.Hour)

	pods, err := impl.CopyPodRun(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	yd, err := yaml.Marshal(pods)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(yd))
}
