package main

import (
	_ "embed"
	"flag"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/paul-nelson-baker/go-git-plumbing/internal/utils"
)

//go:embed pre-commit-config.yaml
var precommitYamlBytes []byte

func main() {
	gitRootDirectory := utils.MustGitProjectRootDirectory()
	initializeYamlFile := flag.Bool("init", false, "Initialize the pre-commit yaml file")
	updateHooks := flag.Bool("update", false, "Update the hooks using mutable references")
	flag.Parse()
	// If specified, create an initial commit hook config but don't overwrite it if it's existing
	if *initializeYamlFile {
		precommitYamlFilename := path.Join(gitRootDirectory, ".pre-commit-config.yaml")
		if _, err := os.Stat(precommitYamlFilename); err == nil {
			os.Stderr.WriteString("Can't initialize, " + precommitYamlFilename + " already exists\n")
			os.Exit(1)
		}
		os.WriteFile(precommitYamlFilename, precommitYamlBytes, 0644)
	}
	// Run the install
	precommitInstallCmd := exec.Command("pre-commit", "install", "--install-hooks")
	os.Stdout.WriteString("$ " + strings.Join(precommitInstallCmd.Args, " ") + "\n")
	precommitInstallCmd.Dir = gitRootDirectory
	precommitInstallCmd.Stdin = os.Stdin
	precommitInstallCmd.Stdout = os.Stdout
	precommitInstallCmd.Stderr = os.Stderr
	if err := precommitInstallCmd.Run(); err != nil {
		os.Exit(1)
	}
	// Update the hooks
	if *updateHooks {
		precommitUpdateCmd := exec.Command("pre-commit", "autoupdate")
		os.Stdout.WriteString("$ " + strings.Join(precommitUpdateCmd.Args, " ") + "\n")
		precommitUpdateCmd.Dir = gitRootDirectory
		precommitUpdateCmd.Stdin = os.Stdin
		precommitUpdateCmd.Stdout = os.Stdout
		precommitUpdateCmd.Stderr = os.Stderr
		if err := precommitInstallCmd.Run(); err != nil {
			os.Exit(1)
		}
	}
}
