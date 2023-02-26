# 任务管理

基于K8s Job的流水线设计

+ 镜像构建: 
   + [如何在Docker 中使用 Docker](https://www.6hu.cc/archives/78414.html)
   + [Build Images In Kubernetes](https://github.com/GoogleContainerTools/kaniko)
   + [kubernetes【工具】kaniko【1】【介绍】-无特权构建镜像](https://blog.csdn.net/xixihahalelehehe/article/details/121659254)
+ 镜像部署: [kubectl 工具镜像](https://hub.docker.com/r/bitnami/kubectl)
+ 虚拟机部署: [在 Docker 容器中配置 Ansible](https://learn.microsoft.com/zh-cn/azure/developer/ansible/configure-in-docker-container?tabs=azure-cli)

## 镜像构建

原生地址:
[Google Kaniko 仓库地址](https://console.cloud.google.com/gcr/images/kaniko-project/GLOBAL/executor)

转化地址:
[如何拉取gcr.io的镜像](https://github.com/anjia0532/gcr.io_mirror/search?p=1&q=kaniko&type=issues)

提前同步好的镜像:
+ [gcr.io/kaniko-project/executor:v1.9.1](https://github.com/anjia0532/gcr.io_mirror/issues/1824)
+ [gcr.io/kaniko-project/executor:v1.9.1-debug](https://github.com/anjia0532/gcr.io_mirror/issues/1906)

拉取最新版本的镜像:
```
docker pull anjia0532/kaniko-project.executor:v1.9.1
docker pull anjia0532/kaniko-project.executor:v1.9.1-debug
```

### 手动操作

启动一个deubg环境, 可以看看里面的工具(二进制可执行文件,工具的用法)
```sh
docker run -it --entrypoint=/busybox/sh docker.io/anjia0532/kaniko-project.executor:v1.9.1-debug

/ # ls -l /kaniko/
total 75448
-rwxr-xr-x    1 0        0         10900549 Sep  8 18:23 docker-credential-acr-env
-rwxr-xr-x    1 0        0          8981984 Sep  8 18:22 docker-credential-ecr-login
-rwxr-xr-x    1 0        0          7814415 Sep  8 18:21 docker-credential-gcr
-rwxr-xr-x    1 0        0         35250176 Sep 26 19:27 executor
drwxr-xr-x    3 0        0             4096 Sep 26 19:27 ssl
-rwxr-xr-x    1 0        0         14303232 Sep 26 19:27 warmer
/ # /kaniko/executor -h
Usage:
  executor [flags]
  executor [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of kaniko

Flags:
      --build-arg multi-arg type                  This flag allows you to pass in ARG values at build time. Set it repeatedly for multiple values.
      --cache                                     Use cache when building image
      --cache-copy-layers                         Caches copy layers
      --cache-dir string                          Specify a local directory to use as a cache. (default "/cache")
      --cache-repo string                         Specify a repository to use as a cache, otherwise one will be inferred from the destination provided
      --cache-run-layers                          Caches run layers (default true)
      --cache-ttl duration                        Cache timeout in hours. Defaults to two weeks. (default 336h0m0s)
      --cleanup                                   Clean the filesystem at the end
      --compressed-caching                        Compress the cached layers. Decreases build time, but increases memory usage. (default true)
  -c, --context string                            Path to the dockerfile build context. (default "/workspace/")
      --context-sub-path string                   Sub path within the given context.
      --custom-platform string                    Specify the build platform if different from the current host
      --customPlatform string                     This flag is deprecated. Please use '--custom-platform'.
  -d, --destination multi-arg type                Registry the final image should be pushed to. Set it repeatedly for multiple destinations.
      --digest-file string                        Specify a file to save the digest of the built image to.
  -f, --dockerfile string                         Path to the dockerfile to be built. (default "Dockerfile")
      --force                                     Force building outside of a container
      --force-build-metadata                      Force add metadata layers to build image
      --git gitoptions                            Branch to clone if build context is a git repository (default branch=,single-branch=false,recurse-submodules=false)
  -h, --help                                      help for executor
      --ignore-path multi-arg type                Ignore these paths when taking a snapshot. Set it repeatedly for multiple paths.
      --ignore-var-run                            Ignore /var/run directory when taking image snapshot. Set it to false to preserve /var/run/ in destination image. (default true)
      --image-fs-extract-retry int                Number of retries for image FS extraction
      --image-name-tag-with-digest-file string    Specify a file to save the image name w/ image tag w/ digest of the built image to.
      --image-name-with-digest-file string        Specify a file to save the image name w/ digest of the built image to.
      --insecure                                  Push to insecure registry using plain HTTP
      --insecure-pull                             Pull from insecure registry using plain HTTP
      --insecure-registry multi-arg type          Insecure registry using plain HTTP to push and pull. Set it repeatedly for multiple registries.
      --kaniko-dir string                         Path to the kaniko directory, this takes precedence over the KANIKO_DIR environment variable. (default "/kaniko")
      --label multi-arg type                      Set metadata for an image. Set it repeatedly for multiple labels.
      --log-format string                         Log format (text, color, json) (default "color")
      --log-timestamp                             Timestamp in log output
      --no-push                                   Do not push the image to the registry
      --no-push-cache                             Do not push the cache layers to the registry
      --oci-layout-path string                    Path to save the OCI image layout of the built image.
      --push-retry int                            Number of retries for the push operation
      --registry-certificate key-value-arg type   Use the provided certificate for TLS communication with the given registry. Expected format is 'my.registry.url=/path/to/the/server/certificate'.
      --registry-mirror multi-arg type            Registry mirror to use as pull-through cache instead of docker.io. Set it repeatedly for multiple mirrors.
      --reproducible                              Strip timestamps out of the image to make it reproducible
      --single-snapshot                           Take a single snapshot at the end of the build.
      --skip-tls-verify                           Push to insecure registry ignoring TLS verify
      --skip-tls-verify-pull                      Pull from insecure registry ignoring TLS verify
      --skip-tls-verify-registry multi-arg type   Insecure registry ignoring TLS verify to push and pull. Set it repeatedly for multiple registries.
      --skip-unused-stages                        Build only used stages if defined to true. Otherwise it builds by default all stages, even the unnecessaries ones until it reaches the target stage / end of Dockerfile
      --snapshot-mode string                      Change the file attributes inspected during snapshotting (default "full")
      --snapshotMode string                       This flag is deprecated. Please use '--snapshot-mode'.
      --tar-path string                           Path to save the image in as a tarball instead of pushing
      --tarPath string                            This flag is deprecated. Please use '--tar-path'.
      --target string                             Set the target build stage to build
      --use-new-run                               Use the experimental run implementation for detecting changes without requiring file system snapshots.
  -v, --verbosity string                          Log level (trace, debug, info, warn, error, fatal, panic) (default "info")

Use "executor [command] --help" for more information about a command.
```

手动挂载并执行构建:
```sh
# 挂在项目到workspace目录下, 注意指定工作目录:/workspace
docker run -it -v ${PWD}/mpaas:/workspace -w /workspace --entrypoint=/busybox/sh docker.io/anjia0532/kaniko-project.executor:v1.9.1-debug
# 执行构建
/kaniko/executor --no-push -v trace
```

### 基于k8s操作

[使用git工具镜像下载依赖](https://hub.docker.com/r/bitnami/git)
```sh
docker pull bitnami/git
```

测试下能否正常使用
```sh
# 挂载secret
docker run -it -v ${HOME}/.ssh/:/workspace -w /workspace bitnami/git
# 测试下载, 关于更多git参数说明请参考看: https://git-scm.com/docs/git-config
git clone git@github.com:infraboard/mpaas.git  --single-branch --branch=master --config core.sshCommand="ssh -i ./id_rsa.pub"
```

创建代码拉取的secret, 可以参考: [use-case-pod-with-ssh-keys](https://kubernetes.io/zh-cn/docs/concepts/configuration/secret/#use-case-pod-with-ssh-keys)
```
kubectl create secret generic git-ssh-key --from-file=id_rsa=${HOME}/.ssh/id_rsa
```

创建镜像推送的secret, [Pushing to Docker Hub](https://github.com/GoogleContainerTools/kaniko#pushing-to-docker-hub) 推送至指定远端镜像仓库须要credential的支持，因此须要将credential以secret的方式挂载到/kaniko/.docker/这个目录下，文件名称为config.json，内容以下:
[](./impl/test/kaniko_config_example.json)

YWRtaW46SGFyYm9yMTIzNDUK是通过registry用户名与密码以下命令获取
```sh
$ echo -n admin:Harbor12345 | base64
YWRtaW46SGFyYm9yMTIzNDUK
```

手动挂载并测试能否推送:
```sh
# 挂在项目到workspace目录下, 注意指定工作目录:/workspace
docker run -it -v ${HOME}/Workspace/mpaas:/workspace -v ${HOME}/Workspace/mpaas/apps/job/impl/test/kaniko_config.json:/kaniko/.docker/config.json -w /workspace --entrypoint=/busybox/sh docker.io/anjia0532/kaniko-project.executor:v1.9.1-debug
# 执行构建
/kaniko/executor -v trace --destination=registry.cn-hangzhou.aliyuncs.com/inforboard/mpaas:v0.0.0
```

最后创建secret
```sh
$ kubectl create secret generic kaniko-secret --from-file=apps/job/impl/test/config.json
secret/kaniko-secret created

$ kubectl  get secret kaniko-secret
NAME            TYPE     DATA   AGE
kaniko-secret   Opaque   1      23s
```

共享配置Job共享Workdir: 
[](./impl/test/build.yml)



## 镜像部署

[k8s 应用部署相关文档](https://kubernetes.io/docs/concepts/workloads/controllers/)

拉取工具镜像
```
docker pull bitnami/kubectl
```

本地测试
```sh
docker run -it  -v ~/.kube/config:/.kube/config bitnami/kubectl get ns
```

k8s支持远程访问部署配置, 比如:
```
kubectl apply -f https://k8s.io/examples/controllers/nginx-deployment.yaml
kubectl apply -f http://localhost:8080/mpaas/api/v1/export/deploys/cfrcv8ea0brnte3v3jc0
```

只更新版本
```
# 更新镜像版本
kubectl set image deployment/nginx busybox=busybox nginx=nginx:1.9.1
# 补充任务标签
kubectl annotate deployments cmdb-deployment task.mpaas.inforboar.io/id="test" --overwrite
```

如果执行失败，也可以收到试试命令
```
kubectl set image deployment/cmdb-deployment cmdb-deployment=busybox:1.30
kubectl annotate deployments cmdb task.mpaas.inforboar.io/id=cfthf85s99bgu9olqr50
```

