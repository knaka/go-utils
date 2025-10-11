package utils

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo struct{}

type Foo interface {
	Bar()
}

type Baz struct{}

var _ Foo = &Baz{}

func (b *Baz) Bar() {}

func TestNewResult(t *testing.T) {
	foo1 := PtrResult((func() (*foo, error) {
		return &foo{}, io.EOF
	})()).NilIf(io.EOF)
	assert.Nil(t, foo1)

	foo2 := Result((func() (*foo, error) { return &foo{}, nil })()).NilIf(io.EOF)
	assert.NotNil(t, foo2)

	foo3 := Result((func() (Foo, error) { return &Baz{}, nil })()).NilIf(io.EOF)
	assert.NotNil(t, foo3)

	foo4 := Result((func() (Foo, error) { return &Baz{}, io.EOF })()).NilIf(io.EOF)
	assert.Nil(t, foo4)
}

func testCatch() (err error) {
	defer Catch(&err)
	Ignore(PtrResult((func() (*foo, error) { return &foo{}, io.EOF })()).NilIf(errors.New("bar")))
	return nil
}

func TestCatch(t *testing.T) {
	err := testCatch()
	assert.NotNil(t, err)
}
