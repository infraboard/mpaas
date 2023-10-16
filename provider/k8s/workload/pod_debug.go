package workload

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/infraboard/mpaas/common/terminal"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	watch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/remotecommand"
	watchtools "k8s.io/client-go/tools/watch"
	"k8s.io/kubectl/pkg/util/interrupt"
)

type DebugPodRequest struct {
	*meta.GetRequest
	EphemeralContainer corev1.EphemeralContainer   `json:""`
	Excutor            *terminal.WebSocketTerminal `json:"-"`
}

func (c *Client) DebugPod(ctx context.Context, req *DebugPodRequest) error {
	// 获取需要进行Debug的Pod
	pod, err := c.GetPod(ctx, req.GetRequest)
	if err != nil {
		return err
	}

	podJS, err := json.Marshal(pod)
	if err != nil {
		return fmt.Errorf("error creating JSON for pod: %v", err)
	}

	// 克隆当前Pod, 并添加临时容器
	debugPod := pod.DeepCopy()
	debugPod.Spec.EphemeralContainers = append(debugPod.Spec.EphemeralContainers, req.EphemeralContainer)

	debugJS, err := json.Marshal(debugPod)
	if err != nil {
		return fmt.Errorf("error creating JSON for debug container: %v", err)
	}

	// 创建Patch请求
	patch, err := strategicpatch.CreateTwoWayMergePatch(podJS, debugJS, pod)
	if err != nil {
		return fmt.Errorf("merge debug match error, %s", err)
	}

	// 执行Pod Patch, 运行临时容器
	_, err = c.corev1.Pods(req.Namespace).Patch(ctx, req.Name, types.StrategicMergePatchType,
		patch, metav1.PatchOptions{}, "ephemeralcontainers")
	if err != nil {
		// The apiserver will return a 404 when the EphemeralContainers feature is disabled because the `/ephemeralcontainers` subresource
		// is missing. Unlike the 404 returned by a missing pod, the status details will be empty.
		if serr, ok := err.(*errors.StatusError); ok && serr.Status().Reason == metav1.StatusReasonNotFound && serr.ErrStatus.Details.Name == "" {
			return fmt.Errorf("ephemeral containers are disabled for this cluster (error from server: %q)", err)
		}

		// The Kind used for the /ephemeralcontainers subresource changed in 1.22. When presented with an unexpected
		// Kind the api server will respond with a not-registered error. When this happens we can optimistically try
		// using the old API.
		if runtime.IsNotRegisteredError(err) {
			c.l.Info().Msgf("Falling back to legacy API because server returned error: %v", err)
			return c.debugByEphemeralContainerLegacy(ctx, pod, &req.EphemeralContainer)
		}

		return err
	}

	// 等待临时容器启动完成
	debugPod, err = c.WaitForPodCondition(ctx, &WaitForContainerRequest{
		Namespace:     req.Namespace,
		PodName:       req.Name,
		ContainerName: req.EphemeralContainer.Name,
		ExitCondition: WaitForContainerRunning(req.EphemeralContainer.Name, req.Excutor),
	})
	if err != nil {
		return err
	}

	// 判断临时容器是否正常启动
	status := GetPodContainerStatusByName(debugPod, req.EphemeralContainer.Name)
	if status == nil {
		return fmt.Errorf("error getting container status of container name %q: %+v", req.EphemeralContainer.Name, err)
	}
	if status.State.Terminated != nil {
		return fmt.Errorf("ephemeral container %s terminated, falling back to logs", req.EphemeralContainer.Name)
	}

	// 登录临时容器的终端
	execReq := c.corev1.RESTClient().Post().Namespace(req.Namespace).
		Resource("pods").Name(req.Name).SubResource("attach").
		VersionedParams(&corev1.PodAttachOptions{
			TypeMeta: metav1.TypeMeta{
				Kind:       "EphemeralContainers",
				APIVersion: "v1",
			},
			Container: req.EphemeralContainer.Name,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(c.restconf, "POST", execReq.URL())
	if err != nil {
		return fmt.Errorf("NewSPDYExecutor error, %s", err)
	}

	return executor.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:             req.Excutor,
		Stdout:            req.Excutor,
		Stderr:            req.Excutor,
		Tty:               true,
		TerminalSizeQueue: req.Excutor,
	})
}

