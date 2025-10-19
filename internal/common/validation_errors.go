package common

import "strings"

type ValidationErrors []error

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}
	msgs := make([]string, len(ve))
	for i, err := range ve {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, "; ")
}
