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
	if pinchfile == "" {
		return ErrNoPinch
	}

	ctx := NewContext(env)

	// add the working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// add the directory as the working directory
	ctx.AddFact("PWD", dir)

	// TODO what about build servers?

	return run(ctx, pinchfile)
}

func run(ctx Context, pinchfile string) error {

	pinch, err := pinchers.Load(pinchfile)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{"pinchfile": pinchfile}).Debug("We have a pinch file")

	// run all the includes first.
	for _, inc := range pinch.Includes {
		err := run(ctx, inc)
		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{"includes": inc}).Debug("What's the include?")
	}

	for key, value := range pinch.Environment.Variables {
		v, err := value.String(ctx.Env)
		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{"key": key, "value": v}).Debug("Adding fact from environment value")
		ctx.AddFact(key, v)
	}

	// foreach fact run.
	for _, pincher := range pinch.Facts.Pinchers {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Services.Pinchers {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Pre.Pinchers {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Tests.Pinchers {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	for _, pincher := range pinch.Post.Pinchers {
		err = handle(ctx, pincher)
		if err != nil {
			return err
		}
	}

	return nil
}

func handle(ctx Context, runner Runner) (err error) {
	result, err := runner.Run(ctx.Facts)
	if err != nil {
		return
	}

	for key, value := range result.Facts {
		logrus.WithFields(logrus.Fields{"key": key, "value": value}).Debug("Adding fact value")
		ctx.AddFact(key, value)
	}

	return
}
