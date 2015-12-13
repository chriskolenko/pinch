package pinchers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Pinch is the parse pinched file
type Pinch struct {
	Language    string           `yaml:"language"`
	Environment PinchEnvironment `yaml:"environment"`
	Facts       FactPinchers     `yaml:"facts"`
	Tests       PinchPinchers    `yaml:"test"`
}

// Load turns a pinch file into a runable pinch
func Load(file string) (*Pinch, error) {
	// load up the config.
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	pinch := &Pinch{}
	err = yaml.Unmarshal([]byte(raw), pinch)
	if err != nil {
		return nil, err
	}

	return pinch, nil
}
