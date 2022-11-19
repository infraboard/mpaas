package k8s_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/provider/k8s"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

var (
	client *k8s.Client
	ctx    = context.Background()
)

func TestServerVersion(t *testing.T) {
	v, err := client.ServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestServerResources(t *testing.T) {
	rs, err := client.ServerResources()
	if err != nil {
		t.Log(err)
	}
	for i := range rs {
		t.Log(rs[i].GroupVersion, rs[i].APIVersion)
		for _, r := range rs[i].APIResources {
			t.Log(r.Name)
		}
	}
}

func TestListNode(t *testing.T) {
	v, err := client.ListNode(ctx, k8s.NewListRequest())
	if err != nil {
		t.Fatal(err)
	}
	for i := range v.Items {
		t.Log(v.Items[i].Name)
	}
}

func TestListNamespace(t *testing.T) {
	v, err := client.ListNamespace(ctx, k8s.NewListRequest())
	if err != nil {
		t.Log(err)
	}
	for i := range v.Items {
		t.Log(v.Items[i].Name)
	}
}

func TestListDeployment(t *testing.T) {
	req := k8s.NewListRequest()
	req.Namespace = "default"
	v, err := client.ListDeployment(ctx, req)
	if err != nil {
		t.Log(err)
	}
	for i := range v.Items {
		item := v.Items[i]
		t.Log(item.Namespace, item.Name)
	}
}

func TestGetDeployment(t *testing.T) {
	req := k8s.NewGetRequest("nginx")
	req.Namespace = "kube-system"
	v, err := client.GetDeployment(ctx, req)
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
		// TypeMeta: metav1.TypeMeta{
		// 	APIVersion: "apps/v1",
		// 	Kind:       "Deployment",
		// },
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
					Annotations: map[string]string{
						"eks.tke.cloud.tencent.com/root-cbs-size":     "20",
						"eks.tke.cloud.tencent.com/security-group-id": "sg-05url5pe",
					},
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
									corev1.ResourceCPU:    resource.MustParse("250m"),
									corev1.ResourceMemory: resource.MustParse("256Mi"),
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

	d, err := client.CreateDeployment(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(d)
}

func TestScaleDeployment(t *testing.T) {
	req := k8s.NewUpdateScaleRequest()
	req.Scale.Namespace = k8s.DEFAULT_NAMESPACE
	req.Scale.Name = "nginx"
	req.Scale.Spec.Replicas = 2
	v, err := client.UpdateDeploymentScale(ctx, req)
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
	req := k8s.NewGetRequest("nginx")
	v, err := client.ReDeploy(ctx, req)
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

func TestGetPod(t *testing.T) {
	req := k8s.NewGetRequest("nginx-84b44d9c9d-q8ftf")

	pods, err := client.GetPod(ctx, req)
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

func TestDeleteDeployment(t *testing.T) {
	req := k8s.NewDeleteRequest("nginx")
	err := client.DeleteDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	zap.DevelopmentSetup()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	kc, err := os.ReadFile(filepath.Join(wd, "kube_config.yml"))
	if err != nil {
		panic(err)
	}

	client, err = k8s.NewClient(string(kc))
	if err != nil {
		panic(err)
	}

}
