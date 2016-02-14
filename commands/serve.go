package commands

import (
	"flag"
	"fmt"
	"strings"

	"github.com/webcanvas/pinch/capabilities"
	"github.com/webcanvas/pinch/ui"
)

type Serve struct {
	ui.Ui
}

func (c Serve) Run(args []string) int {
	var cfgDebug bool

	flags := flag.NewFlagSet("serve", flag.ContinueOnError)
	flags.Usage = func() { c.Ui.Say(c.Help()) }
	flags.BoolVar(&cfgDebug, "debug", false, "")
	if err := flags.Parse(args); err != nil {
		return 1
	}

	// args = flags.Args()
	// if len(args) != 1 {
	// 	flags.Usage()
	// 	return 1
	// }

	// create a new server
	// srv := server.New()

	// load up the capabilities.
	caps, err := capabilities.Load()
	if err != nil {
		// TODO log the error.
		return 1
	}

	fmt.Println(caps)

	// TODO start the server.

	return 0
}

func (Serve) Help() string {
	helpText := `
Usage: pinch serve [options] TEMPLATE
  Will execute multiple builds in parallel as defined in the template.
  The various artifacts created by the template will be outputted.
Options:
  -debug                     Debug mode enabled for builds
  -force                     Force a build to continue if artifacts exist, deletes existing artifacts
  -machine-readable          Machine-readable output
  -except=foo,bar,baz        Build all builds other than these
  -only=foo,bar,baz          Only build the given builds by name
  -parallel=false            Disable parallelization (on by default)
  -var 'key=value'           Variable for templates, can be used multiple times.
  -var-file=path             JSON file containing user variables.
`

	return strings.TrimSpace(helpText)
}

func (Serve) Synopsis() string {
	return "Run the API server"
}
