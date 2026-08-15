//go:build windows
// +build windows

package utils

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func debuggerProcessExists(_ int) (exists bool) {
	output := Value(exec.Command("tasklist.exe").Output())
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "dlv.exe") {
			return true
		}
	}
	return false
}
