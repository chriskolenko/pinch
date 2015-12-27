package pinchers

import (
	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/shared/runtime"
)

// Environment has all the environment variables from the pinch files
type Environment struct {
	Variables runtime.VarMap
}

// UnmarshalYAML parses environment variables from a pinch file
func (ev *Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var strs []string
	if err := unmarshal(&strs); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing strings")
		return err
	}

	// this feels like there is something strange going on with the YAML parser.
	// we don't have to do this with arrays!!
	ev.Variables = make(runtime.VarMap)

	for _, str := range strs {
		vars := runtime.Parse(str)
		for key, val := range vars {
			logrus.WithFields(logrus.Fields{"key": key, "val": val}).Debug("EnvVars add!!!!!!!!!!!\n\n")
			ev.Variables[key] = val
		}
	}

	return nil
}
