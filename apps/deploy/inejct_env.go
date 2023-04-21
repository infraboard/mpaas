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

func NewInjectionEnvGroup() *InjectionEnvGroup {
	return &InjectionEnvGroup{
		Enabled:    true,
		MatchLabel: map[string]string{},
		InjectEnvs: []*InjectionEnv{},
	}
}

func (g *InjectionEnvGroup) AddEnv(env ...*InjectionEnv) {
	g.InjectEnvs = append(g.InjectEnvs, env...)
}

func NewInjectionEnv(key, value string) *InjectionEnv {
	return &InjectionEnv{
		Key:   key,
		Value: value,
	}
}

func (e *InjectionEnv) SetEncrypt(v bool) *InjectionEnv {
	e.Encrypt = v
	return e
}
