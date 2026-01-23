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

// Some asserts that all arguments are non-nil. If any argument is nil, it terminates the program with a fatal error.
func Some(args ...any) {
	for i, arg := range args {
		if arg == nil {
			panic(WithStack(fmt.Errorf("%d-th argument is null", i)))
		}
	}
}
