// Code generated by github.com/infraboard/mcube/v2
// DO NOT EDIT

package event

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseLEVELFromString Parse LEVEL from string
func ParseLEVELFromString(str string) (LEVEL, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := LEVEL_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown LEVEL: %s", str)
	}

	return LEVEL(v), nil
}

// Equal type compare
func (t LEVEL) Equal(target LEVEL) bool {
	return t == target
}

// IsIn todo
func (t LEVEL) IsIn(targets ...LEVEL) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t LEVEL) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *LEVEL) UnmarshalJSON(b []byte) error {
	ins, err := ParseLEVELFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseSTAGEFromString Parse STAGE from string
func ParseSTAGEFromString(str string) (STAGE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := STAGE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown STAGE: %s", str)
	}

	return STAGE(v), nil
}

// Equal type compare
func (t STAGE) Equal(target STAGE) bool {
	return t == target
}

// IsIn todo
func (t STAGE) IsIn(targets ...STAGE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t STAGE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *STAGE) UnmarshalJSON(b []byte) error {
	ins, err := ParseSTAGEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
