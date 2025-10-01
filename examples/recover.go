// Package main is a main package.
package main

import (
	"log"
	"os"

	//revive:disable-next-line:dot-imports
	. "github.com/knaka/go-utils"
)

func foo() (err error) {
	defer Catch(&err)
	V0(os.ReadDir("hoge"))
	//return io.EOF
	return nil
}

func main() {
	if err := foo(); err != nil {
		log.Printf("error: %+v", err)
	}
}
