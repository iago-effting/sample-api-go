# Template go Experience

Author: [@iagoEffting](https://twitter.com/iagoEffting)

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


### Database

This configurations just work with postgres. To run locally you can use docker.
The configuration of access are in `docker-compose.yml` file

```
docker-compose up
```

## CLI

We have a cli to manager your app

```sh
go run ./cmd/cli/main.go
```

There you can:

- Create a migration: `go main.go make migration name_of_migration`
- Run all migrations `go main.go migrate`
- Rollback all migration `go main.go migrate reset`

In the future:

- show all routes

-> We need change the execution of cli. It's too verbose. Maybe a build and send to `bin`. I am thinking in that yet. (makefile doesn't work well for that.)

## Stack

- go lang
- makefile
- asdf (https://asdf-vm.com/)
- golint (https://github.com/golang/lint)
- toml to file configuration (https://toml.io/en/)
- bun ORM (https://bun.uptrace.dev/)
- Database migration (https://github.com/golang-migrate/migrate/)
- apitest (https://apitest.dev/)
- bcrypt (https://pkg.go.dev/golang.org/x/crypto/bcrypt)

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