package pinchers

// RuntimeVarMap used to make it pretty
type RuntimeVarMap map[string]RuntimeVar

// FillRuntimeVarMap parses a string and adds them to a map
func FillRuntimeVarMap(dest RuntimeVarMap, str string) {
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

// // ParseFactPincher will convert a map into a fact pincher
// func ParseFactPincher(source map[string]string) (*FactPincher, error) {
// 	var pincher *FactPincher
//
// 	for key, value := range source {
// 		// if we don't have a pincher
// 		if pincher == nil {
//
// 			// load the fact plugin
// 			plugin, err := SetupFactPlugin(key)
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			// parse opts.
// 			opts := ToMap(value)
//
// 			resolver := func(env environment.Env) (*FactResult, error) {
// 				// TODO check if we have a problem with the opts.
// 				// because we are in a loop etc.
//
// 				// TODO check plugin because we are in a loop.
//
// 				// resolve values.
// 				return nil, nil
// 			}
//
// 			// create the fact pincher
// 			pincher = &FactPincher{
// 				resolver: resolver,
// 			}
// 		} else {
// 			// TODO what about adapter?
// 		}
//
// 	}
//
// 	return pincher, nil
// }
