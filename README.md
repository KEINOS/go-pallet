<!-- markdownlint-disable MD041 -->
[![go1.14+](https://img.shields.io/badge/Go-1.14~17-blue?logo=go)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-pallet.svg)](https://pkg.go.dev/github.com/KEINOS/go-pallet)

# go-pallet<sub><sup><sup>alpha</sup></sup></sub>

`go-pallet` package is a library that simply returns the number of colors used in an image.

## Library Usage

```go
go get "github.com/KEINOS/go-pallet"
```

- Basic Usage

```go
// import "github.com/KEINOS/go-pallet/pallet"

// Load image
imgRGBA, err := pallet.Load("/path/to/image/sample.png")
if err != nil {
    return err
}

// Get list of RGBA combination and thouse number of occurance.
// The elements are sorted by its number of occurances.
pixList := pallet.ByOccurrence(imgRGBA)

// Print used color (RGBA combination) and its occurance.
fmt.Println(util.FmtStructPretty(pixList))

// Output:
// [
//   {"R": 178, "G": 156, "B": 130, "A": 255, "Count": 13},
//   {"R": 15, "G": 13, "B": 11, "A": 255, "Count": 12},
//   {"R": 4, "G": 4, "B": 3, "A": 255, "Count": 11},
//   ** snip **
// ]
```

## Command Usage

Here's a simple implementation of `go-pallet` as a CLI app.

## Install

- Via Homebrew (macOS, Linux)

```bash
brew install KEINOS/apps/go-pallet
```

- Manual Install
    - [releases page](https://github.com/KEINOS/go-pallet/releases/latest)
    - macOS (Intel/AMD64/M1), Windows (AMD64/Intel), Linux (Intel/AMD64, Arm 5,6,7, Arm64)

### Usage

```shellsession
$ pallet ./path/to/image/sample.png
[
  {"R": 178, "G": 156, "B": 130, "A": 255, "Count": 13},
  {"R": 15, "G": 13, "B": 11, "A": 255, "Count": 12},
  {"R": 4, "G": 4, "B": 3, "A": 255, "Count": 11},
  ** snip **
]
```

---

## Statuses

This template adopts the below security measures to start with.

[![go1.14+](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Unit tests")
[![golangci-lint](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml "Static Analysis")
[![codecov](https://codecov.io/gh/KEINOS/go-pallet/branch/main/graph/badge.svg?token=uW30s2bK8M)](https://codecov.io/gh/KEINOS/go-pallet "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-pallet)](https://goreportcard.com/report/github.com/KEINOS/go-pallet "Code Quality")
[![CodeQL](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")

---

<!-- You can use any license to use this template -->

## License

- [MIT](https://github.com/KEINOS/go-pallet/LICENSE.txt) License. Copyright (c) [KEINOS](https://github.com/KEINOS) and [The Contributors](https://github.com/KEINOS/go-pallet/graphs/contributors).
