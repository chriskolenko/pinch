package mysql

import (
	"regexp"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type mysql struct{}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *mysql) Setup(models.PluginType, models.Raw) (interface{}, error) {
	return nil, nil
}

// Ensure setups the service
func (g *mysql) Ensure(opts models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	g := &mysql{}
	plugins.RegisterPlugin("mysql", g)
}
