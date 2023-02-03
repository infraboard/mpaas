package job_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
)

func TestK8SJobRunnerParams(t *testing.T) {
	param := job.NewVersionedRunParam("v0.1")
	param.Add(&job.RunParam{
		Name:     "cluster_id",
		Required: true,
		Value:    "k8s-cluster-01",
	})
	t.Log(param.K8SJobRunnerParams())
}

func TestNewRunParamWithKVPaire(t *testing.T) {
	param := job.NewRunParamWithKVPaire("key1", "value1", "key2", "value2")
	t.Log(param)
}