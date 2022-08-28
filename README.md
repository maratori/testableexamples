# testableexamples <br> [![CI][ci-img]][ci-url] [![Codecov][codecov-img]][codecov-url] [![Codebeat][codebeat-img]][codebeat-url] [![Maintainability][codeclimate-img]][codeclimate-url] [![Go Report Card][goreportcard-img]][goreportcard-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]

Linter checks if examples are testable.

Author of idea is [Jamie Tanna](https://github.com/jamietanna) (see [this issue](https://github.com/golangci/golangci-lint/issues/3084)).


## Usage as standalone linter

### Install
```shell
go install github.com/maratori/testableexamples@latest
```

### Run

```shell
testableexamples ./...
```


[ci-img]: https://github.com/maratori/testableexamples/actions/workflows/ci.yml/badge.svg
[ci-url]: https://github.com/maratori/testableexamples/actions/workflows/ci.yml
[codecov-img]: https://codecov.io/gh/maratori/testableexamples/branch/main/graph/badge.svg?token=VMXc2fc7cJ
[codecov-url]: https://codecov.io/gh/maratori/testableexamples
[codebeat-img]: https://codebeat.co/badges/1b813bf1-336d-4886-b4fa-1d482bedc754
[codebeat-url]: https://codebeat.co/projects/github-com-maratori-testableexamples-main
[codeclimate-img]: https://api.codeclimate.com/v1/badges/47ed5db4a7595d4f95d5/maintainability
[codeclimate-url]: https://codeclimate.com/github/maratori/testableexamples/maintainability
[goreportcard-img]: https://goreportcard.com/badge/github.com/maratori/testableexamples
[goreportcard-url]: https://goreportcard.com/report/github.com/maratori/testableexamples
[license-img]: https://img.shields.io/github/license/maratori/testableexamples.svg
[license-url]: /LICENSE
[godoc-img]: https://pkg.go.dev/badge/github.com/maratori/testableexamples.svg
[godoc-url]: https://pkg.go.dev/github.com/maratori/testableexamples
