package utils

import (
	"fmt"
	"os"
	"time"
)

const duration = 1 * time.Second

// Debugger waits for a debugger to connect if the environment variable $WAIT_DEBUGGER is set or the first argument is "--wait-debugger"
//
//goland:noinspection GoUnusedExportedFunction, GoUnnecessarilyExportedIdentifiers
func Debugger() {
	shouldWaitDebugger := false
	if len(os.Args) > 1 && os.Args[1] == "--wait-debugger" {
		shouldWaitDebugger = true
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
	if !shouldWaitDebugger && os.Getenv("WAIT_DEBUGGER") == "" {
		return
	}
	pid := os.Getpid()
	Must(fmt.Fprintf(os.Stderr, "Process %d is waiting for a debugger to connect.\n", pid))
	for {
		time.Sleep(duration)
		if debuggerProcessExists(pid) {
			break
		}
	}
	Must(fmt.Fprintf(os.Stderr, "Debugger has connected.\n"))
	time.Sleep(duration)
}
