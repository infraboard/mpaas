package job

func NewMapWithKVPaire(kvs ...string) map[string]string {
	if len(kvs)%2 != 0 {
		panic("kvs must paire")
	}

	m := map[string]string{}
	for i := 0; i < len(kvs); i += 2 {
		m[kvs[i]] = kvs[i+1]
	}
	return m
}
