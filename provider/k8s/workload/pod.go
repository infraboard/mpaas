package workload

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
)

func (c *Client) CreatePod(ctx context.Context, pod *v1.Pod, req *meta.CreateRequest) (*v1.Pod, error) {
	return c.corev1.Pods(pod.Namespace).Create(ctx, pod, req.Opts)
}

func (c *Client) ListPod(ctx context.Context, req *meta.ListRequest) (*v1.PodList, error) {
	return c.corev1.Pods(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetPod(ctx context.Context, req *meta.GetRequest) (*v1.Pod, error) {
	return c.corev1.Pods(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Client) DeletePod(ctx context.Context, req *meta.DeleteRequest) error {
	return c.corev1.Pods(req.Namespace).Delete(ctx, req.Name, req.Opts)
}

func InjectPodTemplateSpecAnnotations(pod *v1.PodTemplateSpec, key, value string) {
	if pod == nil {
		return
	}

	if pod.Annotations == nil {
		pod.Annotations = map[string]string{}
	}

	// 注入
	pod.Annotations[key] = value
}

func InjectPodTemplateSpecLabel(pod *v1.PodTemplateSpec, key, value string) {
	if pod == nil {
		return
	}

	if pod.Labels == nil {
		pod.Labels = map[string]string{}
	}

	// 注入
	pod.Labels[key] = value
}

func InjectPodEnvVars(pod *v1.PodSpec, envs []v1.EnvVar) {
	if len(envs) == 0 {
		return
	}

	// 给Init容器注入环境变量
	for i := range pod.InitContainers {
		c := pod.InitContainers[i]
		InjectContainerEnvVars(&c, envs)
		// 替换掉原来的container的值
		pod.InitContainers[i] = c
	}

	// 给用户容器注入环境变量
	for i := range pod.Containers {
		c := pod.Containers[i]
		InjectContainerEnvVars(&c, envs)
		// 替换掉原来的container的值
		pod.Containers[i] = c
	}
}

const (
	ANNOTATION_SECRET_MOUNT    = "secret.mpaas.infraboard.io/mountpath"
	ANNOTATION_CONFIGMAP_MOUNT = "configmap.mpaas.infraboard.io/mountpath"
)

// 把secret注入到Pod中 挂载成卷使用
func InjectPodSecretVolume(pod *v1.PodSpec, ss ...*v1.Secret) {
	vm := []MountVolume{}
	// 注入volume 声明
	for i := range ss {
		secret := ss[i]
		v := NewSecretVolume(secret)
		pod.Volumes = append(pod.Volumes, v)
		vm = append(vm, NewMountVolume(v, secret.Annotations[ANNOTATION_SECRET_MOUNT]))
	}

	// 挂载到Pod中
	for i, c := range pod.Containers {
		c.VolumeMounts = append(c.VolumeMounts, NewVolumeMount(true, vm)...)
		// 替换掉原来的container的值
		pod.Containers[i] = c
	}
}

// 把configmap注入到Pod中 挂载成卷使用
func InjectPodConfigMapVolume(pod *v1.PodSpec, cs ...*v1.ConfigMap) {
	vm := []MountVolume{}
	// 注入volume 声明
	for i := range cs {
		cm := cs[i]
		v := NewConfigMapVolume(cm)
		pod.Volumes = append(pod.Volumes, v)
		vm = append(vm, NewMountVolume(v, cm.Annotations[ANNOTATION_CONFIGMAP_MOUNT]))
	}

	// 挂载到Pod中
	for i, c := range pod.Containers {
		c.VolumeMounts = append(c.VolumeMounts, NewVolumeMount(false, vm)...)
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

func NewConfigMapVolume(cm *v1.ConfigMap) v1.Volume {
	return v1.Volume{
		Name: cm.Name,
		VolumeSource: v1.VolumeSource{
			ConfigMap: &v1.ConfigMapVolumeSource{
				LocalObjectReference: v1.LocalObjectReference{Name: cm.Name},
			},
		},
	}
}

func NewVolumeMount(readonly bool, vs []MountVolume) []v1.VolumeMount {
	vms := []v1.VolumeMount{}
	for _, v := range vs {
		vms = append(vms, v1.VolumeMount{
			Name:      v.Volume.Name,
			ReadOnly:  readonly,
			MountPath: v.Path,
		})
	}
	return vms
}

func GetContainerFromPodTemplate(temp v1.PodTemplateSpec, name string) *v1.Container {
	return GetContainerFromPodSpec(temp.Spec, name)
}

func GetContainerFromPodSpec(spec v1.PodSpec, name string) *v1.Container {
	for i := range spec.Containers {
		c := spec.Containers[i]
		if name == "" || c.Name == name {
			return &c
		}
	}
	return nil
}

func GetPrimaryContainerName(spec v1.PodSpec) string {
	if c := GetContainerFromPodSpec(spec, ""); c != nil {
		return c.Name
	}
	return ""
}

type POD_STATUS int

func (s POD_STATUS) String() string {
	switch s {
	case POD_STATUS_SCHEDULED:
		return "PodScheduled"
	case POD_STATUS_INITIALIZED:
		return "Initialized"
	case POD_STATUS_CONTAINER_READY:
		return "ContainersReady"
	case POD_STATUS_POD_READY:
		return "Ready"
	}

	return fmt.Sprintf("%d", s)
}

const (
	POD_STATUS_PENDDING POD_STATUS = iota
	POD_STATUS_SCHEDULED
	POD_STATUS_INITIALIZED
	POD_STATUS_CONTAINER_READY
	POD_STATUS_POD_READY
)

func GetPodStatus(p *v1.Pod) POD_STATUS {
	status := POD_STATUS_PENDDING
	for i := range p.Status.Conditions {
		cond := p.Status.Conditions[i]
		if cond.Type == v1.PodScheduled && cond.Status == v1.ConditionTrue {
			if status <= POD_STATUS_SCHEDULED {
				status = POD_STATUS_SCHEDULED
			}
		}
		if cond.Type == v1.PodInitialized && cond.Status == v1.ConditionTrue {
			if status <= POD_STATUS_INITIALIZED {
				status = POD_STATUS_INITIALIZED
			}
		}
		if cond.Type == v1.ContainersReady && cond.Status == v1.ConditionTrue {
			if status <= POD_STATUS_CONTAINER_READY {
				status = POD_STATUS_CONTAINER_READY
			}
		}
		if cond.Type == v1.PodReady && cond.Status == v1.ConditionTrue {
			if status <= POD_STATUS_POD_READY {
				status = POD_STATUS_POD_READY
			}
		}
	}
	return status
}
