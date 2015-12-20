package pinchers

import (
	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

// PinchRunner to make it easier to read
type PinchRunner func(map[string]string) (models.Result, error)

// Pincher the encapsulated pincher from the pinch file
type Pincher struct {
	runtimeVars RuntimeVarMap
	runner      PinchRunner
}

// Run will run the pinch
func (p Pincher) Run(vars ...map[string]string) (models.Result, error) {
	opts, err := p.runtimeVars.Resolve(vars...)
	if err != nil {
		return models.Result{}, err
	}

	return p.runner(opts)
}

// NewFactPincher creates a pincher from a fact
func NewFactPincher(src map[string]string) (Pincher, error) {
	pincher := Pincher{}

	for key, value := range src {
		// find the plugin
		plugin, err := plugins.SetupFactPlugin(key)
		if err != nil {
			return pincher, err
		}

		logrus.WithFields(logrus.Fields{"plugin": plugin}).Debug("Plugin found")

		rtm := make(RuntimeVarMap)
		rtm.Fill(value)

		// set up the pincher
		pincher.runtimeVars = rtm
		pincher.runner = plugin.Gather
	}

	return pincher, nil
}

// NewServicePincher creates a pincher from a service
func NewServicePincher(src map[string]string) (Pincher, error) {
	pincher := Pincher{}

	for key, value := range src {
		rtm := make(RuntimeVarMap)
		rtm.Fill(value)

		// find the plugin
		plugin, err := plugins.SetupServicePlugin(key)
		if err != nil {
			return pincher, err
		}

		logrus.WithFields(logrus.Fields{"key": key, "plugin": plugin}).Debug("Plugin found")

		// set up the pincher
		pincher.runtimeVars = rtm
		pincher.runner = plugin.Ensure
	}

	return pincher, nil
}

// NewPinchPincher creates a pincher from a test
func NewPinchPincher(src map[string]string) (Pincher, error) {
	pincher := Pincher{}

	for key, value := range src {
		rtm := make(RuntimeVarMap)
		rtm.Fill(value)

		// find the plugin
		plugin, err := plugins.SetupPinchPlugin(key)
		if err != nil {
			return pincher, err
		}

		logrus.WithFields(logrus.Fields{"key": key, "plugin": plugin}).Debug("Plugin found")

		// set up the pincher
		pincher.runtimeVars = rtm
		pincher.runner = plugin.Exec
	}

	return pincher, nil
}
