package job

import (
	"fmt"
	"net/http"
	"strings"
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

func NewCreateJobRequest() *CreateJobRequest {
	return &CreateJobRequest{
		Domain:         domain.DEFAULT_DOMAIN,
		Namespace:      namespace.DEFAULT_NAMESPACE,
		RunParams:      []*VersionedRunParam{},
		RollbackParams: []*VersionedRunParam{},
		Labels:         make(map[string]string),
	}
}

var (
	INVALIDATE_NAME_CHAR = []rune{'.'}
)

func (req *CreateJobRequest) Validate() error {
	for _, c := range INVALIDATE_NAME_CHAR {
		if strings.ContainsRune(req.Name, c) {
			return fmt.Errorf("名称中不能出现特殊字符: %s", string(INVALIDATE_NAME_CHAR))
		}
	}

	if strings.HasPrefix(req.Name, "#") {
		return fmt.Errorf("名称不能以#开头")
	}

	return validate.Validate(req)
}

func (req *CreateJobRequest) AddVersionParams(item *VersionedRunParam) {
	req.RunParams = append(req.RunParams, item)
}

func NewQueryJobRequestFromHTTP(r *http.Request) *QueryJobRequest {
	return &QueryJobRequest{
		Page: request.NewPageRequestFromHTTP(r),
	}
}

func NewQueryJobRequest() *QueryJobRequest {
	return &QueryJobRequest{
		Page:  request.NewDefaultPageRequest(),
		Ids:   []string{},
		Names: []string{},
		Label: map[string]string{},
	}
}

func ParseDescribeName(name string) (DESCRIBE_BY, string) {
	if name == "" {
		return DESCRIBE_BY_JOB_ID, ""
	}

	switch name[0] {
	case '#':
		return DESCRIBE_BY_JOB_ID, name[1:]
	default:
		return DESCRIBE_BY_JOB_UNIQ_NAME, name
	}
}

func ParseUniqName(name string) (jobname, namespace, domain string) {
	if name == "" {
		return
	}

	nv := strings.Split(name, UNIQ_NAME_SPLITER)
	jobname = nv[0]

	if len(nv) > 1 {
		nd := strings.Split(nv[1], UNIQ_NAMESPACE_SPLITER)
		if len(nd) > 0 {
			namespace = nd[0]
		}
		if len(nd) > 1 {
			domain = nd[1]
		}
	}
	return
}

func NewDescribeJobRequest(name string) *DescribeJobRequest {
	by, v := ParseDescribeName(name)
	return &DescribeJobRequest{
		DescribeBy:    by,
		DescribeValue: v,
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
