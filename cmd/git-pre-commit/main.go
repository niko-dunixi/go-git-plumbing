package main

import (
	_ "embed"
	"path"

	"os"

	"github.com/jessevdk/go-flags"
	"github.com/niko-dunixi/go-git-plumbing/internal/utils"
)

//go:embed pre-commit-config.yaml
var precommitYamlBytes []byte

func main() {
	// Setup the flags with subcommands and do the parsing/validation
	opts := struct {
		Init      struct{} `command:"init" description:"creates a .pre-commit-config.yaml, but fails if one already exists"`
		Exec      struct{} `command:"exec" description:"passes all arguments after -- to pre-commit for execution"`
		Run       struct{} `command:"run" description:"wrapper for 'pre-commit run'"`
		Install   struct{} `command:"install" description:"wrapper for 'pre-commit install --install-hooks'"`
		Update    struct{} `command:"update" description:"wrapper for 'pre-commit autoupdate'"`
		Uninstall struct{} `command:"uninstall" description:"wrapper for 'pre-commit uninstall'"`
	}{}
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	args, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	// Action upon user input
	switch parser.Active.Name {
	// base execution cases
	case "init":
		gitRootDirectory := utils.MustGitProjectRootDirectory()
		precommitYamlFilename := path.Join(gitRootDirectory, ".pre-commit-config.yaml")
		if _, err := os.Stat(precommitYamlFilename); err == nil {
			os.Stderr.WriteString(precommitYamlFilename + " already exists\n")
			os.Exit(1)
		}
		os.WriteFile(precommitYamlFilename, precommitYamlBytes, 0644)
	case "exec":
		utils.MustLazyRunCmd("pre-commit", args...)
	// recursive execution cases
	case "run":
		utils.MustLazyRunCmd("git", "pre-commit", "exec", "--", "run", "--all-files")
	case "install":
		utils.MustLazyRunCmd("git", "pre-commit", "exec", "--", "install", "--install-hooks")
	case "update":
		utils.MustLazyRunCmd("git", "pre-commit", "exec", "--", "autoupdate")
	case "uninstall":
		utils.MustLazyRunCmd("git", "pre-commit", "exec", "--", "uninstall")
	}
}
