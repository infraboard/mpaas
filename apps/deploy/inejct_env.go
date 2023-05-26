package deploy

import (
	"fmt"
	"strings"

	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mpaas/conf"
	v1 "k8s.io/api/core/v1"
)

func NewDdynamicInjection() *DdynamicInjection {
	return &DdynamicInjection{
		SystemEnv: true,
		EnvGroups: []*InjectionEnvGroup{},
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

func (g *InjectionEnvGroup) IsLabelMatched(target map[string]string) bool {
	// 如果没有配置匹配标签, 默认不过滤
	if len(g.MatchLabel) == 0 {
		return true
	}

	// 如果配置了Label，而目录没有Label 则不匹配
	if len(target) == 0 {
		return false
	}

	// 匹配目录是否含有指定的label, 如果有一个不相等则不匹配
	for sk, sv := range g.MatchLabel {
		if target[sk] != sv {
			return false
		}
	}

	return true
}

func (g *InjectionEnvGroup) ToContainerEnvVars() []v1.EnvVar {
	envs := []v1.EnvVar{}

	for i := range g.InjectEnvs {
		env := g.InjectEnvs[i]
		envs = append(envs, v1.EnvVar{
			Name:  env.Key,
			Value: env.Value,
		})
	}
	return envs
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
