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
		return fmt.Errorf("req must be json data")
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

func (i *WebSocketWriter) Failed(err error) {
	i.close(websocket.CloseAbnormalClosure, err.Error())
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
	}
}
