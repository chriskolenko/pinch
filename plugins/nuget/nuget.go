package nuget

import (
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type nuget struct {
	commander *commanders.Commander

	Version string
}

// Setup initializes the NuGet plugin
func (n *nuget) Setup(pluginType models.PluginType, data models.Raw) (interface{}, error) {
	// parse the options
	opts := Options{}
	mapstructure.Decode(data, &opts)

	return CreateRunner(opts)
}

func init() {
	plugins.RegisterPlugin("nuget", &nuget{})
}
