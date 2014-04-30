package validator

import (
	"fmt"
	"strings"

	"github.com/martini-contrib/binding"
)

type matcher struct {
	key string
	val interface{}
	// Set to false the first time there is an error.
	valid bool
	// Set to the result of the previous validation.
	prev bool
	// *binding.Error
	errors binding.Errors
}

type Validator binding.Errors

func (v Validator) Validate(key string, val interface{}) *matcher {
	return &matcher{
		key:    key,
		val:    val,
		valid:  true,
		prev:   true,
		errors: binding.Errors(v),
	}
}

func (m *matcher) Validate(key string, val interface{}) *matcher {
	return Validator(m.errors).Validate(key, val)
}

func (m *matcher) At(index int) *matcher {
	key := fmt.Sprintf("%s.%d", m.key, index)
	var val interface{} = nil
	if a, ok := m.val.([]interface{}); ok && len(a) != 0 {
		val = a[0]
	}
	return m.Validate(key, val)
}

func (m *matcher) Len() int {
	if s, ok := m.val.(string); ok {
		return len(s)
	}
	if a, ok := m.val.([]interface{}); ok {
		return len(a)
	}
	if m, ok := m.val.(map[interface{}]interface{}); ok {
		return len(m)
	}
	return 0
}

func (m *matcher) TrimSpace() *matcher {
	m.val = strings.TrimSpace(m.toString())
	return m
}

func (m *matcher) toString() string {
	if s, ok := m.val.(string); ok {
		return s
	}
	return ""
}

func (m *matcher) Message(msg string) *matcher {
	if !m.prev {
		m.errors[len(m.errors)-1].Message = msg
	}
	return m
}

func (m *matcher) Classification(cls string) *matcher {
	if !m.prev {
		m.errors[len(m.errors)-1].Classification = cls
	}
	return m
}

func (m *matcher) record(cls, msg string) {
	m.errors = append(m.errors, binding.Error{
		FieldNames:     []string{m.key},
		Classification: cls,
		Message:        msg,
	})
	m.valid = false
}
