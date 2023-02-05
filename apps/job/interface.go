package job

import (
	"time"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/common/validate"
	request "github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
)

const (
	AppName = "jobs"
)

type Service interface {
	RPCServer
}

func (req *CreateJobRequest) Validate() error {
	return validate.Validate(req)
}

func NewCreateJobRequest() *CreateJobRequest {
	return &CreateJobRequest{
		Domain:    domain.DEFAULT_DOMAIN,
		Namespace: namespace.DEFAULT_NAMESPACE,
		RunParams: []*VersionedRunParam{},
		Labels:    make(map[string]string),
	}
}

func NewQueryJobRequest() *QueryJobRequest {
	return &QueryJobRequest{
		Page:  request.NewDefaultPageRequest(),
		Ids:   []string{},
		Names: []string{},
	}
}

func NewDescribeJobRequest(id string) *DescribeJobRequest {
	return &DescribeJobRequest{
		DescribeValue: id,
	}
}

func (req *DescribeJobRequest) Validate() error {
	return validate.Validate(req)
}

func NewPutJobRequest(id string) *UpdateJobRequest {
	return &UpdateJobRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().Unix(),
		Spec:       NewCreateJobRequest(),
	}
}

func NewPatchJobRequest(id string) *UpdateJobRequest {
	return &UpdateJobRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().Unix(),
		Spec:       NewCreateJobRequest(),
	}
}
