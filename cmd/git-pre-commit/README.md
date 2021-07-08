# git pre-commit

The laziest way to integrate [pre-commit](https://pre-commit.com/) hooks with your repo

<!-- toc -->

* [Subcommands](#subcommands)
  * [init](#init)
  * [exec](#exec)
  * [Wrapping Commands](#wrapping-commands)

<!-- Regenerate with "pre-commit run -a markdown-toc" -->

<!-- tocstop -->

## Subcommands
### init

```bash
$ git pre-commit init
```

If you don't have a `.pre-commit-config.yaml` this will create one for you.
This is not idempotent, if the file already exist the command will fail.

### exec
```bash
$ git pre-commit exec -- --help
```

Anything after the `--` argument will be passed directly to pre-commit

### Wrapping Commands

| subcommand | wrapped command  |
| ---------: | ---------------- |
| `run`      | `pre-commit run` |
| `install`  | `pre-commit install --install-hooks` |
| `update`   | `pre-commit autoupdate` |
| `uninstall`| `pre-commit uninstall ` |