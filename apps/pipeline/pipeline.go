package pipeline

import job "github.com/infraboard/mpaas/apps/job"

func (p *Pipeline) GetFirstJob() *Job {
	for i := range p.Spec.Stages {
		s := p.Spec.Stages[i]
		if len(s.Jobs) > 0 {
			return s.Jobs[0]
		}
	}
	return nil
}

func (j *Job) JobName() string {
	return j.Name
}

func (j *Job) JobVersion() string {
	return ""
}

func (j *Job) RunParams() []*job.RunParam {
	return j.With
}
