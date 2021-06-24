package main

import (
	"flag"
	"os"
	"os/exec"
)

func main() {
	fullHashFlag := flag.Bool("full", false, "Get the full length hash of current commit")
	flag.Parse()
	gitCmdArgs := []string{"rev-list", "--max-count=1", "--skip=#", "HEAD"}
	if !*fullHashFlag {
		gitCmdArgs = append(gitCmdArgs, "--abbrev-commit")
	}
	gitCmd := exec.Command("git", gitCmdArgs...)
	gitCmd.Stdin = os.Stdin
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		os.Exit(1)
	}
}
