package pinchers

import "github.com/Sirupsen/logrus"

// ServicePinchers is the parsed service pinchers from a pinch file
type ServicePinchers struct {
	Errors   []error
	Pinchers []*Pincher
}

// UnmarshalYAML parses fact pinchers from a pinch file
func (pe *ServicePinchers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var maps []map[string]string
	if err := unmarshal(&maps); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing facts")
		return err
	}

	for _, m := range maps {
		// we have a service pincher.
		pincher, err := NewServicePincher(m)
		if err != nil {
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
