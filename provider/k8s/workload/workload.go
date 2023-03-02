package workload

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"
)

func (c *Client) Run(ctx context.Context, wl *WorkLoad) (*WorkLoad, error) {
	var err error
	switch wl.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		wl.Deployment, err = c.CreateDeployment(ctx, wl.Deployment)
	case WORKLOAD_KIND_STATEFULSET:
		wl.StatefulSet, err = c.CreateStatefulSet(ctx, wl.StatefulSet)
	case WORKLOAD_KIND_DAEMONSET:
		wl.DaemonSet, err = c.CreateDaemonSet(ctx, wl.DaemonSet)
	case WORKLOAD_KIND_CRONJOB:
		wl.CronJob, err = c.CreateCronJob(ctx, wl.CronJob)
	case WORKLOAD_KIND_JOB:
		wl.Job, err = c.CreateJob(ctx, wl.Job)
	}
	if err != nil {
		return nil, err
	}
	return wl, nil
}

func (c *Client) Delete(ctx context.Context, wl *WorkLoad) (*WorkLoad, error) {
	var err error

	m := wl.GetObjectMeta()
	if m == nil {
		return nil, fmt.Errorf("object not found")
	}
	req := meta.NewDeleteRequest(m.Name).WithNamespace(m.Namespace)

	switch wl.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		err = c.DeleteDeployment(ctx, req)
	case WORKLOAD_KIND_STATEFULSET:
		err = c.DeleteStatefulSet(ctx, req)
	case WORKLOAD_KIND_DAEMONSET:
		err = c.DeleteDaemonSet(ctx, req)
	case WORKLOAD_KIND_CRONJOB:
		err = c.DeleteCronJob(ctx, req)
	case WORKLOAD_KIND_JOB:
		err = c.DeleteJob(ctx, req)
	}
	if err != nil {
		return nil, err
	}
	return wl, nil
}

func ParseWorkloadFromYaml(kindStr string, workload string) (w *WorkLoad, err error) {
	w = NewWorkLoad()
	if kindStr == "" {
		return
	}

	kind, err := ParseWorkloadKindFromString(kindStr)
	if err != nil {
		return nil, err
	}
	switch kind {
	case WORKLOAD_KIND_DEPLOYMENT:
		err = yaml.Unmarshal([]byte(workload), w.Deployment)
	case WORKLOAD_KIND_STATEFULSET:
		err = yaml.Unmarshal([]byte(workload), w.StatefulSet)
	case WORKLOAD_KIND_DAEMONSET:
		err = yaml.Unmarshal([]byte(workload), w.DaemonSet)
	case WORKLOAD_KIND_CRONJOB:
		err = yaml.Unmarshal([]byte(workload), w.CronJob)
	case WORKLOAD_KIND_JOB:
		err = yaml.Unmarshal([]byte(workload), w.Job)
	}
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewWorkLoad() *WorkLoad {
	return &WorkLoad{
		Deployment:  &appsv1.Deployment{},
		StatefulSet: &appsv1.StatefulSet{},
		DaemonSet:   &appsv1.DaemonSet{},
		CronJob:     &batchv1.CronJob{},
		Job:         &batchv1.Job{},
	}
}

type WorkLoad struct {
	WorkloadKind WORKLOAD_KIND
	Deployment   *appsv1.Deployment
	StatefulSet  *appsv1.StatefulSet
	DaemonSet    *appsv1.DaemonSet
	CronJob      *batchv1.CronJob
	Job          *batchv1.Job
}

func (w *WorkLoad) SetAnnotations(key, value string) {
	m := w.GetObjectMeta()
	if m != nil {
		if m.Annotations == nil {
			m.Annotations = map[string]string{}
		}
		m.Annotations[key] = value
	}
}

func (w *WorkLoad) SetDefaultNamespace(ns string) {
	switch w.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		w.Deployment.Namespace = ns
	case WORKLOAD_KIND_STATEFULSET:
		w.StatefulSet.Namespace = ns
	case WORKLOAD_KIND_DAEMONSET:
		w.DaemonSet.Namespace = ns
	case WORKLOAD_KIND_CRONJOB:
		w.CronJob.Namespace = ns
	case WORKLOAD_KIND_JOB:
		w.Job.Namespace = ns
	}
}

func (w *WorkLoad) MustToYaml() string {
	var (
		err  error
		data []byte
	)
	switch w.WorkloadKind {
	case WORKLOAD_KIND_DEPLOYMENT:
		data, err = yaml.Marshal(w.Deployment)
	case WORKLOAD_KIND_STATEFULSET:
		data, err = yaml.Marshal(w.StatefulSet)
	case WORKLOAD_KIND_DAEMONSET:
		data, err = yaml.Marshal(w.DaemonSet)
	case WORKLOAD_KIND_CRONJOB:
		data, err = yaml.Marshal(w.CronJob)
	case WORKLOAD_KIND_JOB:
		data, err = yaml.Marshal(w.Job)
	}

	if err != nil {
		panic(err)
	}

	return string(data)
}

type WORKLOAD_KIND int32

const (
	// Deployment无状态部署
	WORKLOAD_KIND_DEPLOYMENT WORKLOAD_KIND = 0
	// StatefulSet
	WORKLOAD_KIND_STATEFULSET WORKLOAD_KIND = 1
	// DaemonSet
	WORKLOAD_KIND_DAEMONSET WORKLOAD_KIND = 2
	// Job
	WORKLOAD_KIND_JOB WORKLOAD_KIND = 3
	// CronJob
	WORKLOAD_KIND_CRONJOB WORKLOAD_KIND = 4
)

// Enum value maps for WORKLOAD_KIND.
var (
	WORKLOAD_KIND_NAME = map[int32]string{
		0: "DEPLOYMENT",
		1: "STATEFULSET",
		2: "DAEMONSET",
		3: "JOB",
		4: "CRONJOB",
	}
	WORKLOAD_KIND_VALUE = map[string]int32{
		"DEPLOYMENT":  0,
		"STATEFULSET": 1,
		"DAEMONSET":   2,
		"JOB":         3,
		"CRONJOB":     4,
	}
)
