package utils

import (
	"fmt"
	"io"
)

var logWriters []io.Writer

// AddLogWriter adds a writer to the list of log writers.
func AddLogWriter(writer io.Writer) {
	logWriters = append(logWriters, writer)
}

// LogPrintf writes formatted output to all registered log writers.
func LogPrintf(format string, v ...interface{}) {
	for _, writer := range logWriters {
		Must(fmt.Fprintf(writer, format, v...))
	}
}

// LogPrint writes output to all registered log writers.
func LogPrint(v ...interface{}) {
	for _, writer := range logWriters {
		Must(fmt.Fprint(writer, v...))
	}
}

// LogPrintln writes output with a newline to all registered log writers.
func LogPrintln(v ...interface{}) {
	for _, writer := range logWriters {
		Must(fmt.Fprintln(writer, v...))
	}
}
