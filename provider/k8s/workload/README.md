# 应用负载

[Workloads 说明](https://kubernetes.io/docs/concepts/workloads/)

[k8s job样例](https://www.cnblogs.com/zouzou-busy/p/16155721.html)


## 关于临时容器

[临时容器](https://kubernetes.io/zh-cn/docs/concepts/workloads/pods/ephemeral-containers/)
[使用 client-go 实现 kubectl debug 命令](https://www.modb.pro/db/137718)
[Kubectl Debug源码](https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/debug/debug.go)


## 关于Hold Pod

要让Kubernetes Pod保持运行状态而不退出，可以采取以下几种方法：
+ 使用无限循环：在容器中运行一个无限循环的进程，例如使用while true命令。这样，Pod将一直运行，直到手动终止。
+ 使用sleep命令：在容器中运行一个sleep命令，例如sleep infinity。这将使容器进入休眠状态，但Pod仍然处于运行状态。
+ 使用sidecar容器：在Pod中添加一个额外的sidecar容器，该容器可以保持运行状态并监控主容器。如果主容器退出，sidecar容器可以重新启动它。
+ 使用init容器：在Pod中添加一个init容器，该容器可以在主容器启动之前运行，并保持运行状态。这样，即使主容器退出，Pod仍然可以保持运行状态。

需要注意的是，Kubernetes会自动重新启动失败的Pod，因此如果Pod退出，Kubernetes将尝试重新启动它。如果您希望Pod保持运行状态而不被重新启动，可以使用上述方法之一。