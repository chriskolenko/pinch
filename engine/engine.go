package engine

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/webcanvas/pinch/environment"
)

// ErrNoPinch no parsed pinch file err
var ErrNoPinch = errors.New("No pinch")

// Engine is the brains of the pinch
type Engine struct {
	Pinch *Pinch
}

// Run will do the pinch
func (e *Engine) Run(env environment.Env) error {
	if e.Pinch == nil {
		return ErrNoPinch
	}

	ctx := NewPinchContext(env)

	for key, value := range e.Pinch.Environment.Variables {
		err := ctx.AddEnvironmentVariable(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Load turns a pinch file into a runable pinch
func Load(file string) (*Engine, error) {
	// load up the config.
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	pinch := &Pinch{}
	err = yaml.Unmarshal([]byte(raw), pinch)
	if err != nil {
		return nil, err
	}

	eng := &Engine{
		Pinch: pinch,
	}

	return eng, nil
}
