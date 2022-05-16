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

## Folders

```
cmd --> programs/binaries
├── main.go --> Ways to start this program (cli mode, server, etc)*
pkg --> Core domain
├── version
│   ├── version.go --> package to use current version in code
scripts --> scripts utils to run/test/cover/check the app
├── build.sh
├── lint.sh
├── run.sh
├── test.sh
.gitlab-ci.yml --> CI configuration
.golangci.yml --> Linter configuration
.tool-versions --> lock language version
go.mod
go.sum
makefile --> automation tool 
README.md --> getting started
```

__*__ This example is just one in cmd. But if we need for example, a cli and a server, we can put the main inside one folder calls server e create another folder bellow cmd calls cli and create a way to execute the staffs we need. (For this work, we need update the **build.sh** to get a array of folders and build each one)

```
...
cmd
├── server
│   ├── main.go
├── cli
│   ├── main.go
...
```
