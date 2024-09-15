package main

import (
	. "github.com/knaka/go-utils"
	"log"
	"os"
)

func Foo() (file *os.File, err error) {
	defer Catch(&err)
	file = V(os.Open("not_exists"))
	return
}

func Bar() (file *os.File, err error) {
	defer Catch(&err)
	file = V(Foo())
	return
}

func main() {
	Debugger()
	for _, arg := range os.Args {
		log.Printf("arg: %s", arg)
	}
	file := V(os.Open("not_exists"))
	file, err := Bar()
	log.Printf("%+v, %+v", file, err)
}
