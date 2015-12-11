package engine

import "github.com/webcanvas/pinch/environment"

// Context holds the information for a run
type Context struct {
	Env environment.Env
}

// NewContext creates a new context
func NewContext(env environment.Env) *Context {
	return &Context{
		Env: env,
	}
}
