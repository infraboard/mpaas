package route

type ROUTE_STATUS int

const (
	// 表示禁用
	ROUTE_STATUS_DISABLED ROUTE_STATUS = iota
	// 表示启用
	ROUTE_STATUS_ENABLED
)
