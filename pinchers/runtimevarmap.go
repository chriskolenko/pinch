package pinchers

import "github.com/Sirupsen/logrus"

// RuntimeVarMap used to make it pretty
type RuntimeVarMap map[string]RuntimeVar

// Resolve a runtime map into a dictionary from many maps.
func (dest RuntimeVarMap) Resolve(vars map[string]string) (map[string]string, error) {
	// create options.
	opts := make(map[string]string)

	// add these to the opts
	for key, value := range vars {
		opts[key] = value
	}

	// go through the runtime vars
	for key, value := range dest {
		logrus.WithFields(logrus.Fields{"key": key, "value": value}).Debug("The runtime key and value")

		resolved, err := value.String(vars)
		if err != nil {
			return opts, err
		}
		opts[key] = resolved
	}

	return opts, nil
}

// Fill parses a string and adds them to a map
func (dest RuntimeVarMap) Fill(str string) {
	var groupingch *rune
	iskey := true
	key := ""
	value := ""

	for _, ch := range str {
		switch {
		case ch == '"' || ch == '\'':
			if groupingch == nil {
				cpy := rune(ch)
				groupingch = &cpy
				continue
			}

			if *groupingch == ch {
				groupingch = nil
				continue
			}

			fallthrough
		case ch == '=':
			if groupingch == nil {
				iskey = false
				continue
			}
			fallthrough
		case ch == ' ':
			if groupingch == nil {
				if key != "" {
					dest[key] = NewRuntimeVar(value)
					key = ""
					value = ""
					iskey = true
				}
				continue
			}
			fallthrough
		default:
			if iskey {
				key += string(ch)
			} else {
				value += string(ch)
			}
		}
	}

	if key != "" {
		dest[key] = NewRuntimeVar(value)
	}
}
