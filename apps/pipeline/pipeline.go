package pipeline

import (
	"strings"

	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
)

func (j *Job) JobName() string {
	_, n := j.ParseName()
	return n
}

func (j *Job) JobVersion() string {
	v, _ := j.ParseName()
	return v
}

// 比如 build@v1
func (j *Job) ParseName() (name, version string) {
	if j.Name != "" {
		nv := strings.Split(j.Name, NAME_VERSION_SPLITER)
		if len(nv) > 1 {
			return nv[0], nv[1]
		}
		return nv[0], ""
	}

	return "", ""
}

func (j *Job) RunParams() []*job.RunParam {
	return j.With
}

func NewQueryPipelineRequest() *QueryPipelineRequest {
	return &QueryPipelineRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
