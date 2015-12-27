package iisexpress

import (
	"regexp"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type iis struct {
	commander *commanders.Commander
	Version   string
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *iis) Setup(models.Raw) (result models.Result, err error) {
	return
}

// Ensure setups the service
func (g *iis) Ensure(opts models.Raw) (models.Result, error) {
	return models.Result{}, nil
}

func init() {
	g := &iis{}
	plugins.RegisterServicePlugin("iis", g)
}
