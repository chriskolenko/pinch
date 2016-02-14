package nunit

import (
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type nunit struct{}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *nunit) Setup(pluginType models.PluginType, data models.Raw) (interface{}, error) {
	opts := Options{}
	mapstructure.Decode(data, &opts)

	return NewRunner(opts)
}

func init() {
	plugins.RegisterPlugin("nunit", &nunit{})
}
