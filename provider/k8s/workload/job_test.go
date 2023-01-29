package workload_test

import (
	"testing"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/test/tools"
	v1 "k8s.io/api/batch/v1"
)

func TestListJob(t *testing.T) {
	req := meta.NewListRequest()
	req.Namespace = "kube-system"
	req.Opts.LabelSelector = "k8s-app=kube-dns"
	list, err := impl.ListJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	for _, v := range list.Items {
		t.Log(v.Namespace, v.Name)
	}
}

func TestCreateJob(t *testing.T) {
	job := &v1.Job{}
	tools.MustReadYamlFile("test/job.yml", job)
	job, err := impl.CreateJob(ctx, job)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(job)
}
