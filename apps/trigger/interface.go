package trigger

import (
	"fmt"
	"path"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mpaas/apps/build"
)

const (
	AppName = "triggers"
)

type Service interface {
	RPCServer
}

func NewGitlabWebHookEvent() *GitlabWebHookEvent {
	return &GitlabWebHookEvent{
		Commits: []*Commit{},
	}
}

func (e *GitlabWebHookEvent) ShortDesc() string {
	return fmt.Sprintf("%s %s [%s]", e.Ref, e.EventName, e.ObjectKind)
}

func (e *GitlabWebHookEvent) GetBaseRef() string {
	return path.Base(e.GetRef())
}

func (e *GitlabWebHookEvent) GetLatestCommit() *Commit {
	count := len(e.Commits)
	if count > 0 {
		return e.Commits[count-1]
	}
	return nil
}

func (e *GitlabWebHookEvent) GetLatestCommitShortId() string {
	cm := e.GetLatestCommit()
	if cm != nil {
		return cm.Short()
	}
	return ""
}

func NewGitlabEvent(event *GitlabWebHookEvent) *Event {
	return &Event{
		Provider:    EVENT_PROVIDER_GITLAB,
		GitlabEvent: event,
	}
}

func NewBuildStatus(bc *build.BuildConfig) *BuildStatus {
	return &BuildStatus{
		BuildConfig: bc,
	}
}

func NewQueryRecordRequest() *QueryRecordRequest {
	return &QueryRecordRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
