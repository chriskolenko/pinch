package git

import (
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type git struct {
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *git) Setup(pluginType models.PluginType, data models.Raw) (interface{}, error) {
	// get the version of git.
	opts := Options{}
	mapstructure.Decode(data, &opts)

	runner, err := CreateRunner(opts)
	return runner, err
}

func init() {
	plugins.RegisterPlugin("git", &git{})
}
