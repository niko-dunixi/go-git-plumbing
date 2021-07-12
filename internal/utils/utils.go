package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

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

func IndexOfStringInSlice(item string, items ...string) int {
	for i := range items {
		if item == items[i] {
			return i
		}
	}
	return -1
}

func MustGitProjectRootDirectory() string {
	stdOutBuffer := bytes.Buffer{}
	gitRootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	gitRootCmd.Stdin = os.Stdin
	gitRootCmd.Stdout = &stdOutBuffer
	gitRootCmd.Stderr = os.Stderr
	if err := gitRootCmd.Run(); err != nil {
		os.Exit(1)
	}
	return strings.TrimSpace(stdOutBuffer.String())
}

func MustLazyRunCmd(command string, args ...string) {
	if err := LazyRunCmd(command, args...); err != nil {
		os.Exit(1)
	}
}

func LazyRunCmd(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = MustGitProjectRootDirectory()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	TraceCmd(cmd)
	return cmd.Run()
}

func TraceCmd(cmd *exec.Cmd) {
	if !ShouldTrace() {
		return
	}
	os.Stdout.WriteString("$ " + strings.Join(cmd.Args, " ") + "\n")
}
