package workload_test

import (
	"io"
	"os"
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/workload"
	"github.com/infraboard/mpaas/test/tools"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestWatchConainterLog(t *testing.T) {
	req := workload.NewWatchConainterLogRequest()
	req.Namespace = "default"
	req.PodName = "cicd-test-8f97775-h78sj"
	stream, err := impl.WatchConainterLog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()
	_, err = io.Copy(os.Stdout, stream)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInjectEnvVars(t *testing.T) {
	obj := new(batchv1.Job)
	tools.MustReadYamlFile("test/job.yml", obj)

	// 给容器注入环境变量
	for i, c := range obj.Spec.Template.Spec.Containers {
		workload.InjectContainerEnvVars(&c, []corev1.EnvVar{
			{
				Name:  "DB_PASS",
				Value: "test",
			},
		})
		obj.Spec.Template.Spec.Containers[i] = c
	}

	t.Log(tools.MustToYaml(obj))
}
