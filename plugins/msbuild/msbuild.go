package msbuild

import (
	"errors"
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")
var registryexp = regexp.MustCompile(`\$\(Registry:HKEY_LOCAL_MACHINE\\(.*?)@(.*)\)`)
var errNotInstalled = errors.New("MSBuild is not installed on your system")

type msbuild struct {
	initialized bool
	versions    Versions
}

func (ms *msbuild) ensure() error {
	if ms.initialized {
		return nil
	}

	var err error
	ms.initialized = true
	ms.versions, err = FindVersions()
	if err != nil {
		return err
	}

	return nil
}

// Setup runs all the pre plugin stuff. IE finding versions
func (ms msbuild) Setup(pluginType models.PluginType, data models.Raw) (interface{}, error) {
	err := ms.ensure()
	if err != nil {
		return nil, err
	}

	if len(ms.versions) == 0 {
		return nil, errNotInstalled
	}

	opts := Options{}
	mapstructure.Decode(data, &opts)

	return CreateRunner(ms.versions, opts)
}

func init() {
	plugins.RegisterPlugin("msbuild", &msbuild{})
}
