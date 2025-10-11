//go:build darwin || linux
// +build darwin linux

package utils

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func debuggerProcessExists(pid int) (exists bool) {
	cmd := exec.Command("ps", "wx")
	cmdOut := Value(cmd.StdoutPipe())
	defer (func() { Ignore(cmdOut.Close()) })()
	scanner := bufio.NewScanner(cmdOut)
	Must(cmd.Start()) // Start() does not wait while Run() does
	// On VSCode, os/exec.Command.Wait() (os/exec.Command.Process.Wait()) does not return after attached.
	// defer (func() { Must(cmd.Wait()) })()
	for scanner.Scan() {
		line := scanner.Text()
		// IntelliJ IDEA, GoLand
		if strings.Contains(line, "dlv") &&
			strings.Contains(line, fmt.Sprintf("attach %d", pid)) {
			return true
		}
		// VSCode. "Debug Adapter Protocol"
		if strings.Contains(line, "/dlv dap") {
			return true
		}
	}
	return false
}
