package task

const (
	AppName = "tasks"
)

type Service interface {
	RPCServer
}

func NewRunJobRequest() *RunJobRequest {
	return &RunJobRequest{}
}
