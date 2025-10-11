package utils

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV(t *testing.T) {
	reader := strings.NewReader("Hello, Reader!")
	bytes := make([]byte, 8)
	count := 0
	for {
		n := Result(reader.Read(bytes)).NilIf(io.EOF)
		if n == 0 {
			break
		}
		assert.True(t, n >= 0)
		count++
	}
	assert.GreaterOrEqual(t, count, 1)
}

func TestExpect(_ *testing.T) {
	Expect((func() error {
		return nil
	})(), nil, io.EOF)
	Expect((func() error {
		return io.EOF
	})(), nil, io.EOF)
	//Expect((func() error {
	//	return io.EOF
	//})(), nil, io.ErrClosedPipe)
}
