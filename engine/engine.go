package engine

import (
	"errors"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/pinchers"
	"github.com/webcanvas/pinch/shared/environment"
)

// ErrNoPinch no parsed pinch file err
var ErrNoPinch = errors.New("No pinch file")

// Run will do the pinch
func Run(env environment.Env, pinchfile string) error {
	// get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// get the directory name so we can use it during a pinch
	info, err := os.Stat(dir)
	if err != nil {
		return err
	}

	// create a new context
	ctx := NewContext(env)

	// add the directory as the working directory
	ctx.AddFact("PWD", dir)

	// add the folder name as the appname
	ctx.AddFact("AppName", info.Name())

	// TODO what about build servers?

	return run(ctx, pinchfile)
}

func run(ctx Context, pinchfile string) error {
	logrus.WithFields(logrus.Fields{"pinchfile": pinchfile}).Debug("The current pinch file")

	pinch, err := pinchers.Load(pinchfile)
	if err != nil {
		return err
	}

	// run all the includes first.
	for _, inc := range pinch.Includes {
		logrus.WithFields(logrus.Fields{"include": inc}).Debug("The included pinch file")
		if err := run(ctx, inc); err != nil {
			return err
		}
	}

	for key, value := range pinch.Environment.Variables {
		v, err := value.String(ctx.Env)
		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{"key": key, "value": v}).Debug("Adding fact from environment value")
		ctx.AddFact(key, v)
	}

	logrus.Debug("Whats the environment vars?", pinch.Environment)

	// foreach fact run.
	for _, pincher := range pinch.Facts {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Services {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Pre {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Tests {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Post {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	return nil
}

func handle(ctx Context, runner Runner) error {
	// run the setup first
	r1, err := runner.Setup(ctx.Facts)
	if err != nil {
		return err
	}

	// add the result back to the context
	ctx.AddResult(r1)

	// TODO Maybe we should think about errors!.

	r2, err := runner.Run(ctx.Facts)
	if err != nil {
		return err
	}

	// add the result back to the context
	ctx.AddResult(r2)

	// TODO Maybe we should think about errors!.

	return nil
}
