package utils

import (
	"fmt"
	"os"
	"time"
)

const duration = 1 * time.Second

// Debugger waits for a debugger to connect if the environment variable $WAIT or $DEBUG is set
//
//goland:noinspection GoUnusedExportedFunction, GoUnnecessarilyExportedIdentifiers
func Debugger() {
	shouldWaitDebugger := false
	for i, arg := range os.Args {
		if arg == "--debugger" {
			shouldWaitDebugger = true
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			break
		}
	}
	if !shouldWaitDebugger &&
		os.Getenv("DEBUG") == "" &&
		os.Getenv("WAIT") == "" {
		return
	}
	pid := os.Getpid()
	V0(fmt.Fprintf(os.Stderr, "Process %d is waiting\n", pid))
	for {
		time.Sleep(duration)
		if debuggerProcessExists(pid) {
			break
		}
	}
	V0(fmt.Fprintf(os.Stderr, "Debugger connected"))
	time.Sleep(duration)
}

var WaitForDebugger = Debugger
