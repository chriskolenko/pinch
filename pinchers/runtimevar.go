package pinchers

import (
	"errors"
	"regexp"
	"strings"
)

var runtimeregexp = regexp.MustCompile(`\$\$([a-zA-Z0-9]+)`)

// ErrMissingValue is the error that occurs when the runtime variable cannot replace placeholders
var ErrMissingValue = errors.New("Missing value")

// RuntimeVar holds a string and allows replacing of placeholders ie $$something
type RuntimeVar struct {
	string
}

// String returns a string that has replaced all the placeholders
func (r *RuntimeVar) String(vars map[string]string) (string, error) {
	str := r.string

	// find all the replaces.
	replaces := runtimeregexp.FindAllString(str, -1)
	for _, r := range replaces {
		// remove the $$
		key := strings.Trim(r, "$")
		val, ok := vars[key]
		if !ok {
			return str, ErrMissingValue
		}
		str = strings.Replace(str, r, val, -1)
	}

	return str, nil
}
