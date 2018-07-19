# go-cleanarch-boilr
[boilr](https://github.com/tmrts/boilr) template to create a clean design 
architecture golang backend

## Introduction

This template is an opinionated clean design architecture boilr template. It has
many automated and opinionated features:

- Uses [dep](https://github.com/golang/dep) for dependency management
- Uses [golangci-lint](ttps://github.com/golangci/golangci-lint) to lint
- Pre-commit hooks that runs the tests and the linter
- Makefile with several targets
- Uses [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper)
  for command-line commands and configuration
- Multi-stage independant Dockerfile:
    - Build using `golang:latest`
    - Deploy in `alpine`
- Version and build number injection in the binary using git history and tags
- Drone as the CI (this can be completely ignored or deleted if another CI is used)
- MIT License

## Usage

### Install Boilr

To install Boilr, follow [this guide](https://github.com/tmrts/boilr/wiki/Installation)
or run the following command:

```
$ go get -u github.com/tmrts/boilr
```

### Install and use the template

```
$ boilr template download leboncoin/go-cleanarch-boilr go-cleanarch
$ cd $GOPATH/src/github.com/myuser/
$ boilr template use go-cleanarch-public myamazingproject
[?] Please choose a value for "Name" [default: "example"]: myamazingproject
[✔] Created .dockerignore
[?] Please choose a value for "GitServer" [default: "github.com"]: 
[?] Please choose a value for "Organization" [default: "username"]: myuser  
[✔] Created .drone.yml
[✔] Created .githooks/pre-commit
[✔] Created .gitignore
[✔] Created Dockerfile
[✔] Created Gopkg.lock
[✔] Created Gopkg.toml
[✔] Created LICENSE
[✔] Created Makefile
[✔] Created README.md
[✔] Created domain/README.md
[✔] Created implem/gin.server/health.go
[✔] Created implem/gin.server/health_test.go
[✔] Created implem/gin.server/router.go
[✔] Created implem/mock.uc/builder.go
[✔] Created infra/conf.go
[✔] Created infra/gin.go
[✔] Created main.go
[✔] Created uc/health.go
[✔] Created uc/interactor.go
[✔] Successfully executed the project template go-cleanarch-public in /home/paul/Prog/Go/src/github.com/myuser/myamazingproject
```

The instructions on how to use your new project are stored in the `README.md`
file in your new projects, including all the options, configuration and Makefile
targets.

## Make Targets

The version is either `0.1.0` if no tag has ever been defined or the latest
tag defined. The build number is the SHA1 of the latest commit.

- **make**: Builds and injects version/build in binary
- **make init**: Sets the pre-commit hook in the repository
- **make docker**: Build docker image and tag it with both `latest` and version
- **make latest**: Build docker image and tag it only with `latest`
- **make test**: Executes the test suite
- **make mock**: Generate the necessary mocks
- **make clean**: Removes the built binary if present