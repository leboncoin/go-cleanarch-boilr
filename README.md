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

### Install the template

```
$ boilr template download leboncoin/go-cleanarch-boilr go-cleanarch
```