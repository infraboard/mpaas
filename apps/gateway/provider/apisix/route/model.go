package route

import (
	"encoding/json"

	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/common"
)

func NewRouteList() *RouteList {
	return &RouteList{
		List: []*Route{},
	}
}

type RouteList struct {
	Total int      `json:"total"`
	List  []*Route `json:"list"`
}

func (l *RouteList) String() string {
	return pretty.ToJSON(l)
}

func (l *RouteList) Add(item json.RawMessage) {
	r := NewRoute()
	err := json.Unmarshal(item, r)
	if err != nil {
		panic(err)
	}
	l.List = append(l.List, r)
	l.Total++
}

func NewRoute() *Route {
	return &Route{
		Meta:               common.NewMeta(),
		CreateRouteRequest: NewCreateRouteRequest(),
	}
}

type Route struct {
	*common.Meta
	*CreateRouteRequest
}

func (r *Route) String() string {
	return pretty.ToJSON(r)
}

func NewCreateRouteRequest() *CreateRouteRequest {
	return &CreateRouteRequest{
		RouteMatchRule: NewRouteMatchRule(),
		Timeout:        common.NewTimeout(),
		Plugins:        map[string]interface{}{},
	}
}

type CreateRouteRequest struct {
	// 匹配规则
	*RouteMatchRule
	// 路由的有效期。超过定义的时间，APISIX 将会自动删除路由，单位为秒
	TTL *int `json:"ttl"`
	// 插件配置, 参考: https://apisix.apache.org/zh/docs/apisix/terminology/plugin/
	Plugins map[string]interface{} `json:"plugins"`
	// Script 配置, 参考: https://apisix.apache.org/zh/docs/apisix/terminology/script/
	Script string `json:"script"`
	// 需要使用的 Upstream id, 参考: https://apisix.apache.org/zh/docs/apisix/terminology/upstream/
	UpstreamId string `json:"upstream_id"`
	// 需要绑定的 Service id 参考: https://apisix.apache.org/zh/docs/apisix/terminology/service/
	ServiceId string `json:"service_id"`
	// 路由名称
	Name string `json:"name"`
	// 路由描述信息
	Desc string `json:"desc"`
	// 路由描述信息
	Status ROUTE_STATUS `json:"status"`
	// 为 Route 设置 Upstream 连接、发送消息和接收消息的超时时间（单位为秒）
	Timeout *common.Timeout `json:"timeout"`
	// 当设置为 true 时，启用 websocket(boolean), 默认值为 false
	EnableWebsocket bool `json:"enable_websocket"`
}

func (r *CreateRouteRequest) ToJSON() string {
	return pretty.ToJSON(r)
}

func NewRouteMatchRule() *RouteMatchRule {
	return &RouteMatchRule{
		Hosts:       []string{},
		URIs:        []string{},
		RemoteAddrs: []string{},
		Methods:     []string{},
		Vars:        []*MatchExpr{},
		Labels:      map[string]string{},
	}
}

type RouteMatchRule struct {
	// 非空列表形态的 host，表示允许有多个不同 host，匹配其中任意一个即可
	Hosts []string `json:"hosts"`
	// 单独
	Host string `json:"host"`
	// URI匹配规则
	URIs []string `json:"uris"`
	// 单独
	URI string `json:"uri"`
	// 非空列表形态的 remote_addr，表示允许有多个不同 IP 地址，符合其中任意一个即可
	RemoteAddrs []string `json:"remote_addrs"`
	// 如果为空或没有该选项，则表示没有任何 method 限制。
	// 你也可以配置一个或多个的组合：GET，POST，PUT，DELETE，PATCH，HEAD，OPTIONS，CONNECT，TRACE，PURGE
	Methods []string `json:"methods"`
	// 如果不同路由包含相同的 uri，则根据属性 priority 确定哪个 route 被优先匹配，值越大优先级越高，默认值为 0
	Priority int `json:"priority"`
	// 由一个或多个[var, operator, val]元素组成的列表，
	// 类似 [[var, operator, val], [var, operator, val], ...]]。
	// 例如：["arg_name", "==", "json"] 则表示当前请求参数 name 是 json。
	// 此处 var 与 NGINX 内部自身变量命名是保持一致的，所以也可以使用 request_uri、host 等。
	// 更多细节请参考 https://github.com/api7/lua-resty-expr
	Vars []*MatchExpr `json:"vars"`
	// 用户自定义的过滤函数。可以使用它来实现特殊场景的匹配要求实现。
	// 该函数默认接受一个名为 vars 的输入参数，可以用它来获取 NGINX 变量
	FilterFunc string `json:"filter_func"`
	// 标识附加属性的键值对
	Labels map[string]string `json:"labels"`
}

type MatchExpr struct {
	Key   string
	Op    string
	Value string
}

// 插件配置 https://apisix.apache.org/zh/docs/apisix/terminology/plugin/
type Plugin struct {
}
