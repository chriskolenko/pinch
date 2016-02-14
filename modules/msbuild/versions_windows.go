package msbuild

import (
	"errors"
	"regexp"

	"golang.org/x/sys/windows/registry"
)

var err32bitNotSupported = errors.New("32 bit not supported")
var registryexp = regexp.MustCompile(`\$\(Registry:HKEY_LOCAL_MACHINE\\(.*?)@(.*)\)`)

// Versions is an array of build versions
type Versions []Version

// Version holds information regarding msbuild paths
type Version struct {
	Path32  string
	Path64  string
	Version string
}

// GetPath returns the msbuild path.
func (version Version) GetPath(is32bit bool) (path string, err error) {
	if is32bit && version.Path32 == "" {
		err = err32bitNotSupported
		return
	}

	if is32bit {
		path = version.Path32
	} else {
		path = version.Path64
	}

	return
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

// FindVersions queries the registry to find the location of the MSBuild
func FindVersions() (Versions, error) {
	regversions, err := getToolsVerions()
	if err != nil {
		return nil, err
	}

	buildversions := make([]Version, len(regversions))

	for index, version := range regversions {
		// look it up
		path64, path32, err := getMsBuildPaths(version)
		if err != nil {
			return nil, err
		}

		buildversion := Version{
			Path32:  path32,
			Path64:  path64,
			Version: version,
		}

		buildversions[index] = buildversion
	}

	return buildversions, nil
}
