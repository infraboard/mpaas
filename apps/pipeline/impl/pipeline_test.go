package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryPipeline(t *testing.T) {
	req := pipeline.NewQueryPipelineRequest()
	set, err := impl.QueryPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestToYaml(t *testing.T) {
	in := pipeline.NewCreatePipelineRequest()
	in.Name = "pipeline_example"
	in.Description = "example"
	in.AddStage(
		&pipeline.Stage{
			Name: "stage_01",
			With: []*job.RunParam{
				{Name: "param1", Value: "value1"},
			},
			Jobs: []*pipeline.RunJobRequest{
				{Job: "job01", Params: &job.VersionedRunParam{
					Version: "v0.1",
					Params: []*job.RunParam{
						{Name: "param1", Value: "value1"},
					},
				}},
			},
		},
		&pipeline.Stage{
			Name: "stage_02",
			With: []*job.RunParam{
				{Name: "param1", Value: "value1"},
			},
			Jobs: []*pipeline.RunJobRequest{
				{Job: "job01", Params: &job.VersionedRunParam{
					Version: "v0.1",
					Params: []*job.RunParam{
						{Name: "param1", Value: "value1"},
					},
				}},
			},
		},
	)
	t.Log(tools.MustToYaml(in))
}
