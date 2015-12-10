package engine

type fakefactplugin struct{}

func init() {
	RegisterFactPlugin("test", &fakefactplugin{})
}
