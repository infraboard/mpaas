package trigger

func (m *Commit) Short() string {
	if len(m.Id) > 8 {
		return m.Id[:8]
	}

	return m.Id
}
