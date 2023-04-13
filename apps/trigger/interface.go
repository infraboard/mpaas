package trigger

import (
	"fmt"
	"strings"

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
		Project: &Project{},
		Commits: []*Commit{},
	}
}

func (e *GitlabWebHookEvent) ShortDesc() string {
	return fmt.Sprintf("%s %s [%s]", e.Ref, e.EventName, e.ObjectKind)
}

func (e *GitlabWebHookEvent) GetBranch() string {
	switch e.EventType {
	case GITLAB_EVENT_TYPE_MERGE_REQUEST:
		return e.ObjectAttributes.TargetBranch
	case GITLAB_EVENT_TYPE_PUSH:
		return strings.TrimPrefix(e.Ref, "refs/heads/")
	case GITLAB_EVENT_TYPE_TAG:
		return e.GetTag()
	default:
		return e.Ref
	}
}

func (e *GitlabWebHookEvent) GetTag() string {
	return strings.TrimPrefix(e.Ref, "refs/tags/")
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

func NewGitlabEvent() *Event {
	return &Event{
		Provider: EVENT_PROVIDER_GITLAB,
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
