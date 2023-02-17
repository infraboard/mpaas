package trigger

import (
	"fmt"
	"path"
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

func (e *GitlabWebHookEvent) GetBranche() string {
	return path.Base(e.GetRef())
}

func NewGitlabEvent(serviceId string, event *GitlabWebHookEvent) *Event {
	return &Event{
		ServiceId:   serviceId,
		Provider:    EVENT_PROVIDER_GITLAB,
		GitlabEvent: event,
	}
}
