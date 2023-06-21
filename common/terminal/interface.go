package terminal

type Logger interface {
	WriteTextln(format string, a ...any)
	WriteText(msg string)
}
