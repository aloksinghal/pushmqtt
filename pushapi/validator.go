package pushapi

import "fmt"

type Validator struct {
	err error
}

func (v *Validator) MustContainKey(input_map map[string]interface{}, value string) bool {
	if v.err != nil {
		return false
	}
	if _, ok := input_map[value]; !ok {
		v.err = fmt.Errorf("Must Contain Key %s", value)
		return true
	}
	return false
}

func (v *Validator) MustNotBeEmptyString(value interface{}) bool {
	if v.err != nil {
		return false
	}
	if value.(string) == "" {
		v.err = fmt.Errorf("Must Not be Empty")
		return true
	}
	return false
}

func (v *Validator) IsValid() bool {
	if v.err != nil {
		return false
	}
	return true
}

func (v *Validator) Error() string {
	return v.err.Error()
}
