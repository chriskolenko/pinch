package engine

import "github.com/webcanvas/pinch/shared/environment"

// Context holds the information for a run
type Context struct {
	Env   environment.Env
	Facts map[string]string
}

// NewContext creates a new context
func NewContext(env environment.Env) Context {
	return Context{
		Env:   env,
		Facts: make(map[string]string),
	}
}
