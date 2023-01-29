package deploy

import (
	context "context"
	"fmt"
	"net/http"
	"time"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
)

const (
	AppName = "deploys"
)

type Service interface {
	CreateDeployConfig(context.Context, *CreateDeployConfigRequest) (*DeployConfig, error)
	DeleteDeployConfig(context.Context, *DeleteDeployConfigRequest) (*DeployConfig, error)
	RPCServer
}

// New 新建一个部署配置
func New(req *CreateDeployConfigRequest) (*DeployConfig, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &DeployConfig{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return d, nil
}

func (req *CreateDeployConfigRequest) Validate() error {
	return validate.Validate(req)
}

func NewQueryDeployConfigRequestFromHttp(r *http.Request) *QueryDeployConfigRequest {
	req := NewQueryDeployConfigRequest()
	req.Page = request.NewPageRequestFromHTTP(r)
	return req
}

func NewQueryDeployConfigRequest() *QueryDeployConfigRequest {
	return &QueryDeployConfigRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewCreateDeployConfigRequest() *CreateDeployConfigRequest {
	return &CreateDeployConfigRequest{}
}

// Validate 校验请求是否合法
func (req *UpdateDeployConfigRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("id required")
	}
	if req.UpdateMode.Equal(pb_request.UpdateMode_PUT) {
		return validate.Validate(req)
	}

	return nil
}

func NewDescribeDeployConfigRequest(id string) *DescribeDeployConfigRequest {
	return &DescribeDeployConfigRequest{
		DescribeValue: id,
	}
}

func (req *DescribeDeployConfigRequest) Validate() error {
	return validate.Validate(req)
}

func NewPutDeployRequest(id string) *UpdateDeployConfigRequest {
	return &UpdateDeployConfigRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		Spec:       NewCreateDeployConfigRequest(),
	}
}

func NewPatchDeployRequest(id string) *UpdateDeployConfigRequest {
	return &UpdateDeployConfigRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Spec:       NewCreateDeployConfigRequest(),
	}
}

func NewDeleteDeployConfigRequest(id string) *DeleteDeployConfigRequest {
	return &DeleteDeployConfigRequest{
		Id: id,
	}
}
