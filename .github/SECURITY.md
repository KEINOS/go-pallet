# Security Policy

## Fail Fast Policy

We update the go.mod versions every week to the latest versions. If the tests fail, we will fix them immediately.

It may break the backward compatibility but we will take it as a priority and try to fix it ASAP.

## Supported Go Versions

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/KEINOS/go-pallet)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Supported versions") 〜 [![go latest](https://img.shields.io/badge/Go-latest-blue?logo=go)](https://github.com/KEINOS/go-pallet/actions/workflows/go-versions.yml "Supported versions")

## Security concern

As a minimum security measure, we keep the following to be green.

[![golangci-lint](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/golangci-lint.yml "Static Analysis")
[![CodeQL](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/go-pallet/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")
[![codecov](https://codecov.io/gh/KEINOS/go-pallet/branch/main/graph/badge.svg?token=uW30s2bK8M)](https://codecov.io/gh/KEINOS/go-pallet "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-pallet)](https://goreportcard.com/report/github.com/KEINOS/go-pallet "Code Quality")

> __Note__: Code coverage and Go Report Card does nothing to do with security but we keep them high to make it easier to maintain the code.

### Security Status

- Check the current "[Security overview](https://github.com/KEINOS/dev-go/security)" status.

## Reporting a Vulnerability

- Please [issue](https://github.com/KEINOS/dev-go/issues) them.
