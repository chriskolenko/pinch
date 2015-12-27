package git

import (
	"regexp"

	"github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

// FactOpts all the options for fact
type FactOpts struct {
	Action string
}

// Git exposted so it can be tested
type Git struct {
	commander *commanders.Commander
	Version   string
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *Git) Setup(models.Raw) (result models.Result, err error) {
	commander, err := commanders.Open("git")
	if err != nil {
		return
	}

	// get the version
	out, err := commander.ExecOutput("version")
	if err != nil {
		return
	}

	vers := versionex.Find(out)
	g.Version = string(vers)

	g.commander = commander

	logrus.WithFields(logrus.Fields{"Version": g.Version}).Debug("Git.Setup: Find version of git")
	return
}

// Gather return facts for git based on options
func (g *Git) Gather(data models.Raw) (models.Result, error) {
	// get the version of git.
	opts := &FactOpts{}
	mapstructure.Decode(data, &opts)

	// TODO validate the opts.

	// Default to shorthash if no action
	action := opts.Action
	if action == "" {
		action = "shorthash"
	}

	switch action {
	case "shorthash":
		return getShortHash(g.commander)
	case "longhash":
		return getLongHash(g.commander)
	}

	return models.Result{}, nil
}

func getShortHash(cmd *commanders.Commander) (result models.Result, err error) {
	args := []string{"log", "--pretty=format:'%h'", "-n", "1"}
	output, err := cmd.ExecOutput(args...)
	if err != nil {
		return
	}

	facts := map[string]string{
		"ShortHash": string(output),
	}
	result.Facts = facts
	return
}

func getLongHash(cmd *commanders.Commander) (result models.Result, err error) {
	args := []string{"log", "--pretty=format:'%H'", "-n", "1"}
	output, err := cmd.ExecOutput(args...)
	if err != nil {
		return
	}

	facts := map[string]string{
		"LongHash": string(output),
	}
	result.Facts = facts
	return
}

func init() {
	plugins.RegisterFactPlugin("git", &Git{})
}
