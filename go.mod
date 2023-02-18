module github.com/KEINOS/go-pallet

go 1.16

require (
	github.com/KEINOS/go-utiles v1.5.3
	github.com/labstack/gommon v0.4.0
	github.com/mkideal/cli v0.2.7
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	github.com/zenizh/go-capturer v0.0.0-20211219060012-52ea6c8fed04
	golang.org/x/image v0.5.0
)

// CVE-2021-43565
require golang.org/x/crypto v0.0.0-20220314234659-1baeb1ce4c0b // indirect
