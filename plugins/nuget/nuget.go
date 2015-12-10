package nuget

import "github.com/webcanvas/pinch/engine"

// NuGet this is exported just for fun. This will be used inside the pinch engine
type NuGet struct{}

func init() {
	engine.RegisterFactPlugin("nuget", &NuGet{})
}
