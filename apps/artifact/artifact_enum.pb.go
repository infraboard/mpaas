// Code generated by github.com/infraboard/mcube/v2
// DO NOT EDIT

package artifact

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseTYPEFromString Parse TYPE from string
func ParseTYPEFromString(str string) (TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TYPE: %s", str)
	}

	return TYPE(v), nil
}

// Equal type compare
func (t TYPE) Equal(target TYPE) bool {
	return t == target
}

// IsIn todo
func (t TYPE) IsIn(targets ...TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseARCHFromString Parse ARCH from string
func ParseARCHFromString(str string) (ARCH, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ARCH_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ARCH: %s", str)
	}

	return ARCH(v), nil
}

// Equal type compare
func (t ARCH) Equal(target ARCH) bool {
	return t == target
}

// IsIn todo
func (t ARCH) IsIn(targets ...ARCH) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ARCH) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ARCH) UnmarshalJSON(b []byte) error {
	ins, err := ParseARCHFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseOSFromString Parse OS from string
func ParseOSFromString(str string) (OS, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := OS_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown OS: %s", str)
	}

	return OS(v), nil
}

// Equal type compare
func (t OS) Equal(target OS) bool {
	return t == target
}

// IsIn todo
func (t OS) IsIn(targets ...OS) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t OS) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *OS) UnmarshalJSON(b []byte) error {
	ins, err := ParseOSFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
