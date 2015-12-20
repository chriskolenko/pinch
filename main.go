package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/webcanvas/pinch/engine"
	"github.com/webcanvas/pinch/shared/environment"

	_ "github.com/denisenkom/go-mssqldb"

	_ "github.com/webcanvas/pinch/plugins/git"
	_ "github.com/webcanvas/pinch/plugins/gitversion"
	_ "github.com/webcanvas/pinch/plugins/golang"
	_ "github.com/webcanvas/pinch/plugins/host"
	_ "github.com/webcanvas/pinch/plugins/iis"
	_ "github.com/webcanvas/pinch/plugins/iisexpress"
	_ "github.com/webcanvas/pinch/plugins/liquibase"
	_ "github.com/webcanvas/pinch/plugins/msbuild"
	_ "github.com/webcanvas/pinch/plugins/mssql"
	_ "github.com/webcanvas/pinch/plugins/mysql"
	_ "github.com/webcanvas/pinch/plugins/nuget"
	_ "github.com/webcanvas/pinch/plugins/nunit"
	_ "github.com/webcanvas/pinch/plugins/print"
	_ "github.com/webcanvas/pinch/plugins/sonar"
)

var version string

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

	env := environment.Load(dotenv)
	err := engine.Run(env, pinchfile)
	if err != nil {
		errored(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "pinch"
	app.Usage = "PINCHIE PINCHIE!!!!!!!!!!!!!"
	app.Version = version
	app.Action = pinch
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Debug mode",
		},
		cli.StringFlag{
			Name:  "pinchfile, p",
			Usage: "Specify an alternate pinch file (default: .pinch.yml)",
			Value: ".pinch.yml",
		},
		cli.StringFlag{
			Name:  "env, e",
			Usage: "Specify additional environment variables (default: .env)",
			Value: ".env",
		},
	}

	app.Run(os.Args)
}
