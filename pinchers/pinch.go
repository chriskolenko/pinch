package pinchers

import (
	"fmt"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/shared/runtime"

	"gopkg.in/yaml.v2"
)

// Pinch is the parse pinched file
type Pinch struct {
	Language    string      `yaml:"language"`
	Includes    Includes    `yaml:"includes"`
	Environment Environment `yaml:"environment"`
	Facts       Facts       `yaml:"facts"`
	Services    Services    `yaml:"services"`
	Setup       Setup       `yaml:"setup"`
	Build       Build       `yaml:"build"`
	Test        Test        `yaml:"test"`
	Post        Post        `yaml:"post"`
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

func parseString(data interface{}) (str string, err error) {
	str, ok := data.(string)
	if !ok {
		err = fmt.Errorf("Could not parse string: %s", data)
	}

	return
}

func parseArrayMapStringVarMap(data interface{}) ([]map[string]runtime.VarMap, error) {
	// it should be a list of strings!.
	items, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Could not parse services: %s", data)
	}

	arr := make([]map[string]runtime.VarMap, len(items))

	// loop througth the items to get the value.
	for index, item := range items {
		m := make(map[string]runtime.VarMap)

		// it could be just a string.
		str, ok := item.(string)
		if ok {
			// we just have a string. create a simple map
			m[str] = make(runtime.VarMap)
			arr[index] = m

			continue
		}

		// or it could be a map.
		itemmap, ok := item.(map[interface{}]interface{})
		if !ok {
			return nil, fmt.Errorf("Not support yaml type. It can be a string or a map of strings: %s", item)
		}

		for key, val := range itemmap {
			// try parse the item.
			kstr, ok := key.(string)
			if !ok {
				return nil, fmt.Errorf("Could not parse key as string in map: %s, %s", key, item)
			}

			// try parse the item.
			vstr, ok := val.(string)
			if !ok {
				return nil, fmt.Errorf("Could not parse value as string in map: %s, %s", val, item)
			}

			// now make the runtime var from vstr.
			m[kstr] = runtime.Parse(vstr)
		}

		// add the service
		arr[index] = m
	}

	return arr, nil
}

func parseArrayOfStrings(data interface{}) ([]string, error) {
	// it should be a list of strings!.
	items, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Could not parse array of strings: %s", data)
	}

	strs := make([]string, len(items))

	for index, item := range items {
		str, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("Could not parse array of strings for item: %s", item)
		}

		strs[index] = str
	}

	return strs, nil
}

func parseEnv(data interface{}) (env Environment, err error) {
	env = make(Environment)

	// it should be a list of strings!.
	items, ok := data.([]interface{})
	if !ok {
		err = fmt.Errorf("Could not parse environment: %s", data)
		return
	}

	for _, item := range items {
		str, ok := item.(string)
		if !ok {
			err = fmt.Errorf("Could not parse environment: %s", item)
			return
		}

		// woot!
		// parse to runtime var.map.
		m := runtime.Parse(str)
		for key, val := range m {
			env[key] = val
		}
	}

	return
}

// UnmarshalYAML parses the pinch yml
func (pinch *Pinch) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing facts")
		return err
	}

	for key, val := range raw {
		// we have a predefined list of supported properties.
		switch key {
		case "language":
			lang, err := parseString(val)
			if err != nil {
				return err
			}
			pinch.Language = lang
		case "includes":
			incs, err := parseArrayOfStrings(val)
			if err != nil {
				return err
			}
			pinch.Includes = incs
		case "environment":
			env, err := parseEnv(val)
			if err != nil {
				return err
			}
			pinch.Environment = env
		case "facts":
			facts, err := parseArrayMapStringVarMap(val)
			if err != nil {
				return err
			}
			pinch.Facts = facts
		case "services":
			sers, err := parseArrayMapStringVarMap(val)
			if err != nil {
				return err
			}
			pinch.Services = sers
		case "setup":
			setup, err := parseArrayMapStringVarMap(val)
			if err != nil {
				return err
			}
			pinch.Setup = setup
		case "build":
			build, err := parseArrayMapStringVarMap(val)
			if err != nil {
				return err
			}
			pinch.Build = build
		case "test":
			test, err := parseArrayMapStringVarMap(val)
			if err != nil {
				return err
			}
			pinch.Test = test
		case "post":
			post, err := parseArrayMapStringVarMap(val)
			if err != nil {
				return err
			}
			pinch.Post = post
		}
	}

	return nil
}
