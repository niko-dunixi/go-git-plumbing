# git pre-commit

The laziest way to integrate [pre-commit](https://pre-commit.com/) hooks with your repo

```bash
$ git pre-commit
```

If you don't have a `.pre-commit-config.yaml`, you can pass the `--init` flag and create the file.
This is not idempotent, if the file already exist the command will fail.

```bash
$ git pre-commit --init
```

You can also update the mutable references by passing the `--update` flag. 

```bash
$ git pre-commit --update
```
