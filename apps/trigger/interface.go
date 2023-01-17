package trigger

import (
	"fmt"
	"path"
)

const (
	AppName = "trigger"
)

type Service interface {
	RPCServer
}

func NewDefaultWebHookEvent() *GitlabWebHookEvent {
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
