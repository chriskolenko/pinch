package nunit

import "github.com/webcanvas/pinch/shared/models"

// Runner holds information running the nunit runner
type Runner struct {
	Consoles Consoles
}

// Run runs the nunit for all given dlls
func (r *Runner) Run() (models.Result, error) {
	result := models.Result{}

	// get the first found path to nunit console.
	console, err := r.Consoles.Find()
	if err != nil {
		return result, err
	}

	err = console.Run()
	return result, err
}

// NewRunner will create a new nunit runner
func NewRunner(opts Options) (*Runner, error) {
	consoles := make([]Console, 2)
	consoles[0] = NewConsole2(opts.WorkingDirectory)
	consoles[1] = NewConsole3(opts.WorkingDirectory)

	return &Runner{
		Consoles: consoles,
	}, nil
}
