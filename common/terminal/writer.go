package terminal

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
)

func (i *WebSocketTerminal) ReadReq(req any) error {
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

func (i *WebSocketTerminal) WriteTo(r io.Reader) (err error) {
	_, err = io.CopyBuffer(i, r, i.writeBuf)
	if err != nil {
		return err
	}
	defer i.ResetWriteBuf()

	_, err = i.Write(i.writeBuf)
	return
}

func (i *WebSocketTerminal) Write(p []byte) (n int, err error) {
	err = i.ws.WriteMessage(websocket.BinaryMessage, p)
	n = len(p)

	i.audit(p)
	return
}

// 命令的返回
func (i *WebSocketTerminal) Response(resp *Response) {
	if resp.Message != "" {
		i.l.Debug().Msgf("response error, %s", resp.Message)
	}

	if err := i.ws.WriteJSON(resp); err != nil {
		i.l.Info().Msgf("write message error, %s", err)
	}
}

func (i *WebSocketTerminal) WriteTextln(format string, a ...any) {
	i.WriteTextf(format, a...)
	i.WriteText("\r\n")
}

func (i *WebSocketTerminal) WriteText(msg string) {
	err := i.ws.WriteMessage(websocket.BinaryMessage, []byte(msg))
	if err != nil {
		i.l.Info().Msgf("write message error, %s", err)
	}
}

func (i *WebSocketTerminal) WriteTextf(format string, a ...any) {
	i.WriteText(fmt.Sprintf(format, a...))
}

func (i *WebSocketTerminal) Failed(err error) {
	i.close(websocket.CloseGoingAway, err.Error())
}

func (i *WebSocketTerminal) Success(msg string) {
	i.close(websocket.CloseNormalClosure, msg)
}

func (i *WebSocketTerminal) ResetWriteBuf() {
	i.writeBuf = make([]byte, DefaultWriteBuf)
}

func (i *WebSocketTerminal) audit(p []byte) {
	if i.auditor == nil {
		return
	}

	_, err := i.auditor.Write(p)
	if err != nil {
		i.l.Error().Msgf("auditor write error, %s", err)
	}
}

func (i *WebSocketTerminal) SetAuditor(rw io.ReadWriter) {
	i.auditor = rw
}

func (i *WebSocketTerminal) close(code int, msg string) {
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
