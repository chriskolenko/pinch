package pinchers

import "github.com/Sirupsen/logrus"

// Services is the parsed service pinchers from a pinch file
type Services []Pincher

func (services Services) add(service Pincher) {
	services = append(services, service)
}

// UnmarshalYAML parses fact pinchers from a pinch file
func (services *Services) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var names []string
	if err := unmarshal(&names); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing services")
		return err
	}

	for _, name := range names {
		service, err := NewService(name)
		if err != nil {
			logrus.WithFields(logrus.Fields{"name": name, "err": err}).Debug("Could not create service")
			return err
		}

		services.add(service)
	}

	return nil
}
