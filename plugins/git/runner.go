package git

import (
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

// Runner holds the information needed for a git operation
type Runner struct {
	cmd  commanders.Commander
	opts Options

	Version string
}

// Run runs the plugin
func (gr *Runner) Run() (models.Result, error) {

	switch gr.opts.Action {
	case "shorthash":
		return getShortHash(gr.cmd)
	case "longhash":
		return getLongHash(gr.cmd)
	}

	return models.Result{}, nil
}

func getShortHash(cmd commanders.Commander) (result models.Result, err error) {
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

func getLongHash(cmd commanders.Commander) (result models.Result, err error) {
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

// CreateRunner creates an object that can do the git operation
func CreateRunner(opts Options) (*Runner, error) {
	cmd, err := commanders.Open("git")
	if err != nil {
		return nil, err
	}

	// get the version
	out, err := cmd.ExecOutput("version")
	if err != nil {
		return nil, err
	}

	bts := versionex.Find(out)
	version := string(bts)

	return &Runner{
		cmd:     *cmd,
		opts:    opts,
		Version: version,
	}, nil
}
