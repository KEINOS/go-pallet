<!-- markdownlint-disable MD041 -->
[![go1.16+](https://img.shields.io/badge/Go-1.16+-blue?logo=go)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-pallet.svg)](https://pkg.go.dev/github.com/KEINOS/go-pallet/pallet)

# go-pallet<sub><sup><sup>beta</sup></sup></sub>

The `go-pallet` package is a Go library that simply gets the number of colors used in an image. For color pallet and/or histogram-base usage.

## Library Usage

```go
go get "github.com/KEINOS/go-pallet"
```

- Basic Usage

```go
/*
This sample yields a color palette with RGBA color combinations
and their number of occurrences as values.
*/

// import "github.com/KEINOS/go-pallet/pallet"

// Load image
imgRGBA, err := pallet.Load("/path/to/image/sample.png")
if err != nil {
  log.Fatal(err)
}

// Get all the color combinations used in an image.
// Returned data are sorted in order of frequency of use.
pixInfoList := pallet.ByOccurrence(imgRGBA)

// Print the first 2 most used colors. Which is:
//   46618 pixcels of RGBA(0,0,0,0) and
//   32505 pixcels of RGBA(208,182,152,255)
fmt.Println(pixInfoList[0:2])

// Output:
// [{0 0 0 0 46618} {208 182 152 255 32505}]
```

```go
/*
This sample yields a histogram of an image which is a color palette
with each channel's shade level (0-255) as a key and their number
of occurrences as values.
*/

// import "github.com/KEINOS/go-pallet/pallet"

// Load image
imgRGBA, err := pallet.Load("/path/to/image/sample.png")
if err != nil {
  log.Fatal(err)
}

// Get the image histogram.
hist := pallet.AsHistogram(imgRGBA)

// Print the occurrences of each color channel's shade level.
//   <channel>[<shade level>] = <occurrence>
// For example if a red pixel with max-opacity (R,G,B,A=255,0,0,255)
// appeared twice in an image then it will be:
//   r[255]=2, g[0]=2, b[0]=2, a[255]=2
fmt.Printf("r[0]=%v, r[255]=%v\n", hist.R[0], hist.R[255])
fmt.Printf("g[0]=%v, g[255]=%v\n", hist.G[0], hist.G[255])
fmt.Printf("b[0]=%v, b[255]=%v\n", hist.B[0], hist.B[255])
fmt.Printf("a[0]=%v, a[255]=%v\n", hist.A[0], hist.A[255])

// Output:
// r[0]=3, r[255]=1
// g[0]=3, g[255]=1
// b[0]=3, b[255]=1
// a[0]=1, a[255]=3
```

## Command Usage

Here's a [simple implementation](./cmd/main.go) of `go-pallet` as a CLI app.

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
$ pallet -h
pallet
  Simply print-outs the number of colors used or the histogram of an image in JSON.

Usage:
  main [options] <file path>

Options:

  -f, --file        file path of an image to analyze
      --histogram   print the histogram of the given image in JSON
  -p, --perline     prints each JSON elements per line
  -v, --version     displays app version
  -h, --help        display help information
```
```shellsession
$ pallet /path/to/image/sample.png
[{"r":255,"g":0,"b":0,"a":255,"count":1},{"r":0,"g":0,"b":255,"a":255,"count":1},{"r":0,"g":255,"b":0,"a":255,"count":1},{"r":0,"g":0,"b":0,"a":0,"count":1}]

$ pallet /path/to/image/sample.png --perline
[
{"r":255,"g":0,"b":0,"a":255,"count":1},
{"r":0,"g":0,"b":255,"a":255,"count":1},
{"r":0,"g":255,"b":0,"a":255,"count":1},
{"r":0,"g":0,"b":0,"a":0,"count":1}
]
```

---

## Statuses

This template adopts the below security measures to start with.

[![go1.16+](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Unit tests")
[![golangci-lint](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml "Static Analysis")
[![codecov](https://codecov.io/gh/KEINOS/go-pallet/branch/main/graph/badge.svg?token=uW30s2bK8M)](https://codecov.io/gh/KEINOS/go-pallet "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-pallet)](https://goreportcard.com/report/github.com/KEINOS/go-pallet "Code Quality")
[![CodeQL](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")

---

## License

- [MIT](https://github.com/KEINOS/go-pallet/LICENSE.txt) License. Copyright (c) [KEINOS](https://github.com/KEINOS) and [The Contributors](https://github.com/KEINOS/go-pallet/graphs/contributors).
