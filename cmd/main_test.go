package main

import (
	"os"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

func Test_main(t *testing.T) {
	// Backup and defer recovery
	oldOsArgs := os.Args
	oldOsExit := util.OsExit

	defer func() {
		os.Args = oldOsArgs
		util.OsExit = oldOsExit
	}()

	var capturedCode int

	util.OsExit = func(code int) {
		capturedCode = code
	}

	{
		out := capturer.CaptureOutput(func() {
			os.Args = []string{t.Name(), "../testdata/r1g2b4a2.png"}

			main()
		})

		assert.Equal(t, 0, capturedCode, "status code should be zero on success")

		expect := "[" +
			"{\"r\":0,\"g\":0,\"b\":0,\"a\":0,\"count\":12}," +
			"{\"r\":255,\"g\":255,\"b\":255,\"a\":255,\"count\":6}," +
			"{\"r\":0,\"g\":0,\"b\":255,\"a\":255,\"count\":4}," +
			"{\"r\":0,\"g\":255,\"b\":0,\"a\":255,\"count\":2}," +
			"{\"r\":255,\"g\":0,\"b\":0,\"a\":255,\"count\":1}]\n"
		actual := out

		assert.Equal(t, expect, actual, "it should be in one line")
	}

	{
		out := capturer.CaptureOutput(func() {
			os.Args = []string{t.Name(), "-f", "../testdata/r1g2b4a2.png", "--perline"}

			main()
		})

		assert.Equal(t, 0, capturedCode, "status code should be zero on success")

		expect := "[\n" +
			"{\"r\":0,\"g\":0,\"b\":0,\"a\":0,\"count\":12},\n" +
			"{\"r\":255,\"g\":255,\"b\":255,\"a\":255,\"count\":6},\n" +
			"{\"r\":0,\"g\":0,\"b\":255,\"a\":255,\"count\":4},\n" +
			"{\"r\":0,\"g\":255,\"b\":0,\"a\":255,\"count\":2},\n" +
			"{\"r\":255,\"g\":0,\"b\":0,\"a\":255,\"count\":1}\n" +
			"]\n"
		actual := out

		assert.Equal(t, expect, actual, "it should print each element per line")
	}
}

func Test_main_as_histogram_default(t *testing.T) {
	// Backup and defer recovery
	oldOsArgs := os.Args
	oldOsExit := util.OsExit

	defer func() {
		os.Args = oldOsArgs
		util.OsExit = oldOsExit
	}()

	var capturedCode int

	util.OsExit = func(code int) {
		capturedCode = code
	}
	out := capturer.CaptureOutput(func() {
		os.Args = []string{t.Name(), "--histogram", "../testdata/r1g2b4a2.png"}

		main()
	})

	assert.Equal(t, 0, capturedCode, "exit status should be zero on success")

	expect := "{" +
		"\"r\":[18,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,7]," +
		"\"g\":[17,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,8]," +
		"\"b\":[15,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,10]," +
		"\"a\":[12,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,13]}\n"
	actual := out

	assert.Equal(t, expect, actual)
}

func Test_main_file_not_found(t *testing.T) {
	// Backup and defer recovery
	oldOsArgs := os.Args
	oldOsExit := util.OsExit

	defer func() {
		os.Args = oldOsArgs
		util.OsExit = oldOsExit
	}()

	var capturedCode int

	util.OsExit = func(code int) {
		capturedCode = code
	}

	out := capturer.CaptureOutput(func() {
		os.Args = []string{t.Name(), "unknown path"}

		main()
	})

	require.Equal(t, 1, capturedCode, "exit status should be 1 on error")
	assert.Contains(t, out, "failed to load image: open unknown path")
}

func Test_main_no_args(t *testing.T) {
	// Backup and defer recovery
	oldOsArgs := os.Args
	oldOsExit := util.OsExit

	defer func() {
		os.Args = oldOsArgs
		util.OsExit = oldOsExit
	}()

	var capturedCode int

	util.OsExit = func(code int) {
		capturedCode = code
	}
	out := capturer.CaptureStderr(func() {
		os.Args = []string{t.Name()}

		main()
	})

	require.Equal(t, 1, capturedCode, "exit status should be 1 on missing args")
	assert.Contains(t, out, "argument missing. At least a file path is required.")
}

func Test_main_show_version(t *testing.T) {
	// Backup and defer recovery
	oldOsArgs := os.Args
	oldOsExit := util.OsExit

	defer func() {
		os.Args = oldOsArgs
		util.OsExit = oldOsExit
	}()

	var capturedCode int

	util.OsExit = func(code int) {
		capturedCode = code
	}

	out := capturer.CaptureOutput(func() {
		os.Args = []string{t.Name(), "--version"}

		main()
	})

	require.Equal(t, 0, capturedCode, "exit status should be zero on version display")
	assert.Contains(t, out, "cmd.test")
	assert.Contains(t, out, "(devel)")
}

func TestGetVersion(t *testing.T) {
	// Backup and defer restore
	oldVersion := version
	oldCommitID := commit

	defer func() {
		version = oldVersion
		commit = oldCommitID
	}()

	// Mock version
	version = "1.2.3"
	commit = "abc"

	// Get
	actual := GetVersion()

	// Assert
	assert.Contains(t, actual, "cmd.test")
	assert.Contains(t, actual, "1.2.3 (abc)")
}
