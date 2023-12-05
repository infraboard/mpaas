package terminal

import (
	"os"

	"github.com/infraboard/mcube/v2/grpc/mock"
	"k8s.io/client-go/tools/remotecommand"
)

func NewStdTerminal(width, height uint16) *StdTerminal {
	return &StdTerminal{
		TerminalResizer: NewTerminalSize(),
	}
}

type StdTerminal struct {
	*mock.ServerStreamBase
	*TerminalResizer
}

func (i *StdTerminal) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (i *StdTerminal) Read(p []byte) (n int, err error) {
	return os.Stdin.Read(p)
}

// Next returns the new terminal size after the terminal has been resized. It returns nil when
// monitoring has been stopped.
func (i *StdTerminal) Next() *remotecommand.TerminalSize {
	select {
	case size := <-i.sizeChan:
		return &size
	case <-i.doneChan:
		return nil
	}
}
