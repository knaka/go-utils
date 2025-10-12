// Fetch command implementation.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/knaka/go-utils/net"
	"github.com/spf13/pflag"
)

var appID = "fetch"

func showUsage(cmdln *pflag.FlagSet, stderr io.Writer) {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] URL\n", appID)
	cmdln.SetOutput(stderr)
	cmdln.PrintDefaults()
}

func fetchEntry(args []string) (err error) {
	flags := pflag.NewFlagSet(appID, pflag.ContinueOnError)
	var shouldPrintHelp bool
	flags.BoolVarP(&shouldPrintHelp, "help", "h", false, "Show help")
	var output string
	flags.StringVarP(&output, "output", "o", "", "Write output to the specified file")
	err = flags.Parse(args)
	if err != nil {
		return
	}
	args = flags.Args()
	if shouldPrintHelp {
		showUsage(flags, os.Stderr)
		return
	}
	var opts net.Options
	if len(output) > 0 {
		opts = append(opts, net.WithOutput(output))
	}
	net.Fetch(args[0], opts...)
	return
}

func main() {
	err := fetchEntry(os.Args[1:])
	if err != nil {
		log.Fatalf("%s: %v\n", appID, err)
	}
}