// debugByEphemeralContainerLegacy adds debugContainer as an ephemeral container using the pre-1.22 /ephemeralcontainers API
// This may be removed when we no longer wish to support releases prior to 1.22.
func (c *Client) debugByEphemeralContainerLegacy(ctx context.Context, pod *corev1.Pod, debugContainer *corev1.EphemeralContainer) error {
	// We no longer have the v1.EphemeralContainers Kind since it was removed in 1.22, but
	// we can present a JSON 6902 patch that the api server will apply.
	patch, err := json.Marshal([]map[string]interface{}{{
		"op":    "add",
		"path":  "/ephemeralContainers/-",
		"value": debugContainer,
	}})
	if err != nil {
		return fmt.Errorf("error creating JSON 6902 patch for old /ephemeralcontainers API: %s", err)
	}

	result := c.corev1.RESTClient().Patch(types.JSONPatchType).
		Namespace(pod.Namespace).
		Resource("pods").
		Name(pod.Name).
		SubResource("ephemeralcontainers").
		Body(patch).
		Do(ctx)
	if err := result.Error(); err != nil {
		return err
	}
	return nil
}

type EventNotifier func(textMsg string)

type WaitForContainerRequest struct {
	Namespace     string `json:"namespace" validate:"required"`
	PodName       string `json:"pod_name" validate:"required"`
	ContainerName string `json:"container_name"`
	TimoutSecond  int    `json:"timeout_second"`
	ExitCondition watchtools.ConditionFunc
}

// WaitForPodCondition watches the given pod until the container is running
func (c *Client) WaitForPodCondition(ctx context.Context, req *WaitForContainerRequest) (*corev1.Pod, error) {
	ctx, cancel := watchtools.ContextWithOptionalTimeout(ctx, time.Duration(req.TimoutSecond)*time.Second)
	defer cancel()

	fieldSelector := fields.OneTermEqualSelector("metadata.name", req.PodName).String()
	lw := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			options.FieldSelector = fieldSelector
			return c.corev1.Pods(req.Namespace).List(ctx, options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			options.FieldSelector = fieldSelector
			return c.corev1.Pods(req.Namespace).Watch(ctx, options)
		},
	}

	intr := interrupt.New(nil, cancel)
	var result *corev1.Pod
	err := intr.Run(func() error {
		ev, err := watchtools.UntilWithSync(ctx, lw, &corev1.Pod{}, nil, req.ExitCondition)
		if ev != nil {
			result = ev.Object.(*corev1.Pod)
		}
		return err
	})

	return result, err
}

func GetPodContainerStatusByName(pod *corev1.Pod, containerName string) *corev1.ContainerStatus {
	allContainerStatus := [][]corev1.ContainerStatus{
		pod.Status.InitContainerStatuses,
		pod.Status.ContainerStatuses,
		pod.Status.EphemeralContainerStatuses,
	}
	for _, statusSlice := range allContainerStatus {
		for i := range statusSlice {
			if statusSlice[i].Name == containerName {
				return &statusSlice[i]
			}
		}
	}
	return nil
}

func FormatContainerState(stage corev1.ContainerState) string {
	if stage.Terminated != nil {
		return fmt.Sprintf("Terminated, %s", stage.Terminated.Reason)
	}
	if stage.Running != nil {
		return "Running"
	}
	if stage.Waiting != nil {
		return fmt.Sprintf("Waiting, %s %s", stage.Waiting.Reason, stage.Waiting.Message)
	}
	return stage.String()
}
