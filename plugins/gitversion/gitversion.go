package gitversion

import (
	"regexp"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type gitversion struct{}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *gitversion) Setup(models.PluginType, models.Raw) (interface{}, error) {
	return g, nil
}

func (g *gitversion) Run() (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	g := &gitversion{}
	plugins.RegisterPlugin("gitversion", g)
}
