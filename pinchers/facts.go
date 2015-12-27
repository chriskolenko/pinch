package pinchers

import "github.com/Sirupsen/logrus"

// Facts is the parsed fact pinchers from a pinch file
type Facts []Pincher

func (facts Facts) add(fact Pincher) {
	facts = append(facts, fact)
}

// UnmarshalYAML parses fact pinchers from a pinch file
func (facts *Facts) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var maps []map[string]string
	if err := unmarshal(&maps); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing facts")
		return err
	}

	for _, m := range maps {
		// we a fact pincher.
		fact, err := NewFact(m)
		if err != nil {
			return err
		}

		facts.add(fact)
	}

	return nil
}
