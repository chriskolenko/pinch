package plugins

type fakepinchplugin struct{}

func init() {
	RegisterPinchPlugin("test", &fakeserviceplugin{})
}
