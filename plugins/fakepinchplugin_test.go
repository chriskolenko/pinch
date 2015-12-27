package plugins

import "github.com/webcanvas/pinch/shared/models"

type fakepinchplugin struct{}

func (f *fakepinchplugin) Setup(models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func (f *fakepinchplugin) Exec(models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	RegisterPinchPlugin("test", &fakepinchplugin{})
}
