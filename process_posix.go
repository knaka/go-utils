//go:build darwin || linux
// +build darwin linux

package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func debuggerProcessExists(pid int) bool {
	output := Value(exec.Command("ps", "wx").Output())
	scanner := bufio.NewScanner(bytes.NewReader(output))
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
