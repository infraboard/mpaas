package k8s

import (
	"encoding/json"
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
func NewWebsocketTerminal(ws *websocket.Conn) *WebsocketTerminal {
	wt := &WebsocketTerminal{
		ws:               ws,
		log:              zap.L().Named("Terminal"),
		maxAuthCount:     6,
		maxMessageSize:   8192,
		writeWait:        3 * time.Second,
		pongWait:         60 * time.Second,
		closeGracePeriod: 10 * time.Second,
		readDeadline:     60 * time.Second,
		writeDeadline:    3 * time.Second,
		sizeChan:         make(chan remotecommand.TerminalSize),
		doneChan:         make(chan struct{}),
	}
	wt.init()
	return wt
}

// Terminal todo
type WebsocketTerminal struct {
	log      logger.Logger
	ws       *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}

	authFailed   int
	maxAuthCount int
	// Maximum message size allowed from peer.
	maxMessageSize int64
	// Time allowed to write a message to the peer.
	writeWait time.Duration
	// Time allowed to read the next pong message from the peer.
	pongWait time.Duration
	// Time to wait before force close on connection.
	closeGracePeriod time.Duration
	// websocket read deadline
	readDeadline time.Duration
	// websocket write deadline
	writeDeadline time.Duration

	sync.Mutex
}

func (t *WebsocketTerminal) init() {
	t.ws.SetReadLimit(t.maxMessageSize)
	t.ws.SetReadDeadline(time.Now().Add(t.pongWait))
	t.ws.SetPongHandler(t.PongHandler)
}

func (t *WebsocketTerminal) isMaxAuthFailed() bool {
	return t.authFailed < t.maxAuthCount
}

// Send pings to peer with this period. Must be less than pongWait.
func (t *WebsocketTerminal) PingPeriod() time.Duration {
	return (t.pongWait * 3) / 10
}

type AuthFunc func(payload string) error

// 等待用户认证
func (t *WebsocketTerminal) Auth(af AuthFunc) {
	for t.isMaxAuthFailed() {
		_, message, err := t.ws.ReadMessage()
		if err != nil {
			t.WriteMessage(NewOperationAuthMessage(fmt.Sprintf("read websocket auth message error, %s", err)))
			t.Close()
			return
		}

		// 读取Token进行认证
		if err := af(string(message)); err != nil {
			t.WriteMessage(NewOperationAuthMessage(err.Error()))
			t.authFailed++
			continue
		}

		if err := t.WriteMessage(NewOperationAuthMessage("auth ok")); err != nil {
			t.log.Errorf("write auth success to websocket error, %s", err)
		}
		return
	}
}

type WebSocketParamer interface {
	Validate() error
}

// 等待用户输入参数
func (t *WebsocketTerminal) ParseParame(param WebSocketParamer) {
	for t.isMaxAuthFailed() {
		_, message, err := t.ws.ReadMessage()
		if err != nil {
			t.WriteMessage(NewOperatinonParamMessage(fmt.Sprintf("read websocket param message error, %s", err)))
			t.Close()
			return
		}

		// 参数解析
		if err := json.Unmarshal(message, param); err != nil {
			t.WriteMessage(NewOperatinonParamMessage(err.Error()))
			t.authFailed++
			continue
		}

		// 参数校验
		if err := param.Validate(); err != nil {
			t.WriteMessage(NewOperatinonParamMessage(err.Error()))
			t.authFailed++
			continue
		}

		if err := t.WriteMessage(NewOperatinonParamMessage("param parse ok")); err != nil {
			t.log.Errorf("write param success to websocket error, %s", err)
		}
		return
	}
}

// 推出终端，关闭socket
func (t *WebsocketTerminal) Close() error {
	return t.ws.Close()
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

// WebSocket 输入
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

// WebSocket 输出
func (t *WebsocketTerminal) Write(p []byte) (int, error) {
	if err := t.WriteMessage(NewBinaryMessage(OperationStdout, p)); err != nil {
		t.log.Debugf("write message err: %v", err)
		return 0, err
	}

	return len(p), nil
}

func (t *WebsocketTerminal) WriteMessage(msg *TerminalMessage) error {
	t.Lock()
	defer t.Unlock()
	t.ws.SetWriteDeadline(time.Now().Add(t.writeWait))
	if err := t.ws.WriteMessage(websocket.BinaryMessage, msg.MarshalToBytes()); err != nil {
		return err
	}
	return nil
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

func (t *WebsocketTerminal) PongHandler(string) error {
	err := t.ws.SetReadDeadline(time.Now().Add(t.pongWait))
	if err != nil {
		t.log.Errorf("set read deadline error, %s", err)
	}
	return nil
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

func NewOperatinonParamMessage(messge string) *TerminalMessage {
	return NewTextMessage(OperatinonParam, messge)
}

func NewOperationAuthMessage(messge string) *TerminalMessage {
	return NewTextMessage(OperationAuth, messge)
}

// NewTerminalMessage todo
func NewBinaryMessage(op TerminalOperation, data []byte) *TerminalMessage {
	return &TerminalMessage{
		Operation: op,
		Data:      data,
	}
}

// NewTextMessage todo
func NewTextMessage(op TerminalOperation, data string) *TerminalMessage {
	return &TerminalMessage{
		Operation: op,
		Data:      []byte(data),
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
	b := make([]byte, 0, len(t.Data)+1)
	b = append(b, byte(t.Operation))
	return append(b, t.Data...)
}
