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

// Pinch runs the commands and gives results
func (p *Pincher) Pinch(maps ...map[string]string) (models.Result, error) {
	// create options.
	opts := make(map[string]string)
	for key, value := range p.runtimeVars {
		var str string
		var err error

		rv := value

		for _, m := range maps {
			str, err = rv.String(m)
			if err == nil {
				break
			}

			rv = NewRuntimeVar(str)
		}

		// if the error is still not nil return
		if err != nil {
			return models.Result{}, err
		}

		// add the value to the opts
		opts[key] = str
	}

	return p.runner(opts)
}

// NewFactPincher creates a pincher from a fact
func NewFactPincher(src map[string]string) (*Pincher, error) {
	var pincher *Pincher

	for key, value := range src {
		// find the plugin
		plugin, err := plugins.SetupFactPlugin(key)
		if err != nil {
			return pincher, err
		}

		logrus.WithFields(logrus.Fields{"plugin": plugin}).Debug("Plugin found")

		rtm := make(RuntimeVarMap)
		FillRuntimeVarMap(rtm, value)

		// create the pincher.
		pincher = &Pincher{
			runtimeVars: rtm,
			runner:      plugin.Gather,
		}
	}

	return pincher, nil
}

// NewServicePincher creates a pincher from a service
func NewServicePincher(src map[string]string) (*Pincher, error) {
	var pincher *Pincher

	for key, value := range src {
		// find the plugin
		plugin, err := plugins.SetupServicePlugin(key)
		if err != nil {
			return pincher, err
		}

		logrus.WithFields(logrus.Fields{"plugin": plugin}).Debug("Plugin found")

		rtm := make(RuntimeVarMap)
		FillRuntimeVarMap(rtm, value)

		// create the pincher.
		pincher = &Pincher{
			runtimeVars: rtm,
			runner:      plugin.Ensure,
		}
	}

	return pincher, nil
}

// NewPinchPincher creates a pincher from a test
func NewPinchPincher(src map[string]string) (*Pincher, error) {
	var pincher *Pincher

	for key, value := range src {
		// find the plugin
		plugin, err := plugins.SetupPinchPlugin(key)
		if err != nil {
			return pincher, err
		}

		logrus.WithFields(logrus.Fields{"plugin": plugin}).Debug("Plugin found")

		rtm := make(RuntimeVarMap)
		FillRuntimeVarMap(rtm, value)

		// create the pincher.
		pincher = &Pincher{
			runtimeVars: rtm,
			runner:      plugin.Exec,
		}
	}

	return pincher, nil
}
