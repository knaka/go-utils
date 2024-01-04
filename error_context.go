package utils

// NewResult returns an error context to ignore specific errors.
//
// noinspection GoExportedFuncWithUnexportedType
func NewResult[T any](first T, rest ...any) *ptrResult[any] {
	var err error
	if len(rest) > 0 {
		if errNew, ok := (rest[len(rest)-1]).(error); ok {
			err = errNew
		}
	} else if errNew, ok := any(first).(error); ok {
		err = errNew
	}
	return &ptrResult[any]{
		Err: err,
	}
}

// NewValueResult returns a value + error context to ignore specific errors.
//
// noinspection GoExportedFuncWithUnexportedType
func NewValueResult[T any](value T, err error) *ptrResult[T] {
	return &ptrResult[T]{
		Ptr: &value,
		Err: err,
	}
}

// NewPtrResult returns a pointer + error context to ignore specific errors.
//
// noinspection GoExportedFuncWithUnexportedType
func NewPtrResult[T any](ptr *T, err error) *ptrResult[T] {
	return &ptrResult[T]{
		Ptr: ptr,
		Err: err,
	}
}

type ptrResult[T any] struct {
	Ptr *T
	Err error
}

func (e *ptrResult[T]) NilIf(errs ...error) *T {
	if e.Err == nil {
		return e.Ptr
	}
	for _, err := range errs {
		if e.Err == err {
			return nil
		}
	}
	panic(e.Err)
}

func (e *ptrResult[T]) TrueIf(errs ...error) bool {
	if e.Err == nil {
		return false
	}
	for _, err := range errs {
		if e.Err == err {
			return true
		}
	}
	panic(e.Err)
}

func (e *ptrResult[T]) FalseIf(errs ...error) bool {
	return !e.TrueIf(errs...)
}
