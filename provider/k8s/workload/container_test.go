package workload_test

import (
	"io"
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/workload"
)

func TestWatchConainterLog(t *testing.T) {
	req := workload.NewWatchConainterLogRequest()
	req.Namespace = "default"
	req.PodName = "test-job-kscwv"
	stream, err := impl.WatchConainterLog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()
	b, err := io.ReadAll(stream)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
