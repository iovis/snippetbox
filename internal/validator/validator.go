package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (self *Validator) Valid() bool {
	return len(self.FieldErrors) == 0
}

func (self *Validator) AddFieldError(key, message string) {
	if self.FieldErrors == nil {
		self.FieldErrors = make(map[string]string)
	}

	if _, exists := self.FieldErrors[key]; !exists {
		self.FieldErrors[key] = message
	}
}

func (self *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		self.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
