package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/paul-nelson-baker/go-git-plumbing/internal/utils"
)

func main() {
	isVerbose := flag.Bool("verbose", false, "Show output for subcommands")
	shouldZipDirty := flag.Bool("dirty", false, "Include all the ignored files as well")
	flag.Parse()

	gitHashOutputBytes, err := exec.Command("git", "hash").CombinedOutput()
	if err != nil {
		os.Stderr.Write(gitHashOutputBytes)
		os.Exit(1)
	}
	gitHash := strings.TrimSpace(string(gitHashOutputBytes))
	gitProjectRootDirectory := utils.MustGitProjectRootDirectory()
	projectName := path.Base(gitProjectRootDirectory)
	archiveFilename := createArchiveFilename(projectName, gitHash, *shouldZipDirty)
	archiveFilepath := path.Join(gitProjectRootDirectory, "..", archiveFilename)

	if err := os.Remove(archiveFilepath); err != nil && !os.IsNotExist(err) {
		os.Stderr.WriteString(`Could not remove already created file: ` + archiveFilepath)
		os.Exit(1)
	}

	zipCommand := createZipCommand(archiveFilepath, *shouldZipDirty)
	zipCommand.Dir = gitProjectRootDirectory
	if *isVerbose {
		fmt.Println("$ " + strings.Join(zipCommand.Args, " "))
		zipCommand.Stdout = os.Stderr
		zipCommand.Stderr = os.Stderr
	}
	if err := zipCommand.Run(); err != nil {
		os.Exit(1)
	}
	fmt.Println(archiveFilepath)
}

func createArchiveFilename(projectName, gitHash string, isDirty bool) (archiveFilename string) {
	projectName = strings.TrimSpace(projectName)
	gitHash = strings.TrimSpace(gitHash)
	archiveFilename += projectName + "_" + gitHash
	if isDirty {
		archiveFilename += "_dirty"
	}
	archiveFilename += ".zip"
	return
}

func createZipCommand(archiveFilepath string, isDirty bool) *exec.Cmd {
	if isDirty {
		return exec.Command("zip", "-r", archiveFilepath, ".")
	}
	return exec.Command("git", "archive",
		"-v",
		"--format=zip",
		"--output="+archiveFilepath,
		"HEAD",
	)
}
