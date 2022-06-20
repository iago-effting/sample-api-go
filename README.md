# Template go Experience

## Getting started

### Instal the language

```sh
asdf plugin-add golang
asdf install
```

You can run any command below:

- `make build` compile the projet to `bin/` folder
- `make run` execute bin app
- `make test` test the code
- `make lint` execute golint

## Stack

- asdf to lock the Go lang
- go lang
- golint
- makefile
- toml to file configuration

## Folders

```
...
configs
├── config.go -> code about set envs
├── config.toml -> configs common in all envs
├── dev.toml -> variables in the dev env
├── test.toml -> variables in the test env
pkg --> Core domain
├── http --> transport layer
│   ├── rest --> handlers/middlewares/input validations/output
├── version
│   ├── version.go --> package to use current version in code
...
```
