package golang

import (
	"regexp"

	"github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type golang struct {
	commander *commanders.Commander
	Version   string
}

type buildOpts struct {
	LDFlags   string
	OutputDir string
	AppName   string
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *golang) Setup() error {
	commander, err := commanders.Open("go")
	if err != nil {
		return err
	}

	// get the version
	out, err := commander.ExecOutput("version")
	if err != nil {
		return err
	}

	vers := versionex.Find(out)
	g.Version = string(vers)

	g.commander = commander

	logrus.WithFields(logrus.Fields{"Version": g.Version}).Debug("Git.Setup: Find version of git")
	return nil
}

// Exec runs the pinch
func (g *golang) Exec(data map[string]string) (models.Result, error) {
	// Default to shorthash if no action
	action, ok := data["action"]
	if !ok {
		action = "build"
	}

	switch action {
	case "build":
		// get the version of git.
		opts := buildOpts{}
		mapstructure.Decode(data, &opts)
		return g.build(opts)
	}

	// get the output directory if supplied or use our default one.
	return models.Result{}, nil
}

func (g golang) build(opts buildOpts) (models.Result, error) {
	logrus.WithFields(logrus.Fields{"LDFlags": opts.LDFlags}).Debug("At build")

	ldflags := opts.LDFlags
	outputdir := opts.OutputDir
	appname := opts.AppName

	if outputdir == "" {
		outputdir = ".build"
	}

	if appname == "" {
		appname = ""
	}

	args := []string{"build"}

	if ldflags != "" {
		// add the ldflags option.
		args = append(args, "-ldflags")
		args = append(args, opts.LDFlags)
	}

	// Now run it.
	output, err := g.commander.ExecOutput(args...)
	if err != nil {
		// What's the error?
		logrus.WithFields(logrus.Fields{"output": string(output), "err": err}).Debug("What's the output?")
	}
	logrus.Debug(string(output))

	return models.Result{}, err
}

func init() {
	g := &golang{}
	plugins.RegisterPinchPlugin("go", g)
}
