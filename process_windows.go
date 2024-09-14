//go:build windows
// +build windows

package utils

import (
	"bufio"
	"os/exec"
	"strings"
)

// This function can be platform specific.
func debuggerProcessExists(_pid int) (exists bool) {
	cmd := exec.Command("tasklist.exe")
	cmdOut := V(cmd.StdoutPipe())
	defer (func() { Ignore(cmdOut.Close()) })()
	scanner := bufio.NewScanner(cmdOut)
	V0(cmd.Start()) // Start() does not wait while Run() does
	defer (func() { V0(cmd.Wait()) })()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "dlv.exe") {
			return true
		}
	}
	return false
}
