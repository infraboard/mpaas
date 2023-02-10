package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/test/tools"
)

func TestCreatePipeline(t *testing.T) {
	req := pipeline.NewCreatePipelineRequest()
	tools.MustReadYamlFile("test/create.yml", req)
	ins, err := impl.CreatePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdatePipeline(t *testing.T) {
	req := pipeline.NewPutPipelineRequest("cfi9s16a0brmn92t1i7g")
	tools.MustReadYamlFile("test/create.yml", req)
	ins, err := impl.UpdatePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribePipeline(t *testing.T) {
	req := pipeline.NewDescribePipelineRequest("cfi9s16a0brmn92t1i7g")
	ins, err := impl.DescribePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestQueryPipeline(t *testing.T) {
	req := pipeline.NewQueryPipelineRequest()
	set, err := impl.QueryPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestNewCreatePipelineRequestFromYAML(t *testing.T) {
	yml := tools.MustReadContentFile("test/create.yml")

	obj, err := pipeline.NewCreatePipelineRequestFromYAML(yml)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(obj)
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
	t.Log(in.ToYAML())
}
