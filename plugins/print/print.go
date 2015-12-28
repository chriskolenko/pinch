package print

import (
	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

type print struct{}

// Setup initializes the NuGet plugin
func (n *print) Setup(models.PluginType, models.Raw) (interface{}, error) {
	return n, nil
}

// Gather gathers all the facts for a pinch
func (n *print) Gather(opts models.Raw) (models.Result, error) {
	msg, ok := opts["msg"]
	if ok {
		logrus.Info("msg: ", msg)
	}
	return models.Result{}, nil
}

// Exec runs the pinch
func (n *print) Exec(opts models.Raw) (models.Result, error) {
	msg, ok := opts["msg"]
	if ok {
		logrus.Info("msg: ", msg)
	}
	return models.Result{}, nil
}

func init() {
	plugins.RegisterPlugin("print", &print{})
}
