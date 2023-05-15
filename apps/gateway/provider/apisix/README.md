# 基于APISIX的网关管理


## 环境搭建

```sh
git clone https://github.com/apache/apisix-docker.git
cd apisix-docker/example
# x86
docker-compose -p docker-apisix up -d
# arm/m1
docker-compose -p docker-apisix -f docker-compose-arm64.yml up -d
```


## 参考

+ [API Six 官方文档](https://apisix.apache.org/zh/docs/apisix/getting-started/)
+ [API Six Go语言插件](https://apisix.apache.org/zh/docs/go-plugin-runner/getting-started/)
+ [apache apisix](https://www.jianshu.com/p/16f0b621815c)