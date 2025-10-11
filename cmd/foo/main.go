// Entry
package main

import (
	"log"
	"os"

	_ "github.com/knaka/go-utils/initwait"

	//revive:disable-next-line:dot-imports
	. "github.com/knaka/go-utils"
)

// Foo is a function.
func Foo() (file *os.File, err error) {
	defer Catch(&err)
	file = Value(os.Open("not_exists"))
	return
}

// Bar is also a function.
func Bar() (file *os.File, err error) {
	defer Catch(&err)
	file = Value(Foo())
	return
}

func main() {
	for _, arg := range os.Args {
		log.Printf("arg: %s", arg)
	}
	file := Value(os.Open("not_exists"))
	file, err := Bar()
	log.Printf("%+v, %+v", file, err)
}
