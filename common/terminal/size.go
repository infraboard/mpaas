package terminal

import "k8s.io/client-go/tools/remotecommand"

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
