// Entry
package main

import (
	"log"
	"os"

	//revive:disable-next-line dot-import
	. "github.com/knaka/go-utils"
)

// Foo is a function.
func Foo() (file *os.File, err error) {
	defer Catch(&err)
	file = V(os.Open("not_exists"))
	return
}

// Bar is also a function.
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
