package pinchers

import "github.com/Sirupsen/logrus"

// PinchPinchers is the parsed test pinchers from a pinch file
type PinchPinchers struct {
	Errors   []error
	Pinchers []Pincher
}

// UnmarshalYAML parses test pinchers from a pinch file
func (pe *PinchPinchers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var maps []map[string]string
	if err := unmarshal(&maps); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing facts")
		return err
	}

	for _, m := range maps {
		// we a fact pincher.
		pincher, err := NewPinchPincher(m)
		if err != nil {
			logrus.WithFields(logrus.Fields{"err": err}).Debug("NewPinchPincher: had an error when newing up")
			pe.Errors = append(pe.Errors, err)
		} else {
			pe.Pinchers = append(pe.Pinchers, pincher)
		}
	}

	if len(pe.Errors) > 0 {
		return pe.Errors[0]
	}

	return nil
}
