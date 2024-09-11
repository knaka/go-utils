//go:build darwin || linux
// +build darwin linux

package utils

// This function can be platform specific.
func debuggerProcessExists(pid int) (exists bool) {
	cmd := exec.Command("ps", "w")
	cmdOut := V(cmd.StdoutPipe())
	defer (func() { Ignore(cmdOut.Close()) })()
	scanner := bufio.NewScanner(cmdOut)
	V0(cmd.Start()) // Start() does not wait while Run() does
	defer (func() { V0(cmd.Wait()) })()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "dlv") &&
			strings.Contains(line, fmt.Sprintf("attach %d", pid)) {
			return true
		}
	}
	return false
}
