package engine

type fakeserviceplugin struct{}

func init() {
	RegisterServicePlugin("test", &fakeserviceplugin{})
}
