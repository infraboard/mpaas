package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
)

func (c *Workload) CreatePod(ctx context.Context, pod *v1.Pod, req *meta.CreateRequest) (*v1.Pod, error) {
	return c.corev1.Pods(req.Namespace).Create(ctx, pod, req.Opts)
}

func (c *Workload) ListPod(ctx context.Context, req *meta.ListRequest) (*v1.PodList, error) {
	return c.corev1.Pods(req.Namespace).List(ctx, req.Opts)
}

func (c *Workload) GetPod(ctx context.Context, req *meta.GetRequest) (*v1.Pod, error) {
	return c.corev1.Pods(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Workload) DeletePod(ctx context.Context, req *meta.DeleteRequest) error {
	return c.corev1.Pods("").Delete(ctx, "", req.Opts)
}

func InjectPodEnvVars(pod *v1.PodSpec, envs []v1.EnvVar) {
	// 给Init容器注入环境变量
	for i, c := range pod.InitContainers {
		InjectContainerEnvVars(&c, envs)
		// 替换掉原来的container的值
		pod.InitContainers[i] = c
	}

	// 给用户容器注入环境变量
	for i, c := range pod.Containers {
		InjectContainerEnvVars(&c, envs)
		// 替换掉原来的container的值
		pod.Containers[i] = c
	}
}

const (
	SECRET_MOUNT_ANNOTATION_KEY = "inforboard.io/mpaas/mount_path"
)

// 把secret注入到Pod中 挂载成卷使用
func InjectPodSecretVolume(pod *v1.PodSpec, ss ...*v1.Secret) {
	vm := []MountVolume{}
	// 注入volume 声明
	for i := range ss {
		secret := ss[i]
		v := NewSecretVolume(secret)
		pod.Volumes = append(pod.Volumes, v)
		vm = append(vm, NewMountVolume(v, secret.Annotations[SECRET_MOUNT_ANNOTATION_KEY]))
	}

	// 挂载到Pod中
	for i, c := range pod.Containers {
		c.VolumeMounts = append(c.VolumeMounts, NewVolumeMount(vm)...)
		// 替换掉原来的container的值
		pod.Containers[i] = c
	}
}

func NewMountVolume(v v1.Volume, path string) MountVolume {
	return MountVolume{
		Path:   path,
		Volume: v,
	}
}

type MountVolume struct {
	Path   string
	Volume v1.Volume
}

func NewSecretVolume(secret *v1.Secret) v1.Volume {
	return v1.Volume{
		Name: secret.Name,
		VolumeSource: v1.VolumeSource{
			Secret: &v1.SecretVolumeSource{
				SecretName: secret.Name,
			},
		},
	}
}

func NewVolumeMount(vs []MountVolume) []v1.VolumeMount {
	vms := []v1.VolumeMount{}
	for _, v := range vs {
		vms = append(vms, v1.VolumeMount{
			Name:      v.Volume.Name,
			ReadOnly:  true,
			MountPath: v.Path,
		})
	}
	return vms
}
