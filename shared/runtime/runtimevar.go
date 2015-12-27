package runtime

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Sirupsen/logrus"
)

var runtimeregexp = regexp.MustCompile(`\$\$([a-zA-Z0-9_]+)`)

// ErrMissingRuntimeVal is the error that occurs when the runtime variable cannot replace placeholders
var ErrMissingRuntimeVal = errors.New("Missing value")

// Var holds a string and allows replacing of placeholders ie $$something
type Var struct {
	raw      string
	replaces []string
}

// String returns a string that has replaced all the placeholders
func (r Var) String(vars map[string]string) (string, error) {
	str := r.raw

	if len(r.replaces) == 0 {
		return str, nil
	}

	var err error

	// find all the replaces.
	for _, r := range r.replaces {
		logrus.WithFields(logrus.Fields{"r": r}).Debug("Replacement values")
		// remove the $$
		key := strings.Trim(r, "$$")
		val, ok := vars[key]
		if !ok {
			err = ErrMissingRuntimeVal
			logrus.WithFields(logrus.Fields{"key": key}).Debug("Could not find key inside vars")
		} else {
			str = strings.Replace(str, r, val, -1)
		}
	}

	return str, err
}

// NewVar returns a new runtime variable
func NewVar(raw string) Var {
	replaces := runtimeregexp.FindAllString(raw, -1)
	return Var{
		raw:      raw,
		replaces: replaces,
	}
}
