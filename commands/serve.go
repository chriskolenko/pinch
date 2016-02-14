package commands

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/signal"
	"github.com/webcanvas/pinch/api/server"
	"github.com/webcanvas/pinch/capabilities"
	"github.com/webcanvas/pinch/ui"

	_ "github.com/webcanvas/pinch/modules/docker"
	_ "github.com/webcanvas/pinch/modules/git"
	_ "github.com/webcanvas/pinch/modules/msbuild"
)

const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

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

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: RFC3339NanoFixed,
		DisableColors:   false,
	})

	// load up the capabilities.
	caps, err := capabilities.Load()
	if err != nil {
		// TODO log the error.
		return 1
	}

	fmt.Println(caps)

	cfg := &server.Config{
		Logging: cfgDebug,
		Version: "11.33.22",
	}

	// create a new server
	api := server.New(cfg)

	logrus.Info("Daemon has completed initialization")

	// The serve API routine never exits unless an error occurs
	// We need to start it as a goroutine and wait on it so
	// daemon doesn't exit
	serveAPIWait := make(chan error)
	go api.Wait(serveAPIWait)

	signal.Trap(func() {
		api.Close()
		<-serveAPIWait
		// shutdownDaemon(d, 15)
	})

	return 0
}

func (Serve) Help() string {
	helpText := `
Usage: pinch serve [options]
  Will execute multiple builds in parallel as defined in the template.
  The various artifacts created by the template will be outputted.
Options:
  -debug                     Debug mode enabled for builds
`

	return strings.TrimSpace(helpText)
}

func (Serve) Synopsis() string {
	return "Run the API server"
}
