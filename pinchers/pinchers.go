package pinchers

import "github.com/Sirupsen/logrus"

// Pinchers is the parsed test pinchers from a pinch file
type Pinchers []Pincher

func (pinchers Pinchers) add(pincher Pincher) {
	pinchers = append(pinchers, pincher)
}

// UnmarshalYAML parses test pinchers from a pinch file
func (pinchers *Pinchers) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var maps []map[string]string
	if err := unmarshal(&maps); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing facts")
		return err
	}

	for _, m := range maps {
		// we a fact pincher.
		pincher, err := NewPincher(m)
		if err != nil {
			logrus.WithFields(logrus.Fields{"err": err}).Debug("NewPinchPincher: had an error when newing up")
			return err
		}

		pinchers.add(pincher)
	}

	return nil
}
