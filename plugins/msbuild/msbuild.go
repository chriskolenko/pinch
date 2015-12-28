package msbuild

import (
	"fmt"
	"regexp"

	"golang.org/x/sys/windows/registry"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")
var registryexp = regexp.MustCompile(`\$\(Registry:HKEY_LOCAL_MACHINE\\(.*?)@(.*)\)`)

type msbuildversion struct {
	path32  string
	path64  string
	version string
}

type msbuild struct {
	versions []msbuildversion
}

func getToolsVerions() ([]string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\MSBuild\ToolsVersions`, registry.READ)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	sub, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func getKeyStringValue(path, name string) (val string, err error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.READ)
	if err != nil {
		return
	}
	defer k.Close()

	val, _, err = k.GetStringValue(name)
	if err != nil {
		return
	}

	// get the values of the match
	strs := registryexp.FindStringSubmatch(val)
	if strs == nil {
		return
	}

	// recurse until we don't need too.
	val, err = getKeyStringValue(strs[1], strs[2])

	return
}

func getMsBuildPaths(version string) (path64 string, path32 string, err error) {
	keypath := `SOFTWARE\Microsoft\MSBuild\ToolsVersions\` + version

	path64, err = getKeyStringValue(keypath, "MSBuildToolsPath")
	if err != nil {
		return
	}

	path32, _ = getKeyStringValue(keypath, "MSBuildToolsPath32")

	return
}

func getMsBuildVersions() ([]msbuildversion, error) {
	regversions, err := getToolsVerions()
	if err != nil {
		return nil, err
	}

	buildversions := make([]msbuildversion, len(regversions))

	for index, version := range regversions {
		// look it up
		path64, path32, err := getMsBuildPaths(version)
		if err != nil {
			return nil, err
		}

		buildversion := msbuildversion{
			path32:  path32,
			path64:  path64,
			version: version,
		}

		buildversions[index] = buildversion
	}

	return buildversions, nil
}

// Setup runs all the pre plugin stuff. IE finding versions
func (ms msbuild) Setup(models.PluginType, models.Raw) (interface{}, error) {
	versions, err := getMsBuildVersions()
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		err = fmt.Errorf("MSBuild is not installed on your system")
	}

	ms.versions = versions
	return nil, nil
}

// Ensure setups the service
func (ms msbuild) Exec(opts models.Raw) (result models.Result, err error) {

	// if we don't have any versions

	return
}

func init() {
	plugins.RegisterPlugin("msbuild", &msbuild{})
}
