package net

import (
	"fmt"
	"io"
	"net/http"
	"os"

	//revive:disable-next-line:dot-imports
	. "github.com/knaka/go-utils"
	"github.com/knaka/go-utils/funcopt"
)

// fetchParams holds configuration parameters for fetch operations.
type fetchParams struct {
	stdout    io.Writer
	stderr    io.Writer
	verbose   bool
	maxredirs int
	output    string
}

// Options is functional options type
type Options []funcopt.Option[fetchParams]

// WithStderr sets the stderr stream for the fetch operation.
var WithStderr = funcopt.New(func(params *fetchParams, stderr io.Writer) {
	params.stderr = stderr
})

// WithVerbose sets the verbosity.
var WithVerbose = funcopt.New(func(params *fetchParams, verbose bool) {
	params.verbose = verbose
})

// WithMaxRedirs sets the maximum number of redirections to follow. Zero does not follow redirections. Negative number uses the default of the underlying library.
var WithMaxRedirs = funcopt.New(func(params *fetchParams, maxredirs int) {
	params.maxredirs = maxredirs
})

// WithOutput writes output to the specified file instead of stdout.
var WithOutput = funcopt.New(func(params *fetchParams, output string) {
	params.output = output
})

// Fetch downloads a file from the given URL to the local filesystem.
func Fetch(url string, opts ...funcopt.Option[fetchParams]) (err error) {
	defer Catch(&err)
	params := fetchParams{
		verbose:   false,
		stdout:    os.Stdout,
		stderr:    os.Stderr,
		maxredirs: -1,
		output:    "",
	}
	Must(funcopt.Apply(&params, opts))
	if params.verbose {
		fmt.Fprintf(params.stderr, "Fetching %s ...\n", url)
	}
	client := http.Client{}
	client.CheckRedirect = func(_ *http.Request, via []*http.Request) error {
		if params.maxredirs >= 0 && len(via) >= params.maxredirs {
			return fmt.Errorf("stopped after %d redirects", params.maxredirs)
		}
		return nil
	}
	resp := Value(client.Get(url))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		Throw(fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status))
	}
	var writer io.Writer
	if len(params.output) > 0 {
		writeCloser := Value(os.Create(params.output))
		defer writeCloser.Close()
		writer = writeCloser
	} else {
		writer = params.stdout
	}
	Must(io.Copy(writer, resp.Body))
	if params.verbose && len(params.output) > 0 {
		fmt.Fprintf(params.stderr, "Saved to %s\n", params.output)
	}
	return
}
