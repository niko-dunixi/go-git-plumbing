package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("You must specify at least one file or folder")
		os.Exit(1)
	}
	for _, currentFile := range files {
		gitCmd := exec.Command("git", "--no-pager", "log", `--pretty=tformat:%h`, "-n1", currentFile)
		gitCmd.Stdin = os.Stdin
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			os.Exit(1)
		}
	}
}
