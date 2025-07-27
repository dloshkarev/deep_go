package hw8

import (
	"errors"
	"strconv"
	"strings"
)

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	sb := new(strings.Builder)
	if len(e.errors) > 0 {
		sb.WriteString(strconv.Itoa(len(e.errors)))
		sb.WriteString(" errors occured:\n")
	}
	for _, err := range e.errors {
		sb.WriteString("\t* ")
		sb.WriteString(err.Error())
	}
	if len(e.errors) > 0 {
		sb.WriteString("\n")
	}
	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	var out *MultiError

	if !errors.As(err, &out) {
		out = &MultiError{}
	}

	var errLocal *MultiError
	for _, errParam := range errs {
		if errors.As(errParam, &errLocal) {
			for _, ex := range errLocal.errors {
				out.errors = append(out.errors, ex)
			}
		} else {
			out.errors = append(out.errors, errParam)
		}
	}

	return out
}
