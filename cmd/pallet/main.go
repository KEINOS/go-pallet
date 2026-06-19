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

//nolint:gochecknoglobals // Allow for app info.
var (
	version string // The app version assigned via build flag.
	commit  string // The commit ID assigned via build flag.
)

type argT struct {
	cli.Helper

	PathFileImg  string `cli:"f,file"    usage:"file path of an image to analyze"`
	AsHistogram  bool   `cli:"histogram" usage:"print the histogram of the given image in JSON"`
	PrintJSONL   bool   `cli:"jsonl"     usage:"print each color as a JSON object on its own line"`
	PrintPerLine bool   `cli:"p,perline" usage:"print each JSON element on its own line"`
	ShowVersion  bool   `cli:"v,version" usage:"print the app version"`
}

func main() {
	util.OsExit(cli.Run(
		new(argT),
		PreRun,
		color.Bold(util.GetNameBin()),
		"  Print the colors used in an image, or print its histogram as JSON.\n",
		color.Bold("Usage:"),
		fmt.Sprintf("  %s [options] <file path>", util.GetNameBin()),
	))
}

// PreRun selects the action for the given flags.
func PreRun(ctx *cli.Context) error {
	argv, _ := ctx.Argv().(*argT)
	args := ctx.Args()

	// Show app version
	if argv.ShowVersion {
		_ = ctx.String("%v\n", GetVersion())

		return nil
	}

	if argv.PrintJSONL && argv.PrintPerLine {
		return errors.New("--jsonl and --perline cannot be used together")
	}

	if argv.PrintJSONL && argv.AsHistogram {
		return errors.New("--jsonl and --histogram cannot be used together")
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
	result, err := Run(pathFileImg, argv.AsHistogram, argv.PrintPerLine, argv.PrintJSONL)
	if err != nil {
		return err
	}

	// Print results
	if argv.PrintJSONL {
		_ = ctx.String("%v", result)

		return nil
	}

	_ = ctx.String("%v\n", result)

	return nil
}

// Run returns either the color list or the histogram of an image.
func Run(pathFile string, asHistogram bool, printPerLine bool, printJSONL bool) (string, error) {
	imgRGBA, err := pallet.Load(pathFile)
	if err != nil {
		return "", errors.Wrap(err, "failed to load image")
	}

	if asHistogram {
		pl := pallet.AsHistogram(imgRGBA)
		result, err := pl.InJSON(printPerLine)

		return result, errors.Wrap(err, "failed to format the histogram pallet to JSON (main.Run())")
	}

	pixInfoList := pallet.ByOccurrence(imgRGBA)

	if printJSONL {
		result, err := pixInfoList.InJSONL()

		return result, errors.Wrap(err, "failed to format the occurrence pallet to JSON Lines (main.Run())")
	}

	result, err := pixInfoList.InJSON(printPerLine)

	return result, errors.Wrap(err, "failed to format the occurrence pallet to JSON (main.Run())")
}

// GetVersion returns the app version to display.
//
// It uses the `version` and `commit` values set during build.
// It also works with `go install`, which uses the Git tag version when it is
// available.
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
