package engine

import "github.com/webcanvas/pinch/shared/environment"

// Context holds the information for a run
type Context struct {
	Env   environment.Env
	Facts map[string]string
}

// AddFact adds a key and value to existing facts
func (c Context) AddFact(key, val string) {
	c.Facts[key] = val
}

// NewContext creates a new context
func NewContext(env environment.Env) Context {
	return Context{
		Env:   env,
		Facts: make(map[string]string),
	}
}
