package apisix

type Meta struct {
	// 唯一Id
	Id string `json:"id"`
	// epoch 时间戳，单位为秒。如果不指定则自动创建
	CreateTime int64 `json:"create_time"`
	// epoch 时间戳，单位为秒。如果不指定则自动创建
	UpdateTime int64 `json:"update_time"`
}
