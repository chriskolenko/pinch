package engine

import (
	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/shared/environment"
	"github.com/webcanvas/pinch/shared/models"
)

// Context holds the information for a run
type Context struct {
	Env   environment.Env
	Facts map[string]string
}

// AddFact adds a key and value to existing facts
func (ctx Context) AddFact(key, val string) {
	ctx.Facts[key] = val
}

// AddResult adds the result back into the context for later reuse.
func (ctx Context) AddResult(result models.Result) {
	for key, value := range result.Facts {
		logrus.WithFields(logrus.Fields{"key": key, "value": value}).Debug("Adding fact value")
		ctx.AddFact(key, value)
	}
}

// NewContext creates a new context
func NewContext(env environment.Env) Context {
	return Context{
		Env:   env,
		Facts: make(map[string]string),
	}
}
