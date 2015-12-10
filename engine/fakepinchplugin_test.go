package engine

type fakepinchplugin struct{}

func init() {
	RegisterPinchPlugin("test", &fakeserviceplugin{})
}
