package plugins

type fakeserviceplugin struct{}

func init() {
	RegisterServicePlugin("test", &fakeserviceplugin{})
}
