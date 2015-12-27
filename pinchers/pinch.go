package pinchers

import (
	"io/ioutil"

	"github.com/Sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

// Pinch is the parse pinched file
type Pinch struct {
	Language    string      `yaml:"language"`
	Includes    Includes    `yaml:"includes"`
	Environment Environment `yaml:"environment"`
	Facts       Facts       `yaml:"facts"`
	Services    Services    `yaml:"services"`
	Pre         Pinchers    `yaml:"pre"`
	Tests       Pinchers    `yaml:"test"`
	Post        Pinchers    `yaml:"post"`
}

// Load turns a pinch file into a runable pinch
func Load(file string) (pinch Pinch, err error) {
	// load up the config.
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Could not load the pinch file")
		return
	}

	err = yaml.Unmarshal([]byte(raw), &pinch)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Could not parse pinch yaml")
		return
	}

	return
}
