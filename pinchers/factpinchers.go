package pinchers

import "github.com/Sirupsen/logrus"

// FactPinchers is the parsed fact pinchers from a pinch file
type FactPinchers struct {
	Errors   []error
	Pinchers []Pincher
}

// UnmarshalYAML parses fact pinchers from a pinch file
func (pe *FactPinchers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var maps []map[string]string
	if err := unmarshal(&maps); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing facts")
		return err
	}

	for _, m := range maps {
		// we a fact pincher.
		pincher, err := NewFactPincher(m)
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
