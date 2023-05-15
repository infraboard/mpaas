package apisix

type UPSTREAM_TYPE int

const (
	// 带权重的 Round Robin。
	UPSTREAM_TYPE_ROUNDROBIN UPSTREAM_TYPE = iota
	// 一致性哈希。
	UPSTREAM_TYPE_CHASH
	// 选择延迟最小的节点，请参考 EWMA_chart。
	UPSTREAM_TYPE_EWMA
	// 选择 (active_conn + 1) / weight
	UPSTREAM_TYPE_LEAST_CONN
)

type Upstream struct {
	// 唯一Id
	Id string `json:"id"`
	// epoch 时间戳，单位为秒。如果不指定则自动创建
	CreateTime int64 `json:"create_time"`
	// epoch 时间戳，单位为秒。如果不指定则自动创建
	UpdateTime int64 `json:"update_time"`
	// 具体参数
	*CreateUpstreamRequeset
}

type CreateUpstreamRequeset struct {
	// 负载均衡算法，默认值是roundrobin
	Type UPSTREAM_TYPE `json:"type"`
	// 后端服务地址
	Nodes []*Node `json:"nodes"`
	// 采用注册中心时的配置, 与 nodes 二选一
	DiscoverNodes
	// 该选项只有类型是 chash 才有效。根据 key 来查找对应的节点 id，相同的 key 在同一个对象中，则返回相同 id。
	// 目前支持的 NGINX 内置变量有 uri, server_name, server_addr, request_uri, remote_port, remote_addr, query_string, host, hostname, arg_***，
	// 其中 arg_*** 是来自 URL 的请求参数，详细信息请参考 http://nginx.org/en/docs/varindex.html
	Key string `json:"key"`
	// 配置健康检查的参数，详细信息请参考 https://apisix.apache.org/zh/docs/apisix/tutorials/health-check/
	Checks HealthyCheck `json:"checks"`
	// 使用 NGINX 重试机制将请求传递给下一个上游，默认启用重试机制且次数为后端可用的节点数量。
	// 如果指定了具体重试次数，它将覆盖默认值。当设置为 0 时，表示不启用重试机制
	Retries int `json:"retries"`
	// 限制是否继续重试的时间，若之前的请求和重试请求花费太多时间就不再继续重试。
	// 当设置为 0 时，表示不启用重试超时机制
	RetryTimeout int `json:"retry_timeout"`
	// 设置连接、发送消息、接收消息的超时时间，以秒为单位
	Timeout Timeout `json:"timeout"`
	// hash_on 支持的类型有 vars（NGINX 内置变量），header（自定义 header），cookie，consumer，默认值为 vars
	HashOn string `json:"hash_on"`
	// 标识上游服务名称、使用场景等
	Name string `json:"name"`
	// 上游服务描述、使用场景等
	Desc string `json:"desc"`
	// 跟上游通信时使用的 scheme。
	// 对于 7 层代理，可选值为 [http, https, grpc, grpcs]。
	// 对于 4 层代理，可选值为 [tcp, udp, tls]。
	// 默认值为 http，详细信息请参考下文。
	Scheme UPSTREAM_SCHEME `json:"scheme"`
	// 标识附加属性的键值对
	// {"version":"v2","build":"16","env":"production"}
	Labels map[string]string `json:"labels"`
	// TLS 配置
	TLSConfig TLSConfig `json:"tls"`
	// 允许 Upstream 有自己单独的连接池。它下属的字段，比如 requests，可以用于配置上游连接保持的参数
	KeepalivePool KeepalivePool `json:"keepalive_pool"`
}

// 动态设置 keepalive 指令，详细信息请参考下文
type KeepalivePool struct {
	Size        int `json:"size"`
	IdleTimeout int `json:"idle_timeout"`
	Requests    int `json:"requests"`
}

type TLSConfig struct {
	// https 证书
	ClientCert string `json:"client_cert"`
	// https 证书私钥
	ClientKey string `json:"client_key"`
	// 设置引用的 SSL id，详见 https://apisix.apache.org/zh/docs/apisix/admin-api/#ssl
	// 不能和 client_cert、client_key 一起使用
	ClientCertId string `json:"client_cert_id"`
}

