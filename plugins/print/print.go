package print

import (
	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

type print struct{}

// Setup initializes the NuGet plugin
func (n *print) Setup() error {
	return nil
}

// Gather gathers all the facts for a pinch
func (n *print) Gather(opts map[string]string) (models.Result, error) {
	msg, ok := opts["msg"]
	if ok {
		logrus.Info("msg: ", msg)
	}
	return models.Result{}, nil
}

// Exec runs the pinch
func (n *print) Exec(opts map[string]string) (models.Result, error) {
	msg, ok := opts["msg"]
	if ok {
		logrus.Info("msg: ", msg)
	}
	return models.Result{}, nil
}

func init() {
	p := &print{}
	plugins.RegisterFactPlugin("print", p)
	// plugins.RegisterServicePlugin("print", p)
	plugins.RegisterPinchPlugin("print", p)
}
