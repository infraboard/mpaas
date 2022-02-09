package k8s

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/client-go/tools/remotecommand"
)

// NewWebsocketTerminal todo
func NewWebsocketTerminal(ws *websocket.Conn) ContainerExecutor {
	return &WebsocketTerminal{
		ws:            ws,
		log:           zap.L().Named("Terminal"),
		readDeadline:  60 * time.Second,
		writeDeadline: 3 * time.Second,
		sizeChan:      make(chan remotecommand.TerminalSize),
		doneChan:      make(chan struct{}),
	}
}

// Terminal todo
type WebsocketTerminal struct {
	log      logger.Logger
	ws       *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}

	readDeadline  time.Duration
	writeDeadline time.Duration

	sync.Mutex
}

// Next called in a loop from remotecommand as long as the process is running
func (t *WebsocketTerminal) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// Done done, must call Done() before connection close, or Next() would not exits.
func (t *WebsocketTerminal) Done() {
	close(t.doneChan)
}

func (t *WebsocketTerminal) Read(p []byte) (int, error) {
	_, message, err := t.ws.ReadMessage()
	if err != nil {
		t.log.Errorf("read message err: %s", err)
		return copy(p, TerminalEnd), err
	}
	msg, err := ParseTerminalMessage(message)
	if err != nil {
		return copy(p, []byte(err.Error())), nil
	}
	switch msg.Operation {
	case OperationStdin:
		return copy(p, msg.Data), nil
	case OperationResize:
		t.log.Debugf("resize terminal request: %s", msg)
		width, height := msg.GetTermianlSize()
		t.sizeChan <- remotecommand.TerminalSize{Width: width, Height: height}
		t.log.Debugf("send resize to channel")
		return 0, nil
	default:
		t.log.Errorf("unknown message type '%s'", msg.Operation)
		return copy(p, TerminalEnd), fmt.Errorf("unknown message type '%d'", msg.Operation)
	}
}

func (t *WebsocketTerminal) Write(p []byte) (int, error) {
	t.Lock()
	defer t.Unlock()
	msg := NewTerminalMessage(OperationStdout, p)
	t.ws.SetWriteDeadline(time.Now().Add(t.writeDeadline))
	if err := t.ws.WriteMessage(websocket.BinaryMessage, msg.MarshalToBytes()); err != nil {
		t.log.Debugf("write message err: %v", err)
		return 0, err
	}
	return len(p), nil
}

// Ping 用户检测客户端状态, 如果客户端不在线则关闭连接
func (t *WebsocketTerminal) Ping(pingPeriod, writeWait time.Duration) {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		t.log.Debug("pingger exit close websocket")
		pingTicker.Stop()
		t.ws.Close()
	}()
	t.log.Info("start websocket ping")
	for {
		<-pingTicker.C
		t.Lock()
		t.ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := t.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			t.log.Errorf("write ping message error, %s", err)
			t.Unlock()
			return
		}
		t.Unlock()
	}
}

var (
	// TerminalEnd 中断关闭
	TerminalEnd = []byte("\u0004")
)

// TerminalOperation 终端操作类型
type TerminalOperation byte

const (
	// OperationAuth 校验
	OperationAuth TerminalOperation = iota + 1
	// OperatinonParam 操作参数
	OperatinonParam
	// OperationStdin 输入
	OperationStdin
	// OperationStdout todo
	OperationStdout
	// OperationResize resize
	OperationResize
	// OperationUnknown
	OperationUnknown
)

func parseTerminalOperation(op byte) TerminalOperation {
	switch op {
	case 1:
		return OperationAuth
	case 2:
		return OperatinonParam
	case 3:
		return OperationStdin
	case 4:
		return OperationStdout
	case 5:
		return OperationResize
	default:
		return OperationUnknown
	}
}

var (
	ErrMessageFormat = fmt.Errorf("message format error, must <op>,<message>")
)

// ParseTerminalMessage todo
func ParseTerminalMessage(data []byte) (*TerminalMessage, error) {
	if len(data) < 2 {
		return nil, ErrMessageFormat
	}

	op := parseTerminalOperation(data[0])
	if op == OperationUnknown {
		return nil, ErrMessageFormat
	}
	return &TerminalMessage{
		Operation: op,
		Data:      data[1:],
	}, nil
}

// NewTerminalMessage todo
func NewTerminalMessage(op TerminalOperation, data []byte) *TerminalMessage {
	return &TerminalMessage{
		Operation: op,
		Data:      data,
	}
}

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
type TerminalMessage struct {
	Operation TerminalOperation `json:"op"`
	Data      []byte            `json:"data"`
}

// GetTermianlSize todo
func (t *TerminalMessage) GetTermianlSize() (uint16, uint16) {
	var (
		cols uint64
		rows uint64
	)
	wh := strings.Split(string(t.Data), ",")
	if len(wh) == 2 {
		cols, _ = strconv.ParseUint(wh[0], 10, 16)
		rows, _ = strconv.ParseUint(wh[1], 10, 16)
	}
	return uint16(cols), uint16(rows)
}

// MarshalToBytes todo
func (t *TerminalMessage) MarshalToBytes() []byte {
	bin := make([]byte, 0, len(t.Data)+1)
	bin = append(bin, byte(t.Operation))
	return append(bin, t.Data...)
}
