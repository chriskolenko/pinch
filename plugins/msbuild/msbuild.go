package msbuild

import (
	"regexp"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type msbuild struct {
	commander *commanders.Commander
	Version   string
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *msbuild) Setup() error {
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

// Ensure setups the service
func (g *msbuild) Exec(opts map[string]string) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	g := &msbuild{}
	plugins.RegisterPinchPlugin("msbuild", g)
}