package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

func Test_main(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Contains(t, out, "Hello, Gopher!\n")
}

func TestGetVersion(t *testing.T) {
	ver := GetVersion()

	assert.Contains(t, ver, "")
}
