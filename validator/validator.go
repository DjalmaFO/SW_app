package validator

import (
	"fmt"
	"strings"
)

type Validator struct {
	Value    interface{}
	Esperado bool
}

type ValidatorTest struct {
	ValidatorTest []Validator
}

func ExecValidate(validator ValidatorTest, f func(interface{}) bool) bool {
	for _, v := range validator.ValidatorTest {
		if v.Esperado != f(v.Value) {
			return false
		}
	}
	return true
}

func HaveContent(value interface{}) bool {
	if len(strings.TrimSpace(fmt.Sprintf("%v", value))) == 0 {
		return false
	}

	return true
}
