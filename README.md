<!-- markdownlint-disable MD041 MD033 -->
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/KEINOS/go-pallet)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-pallet.svg)](https://pkg.go.dev/github.com/KEINOS/go-pallet/pallet)

# go-pallet

`go-pallet` is a Go package for analyzing colors in images.
It also includes the `pallet` command, which prints color occurrence data or per-channel histograms as JSON.

## Library Usage

```console
go get github.com/KEINOS/go-pallet/pallet
```

- Basic Usage

```go
/*
  This example returns the unique RGBA colors in an image
  and their occurrence counts.
*/

// import "github.com/KEINOS/go-pallet/pallet"

// Load image
imgRGBA, err := pallet.Load("/path/to/image/sample.png")
if err != nil {
  log.Fatal(err)
}

// Get all unique colors used in the image.
// Results are sorted by descending occurrence count, with RGBA values used
// as a deterministic ascending tie-breaker.
pixInfoList := pallet.ByOccurrence(imgRGBA)

// Print the two most common colors:
//   46618 pixels of RGBA(0,0,0,0) and
//   32505 pixels of RGBA(208,182,152,255)
fmt.Println(pixInfoList[0:2])

// Output:
// [{0 0 0 0 46618} {208 182 152 255 32505}]
```

```go
/*
  This example returns an image histogram.
  Each channel uses shade levels (0-255) as indexes,
  and each value is the number of pixels at that level.
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

The repository also includes the [`pallet` command](./cmd/pallet/main.go).

## Install

- Via Homebrew (macOS, Linux)

```bash
brew install KEINOS/apps/go-pallet
```

- Via Go install (all platforms)

```bash
go install "github.com/KEINOS/go-pallet/cmd/pallet@latest"
```

- Prebuilt binaries
  - Download from the [latest release](https://github.com/KEINOS/go-pallet/releases/latest).

### Usage

```shellsession
$ pallet -h
pallet
  Print the colors used in an image, or print its histogram as JSON.

Usage:
  pallet [options] <file path>

Options:

  -f, --file        file path of an image to analyze
      --histogram   print the histogram of the given image in JSON
      --jsonl       print each color as a JSON object on its own line
  -p, --perline     print each JSON element on its own line
  -v, --version     print the app version
  -h, --help        display help information
```

```shellsession
$ pallet /path/to/image/sample.png
[{"r":0,"g":0,"b":0,"a":0,"count":1},{"r":0,"g":0,"b":255,"a":255,"count":1},{"r":0,"g":255,"b":0,"a":255,"count":1},{"r":255,"g":0,"b":0,"a":255,"count":1}]

$ pallet /path/to/image/sample.png --perline
[
{"r":0,"g":0,"b":0,"a":0,"count":1},
{"r":0,"g":0,"b":255,"a":255,"count":1},
{"r":0,"g":255,"b":0,"a":255,"count":1},
{"r":255,"g":0,"b":0,"a":255,"count":1}
]

$ pallet /path/to/image/sample.png --jsonl
{"r":0,"g":0,"b":0,"a":0,"count":1}
{"r":0,"g":0,"b":255,"a":255,"count":1}
{"r":0,"g":255,"b":0,"a":255,"count":1}
{"r":255,"g":0,"b":0,"a":255,"count":1}
```

The `--jsonl` option cannot be combined with `--perline` or `--histogram`.

---

## Statuses

This project uses the following status checks.

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/KEINOS/go-pallet)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Supported versions")
[![golangci-lint](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml "Static Analysis")
[![codecov](https://codecov.io/gh/KEINOS/go-pallet/branch/main/graph/badge.svg?token=uW30s2bK8M)](https://codecov.io/gh/KEINOS/go-pallet "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-pallet)](https://goreportcard.com/report/github.com/KEINOS/go-pallet "Code Quality")
[![CodeQL](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")

---

## License

- [MIT](./LICENSE) License. Copyright (c) [KEINOS](https://github.com/KEINOS) and [The Contributors](https://github.com/KEINOS/go-pallet/graphs/contributors).
