package engine

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/environment"
)

var runtimeregexp = regexp.MustCompile(`\$\$([a-zA-Z0-9]+)`)

// ErrMissingValue is the error that occurs when the runtime variable cannot replace placeholders
var ErrMissingValue = errors.New("Missing value")

// Pinch is the parse pinched file
type Pinch struct {
	Language    string           `yaml:"language"`
	Environment PinchEnvironment `yaml:"environment"`
}

// PinchEnvironment has all the environment variables from the pinch files
type PinchEnvironment struct {
	Variables map[string]RuntimeVar
}

// UnmarshalYAML parses environment variables from a pinch file
func (pe *PinchEnvironment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var strs []string
	if err := unmarshal(&strs); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Debug("Has err when parsing strings")
		return err
	}

	variables := make(map[string]RuntimeVar)
	for _, str := range strs {
		MergeMap(variables, str)
	}

	pe.Variables = variables
	return nil
}

// PinchContext contains all the information regarding the context
type PinchContext struct {
	env environment.Env
}

// AddEnvironmentVariable adds an environment variable to the context
func (ctx *PinchContext) AddEnvironmentVariable(key string, value RuntimeVar) error {
	v, err := value.String(ctx.env)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{"key": key, "value": v}).Debug("Adding environment value")

	ctx.env[key] = v
	return nil
}

// NewPinchContext creates a pinch context for a pinch run
func NewPinchContext(env environment.Env) *PinchContext {
	return &PinchContext{
		env: env,
	}
}

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
