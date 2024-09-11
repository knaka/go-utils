//go:build windows
// +build windows

package utils

import (
	"bufio"
	"os/exec"
	"strings"
)

// This function can be platform specific.
func debuggerProcessExists(pid int) (exists bool) {
	cmd := exec.Command("tasklist.exe")
	cmdOut := V(cmd.StdoutPipe())
	defer (func() { Ignore(cmdOut.Close()) })()
	scanner := bufio.NewScanner(cmdOut)
	V0(cmd.Start()) // Start() does not wait while Run() does
	defer (func() { V0(cmd.Wait()) })()
	for scanner.Scan() {
		line := scanner.Text()
		// No way to get the arguments of a process?
		if strings.Contains(line, "dlv.exe") {
			//&&strings.Contains(line, fmt.Sprintf("attach %d", pid)) {
			return true
		}
	}
	return false
}
