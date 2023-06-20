package terminal

import (
	"encoding/json"

	"github.com/infraboard/mcube/tools/pretty"
)

var handleFuncs = map[string]HandleFunc{}

// 注册请求处理函数
func RegistryCmdHandleFunc(command string, fn HandleFunc) {
	handleFuncs[command] = fn
}

func GetCmdHandleFunc(command string) HandleFunc {
	return handleFuncs[command]
}

// 基于Websocket的请求响应模式, 用于websocket的指令控制设计
type HandleFunc func(*Request, *Response)

func ParseRequest(payload []byte) (*Request, error) {
	req := NewRequest()
	err := json.Unmarshal(payload, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func NewRequest() *Request {
	return &Request{}
}

type Request struct {
	// 指令名称
	Command string `json:"command"`
	// 指令参数
	Params json.RawMessage `json:"params"`
}

func NewResponse() *Response {
	return &Response{}
}

type Response struct {
	// 异常信息
	Message string `json:"message"`
	// 处理成功后的数据
	Data any `json:"data"`
}

func (resp *Response) ToJSON() string {
	return pretty.ToJSON(resp)
}
