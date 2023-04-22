package deploy

import (
	"fmt"
	"strings"

	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mpaas/conf"
)

func NewDdynamicInjection() *DdynamicInjection {
	return &DdynamicInjection{
		SystemEnv:  true,
		EnvGroups:  []*InjectionEnvGroup{},
		AccessStat: NewAccessStat(),
	}
}

func NewInjectionEnvGroupSet() *InjectionEnvGroupSet {
	return &InjectionEnvGroupSet{
		EnvGroups: []*InjectionEnvGroup{},
	}
}

func (s *InjectionEnvGroupSet) Add(items ...*InjectionEnvGroup) {
	s.EnvGroups = append(s.EnvGroups, items...)
}

func (s *InjectionEnvGroupSet) Encrypt(key string) {
	for m := range s.EnvGroups {
		group := s.EnvGroups[m]
		for n := range group.InjectEnvs {
			env := group.InjectEnvs[n]
			env.MakeEncrypt(key)
		}
	}
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

func (e *InjectionEnv) MakeEncrypt(key string) {
	if strings.HasPrefix(e.Value, conf.CIPHER_TEXT_PREFIX) {
		return
	}

	if !e.Encrypt {
		return
	}

	if key == "" {
		e.EncryptFailed = "加密key为空"
		return
	}

	encrypt, err := cbc.EncryptToString(e.Value, []byte(key))
	if err != nil {
		e.EncryptFailed = err.Error()
		return
	}

	e.Value = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, encrypt)
}

func NewAccessStat() *AccessStat {
	return &AccessStat{}
}
