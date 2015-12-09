package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/webcanvas/pinch/engine"
	"github.com/webcanvas/pinch/environment"
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

	env := environment.Load(dotenv)

	eng, err := engine.Load(pinchfile)
	if err != nil {
		errored(err)
		return
	}

	err = eng.Run(env)
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
