//go:build windows
// +build windows

package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func debuggerProcessExists(pid int) (exists bool) {
	// tasklist.exe cannot show command-line arguments; Win32_Process via CIM can,
	// which lets us tell an "attach <pid>" (IntelliJ/GoLand) session from a "dap" (VSCode) one.
	output := Value(exec.Command(
		"powershell.exe", "-NoProfile", "-NonInteractive", "-Command",
		`Get-CimInstance Win32_Process -Filter "Name='dlv.exe'" | Select-Object -ExpandProperty CommandLine`,
	).Output())
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		// IntelliJ IDEA, GoLand
		if strings.Contains(line, fmt.Sprintf("attach %d", pid)) {
			return true
		}
		// VSCode. "Debug Adapter Protocol"
		if strings.Contains(line, " dap") {
			return true
		}
	}
	Must(scanner.Err())
	return false
}
