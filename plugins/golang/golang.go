package golang

import (
	"path/filepath"
	"regexp"
	"runtime"

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

type testOpts struct {
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *golang) Setup(models.PluginType, models.Raw) (interface{}, error) {
	return nil, nil
}

// Exec runs the pinch
func (g *golang) Exec(data models.Raw) (models.Result, error) {
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
	case "test":
		opts := testOpts{}
		mapstructure.Decode(data, &opts)
		return g.test(opts)
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

	if runtime.GOOS == "windows" {
		appname = appname + ".exe"
	}

	output := filepath.Join(outputdir, appname)
	args := []string{"build", "-o", output}

	if ldflags != "" {
		// add the ldflags option.
		args = append(args, "-ldflags")
		args = append(args, opts.LDFlags)
	}

	// Now run it.
	o, err := g.commander.ExecOutput(args...)
	if err != nil {
		// What's the error?
		logrus.WithFields(logrus.Fields{"output": string(o), "err": err}).Debug("What's the output?")
	}
	logrus.Debug(string(o))

	// TODO check the result.
	return models.Result{}, err
}

func (g golang) test(opts testOpts) (models.Result, error) {
	args := []string{"test", "./..."}

	// Now run it.
	o, err := g.commander.ExecOutput(args...)
	if err != nil {
		// What's the error?
		logrus.WithFields(logrus.Fields{"output": string(o), "err": err}).Debug("What's the output?")
	}
	logrus.Debug(string(o))

	// TODO check the result.
	return models.Result{}, err
}

func init() {
	g := &golang{}
	plugins.RegisterPlugin("go", g)
}
