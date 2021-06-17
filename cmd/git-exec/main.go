package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/paul-nelson-baker/go-git-plumbing/internal/utils"
)

func main() {
	// get the git project root directory
	gitRootCmd := exec.Command("git", "root")
	gitRootOutputBytes, err := gitRootCmd.Output()
	if err != nil {
		_, _ = os.Stderr.Write(gitRootOutputBytes)
		os.Exit(1)
	}
	gitRootDirectory := strings.TrimSpace(string(gitRootOutputBytes))
	// find out if an executable command is actually present
	commandIndex := utils.IndexOfStringInSlice("--", os.Args...)
	if commandIndex < 0 {
		_, _ = os.Stderr.WriteString(`You must specify a command to run after a "--"`)
		os.Exit(1)
	} else if commandIndex+2 > len(os.Args) {
		_, _ = os.Stderr.WriteString(`No executable command was given after "--"`)
		os.Exit(1)
	}
	// chop up the slice into smaller slices
	executableCommandSlice := os.Args[commandIndex+1:]
	executableCommand := executableCommandSlice[0]
	executableArgs := executableCommandSlice[1:]
	// execute the command that the user specifies
	executableCmd := exec.Command(executableCommand, executableArgs...)
	executableCmd.Stdin = os.Stdin
	executableCmd.Stdout = os.Stdout
	executableCmd.Stderr = os.Stderr
	executableCmd.Dir = gitRootDirectory
	if err := executableCmd.Run(); err != nil {
		os.Exit(1)
	}
}
