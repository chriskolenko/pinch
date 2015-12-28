package liquibase

import (
	"regexp"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type liquibase struct{}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *liquibase) Setup(models.PluginType, models.Raw) (interface{}, error) {
	return nil, nil
}

// Ensure setups the service
func (g *liquibase) Exec(opts models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	g := &liquibase{}
	plugins.RegisterPlugin("liquibase", g)
}
