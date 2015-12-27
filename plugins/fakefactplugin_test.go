package plugins

import "github.com/webcanvas/pinch/shared/models"

type fakefactplugin struct{}

func (p fakefactplugin) Setup(data models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func (p fakefactplugin) Gather(data models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	RegisterFactPlugin("test", &fakefactplugin{})
}
