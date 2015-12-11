package nuget

import "github.com/webcanvas/pinch/plugins"

// NuGet this is exported just for fun. This will be used inside the pinch engine
type NuGet struct{}

func init() {
	plugins.RegisterFactPlugin("nuget", &NuGet{})
}
