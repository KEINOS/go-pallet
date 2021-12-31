package main

import (
	"fmt"
	"strings"

	"github.com/KEINOS/go-utiles/util"
)

var (
	version string
	commit  string
)

func main() {
	util.ExitOnErr(Run())
}

// Run is the actual function of the app.
func Run() error {
	fmt.Println("Hello, Gopher!")
	fmt.Println(GetVersion())

	return nil
}

func GetVersion() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", version, commit))
}
