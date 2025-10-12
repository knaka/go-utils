package net

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	//revive:disable-next-line:dot-imports
	. "github.com/knaka/go-utils"
	"github.com/knaka/go-utils/funcopt"
)

// fetchParams holds configuration parameters for fetch operations.
type fetchParams struct {
	stderr    io.Writer
	dir       string
	verbose   bool
	base      string
	maxredirs int
}

// WithDir sets the working directory for the fetch operation.
var WithDir = funcopt.NewFailable(func(params *fetchParams, dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dir)
	}
	params.dir = dir
	return nil
})

// WithStderr sets the stderr stream for the fetch operation.
var WithStderr = funcopt.New(func(params *fetchParams, stderr io.Writer) {
	params.stderr = stderr
})

// WithVerbose sets the verbosity.
var WithVerbose = funcopt.New(func(params *fetchParams, verbose bool) {
	params.verbose = verbose
})

// WithBase sets the basename of the output file.
var WithBase = funcopt.New(func(params *fetchParams, base string) {
	params.base = base
})

// WithMaxRedirs sets the maximum number of redirections to follow. Zero does not follow redirections. Negative number uses the default of the underlying library.
var WithMaxRedirs = funcopt.New(func(params *fetchParams, maxredirs int) {
	params.maxredirs = maxredirs
})

// Fetch downloads a file from the given URL to the local filesystem.
func Fetch(url string, opts ...funcopt.Option[fetchParams]) (err error) {
	defer Catch(&err)
	params := fetchParams{
		dir:       ".",
		verbose:   false,
		stderr:    os.Stderr,
		base:      path.Base(url),
		maxredirs: -1,
	}
	Must(funcopt.Apply(&params, opts))
	if params.verbose {
		fmt.Fprintf(params.stderr, "Fetching %s ...\n", url)
	}
	outPath := filepath.Join(params.dir, params.base)
	clt := http.Client{}
	clt.CheckRedirect = func(_ *http.Request, via []*http.Request) error {
		if params.maxredirs >= 0 && len(via) >= params.maxredirs {
			return fmt.Errorf("stopped after %d redirects", params.maxredirs)
		}
		return nil
	}
	resp := Value(clt.Get(url))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		Throw(fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status))
	}
	outFile := Value(os.Create(outPath))
	defer outFile.Close()
	Must(io.Copy(outFile, resp.Body))
	if params.verbose {
		fmt.Fprintf(params.stderr, "Saved to %s\n", outPath)
	}
	return
}
