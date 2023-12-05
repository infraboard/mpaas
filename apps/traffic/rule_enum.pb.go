// Code generated by github.com/infraboard/mcube/v2
// DO NOT EDIT

package traffic

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseTARGET_TYPEFromString Parse TARGET_TYPE from string
func ParseTARGET_TYPEFromString(str string) (TARGET_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TARGET_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TARGET_TYPE: %s", str)
	}

	return TARGET_TYPE(v), nil
}

// Equal type compare
func (t TARGET_TYPE) Equal(target TARGET_TYPE) bool {
	return t == target
}

// IsIn todo
func (t TARGET_TYPE) IsIn(targets ...TARGET_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TARGET_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TARGET_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTARGET_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
