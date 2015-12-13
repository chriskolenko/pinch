package nuget

import (
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

// NuGet this is exported just for fun. This will be used inside the pinch engine
type NuGet struct{}

// Setup initializes the NuGet plugin
func (n *NuGet) Setup() error {
	return nil
}

// Gather gathers all the facts for a pinch
func (n *NuGet) Gather(opts map[string]string) (*models.Result, error) {
	return nil, nil
}

func init() {
	plugins.RegisterFactPlugin("nuget", &NuGet{})
}
