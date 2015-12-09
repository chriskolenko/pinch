package engine

// ToMap parses a string to a map
func ToMap(str string) map[string]RuntimeVar {
	m := make(map[string]RuntimeVar)
	MergeMap(m, str)
	return m
}

// MergeMap parses a string and adds them to a map
func MergeMap(dest map[string]RuntimeVar, str string) {
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
					dest[key] = RuntimeVar{value}
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
		dest[key] = RuntimeVar{value}
	}
}
