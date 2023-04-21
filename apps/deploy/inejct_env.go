package deploy

func NewInjectionEnvGroupSet() *InjectionEnvGroupSet {
	return &InjectionEnvGroupSet{
		EnvGroups: []*InjectionEnvGroup{},
	}
}

func (s *InjectionEnvGroupSet) Add(items ...*InjectionEnvGroup) {
	s.EnvGroups = append(s.EnvGroups, items...)
}

func (e *DdynamicInjection) AddEnabledGroupTo(set *InjectionEnvGroupSet) {
	for i := range e.EnvGroups {
		g := e.EnvGroups[i]
		if g.Enabled {
			set.Add(g)
		}
	}
}
