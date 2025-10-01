package utils

import (
	"errors"
	"fmt"
)

// Assertf panics with the given message if the condition is false.
func Assertf(b bool, format string, a ...any) {
	if b {
		return
	}
	panic(WithStack(fmt.Errorf(format, a...)))
}

// Assert panics with the given message if the condition is false.
func Assert(b bool, a ...any) {
	if b {
		return
	}
	panic(WithStack(errors.New(TernaryF(len(a) == 0,
		func() string { return "assertion failed" },
		func() string { return fmt.Sprint(a...) },
	))))
}

// A is a shorthand for Assert.
var A = Assert
