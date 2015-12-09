package engine

import "github.com/webcanvas/pinch/environment"

// Engine is the brains of the pinch
type Engine struct{}

// Run will do the pinch
func (e *Engine) Run(env environment.Env) error {
	return nil
}

// Load turns a pinch file into a runable pinch
func Load(file string) (*Engine, error) {
	return &Engine{}, nil
}
