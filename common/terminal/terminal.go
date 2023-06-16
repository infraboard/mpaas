package terminal

import (
	"github.com/gorilla/websocket"
	"k8s.io/client-go/tools/remotecommand"
)

func NewWebSocketTerminal(conn *websocket.Conn, width, height uint16) *WebSocketTerminal {
	return &WebSocketTerminal{
		WebSocketWriter: NewWebSocketWriter(conn),
		TerminalSize:    NewTerminalSize(width, height),
	}
}

type WebSocketTerminal struct {
	*WebSocketWriter
	*TerminalSize
}

func (t *WebSocketTerminal) Close() {
	close(t.doneChan)
}

func (t *WebSocketTerminal) Read(p []byte) (n int, err error) {
	_, m, err := t.ws.ReadMessage()
	if err != nil {
		return 0, err
	}

	n = copy(p, m)
	return n, nil
}

func NewTerminalSize(width, height uint16) *TerminalSize {
	size := &TerminalSize{
		sizeChan: make(chan remotecommand.TerminalSize, 1),
		doneChan: make(chan struct{}),
	}
	size.sizeChan <- remotecommand.TerminalSize{Width: width, Height: height}
	return size
}

type TerminalSize struct {
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

// Next returns the new terminal size after the terminal has been resized. It returns nil when
// monitoring has been stopped.
func (i *TerminalSize) Next() *remotecommand.TerminalSize {
	select {
	case size := <-i.sizeChan:
		return &size
	case <-i.doneChan:
		return nil
	}
}
