package workload

import (
	"bytes"
	"fmt"
	"strings"
)

func ParseWorkloadKindFromString(str string) (WORKLOAD_KIND, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := WORKLOAD_KIND_VALUE[key]
	if !ok {
		return 0, fmt.Errorf("unknown WORKLOAD_KIND: %s", str)
	}

	return WORKLOAD_KIND(v), nil
}

// Equal type compare
func (t WORKLOAD_KIND) Equal(target WORKLOAD_KIND) bool {
	return t == target
}

// IsIn todo
func (t WORKLOAD_KIND) IsIn(targets ...WORKLOAD_KIND) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

func (t WORKLOAD_KIND) String() string {
	return WORKLOAD_KIND_NAME[int32(t)]
}

// MarshalJSON todo
func (t WORKLOAD_KIND) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(t.String())
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *WORKLOAD_KIND) UnmarshalJSON(b []byte) error {
	ins, err := ParseWorkloadKindFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
