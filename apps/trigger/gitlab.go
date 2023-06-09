package trigger

import "strings"

func (m *Commit) Short() string {
	if len(m.Id) > 8 {
		return m.Id[:8]
	}

	return m.Id
}

// "GitLab/15.5.0-pre"
func ParseGitLabServerVersion(ua string) string {
	if ua == "" {
		return ""
	}

	kv := strings.Split(ua, "/")
	if kv[0] != "GitLab" {
		return ua
	}

	if len(kv) > 1 {
		return kv[1]
	}

	return kv[0]
}
