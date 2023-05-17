package apisix

func NewMeta() *Meta {
	return &Meta{}
}

type Meta struct {
	// 唯一Id
	Id string `json:"id"`
	// epoch 时间戳，单位为秒。如果不指定则自动创建
	CreateTime int64 `json:"create_time"`
	// epoch 时间戳，单位为秒。如果不指定则自动创建
	UpdateTime int64 `json:"update_time"`
}

func NewTimeout() *Timeout {
	return &Timeout{}
}

// 设置连接、发送消息、接收消息的超时时间，每项都为 15 秒
type Timeout struct {
	Connect int `json:"connect"`
	Send    int `json:"send"`
	Read    int `json:"read"`
}
