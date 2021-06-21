package main

import (
	"os"

	"github.com/paul-nelson-baker/go-git-plumbing/internal/utils"
)

func main() {
	gitRootDirectory := utils.MustGitProjectRootDirectory()
	_, _ = os.Stdout.WriteString(gitRootDirectory)
}
