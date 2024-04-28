package terminal

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/v2/grpc/mock"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
)

var (
	// 4K
	DefaultWriteBuf = 4 * 1024
)

func NewWebSocketWriter(conn *websocket.Conn) *WebSocketWriter {
	return &WebSocketWriter{
		ServerStreamBase: mock.NewServerStreamBase(),
		ws:               conn,
		timeout:          3 * time.Second,
		l:                log.Sub("tasklog.term"),
		writeBuf:         make([]byte, DefaultWriteBuf),
	}
}

type WebSocketWriter struct {
	*mock.ServerStreamBase
	ws      *websocket.Conn
	timeout time.Duration
	l       *zerolog.Logger

	// 写入websocket时 buffer大小
	writeBuf []byte
	// terminal镜像数据
	auditor io.ReadWriter
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
	_, err = io.CopyBuffer(i, r, i.writeBuf)
	if err != nil {
		return err
	}
	defer i.ResetWriteBuf()

	_, err = i.Write(i.writeBuf)
	return
}

func (i *WebSocketWriter) Write(p []byte) (n int, err error) {
	err = i.ws.WriteMessage(websocket.BinaryMessage, p)
	n = len(p)

	i.audit(p)
	return
}

// 命令的返回
func (i *WebSocketWriter) Response(resp *Response) {
	if resp.Message != "" {
		i.l.Debug().Msgf("response error, %s", resp.Message)
	}

	err := i.ws.WriteJSON(resp)
	if err != nil {
		i.l.Info().Msgf("write message error, %s", err)
	}
}

func (i *WebSocketWriter) WriteTextln(format string, a ...any) {
	i.WriteTextf(format, a...)
	i.WriteText("\r\n")
}

func (i *WebSocketWriter) WriteText(msg string) {
	err := i.ws.WriteMessage(websocket.BinaryMessage, []byte(msg))
	if err != nil {
		i.l.Info().Msgf("write message error, %s", err)
	}
}

func (i *WebSocketWriter) WriteTextf(format string, a ...any) {
	i.WriteText(fmt.Sprintf(format, a...))
}

func (i *WebSocketWriter) Failed(err error) {
	i.WriteTextf("[失败] %s", err.Error())
	i.close(websocket.CloseGoingAway, err.Error())
}

func (i *WebSocketWriter) Success(msg string) {
	i.close(websocket.CloseNormalClosure, msg)
}

func (i *WebSocketWriter) ResetWriteBuf() {
	i.writeBuf = make([]byte, DefaultWriteBuf)
}

func (i *WebSocketWriter) audit(p []byte) {
	if i.auditor == nil {
		return
	}

	_, err := i.auditor.Write(p)
	if err != nil {
		i.l.Error().Msgf("auditor write error, %s", err)
	}
}

func (i *WebSocketWriter) SetAuditor(rw io.ReadWriter) {
	i.auditor = rw
}

func (i *WebSocketWriter) close(code int, msg string) {
	i.l.Debug().Msgf("close code: %d, msg: %s", code, msg)
	err := i.ws.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, msg),
		time.Now().Add(i.timeout),
	)
	if err != nil {
		i.l.Error().Msgf("close error, %s", err)
		i.WriteText("\n" + msg)
	}
}
