package task

import "fmt"

const (
	// 该注解标签 用于任务执行成功后 回调更新任务状态, 具体功能由operator实现
	ANNOTATION_TASK_ID = "task.mpaas.inforboar.io/id"
)

const (
	CONFIG_MAP_RUNTIME_ENV_KEY        = "task.env"
	CONFIG_MAP_RUNTIME_ENV_MOUNT_PATH = "/workspace/runtime"
)

const (
	EXTENSION_FOR_TASK_POD_ARRAY = "pod"
)

func TaskPodArrayKey(index int8) string {
	return fmt.Sprintf("%s.%d", EXTENSION_FOR_TASK_POD_ARRAY, index)
}
