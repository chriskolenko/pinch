package plugins

import "github.com/webcanvas/pinch/shared/models"

type fakeserviceplugin struct{}

func (f *fakeserviceplugin) Setup(models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func (f *fakeserviceplugin) Ensure(models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	RegisterServicePlugin("test", &fakeserviceplugin{})
}
