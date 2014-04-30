package validator

import (
	"fmt"
)

const (
	msgInvalid  = "Invalid."
	msgTooShort = "Too short, minimum %d characters."
	msgTooLong  = "Too long, maximum %d characters."
)

func (m *matcher) Valid() bool {
	return m.valid
}

func (m *matcher) Invalid(key string) *matcher {
	m = m.Validate(key, nil)
	m.record("", msgInvalid)
	return m
}

func (m *matcher) MinLength(min int) *matcher {
	if m.prev = m.Len() < min; !m.prev {
		m.record("", fmt.Sprintf(msgTooShort, min))
	}
	return m
}

func (m *matcher) MaxLength(max int) *matcher {
	if m.prev = m.Len() > max; !m.prev {
		m.record("", fmt.Sprintf(msgTooLong, max))
	}
	return m
}
