package terminal

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"k8s.io/client-go/tools/remotecommand"
)

func NewWebSocketTerminal(conn *websocket.Conn) *WebSocketTerminal {
	return &WebSocketTerminal{
		WebSocketWriter: NewWebSocketWriter(conn),
		TerminalResizer: NewTerminalSize(),
	}
}

type WebSocketTerminal struct {
	*WebSocketWriter
	*TerminalResizer
}

func (t *WebSocketTerminal) Close() {
	close(t.doneChan)
}

func (t *WebSocketTerminal) Read(p []byte) (n int, err error) {
	mt, m, err := t.ws.ReadMessage()
	if err != nil {
		return 0, err
	}

	// 注意文本消息和关闭消息专门被设计为了指令通道
	switch mt {
	case websocket.TextMessage:
		t.HandleCmd(m)
	case websocket.CloseMessage:
		t.l.Debugf("receive client close: %s", m)
	default:
		n = copy(p, m)
	}

	return n, nil
}

func (t *WebSocketTerminal) HandleCmd(m []byte) {
	resp := NewResponse()
	defer t.Response(resp)

	if !json.Valid(m) {
		resp.Message = "command must be json"
		return
	}

	req, err := ParseRequest(m)
	if err != nil {
		resp.Message = err.Error()
		return
	}
	resp.Request = req

	// 单独处理Resize请求
	switch req.Command {
	case "resize":
		payload := NewTerminalSzie()
		err := json.Unmarshal(req.Params, payload)
		if err != nil {
			resp.Message = err.Error()
			return
		}
		t.SetSize(*payload)
		t.l.Debugf("resize add to queue success: %s", req)
		return
	}

	// 其他业务请求
	fn := GetCmdHandleFunc(req.Command)
	if fn == nil {
		resp.Message = "command not found"
		return
	}

	fn(req, resp)
}

func NewTerminalSize() *TerminalResizer {
	size := &TerminalResizer{
		sizeChan: make(chan remotecommand.TerminalSize, 10),
		doneChan: make(chan struct{}),
	}

	return size
}

func NewTerminalSzie() *TerminalSize {
	return &TerminalSize{}
}

type TerminalSize struct {
	// 终端的宽度
	// @gotags: json:"width"
	Width uint16 `json:"width"`
	// 终端的高度
	// @gotags: json:"heigh"
	Heigh uint16 `json:"heigh"`
}

type TerminalResizer struct {
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

func (i *TerminalResizer) SetSize(ts TerminalSize) {
	i.sizeChan <- remotecommand.TerminalSize{Width: ts.Width, Height: ts.Heigh}
}

// Next returns the new terminal size after the terminal has been resized. It returns nil when
// monitoring has been stopped.
func (i *TerminalResizer) Next() *remotecommand.TerminalSize {
	select {
	case size := <-i.sizeChan:
		return &size
	case <-i.doneChan:
		return nil
	}
}
