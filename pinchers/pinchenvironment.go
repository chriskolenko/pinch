package pinchers

import "github.com/Sirupsen/logrus"

// PinchEnvironment has all the environment variables from the pinch files
type PinchEnvironment struct {
	Variables map[string]RuntimeVar
}

// UnmarshalYAML parses environment variables from a pinch file
func (pe *PinchEnvironment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var strs []string
	if err := unmarshal(&strs); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing strings")
		return err
	}

	variables := make(map[string]RuntimeVar)
	for _, str := range strs {
		FillRuntimeVarMap(variables, str)
	}

	pe.Variables = variables
	return nil
}