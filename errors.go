package errors

import (
	"fmt"
)

// An Error is a stack errors.
type Error []error

// Error returns a string containing the indented error stack.
func (e Error) Error() string {
	n := len(e) - 1
	s := "Error stack:"
	for i := range e {
		s += "\n"
		for j := 0; j < i; j++ {
			s += "  "
		}
		s += e[n-i].Error()
	}
	return s
}

// Stack stacks an error on top of previous errors.
// If the original error is nil, it returns nil.
func Stack(e error, s string, a ...interface{}) error {
	if e == nil {
		return nil
	}
	if e, ok := e.(Error); ok {
		return append(e, fmt.Errorf(s, a...))
	}
	return Error{e, fmt.Errorf(s, a...)}
}
