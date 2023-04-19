package job_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
)

func TestParseDescribeName(t *testing.T) {
	t.Log(job.ParseDescribeName("#xxxx"))
	t.Log(job.ParseDescribeName("job_name@namespace.domain"))
}

func TestParseUniqName(t *testing.T) {
	t.Log(job.ParseUniqName("job_name@namespace.domain:v1"))
}
