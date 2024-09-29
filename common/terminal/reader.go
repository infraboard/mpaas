package terminal

import (
	"encoding/json"
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
	// 1K
	DefaultReadBuf = 1 * 1024
)

func NewWebSocketTerminal(conn *websocket.Conn) *WebSocketTerminal {
	return &WebSocketTerminal{
		ServerStreamBase: mock.NewServerStreamBase(),
		ws:               conn,
		timeout:          3 * time.Second,
		l:                log.Sub("tasklog.term"),
		readBuf:          make([]byte, DefaultReadBuf),
		writeBuf:         make([]byte, DefaultWriteBuf),
		TerminalResizer:  NewTerminalSize(),
	}
}

type WebSocketTerminal struct {
	ws       *websocket.Conn
	timeout  time.Duration
	l        *zerolog.Logger
	readBuf  []byte
	writeBuf []byte
	auditor  io.ReadWriter

	*TerminalResizer
	*mock.ServerStreamBase
}

func (t *WebSocketTerminal) Close() error {
	close(t.doneChan)
	return nil
}

// 读取二进制数据, 而读取过程中指令会自动处理
func (t *WebSocketTerminal) ReadBinData(success func([]byte), failed func(error)) {
	for {
		n, err := t.Read(t.readBuf)
		if err != nil {
			failed(err)
			return
		}

		// 用户输入的指令会被自动处理, 因此指令模式读取不到数据
		if n == 0 {
			continue
		}

		success(t.readBuf[:n])
		t.readBuf = make([]byte, DefaultReadBuf)
	}
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
		t.l.Debug().Msgf("receive client close: %s", m)
	default:
		n = copy(p, m)
		t.audit(p)
	}

	return n, nil
}

func (t *WebSocketTerminal) HandleCmd(m []byte) {
	resp := NewResponse()
	defer t.Response(resp)

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
		t.l.Debug().Msgf("resize add to queue success: %s", req)
		return
	}

	// 处理自定义指令
	fn := GetCmdHandleFunc(req.Command)
	if fn == nil {
		resp.Message = "command not found"
		return
	}

	fn(req, resp)
}
