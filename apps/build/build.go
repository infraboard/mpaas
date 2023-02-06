package build

import (
	"time"

	"github.com/rs/xid"
)

// New 新建一个domain
func New(req *CreateBuildConfigRequest) (*BuildConfig, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &BuildConfig{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
		Spec:     req,
	}

	return d, nil
}

func NewBuildConfigSet() *BuildConfigSet {
	return &BuildConfigSet{
		Items: []*BuildConfig{},
	}
}

func (s *BuildConfigSet) Add(item *BuildConfig) {
	s.Items = append(s.Items, item)
}

func NewDefaultDeploy() *BuildConfig {
	return &BuildConfig{
		Spec: NewCreateDeployConfigRequest(),
	}
}

func NewCreateDeployConfigRequest() *CreateBuildConfigRequest {
	return &CreateBuildConfigRequest{}
}
