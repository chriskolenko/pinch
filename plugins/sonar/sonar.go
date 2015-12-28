package sonar

import (
	"regexp"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type sonar struct{}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *sonar) Setup(models.PluginType, models.Raw) (interface{}, error) {
	return nil, nil
}

// Ensure setups the service
func (g *sonar) Exec(opts models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	plugins.RegisterPlugin("sonar", &sonar{})
}
