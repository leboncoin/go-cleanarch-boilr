# {{Name}}

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://{{GitServer}}/{{Organization}}/{{Name}}/blob/master/LICENSE)

## Install and Hooks

Make sure you installed [dep](https://github.com/golang/dep) and 
[golangci-lint](https://github.com/golangci/golangci-lint). Then install the
dependencies using the `dep` command:

```
$ go get -u github.com/golang/dep/cmd/dep
$ go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
$ dep ensure
```

To enable the pre-commit hook:

```
$ git config core.hooksPath .githooks
```

## Mocks Generation

```
$ make mock
```

## CI and Notifications

Please check out the [documentation about Drone](https://docs.lelab.io/services/tech/drone/)
to see how to activate and add Slack notifications to this repository