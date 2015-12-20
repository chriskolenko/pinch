package pinchers

import (
	"io/ioutil"

	"github.com/Sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

// Pinch is the parse pinched file
type Pinch struct {
	Language    string           `yaml:"language"`
	Includes    PinchIncludes    `yaml:"includes"`
	Environment PinchEnvironment `yaml:"environment"`
	Facts       FactPinchers     `yaml:"facts"`
	Services    ServicePinchers  `yaml:"services"`
	Pre         PinchPinchers    `yaml:"pre"`
	Tests       PinchPinchers    `yaml:"test"`
	Post        PinchPinchers    `yaml:"post"`
}

// Load turns a pinch file into a runable pinch
func Load(file string) (Pinch, error) {
	pinch := Pinch{}

	// load up the config.
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Could not load the pinch file")
		return pinch, err
	}

	err = yaml.Unmarshal([]byte(raw), &pinch)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Could not parse pinch yaml")
		return pinch, err
	}

	return pinch, nil
}
