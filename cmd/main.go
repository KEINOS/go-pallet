package main

import (
	"fmt"
	"runtime/debug"

	"github.com/KEINOS/go-pallet/pallet"
	"github.com/KEINOS/go-utiles/util"
	"github.com/labstack/gommon/color"
	"github.com/mkideal/cli"
	"github.com/pkg/errors"
)

var (
	version string // The app version assigned via build flag.
	commit  string // The commit ID assigned via build flag.
)

type argT struct {
	PathFileImg  string `cli:"f,file" usage:"file path of an image to analyze"`
	AsHistogram  bool   `cli:"histogram" usage:"print the histogram of the given image in JSON"`
	PrintPerLine bool   `cli:"p,perline" usage:"prints each JSON elements per line"`
	ShowVersion  bool   `cli:"v,version" usage:"displays app version"`
	cli.Helper
}

func main() {
	util.OsExit(cli.Run(
		new(argT),
		PreRun,
		color.Bold(util.GetNameBin()),
		"  Simply print-outs the number of colors used or the histogram of an image in JSON.\n",
		color.Bold("Usage:"),
		fmt.Sprintf("  %s [options] <file path>", util.GetNameBin()),
	))
}

// PreRun allots function according to the flag options.
func PreRun(ctx *cli.Context) error {
	argv, _ := ctx.Argv().(*argT)
	args := ctx.Args()

	// Show app version
	if argv.ShowVersion {
		_ = ctx.String("%v\n", GetVersion())

		return nil
	}

	// Get target file path
	pathFileImg := argv.PathFileImg

	if pathFileImg == "" {
		// Missing args
		if len(args) == 0 {
			return errors.New("argument missing. At least a file path is required.\n" + ctx.Usage())
		}

		pathFileImg = args[0]
	}

	// Run
	result, err := Run(pathFileImg, argv.AsHistogram, argv.PrintPerLine)
	if err != nil {
		return err
	}

	// Print results
	_ = ctx.String("%v\n", result)

	return nil
}

// Run is the actual function of the app. It will print the color combination
// used or the histogram of an image.
func Run(pathFile string, asHistogram bool, printPerLine bool) (result string, err error) {
	imgRGBA, err := pallet.Load(pathFile)
	if err != nil {
		return "", errors.Wrap(err, "failed to load image")
	}

	if asHistogram {
		pl := pallet.AsHistogram(imgRGBA)

		return pl.InJSON(printPerLine)
	}

	pl := pallet.ByOccurrence(imgRGBA)

	return pl.InJSON(printPerLine)
}

// GetVersion returns the app version to display.
//
// It will use the `version` and `commit` value set during build.
// Though, this function works with "go install" which fetches the git tagged
// version if set.
func GetVersion() string {
	nameBin := util.GetNameBin()
	verBin := "(unknown)"
	idCommit := ""

	if commit != "" {
		idCommit = fmt.Sprintf(" (%v)", commit)
	}

	if version != "" {
		verBin = version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		verBin = buildInfo.Main.Version
	}

	return fmt.Sprintf("%s %s%s", nameBin, verBin, idCommit)
}
