package main

import (
	"os"
	"os/exec"
)

func main() {
	gitRootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	gitRootCmd.Stdin = os.Stdin
	gitRootCmd.Stdout = os.Stdout
	gitRootCmd.Stderr = os.Stderr
	if err := gitRootCmd.Run(); err != nil {
		os.Exit(1)
	}
}
