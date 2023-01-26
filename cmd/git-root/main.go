package main

import (
	"os"

	"github.com/niko-dunixi/go-git-plumbing/internal/utils"
)

func main() {
	gitRootDirectory := utils.MustGitProjectRootDirectory()
	_, _ = os.Stdout.WriteString(gitRootDirectory)
}
