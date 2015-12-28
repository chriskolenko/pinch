package plugins

import "github.com/webcanvas/pinch/shared/models"

type fakeplugin struct{}

func (p fakeplugin) Setup(data models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	RegisterPlugin("test", &fakeplugin{})
}
