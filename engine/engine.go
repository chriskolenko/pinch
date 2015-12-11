package engine

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/environment"
	"github.com/webcanvas/pinch/pinchers"
)

// ErrNoPinch no parsed pinch file err
var ErrNoPinch = errors.New("No pinch")

// Run will do the pinch
func Run(env environment.Env, pinch *pinchers.Pinch) error {
	if pinch == nil {
		return ErrNoPinch
	}

	ctx := NewContext(env)

	for key, value := range pinch.Environment.Variables {
		v, err := value.String(ctx.Env)
		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{"key": key, "value": v}).Debug("Adding environment value")
		ctx.Env[key] = v
	}

	// foreach fact run.
	for _, fact := range pinch.FactPinchers.Pinchers {
		err := ctx.RunFact(fact)
		if err != nil {
			return err
		}
	}

	return nil
}
