package upstream

type UPSTREAM_TYPE string

const (
	// 带权重的 Round Robin。
	UPSTREAM_TYPE_ROUNDROBIN UPSTREAM_TYPE = "roundrobin"
	// 一致性哈希。
	UPSTREAM_TYPE_CHASH UPSTREAM_TYPE = "chash"
	// 选择延迟最小的节点，请参考 EWMA_chart。
	UPSTREAM_TYPE_EWMA UPSTREAM_TYPE = "ewma"
	// 选择 (active_conn + 1) / weight
	UPSTREAM_TYPE_LEAST_CONN UPSTREAM_TYPE = "least_conn"
)
