package terminal

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/grpc/mock"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewWebSocketWriter(conn *websocket.Conn) *WebSocketWriter {
	return &WebSocketWriter{
		ServerStreamBase: mock.NewServerStreamBase(),
		ws:               conn,
		timeout:          3 * time.Second,
		l:                zap.L().Named("tasklog.term"),
	}
}

type WebSocketWriter struct {
	*mock.ServerStreamBase
	ws      *websocket.Conn
	timeout time.Duration
	l       logger.Logger
}

func (i *WebSocketWriter) ReadReq(req any) error {
	mt, data, err := i.ws.ReadMessage()
	if err != nil {
		return err
	}
	if mt != websocket.TextMessage {
		return fmt.Errorf("req must be TextMessage, but now not, is %d", mt)
	}
	if !json.Valid(data) {
		return fmt.Errorf("req must be json data, but %s", string(data))
	}

	return json.Unmarshal(data, req)
}

func (i *WebSocketWriter) WriteTo(r io.Reader) (err error) {
	_, err = io.Copy(i, r)
	return
}

func (i *WebSocketWriter) Write(p []byte) (n int, err error) {
	err = i.ws.WriteMessage(websocket.BinaryMessage, p)
	n = len(p)
	return
}

// 命令的返回
func (i *WebSocketWriter) Response(resp *Response) {
	if resp.Message != "" {
		i.l.Debugf("response error, %s", resp.Message)
	}

	err := i.ws.WriteJSON(resp)
	if err != nil {
		i.l.Infof("write message error, %s", err)
	}
}

func (i *WebSocketWriter) WriteTextln(format string, a ...any) {
	i.WriteTextf(format, a...)
	i.WriteText("\r\n")
}

func (i *WebSocketWriter) WriteText(msg string) {
	err := i.ws.WriteMessage(websocket.BinaryMessage, []byte(msg))
	if err != nil {
		i.l.Infof("write message error, %s", err)
	}
}

func (i *WebSocketWriter) WriteTextf(format string, a ...any) {
	i.WriteText(fmt.Sprintf(format, a...))
}

func (i *WebSocketWriter) Failed(err error) {
	i.close(websocket.CloseGoingAway, err.Error())
}

func (i *WebSocketWriter) Success(msg string) {
	i.close(websocket.CloseNormalClosure, msg)
}

func (i *WebSocketWriter) close(code int, msg string) {
	i.l.Debugf("close code: %d, msg: %s", code, msg)
	err := i.ws.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, msg),
		time.Now().Add(i.timeout),
	)
	if err != nil {
		i.l.Errorf("close error, %s", err)
		i.WriteText("\n" + msg)
	}
}
