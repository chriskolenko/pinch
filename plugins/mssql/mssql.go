package mssql

import (
	"errors"
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")
var errNotSupported = errors.New("Not supported")

type mssql struct {
	commander *commanders.Commander
	Version   string
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *mssql) Setup(pluginType models.PluginType, data models.Raw) (interface{}, error) {
	// we only support running this as a service.
	if pluginType != models.ServicePluginType {
		return nil, errNotSupported
	}

	// get the options for the plugin.
	opts := Options{}
	mapstructure.Decode(data, &opts)

	return CreateRunner(opts)
}

func init() {
	g := &mssql{}
	plugins.RegisterPlugin("mssql", g)
}
