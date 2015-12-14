package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/webcanvas/pinch/engine"
	"github.com/webcanvas/pinch/pinchers"
	"github.com/webcanvas/pinch/shared/environment"

	_ "github.com/webcanvas/pinch/plugins/git"
	_ "github.com/webcanvas/pinch/plugins/gitversion"
	_ "github.com/webcanvas/pinch/plugins/golang"
	_ "github.com/webcanvas/pinch/plugins/host"
	_ "github.com/webcanvas/pinch/plugins/iis"
	_ "github.com/webcanvas/pinch/plugins/iisexpress"
	_ "github.com/webcanvas/pinch/plugins/liquibase"
	_ "github.com/webcanvas/pinch/plugins/msbuild"
	_ "github.com/webcanvas/pinch/plugins/mssql"
	_ "github.com/webcanvas/pinch/plugins/nuget"
	_ "github.com/webcanvas/pinch/plugins/nunit"
	_ "github.com/webcanvas/pinch/plugins/print"
	_ "github.com/webcanvas/pinch/plugins/sonar"
)

func errored(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func pinch(c *cli.Context) {
	debug := c.Bool("debug")
	pinchfile := c.String("pinchfile")
	dotenv := c.String("env")

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	pinch, err := pinchers.Load(pinchfile)
	if err != nil {
		errored(err)
		return
	}

	env := environment.Load(dotenv)

	err = engine.Run(env, pinch)
	if err != nil {
		errored(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "pinch"
	app.Action = pinch
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Debug mode",
		},
		cli.StringFlag{
			Name:  "pinchfile, p",
			Usage: "Specify an alternate pinch file (default: .pinch.yml)",
		},
		cli.StringFlag{
			Name:  "env, e",
			Usage: "Specify additional environment variables (default: .env)",
		},
	}

	app.Run(os.Args)
}
