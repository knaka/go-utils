package utils

import "errors"

// valueResult is not exported because the member values are not expected to be used by the user directly.
type valueResult[T any] struct {
	value T
	err   error
}

// R returns a value + error result context.
//
//revive:disable-next-line:unexported-return
//goland:noinspection GoExportedFuncWithUnexportedType
func R[T any](value T, err error) *valueResult[T] {
	return &valueResult[T]{
		value: value,
		err:   err,
	}
}

// PR returns a pointer + error result context.
//
//revive:disable-next-line:unexported-return
//goland:noinspection GoExportedFuncWithUnexportedType
func PR[T any](value T, err error) *valueResult[*T] {
	return &valueResult[*T]{
		value: &value,
		err:   err,
	}
}

// NilIf returns nil if the error is one of the given errors and returns the value otherwise. This is useful against functions which use some errors for success results.
func (e *valueResult[T]) NilIf(errs ...error) T {
	if e.err == nil {
		return e.value
	}
	for _, err := range errs {
		if errors.Is(e.err, err) {
			return Nil[T]()
		}
	}
	panic(WithStack(e.err))
}

// NilIfF returns nil if any of the given functions return true for the error, otherwise panics.
func (e *valueResult[T]) NilIfF(fn ...func(error) bool) T {
	if e.err == nil {
		return e.value
	}
	for _, f := range fn {
		if f(e.err) {
			return Nil[T]()
		}
	}
	panic(WithStack(e.err))
}

// TrueIf returns true if the error is one of the given errors, otherwise panics.
func (e *valueResult[T]) TrueIf(errs ...error) bool {
	if e.err == nil {
		return false
	}
	for _, err := range errs {
		if errors.Is(e.err, err) {
			return true
		}
	}
	panic(WithStack(e.err))
}

// FalseIf returns false if the error is one of the given errors, otherwise panics.
func (e *valueResult[T]) FalseIf(errs ...error) bool {
	return !e.TrueIf(errs...)
}
