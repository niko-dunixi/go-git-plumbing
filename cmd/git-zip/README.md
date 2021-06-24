# git zip

The simplest way to create an archive of your current project state.

```bash
$ git zip
> /Users/paulfreaknbaker/go-git-plumbing_83cc32e.zip
# Functionally equivalent to calling `git archive` and will create a
# new zip archive in the folder above the project directory
$ git zip --dirty
> /Users/paulfreaknbaker/go-git-plumbing_83cc32e_dirty.zip
# Instead of calling `git archive`, will execute the actual `zip`
# command instead. This can be used when you want to include files
# that aren't tracked with git, are .gitignored, or want to include
# the .git directory as well
```