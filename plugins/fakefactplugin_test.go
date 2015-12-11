package plugins

type fakefactplugin struct{}

func init() {
	RegisterFactPlugin("test", &fakefactplugin{})
}
