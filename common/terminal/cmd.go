package terminal

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
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
	if !json.Valid(payload) {
		return nil, fmt.Errorf("command must be json")
	}

	req := NewRequest()
	err := json.Unmarshal(payload, req)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validate cmd request error, %s", err)
	}

	return req, nil
}

func NewRequest() *Request {
	return &Request{}
}

type Request struct {
	// 请求Id
	Id string `json:"id"`
	// 指令名称
	Command string `json:"command"`
	// 指令参数
	Params json.RawMessage `json:"params"`
}

func (r *Request) Validate() error {
	return validator.Validate(r)
}

func (r *Request) String() string {
	return pretty.ToJSON(r)
}

func NewResponse() *Response {
	return &Response{
		Request: NewRequest(),
	}
}

type Response struct {
	Request *Request `json:"request"`
	// 异常信息
	Message string `json:"message"`
	// 处理成功后的数据
	Data any `json:"data"`
}

func (resp *Response) SetMessage(msg string) *Response {
	resp.Message = msg
	return resp
}

func (resp *Response) SetData(data any) *Response {
	resp.Data = data
	return resp
}

func (resp *Response) ToJSON() string {
	return pretty.ToJSON(resp)
}

// 处理来源客户端实现的自定义Ping, 因为浏览器并没有实现客户端Ping功能
func PingHandleFunc(r *Request, w *Response) {
	w.Data = "pong"
}

func init() {
	RegistryCmdHandleFunc("ping", PingHandleFunc)
}
