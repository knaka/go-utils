package main

import (
	"errors"
	"io"
	"log"
	"strings"

	. "github.com/knaka/go-utils"
)

func Bar() (err error) {
	defer Catch(&err)
	reader := strings.NewReader("Hello, Reader!")
	bytes := make([]byte, 8)
	for {
		_ = Value(reader.Read(bytes))
	}
}

func Foo() (err error) {
	defer Catch(&err)
	Must(Bar())
	return
}

func Main() (err error) {
	defer Catch(&err)
	Must(Foo())
	return
}

func main() {
	Must((func() error { return nil })())
	//err := Bar()
	err := Main()
	Assertf(errors.Is(err, io.EOF), "should be io.EOF: %v", err)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//if err != nil {
	//	panic(err)
	//}
}
