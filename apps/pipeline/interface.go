package pipeline

const (
	AppName = "pipelines"
)

type Service interface {
	RPCServer
}

func NewDescribePipelineRequest(id string) *DescribePipelineRequest {
	return &DescribePipelineRequest{
		Id: id,
	}
}
