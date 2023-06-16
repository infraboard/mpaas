package terminal

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/grpc/mock"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/client-go/tools/remotecommand"
)

func NewWebSocketTerminal(conn *websocket.Conn) *WebSocketTerminal {
	return &WebSocketTerminal{
		ServerStreamBase: mock.NewServerStreamBase(),
		ws:               conn,
		timeout:          3 * time.Second,
		sizeChan:         make(chan remotecommand.TerminalSize),
		doneChan:         make(chan struct{}),
	}
}

type WebSocketTerminal struct {
	*mock.ServerStreamBase
	ws       *websocket.Conn
	timeout  time.Duration
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

func (i *WebSocketTerminal) ReadReq(req any) error {
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

func (i *WebSocketTerminal) WriteTo(r io.Reader) (err error) {
	_, err = io.Copy(i, r)
	return
}

func (i *WebSocketTerminal) Write(p []byte) (n int, err error) {
	err = i.ws.WriteMessage(websocket.BinaryMessage, p)
	n = len(p)
	return
}

func (i *WebSocketTerminal) Failed(err error) error {
	return i.close(websocket.CloseAbnormalClosure, err.Error())
}

func (i *WebSocketTerminal) Success(msg string) error {
	return i.close(websocket.CloseNormalClosure, msg)
}

func (i *WebSocketTerminal) close(code int, msg string) error {
	zap.L().Named("tasklog.term").Debugf("close code: %d, msg: %s", code, msg)
	close(i.doneChan)
	return i.ws.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, msg),
		time.Now().Add(i.timeout),
	)
}

// Next returns the new terminal size after the terminal has been resized. It returns nil when
// monitoring has been stopped.
func (i *WebSocketTerminal) Next() *remotecommand.TerminalSize {
	select {
	case size := <-i.sizeChan:
		return &size
	case <-i.doneChan:
		return nil
	}
}
