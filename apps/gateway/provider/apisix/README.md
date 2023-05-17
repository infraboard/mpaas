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

验证联通性: 使用host.docker.internal或者gateway.docker.internal来替换localhost 这种方法 在macos上测试了 可以通
```sh
# 在本地启动后端服务, 并测试可以访问
curl http://localhost:8010/apidocs.json

# 进入
docker exec -it docker-apisix_apisix_1 /bin/bash
# 确保能访问到宿主机网络
curl gateway.docker.internal:8010/apidocs.json
```

## 验证环境

获取 API Key
```sh
docker exec -it docker-apisix_apisix_1 cat /usr/local/apisix/conf/config-default.yaml | grep admin_key -A 10
# 下面是我获取的本地的key
admin_key:
    -
    name: admin
    key: edd1c9f034335f136f87ad84b625c8f1
    role: admin                 # admin: manage all configuration data
                                # viewer: only can view configuration data
```

下面的流程是参考: [发布API](https://apisix.apache.org/zh/docs/apisix/tutorials/expose-api/)

创建上游
```sh
# 本地后端服务 http://localhost:8010/apidocs.json
curl "http://127.0.0.1:9180/apisix/admin/upstreams/1" \
-H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
{
  "type": "roundrobin",
  "nodes": {
    "gateway.docker.internal:8010": 1
  }
}'
```
返回结果
```json
{
	"key": "/apisix/upstreams/1",
	"value": {
		"pass_host": "pass",
		"id": "1",
		"create_time": 1684288268,
		"type": "roundrobin",
		"update_time": 1684289950,
		"scheme": "http",
		"nodes": {
			"gateway.docker.internal:8010": 1
		},
		"hash_on": "vars"
	}
}
```

创建路由
```sh
curl "http://127.0.0.1:9180/apisix/admin/routes/1" \
-H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
{
  "methods": ["GET"],
  "host": "devcloud.com",
  "uri": "/*",
  "upstream_id": "1"
}'
```

返回结果:
```json
{
	"key": "/apisix/routes/1",
	"value": {
		"id": "1",
		"create_time": 1684288510,
		"update_time": 1684290057,
		"uri": "/*",
		"status": 1,
		"methods": ["GET"],
		"priority": 0,
		"upstream_id": "1",
		"host": "devcloud.com"
	}
}
```

测试路由
```
curl -i -X GET "http://127.0.0.1:9080/apidocs.json" -H "Host: devcloud.com"
```

## 开发SDK



## 参考

+ [API Six 官方文档](https://apisix.apache.org/zh/docs/apisix/getting-started/)
+ [API Six Go语言插件](https://apisix.apache.org/zh/docs/go-plugin-runner/getting-started/)
+ [apache apisix](https://www.jianshu.com/p/16f0b621815c)
+ [为什么 APISIX Ingress 是比 Traefik 更好的选择？](https://cloud.tencent.com/developer/article/2207539)