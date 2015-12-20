package pinchers

import "github.com/Sirupsen/logrus"

// RuntimeVarMap used to make it pretty
type RuntimeVarMap map[string]RuntimeVar

// Resolve a runtime map into a dictionary from many maps.
func (dest RuntimeVarMap) Resolve(maps ...map[string]string) (map[string]string, error) {
	// create options.
	opts := make(map[string]string)
	for key, value := range dest {
		var str string
		var err error

		rv := value

		logrus.WithFields(logrus.Fields{"key": key, "value": value}).Debug("The runtime key and value")
		for _, m := range maps {
			str, err = rv.String(m)
			if err == nil {
				break
			}

			rv = NewRuntimeVar(str)
		}

		// if the error is still not nil return
		if err != nil {
			return opts, err
		}

		// add the value to the opts
		opts[key] = str
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