type UPSTREAM_SCHEME int

const (
	UPSTREAM_SCHEME_HTTP UPSTREAM_SCHEME = iota
	UPSTREAM_SCHEME_HTTPS
	UPSTREAM_SCHEME_GRPC
	UPSTREAM_SCHEME_GRPCS
	UPSTREAM_SCHEME_TCP
	UPSTREAM_SCHEME_UDP
	UPSTREAM_SCHEME_TLS
)

// 设置连接、发送消息、接收消息的超时时间，每项都为 15 秒
type Timeout struct {
	Connect int `json:"connect"`
	Send    int `json:"send"`
	Read    int `json:"read"`
}

type Node struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Weight   int    `json:"weight"`
	Priority int    `json:"priority"`
}

type DiscoverNodes struct {
	// 服务名称, 与 nodes 二选一
	ServiceName string `json:"service_name"`
	// 与 service_name 配合使用
	Discovery_type string `json:"discovery_type"`
}

// 健康检查 https://apisix.apache.org/zh/docs/apisix/tutorials/health-check/
type HealthyCheck struct {
	// 主动检查配置
	Active ActiveHealthCheck `json:"active"`
	// 被动检查配置
	Passive PassiveHealthCheck `json:"passive"`
}

// 主动检查配置
type ActiveHealthCheck struct {
	// 主动检查的类型
	Type CHECK_PROTOCOL `json:"type"`
	// 主动检查的超时时间（单位为秒）。
	Timeout int `json:"timeout"`
	// 主动检查时同时检查的目标数。
	Concurrency int `json:"concurrency"`
	// 主动检查的 HTTP 请求路径
	HttpPath string `json:"http_path"`
	// 主动检查的 HTTP 请求主机名
	// ${upstream.node.host}
	Host string `json:"host"`
	// 主动检查的 HTTP 请求主机端口
	// ${upstream.node.port}
	Port int `json:"port"`
	// 主动检查使用 HTTPS 类型检查时，是否检查远程主机的 SSL 证书
	HttpsVerifyCertificate bool `json:"https_verify_certificate"`
	// 主动检查使用 HTTP 或 HTTPS 类型检查时，设置额外的请求头信息
	ReqHeaders []string `json:"req_headers"`
	// 主动检查（健康节点）检查的间隔时间（单位为秒）
	Interval int `json:"interval"`
	// 被动检查（健康节点）HTTP 或 HTTPS 类型检查时，健康节点的 HTTP 状态码
	// [200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308]
	HttpStatuses []int `json:"http_statuses"`
	// 被动检查（健康节点）确定节点健康的次数。
	Successes int `json:"successes"`
}

type CHECK_PROTOCOL int

const (
	// http
	CHECK_PROTOCOL_HTTP CHECK_PROTOCOL = iota
	// https
	CHECK_PROTOCOL_HTTPS
	// tcp
	CHECK_PROTOCOL_TCP
)

// 被动检查配置
type PassiveHealthCheck struct {
	// 检查成功条件
	Healthy PassiveHealthyCheck `json:"healthy"`
	// 检查失败条件
	Unhealthy PassiveUnHealthyCheck `json:"unhealthy"`
}

type PassiveHealthyCheck struct {
	// 被动检查（健康节点）HTTP 或 HTTPS 类型检查时，健康节点的 HTTP 状态码
	// [200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308]
	HttpStatuses []int `json:"http_statuses"`
	// 被动检查（健康节点）确定节点健康的次数。
	Successes int `json:"successes"`
}

type PassiveUnHealthyCheck struct {
	// 被动检查（非健康节点）HTTP 或 HTTPS 类型检查时，非健康节点的 HTTP 状态码。
	// [429, 500, 503]
	HttpStatuses []int `json:"http_statuses"`
	// 被动检查（非健康节点）TCP 类型检查时，确定节点非健康的次数
	TcpFailures int `json:"tcp_failures"`
	// 被动检查（非健康节点）HTTP 或 HTTPS 类型检查时，确定节点非健康的次数
	HttpFailures int `json:"http_failures"`
	// 被动检查（非健康节点）确定节点非健康的超时次数
	Timeouts int `json:"timeouts"`
}
