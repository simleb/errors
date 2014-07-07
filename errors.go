package errors

import (
	"fmt"
)

type Error []error

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

func Stack(e error, s string) error {
	if e, ok := e.(Error); ok {
		return append(e, fmt.Errorf(s))
	}
	return Error{e, fmt.Errorf(s)}
}

func Stackf(e error, s string, a ...interface{}) error {
	return Stack(e, fmt.Sprintf(s, a...))
}
