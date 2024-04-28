package workload_test

import (
	"fmt"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func TestListDeployment(t *testing.T) {
	req := meta.NewListRequest()
	req.Namespace = "go8"
	v, err := impl.ListDeployment(ctx, req)
	if err != nil {
		t.Log(err)
	}
	for i := range v.Items {
		item := v.Items[i]
		t.Log(item.Namespace, item.Name)
	}
}

func TestGetDeployment(t *testing.T) {
	req := meta.NewGetRequest("coredns")
	req.Namespace = "kube-system"
	v, err := impl.GetDeployment(ctx, req)
	if err != nil {
		t.Log(err)
	}

	// 序列化
	yd, err := yaml.Marshal(v)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(yd))
}

func TestCreateDeployment(t *testing.T) {
	req := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
		},
		Spec: v1.DeploymentSpec{
			Replicas: tea.Int32(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"k8s-app": "nginx"},
			},
			Strategy: v1.DeploymentStrategy{
				Type: v1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1.RollingUpdateDeployment{
					MaxSurge:       k8s.NewIntStr(1),
					MaxUnavailable: k8s.NewIntStr(0),
				},
			},
			// Pod模板参数
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{},
					Labels: map[string]string{
						"k8s-app": "nginx",
					},
				},
				Spec: corev1.PodSpec{
					// Pod参数
					DNSPolicy:                     corev1.DNSClusterFirst,
					RestartPolicy:                 corev1.RestartPolicyAlways,
					SchedulerName:                 "default-scheduler",
					TerminationGracePeriodSeconds: tea.Int64(30),
					// Container参数
					Containers: []corev1.Container{
						{
							Name:            "nginx",
							Image:           "nginx:latest",
							ImagePullPolicy: corev1.PullAlways,
							Env: []corev1.EnvVar{
								{Name: "APP_NAME", Value: "nginx"},
								{Name: "APP_VERSION", Value: "v1"},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("500m"),
									corev1.ResourceMemory: resource.MustParse("1Gi"),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("50m"),
									corev1.ResourceMemory: resource.MustParse("50Mi"),
								},
							},
						},
					},
				},
			},
		},
	}

	yamlReq, err := yaml.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(yamlReq))

	d, err := impl.CreateDeployment(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(d)
}

func TestScaleDeployment(t *testing.T) {
	req := meta.NewScaleRequest()
	req.Scale.Namespace = "go8"
	req.Scale.Name = "nginx"
	req.Scale.Spec.Replicas = 2
	v, err := impl.ScaleDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	yd, err := yaml.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(yd))
}

func TestReDeployment(t *testing.T) {
	req := meta.NewGetRequest("nginx")
	req.Namespace = "go8"
	v, err := impl.ReDeploy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 序列化
	yd, err := yaml.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(yd))
}

func TestDeleteDeployment(t *testing.T) {
	req := meta.NewDeleteRequest("nginx")
	err := impl.DeleteDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
