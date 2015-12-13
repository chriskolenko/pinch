package pinchers

import (
	"errors"
	"regexp"
	"strings"
)

var runtimeregexp = regexp.MustCompile(`\$\$([a-zA-Z0-9]+)`)

// ErrMissingValue is the error that occurs when the runtime variable cannot replace placeholders
var ErrMissingRuntimeVal = errors.New("Missing value")

// RuntimeVar holds a string and allows replacing of placeholders ie $$something
type RuntimeVar struct {
	raw      string
	replaces []string
}

// String returns a string that has replaced all the placeholders
func (r *RuntimeVar) String(vars map[string]string) (string, error) {
	str := r.raw

	if len(r.replaces) == 0 {
		return str, nil
	}

	var err error

	// find all the replaces.
	for _, r := range r.replaces {
		// remove the $$
		key := strings.Trim(r, "$")
		val, ok := vars[key]
		if !ok {
			err = ErrMissingRuntimeVal
		} else {
			str = strings.Replace(str, r, val, -1)
		}
	}

	return str, err
}

// NewRuntimeVar returns a new runtimevar
func NewRuntimeVar(raw string) RuntimeVar {
	replaces := runtimeregexp.FindAllString(raw, -1)
	return RuntimeVar{
		raw:      raw,
		replaces: replaces,
	}
}
