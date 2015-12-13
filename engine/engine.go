package engine

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/pinchers"
	"github.com/webcanvas/pinch/shared/environment"
)

// ErrNoPinch no parsed pinch file err
var ErrNoPinch = errors.New("No pinch")

// Run will do the pinch
func Run(env environment.Env, pinch *pinchers.Pinch) error {
	if pinch == nil {
		return ErrNoPinch
	}

	// TODO what about build servers?

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
	for _, pincher := range pinch.Facts.Pinchers {
		// lets run it.
		result, err := pincher.Pinch(ctx.Env, ctx.Facts)
		if err != nil {
			return err
		}

		for key, value := range result.Facts {
			logrus.WithFields(logrus.Fields{"key": key, "value": value}).Debug("Adding fact value")
			ctx.Facts[key] = value
		}
	}

	// TODO Service
	// TODO Pre

	for _, pincher := range pinch.Tests.Pinchers {
		// lets run it.
		result, err := pincher.Pinch(ctx.Facts, ctx.Env)
		if err != nil {
			return err
		}

		for key, value := range result.Facts {
			ctx.Facts[key] = value
		}
	}

	// TODO Post

	return nil
}
