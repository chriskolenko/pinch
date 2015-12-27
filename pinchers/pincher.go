package pinchers

import (
	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
	"github.com/webcanvas/pinch/shared/runtime"
)

// PinchRunner to make it easier to read
type PinchRunner func(models.Raw) (models.Result, error)

// Pincher the encapsulated pincher from the pinch file
type Pincher struct {
	runtimeVars runtime.VarMap
	runner      PinchRunner
	setup       PinchRunner
}

// Setup is appart of the plugin lifecycle. This should happen before we run anything.
func (p Pincher) Setup(vars models.Raw) (models.Result, error) {
	// todo this should replace the way it's being done now.!
	return p.setup(vars)
}

// Run will run the pinch
func (p Pincher) Run(vars models.Raw) (models.Result, error) {
	opts, err := p.runtimeVars.Resolve(vars)
	if err != nil {
		return models.Result{}, err
	}

	return p.runner(opts)
}

// NewFact creates a pincher from a fact
func NewFact(src map[string]string) (Pincher, error) {
	pincher := Pincher{}

	for key, value := range src {
		// find the plugin
		plugin, err := plugins.LoadFactPlugin(key)
		if err != nil {
			logrus.WithFields(logrus.Fields{"key": key, "type": "fact"}).Debug("Plugin not found")
			return pincher, err
		}

		rtm := runtime.Parse(value)

		// set up the pincher
		pincher.runtimeVars = rtm
		pincher.runner = plugin.Gather
		pincher.setup = plugin.Setup
	}

	return pincher, nil
}

// NewService creates a pincher from a service
func NewService(name string) (pincher Pincher, err error) {
	// find the plugin
	plugin, err := plugins.LoadServicePlugin(name)
	if err != nil {
		logrus.WithFields(logrus.Fields{"key": name, "type": "service"}).Debug("Plugin not found")
		return
	}

	// set up the pincher
	pincher.runner = plugin.Ensure
	pincher.setup = plugin.Setup

	return
}

// NewPincher creates a pincher from a test
func NewPincher(src map[string]string) (Pincher, error) {
	pincher := Pincher{}

	for key, value := range src {
		rtm := runtime.Parse(value)

		// find the plugin
		plugin, err := plugins.LoadPinchPlugin(key)
		if err != nil {
			logrus.WithFields(logrus.Fields{"key": key, "type": "pincher"}).Debug("Plugin not found")
			return pincher, err
		}

		// set up the pincher
		pincher.runtimeVars = rtm
		pincher.runner = plugin.Exec
		pincher.setup = plugin.Setup
	}

	return pincher, nil
}
