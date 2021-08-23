package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

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
	PrintCmd(cmd)
	return cmd.Run()
}
