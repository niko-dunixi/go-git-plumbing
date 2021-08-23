package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	Flags = log.Ltime | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix
)

var (
	logger = NewTraceLogger()
)

func NewTraceLogger() *log.Logger {
	return log.New(os.Stderr, "\t\ttrace:", Flags)
}

// Checks the environment variables for GIT_TRACE, which is used by git
// to indicate the user wants debug output to print to STDERR.
//
// See: https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables
func ShouldTrace() bool {
	if gitTraceValue, isPresent := os.LookupEnv("GIT_TRACE"); isPresent {
		gitTraceValue = strings.TrimSpace(gitTraceValue)
		gitTraceValue = strings.ToLower(gitTraceValue)
		if isTrace := IndexOfStringInSlice(gitTraceValue, "true", "1", "2") >= 0; isTrace {
			return true
		}
	}
	return false
}

func PrintCmd(cmd *exec.Cmd) {
	if !ShouldTrace() {
		return
	}
	logger.Printf("$ %s \n", strings.Join(cmd.Args, " "))
}
