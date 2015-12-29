package engine

import (
	"errors"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/pinchers"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/environment"
	"github.com/webcanvas/pinch/shared/models"
	"github.com/webcanvas/pinch/shared/runtime"
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
	ctx.AddFact("WD", dir)

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

	for key, value := range pinch.Environment {
		v, err := value.String(ctx.Env)
		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{"key": key, "value": v}).Debug("Adding fact from environment value")
		ctx.AddFact(key, v)
	}

	logrus.Debug("Whats the environment vars?", pinch.Environment)
	logrus.Debug("Whats the services?", pinch.Services)

	// foreach fact run.
	for _, item := range pinch.Facts {
		fact, err := loadPlugin(ctx, models.FactPluginType, item)
		if err != nil {
			return err
		}

		result, err := fact.Run()
		if err != nil {
			return err
		}

		ctx.AddResult(result)
	}

	for _, item := range pinch.Services {
		service, err := loadPlugin(ctx, models.ServicePluginType, item)
		if err != nil {
			return err
		}

		result, err := service.Run()
		if err != nil {
			return err
		}

		ctx.AddResult(result)
	}

	for _, item := range pinch.Setup {
		setup, err := loadPlugin(ctx, models.SetupPluginType, item)
		if err != nil {
			return err
		}

		result, err := setup.Run()
		if err != nil {
			return err
		}

		ctx.AddResult(result)
	}

	for _, item := range pinch.Build {
		build, err := loadPlugin(ctx, models.BuildPluginType, item)
		if err != nil {
			return err
		}

		result, err := build.Run()
		if err != nil {
			return err
		}

		ctx.AddResult(result)
	}

	for _, item := range pinch.Test {
		test, err := loadPlugin(ctx, models.TestPluginType, item)
		if err != nil {
			return err
		}

		result, err := test.Run()
		if err != nil {
			return err
		}

		ctx.AddResult(result)
	}

	for _, item := range pinch.Post {
		post, err := loadPlugin(ctx, models.PostPluginType, item)
		if err != nil {
			return err
		}

		result, err := post.Run()
		if err != nil {
			return err
		}

		ctx.AddResult(result)
	}

	return nil
}

func loadPlugin(ctx Context, typ models.PluginType, pluginDetails map[string]runtime.VarMap) (Runner, error) {
	if pluginDetails == nil {
		return nil, fmt.Errorf("No plugin defined")
	}

	wrapped := new(wrapped)

	for key, val := range pluginDetails {
		// TODO support adapters!.

		// load the plugin
		plugin, err := plugins.GetPlugin(key)
		if err != nil {
			return nil, err
		}

		// create the options for the plugin
		opts, err := val.Resolve(ctx.Facts)
		if err != nil {
			return nil, err
		}

		// setup the plugin.
		loaded, err := plugin.Setup(typ, opts)
		if err != nil {
			return nil, err
		}

		// This could cause a panic!.
		runner, ok := loaded.(Runner)
		if !ok {
			return nil, fmt.Errorf("Plugin not supported for type: %s, %v", key, typ)
		}

		wrapped.inner = runner
	}

	return wrapped, nil
}

type wrapped struct {
	inner Runner
}

func (w *wrapped) Run() (models.Result, error) {
	if w.inner == nil {
		return models.Result{}, fmt.Errorf("We don't have anything to run, the plugin wasn't created correctly")
	}

	return w.inner.Run()
}
